package local

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/convox/praxis/helpers"
	"github.com/convox/praxis/types"
	"github.com/kr/pty"
)

func (p *Provider) ProcessExec(app, pid, command string, opts types.ProcessExecOptions) (int, error) {
	if _, err := p.AppGet(app); err != nil {
		return 0, err
	}

	cmd := exec.Command("docker", "exec", "-it", pid, "sh", "-c", command)

	fd, err := pty.Start(cmd)
	if err != nil {
		return 0, err
	}

	go io.Copy(fd, opts.Input)
	go io.Copy(opts.Output, fd)

	if err := cmd.Wait(); err != nil {
		return 0, err
	}

	return 0, nil
}

func (p *Provider) ProcessGet(app, pid string) (*types.Process, error) {
	data, err := exec.Command("docker", "inspect", pid, "--format", "{{.ID}}").CombinedOutput()
	if err != nil {
		return nil, err
	}

	fpid := strings.TrimSpace(string(data))

	filters := []string{
		fmt.Sprintf("label=convox.app=%s", app),
		fmt.Sprintf("label=convox.rack=%s", p.Name),
		fmt.Sprintf("id=%s", fpid),
	}

	pss, err := processList(filters, true)
	if err != nil {
		return nil, err
	}

	if len(pss) != 1 {
		return nil, fmt.Errorf("no such process: %s", pid)
	}

	return &pss[0], nil
}

func (p *Provider) ProcessList(app string, opts types.ProcessListOptions) (types.Processes, error) {
	filters := []string{
		fmt.Sprintf("label=convox.app=%s", app),
		fmt.Sprintf("label=convox.rack=%s", p.Name),
	}

	if opts.Service != "" {
		filters = append(filters, fmt.Sprintf("label=convox.type=service"))
		filters = append(filters, fmt.Sprintf("label=convox.service=%s", opts.Service))
	}

	return processList(filters, false)
}

func (p *Provider) ProcessLogs(app, pid string, opts types.LogsOptions) (io.ReadCloser, error) {
	_, err := p.AppGet(app)
	if err != nil {
		return nil, err
	}

	ps, err := p.ProcessGet(app, pid)
	if err != nil {
		return nil, err
	}

	cr, cw := io.Pipe()

	args := []string{"logs"}

	if opts.Follow {
		args = append(args, "-f")
	}

	args = append(args, pid)

	cmd := exec.Command("docker", args...)

	cmd.Stdout = cw
	cmd.Stderr = cw

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	s := bufio.NewScanner(cr)

	rr, rw := io.Pipe()

	go func() {
		defer rw.Close()
		for s.Scan() {
			if opts.Prefix {
				fmt.Fprintf(rw, "%s %s/%s/%s %s\n", time.Now().Format(helpers.PrintableTime), ps.App, ps.Service, ps.Id, s.Text())
			} else {
				fmt.Fprintf(rw, "%s\n", s.Text())
			}
		}
	}()

	go func() {
		defer cw.Close()
		cmd.Wait()
	}()

	return rr, nil
}

func (p *Provider) ProcessProxy(app, pid string, port int, in io.Reader) (io.ReadCloser, error) {
	_, err := p.AppGet(app)
	if err != nil {
		return nil, err
	}

	data, err := exec.Command("docker", "inspect", pid, "--format", "{{.NetworkSettings.IPAddress}}").CombinedOutput()
	if err != nil {
		return nil, err
	}

	ip := strings.TrimSpace(string(data))

	cn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return nil, err
	}

	go io.Copy(cn, in)

	return cn, nil
}

func (p *Provider) ProcessRun(app string, opts types.ProcessRunOptions) (int, error) {
	if opts.Name != "" {
		exec.Command("docker", "rm", "-f", opts.Name).Run()
	}

	oargs, err := p.argsFromOpts(app, opts)
	if err != nil {
		return 0, err
	}

	cmd := exec.Command("docker", oargs...)

	if opts.Input != nil {
		rw, err := pty.Start(cmd)
		if err != nil {
			return 0, err
		}
		defer rw.Close()

		go io.Copy(rw, opts.Input)
		go io.Copy(opts.Output, rw)
	} else {
		cmd.Stdout = opts.Output
		cmd.Stderr = opts.Output

		if err := cmd.Start(); err != nil {
			return 0, err
		}
	}

	if err := cmd.Wait(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			if status, ok := ee.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus(), nil
			}
		}
	}

	return 0, err
}

func (p *Provider) ProcessStart(app string, opts types.ProcessRunOptions) (string, error) {
	if opts.Name != "" {
		exec.Command("docker", "rm", "-f", opts.Name).Run()
	}

	if opts.Name == "" {
		rs, err := types.Key(6)
		if err != nil {
			return "", err
		}

		opts.Name = fmt.Sprintf("%s.%s.process.%s.%s", p.Name, app, opts.Service, rs)
	}

	oargs, err := p.argsFromOpts(app, opts)
	if err != nil {
		return "", err
	}

	args := append(oargs[0:1], "--detach")
	args = append(args, oargs[1:]...)

	data, err := exec.Command("docker", args...).CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}

func (p *Provider) ProcessStop(app, pid string) error {
	return exec.Command("docker", "stop", "-t", "2", pid).Run()
}

func (p *Provider) argsFromOpts(app string, opts types.ProcessRunOptions) ([]string, error) {
	args := []string{"run", "--rm", "-i"}

	if opts.Input != nil {
		args = append(args, "-t")
	}

	image := opts.Image

	if image == "" {
		m, r, err := helpers.ReleaseManifest(p, app, opts.Release)
		if err != nil {
			return nil, err
		}

		s, err := m.Service(opts.Service)
		if err != nil {
			return nil, err
		}

		for _, v := range s.Volumes {
			args = append(args, "-v", v)
		}

		image = fmt.Sprintf("%s/%s/%s:%s", p.Name, app, opts.Service, r.Build)
	}

	if p.Frontend != "none" {
		args = append(args, "--dns", p.Frontend)
	}

	env, err := helpers.AppEnvironment(p, app)
	if err != nil {
		return nil, err
	}

	for k, v := range env {
		args = append(args, "-e", fmt.Sprintf("%s=%s", k, v))
	}

	for k, v := range opts.Environment {
		args = append(args, "-e", fmt.Sprintf("%s=%s", k, v))
	}

	for _, l := range opts.Links {
		args = append(args, "--link", l)
	}

	if opts.Memory > 0 {
		args = append(args, "--memory", fmt.Sprintf("%dM", opts.Memory))
	}

	if opts.Name != "" {
		args = append(args, "--name", opts.Name)
	}

	for from, to := range opts.Ports {
		args = append(args, "-p", fmt.Sprintf("%d:%d", from, to))
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	args = append(args, "-e", fmt.Sprintf("APP=%s", app))
	args = append(args, "-e", fmt.Sprintf("RACK_URL=https://%s:3000", hostname))
	args = append(args, "-e", fmt.Sprintf("RELEASE=%s", opts.Release))

	args = append(args, "--link", hostname)

	args = append(args, "--label", fmt.Sprintf("convox.app=%s", app))
	args = append(args, "--label", fmt.Sprintf("convox.rack=%s", p.Name))
	args = append(args, "--label", fmt.Sprintf("convox.release=%s", opts.Release))
	args = append(args, "--label", fmt.Sprintf("convox.service=%s", opts.Service))
	args = append(args, "--label", fmt.Sprintf("convox.type=%s", "process"))

	for from, to := range opts.Volumes {
		args = append(args, "-v", fmt.Sprintf("%s:%s", from, to))
	}

	args = append(args, image)

	if opts.Command != "" {
		args = append(args, "sh", "-c", opts.Command)
	}

	return args, nil
}

func processList(filters []string, all bool) (types.Processes, error) {
	args := []string{"ps"}

	if all {
		args = append(args, "-a")
	}

	for _, f := range filters {
		args = append(args, "--filter", f)
	}

	args = append(args, "--format", "{{json .}}")

	data, err := exec.Command("docker", args...).CombinedOutput()
	if err != nil {
		return nil, err
	}

	ps := types.Processes{}

	jd := json.NewDecoder(bytes.NewReader(data))

	for jd.More() {
		var dps struct {
			CreatedAt string
			Command   string
			ID        string
			Labels    string
		}

		if err := jd.Decode(&dps); err != nil {
			return nil, err
		}

		labels := map[string]string{}

		for _, kv := range strings.Split(dps.Labels, ",") {
			parts := strings.SplitN(kv, "=", 2)

			if len(parts) == 2 {
				labels[parts[0]] = parts[1]
			}
		}

		if labels["convox.service"] == "" {
			continue
		}

		started, err := time.Parse("2006-01-02 15:04:05 -0700 MST", dps.CreatedAt)
		if err != nil {
			return nil, err
		}

		ps = append(ps, types.Process{
			Id:      dps.ID,
			App:     labels["convox.app"],
			Command: strings.Trim(dps.Command, `"`),
			Release: labels["convox.release"],
			Service: labels["convox.service"],
			Started: started,
			Type:    labels["convox.type"],
		})
	}

	return ps, nil
}
