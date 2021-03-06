package local

import (
	"fmt"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/convox/praxis/helpers"
	"github.com/convox/praxis/manifest"
	"github.com/pkg/errors"
)

var convergeLock sync.Mutex

func (p *Provider) converge(app string) error {
	convergeLock.Lock()
	defer convergeLock.Unlock()

	log := p.logger("converge").Append("app=%q", app)

	m, r, err := helpers.AppManifest(p, app)
	if err != nil {
		return log.Error(err)
	}

	desired := []container{}

	c, err := p.balancerContainers(m.Balancers, app, r.Id, r.Stage)
	if err != nil {
		return errors.WithStack(log.Error(err))
	}

	desired = append(desired, c...)

	c, err = p.resourceContainers(m.Resources, app, r.Id)
	if err != nil {
		return errors.WithStack(log.Error(err))
	}

	desired = append(desired, c...)

	c, err = p.serviceContainers(m.Services, app, r.Id, r.Stage)
	if err != nil {
		return errors.WithStack(log.Error(err))
	}

	desired = append(desired, c...)

	// TODO: timers

	current, err := containersByLabels(map[string]string{
		"convox.rack": p.Name,
		"convox.app":  app,
	})
	if err != nil {
		return errors.WithStack(log.Error(err))
	}

	needed := []container{}

	for _, c := range desired {
		found := false

		for _, d := range current {
			if reflect.DeepEqual(c.Labels, d.Labels) {
				found = true
				break
			}
		}

		if !found {
			needed = append(needed, c)
		}
	}

	for _, c := range needed {
		p.storageLogWrite(fmt.Sprintf("apps/%s/releases/%s/log", app, r.Id), []byte(fmt.Sprintf("starting: %s\n", c.Name)))

		id, err := p.containerStart(c, app, r.Id)
		if err != nil {
			return errors.WithStack(log.Error(err))
		}

		c.Id = id

		if err := p.containerRegister(c); err != nil {
			return errors.WithStack(log.Error(err))
		}
	}

	for _, c := range current {
		if err := p.containerRegister(c); err != nil {
			return errors.WithStack(log.Error(err))
		}
	}

	return log.Success()
}

func (p *Provider) prune() error {
	convergeLock.Lock()
	defer convergeLock.Unlock()

	log := p.logger("prune")

	apps, err := p.AppList()
	if err != nil {
		return errors.WithStack(log.Error(err))
	}

	all, err := containersByLabels(map[string]string{
		"convox.rack": p.Name,
	})
	if err != nil {
		return errors.WithStack(log.Error(err))
	}

	for _, c := range all {
		found := false

		for _, a := range apps {
			if a.Name == c.Labels["convox.app"] {
				found = true
				break
			}
		}

		if !found {
			log.Successf("action=kill id=%s", c.Id)
			exec.Command("docker", "stop", c.Id).Run()
		}
	}

	return log.Success()
}

func resourcePort(kind string) (int, error) {
	switch kind {
	case "postgres":
		return 5432, nil
	case "redis":
		return 6379, nil
	}

	return 0, fmt.Errorf("unknown resource type: %s", kind)
}

func resourceURL(app, kind, name string) (string, error) {
	switch kind {
	case "postgres":
		return fmt.Sprintf("postgres://postgres:password@%s.resource.%s.convox:5432/app?sslmode=disable", name, app), nil
	case "redis":
		return fmt.Sprintf("redis://%s.resource.%s.convox:6379/0", name, app), nil
	}

	return "", fmt.Errorf("unknown resource type: %s", kind)
}

func resourceVolumes(app, kind, name string) ([]string, error) {
	switch kind {
	case "postgres":
		return []string{fmt.Sprintf("/var/convox/%s/resource/%s:/var/lib/postgresql/data", app, name)}, nil
	case "redis":
		return []string{}, nil
	}

	return []string{}, fmt.Errorf("unknown resource type: %s", kind)
}

func (p *Provider) balancerContainers(balancers manifest.Balancers, app, release string, stage int) ([]container, error) {
	cs := []container{}

	// don't run balancers in test stage
	if stage == manifest.StageTest {
		return cs, nil
	}

	sys, err := p.SystemGet()
	if err != nil {
		return nil, err
	}

	for _, b := range balancers {
		for _, e := range b.Endpoints {
			command := []string{}

			switch {
			case e.Redirect != "":
				command = []string{"balancer", e.Protocol, "redirect", e.Redirect}
			case e.Target != "":
				command = []string{"balancer", e.Protocol, "target", e.Target}
			default:
				return nil, fmt.Errorf("invalid balancer endpoint: %s:%s", b.Name, e.Port)
			}

			cs = append(cs, container{
				Name:     fmt.Sprintf("%s.%s.balancer.%s", p.Name, app, b.Name),
				Hostname: fmt.Sprintf("%s.balancer.%s.%s", b.Name, app, p.Name),
				Port: containerPort{
					Host:      443,
					Container: 3000,
				},
				Memory:  64,
				Image:   sys.Image,
				Command: command,
				Labels: map[string]string{
					"convox.rack":    p.Name,
					"convox.version": p.Version,
					"convox.app":     app,
					"convox.release": release,
					"convox.type":    "balancer",
					"convox.name":    b.Name,
					"convox.port":    e.Port,
				},
			})
		}
	}

	return cs, nil
}

func (p *Provider) resourceContainers(resources manifest.Resources, app, release string) ([]container, error) {
	cs := []container{}

	for _, r := range resources {
		rp, err := resourcePort(r.Type)
		if err != nil {
			return nil, err
		}

		vs, err := resourceVolumes(app, r.Type, r.Name)
		if err != nil {
			return nil, err
		}

		cs = append(cs, container{
			Name:     fmt.Sprintf("%s.%s.resource.%s", p.Name, app, r.Name),
			Hostname: fmt.Sprintf("%s.resource.%s.%s", r.Name, app, p.Name),
			Port: containerPort{
				Host:      rp,
				Container: rp,
			},
			Image:   fmt.Sprintf("convox/%s", r.Type),
			Volumes: vs,
			Labels: map[string]string{
				"convox.rack":     p.Name,
				"convox.version":  p.Version,
				"convox.app":      app,
				"convox.release":  release,
				"convox.type":     "resource",
				"convox.name":     r.Name,
				"convox.resource": r.Type,
			},
		})
	}

	return cs, nil
}

func (p *Provider) serviceContainers(services manifest.Services, app, release string, stage int) ([]container, error) {
	cs := []container{}

	// don't run background services in test stage
	if stage == manifest.StageTest {
		return cs, nil
	}

	sys, err := p.SystemGet()
	if err != nil {
		return nil, err
	}

	m, r, err := helpers.ReleaseManifest(p, app, release)
	if err != nil {
		return nil, err
	}

	for _, s := range services {
		if s.Port.Port > 0 {
			cs = append(cs, container{
				Name:     fmt.Sprintf("%s.%s.endpoint.%s", p.Name, app, s.Name),
				Hostname: fmt.Sprintf("%s.%s.%s", s.Name, app, p.Name),
				Port: containerPort{
					Host:      443,
					Container: 3000,
				},
				Memory:  64,
				Image:   sys.Image,
				Command: []string{"balancer", "https", "target", fmt.Sprintf("%s://%s:%d", s.Port.Scheme, s.Name, s.Port.Port)},
				Labels: map[string]string{
					"convox.rack":    p.Name,
					"convox.version": p.Version,
					"convox.app":     app,
					"convox.release": release,
					"convox.type":    "endpoint",
					"convox.name":    s.Name,
					"convox.port":    strconv.Itoa(s.Port.Port),
				},
			})
		}

		var command string

		switch stage {
		case manifest.StageDevelopment:
			command = s.Command.Development
		case manifest.StageTest:
			return nil, fmt.Errorf("can not run background services in test")
		case manifest.StageProduction:
			command = s.Command.Production
		default:
			return nil, fmt.Errorf("unknown stage: %d", stage)
		}

		cmd := []string{}

		if c := strings.TrimSpace(command); c != "" {
			cmd = append(cmd, "sh", "-c", c)
		}

		env, err := m.ServiceEnvironment(s.Name)
		if err != nil {
			return nil, err
		}

		// copy the map so we can hold on to it
		e := map[string]string{}

		for k, v := range env {
			e[k] = v
		}

		// add resources
		for _, sr := range s.Resources {
			for _, r := range m.Resources {
				if r.Name == sr {
					u, err := resourceURL(app, r.Type, r.Name)
					if err != nil {
						return nil, err
					}

					e[fmt.Sprintf("%s_URL", strings.ToUpper(sr))] = u
				}
			}
		}

		for i := 1; i <= s.Scale.Count.Min; i++ {
			cs = append(cs, container{
				Name:    fmt.Sprintf("%s.%s.service.%s.%d", p.Name, app, s.Name, i),
				Image:   fmt.Sprintf("%s/%s/%s:%s", p.Name, app, s.Name, r.Build),
				Command: cmd,
				Env:     e,
				Memory:  s.Scale.Memory,
				Volumes: s.Volumes,
				Labels: map[string]string{
					"convox.rack":    p.Name,
					"convox.version": p.Version,
					"convox.app":     app,
					"convox.release": release,
					"convox.type":    "service",
					"convox.name":    s.Name,
					"convox.service": s.Name,
					"convox.index":   fmt.Sprintf("%d", i),
				},
			})
		}
	}

	return cs, nil
}
