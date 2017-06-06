package manifest

import (
	"crypto/sha1"
	"fmt"
)

type Service struct {
	Name string

	Build       ServiceBuild   `yaml:"build,omitempty"`
	Certificate string         `yaml:"certificate,omitempty"`
	Command     ServiceCommand `yaml:"command,omitempty"`
	Environment []string       `yaml:"environment,omitempty"`
	Health      ServiceHealth  `yaml:"health,omitempty"`
	Image       string         `yaml:"image,omitempty"`
	Port        ServicePort    `yaml:"port,omitempty"`
	Resources   []string       `yaml:"resources,omitempty"`
	Scale       ServiceScale   `yaml:"scale,omitempty"`
	Volumes     []string       `yaml:"volumes,omitempty"`
}

type Services []Service

type ServiceBuild struct {
	Args []string `yaml:"args,omitempty"`
	Path string   `yaml:"path,omitempty"`
}

type ServiceCommand struct {
	Development string
	Test        string
	Production  string
}

type ServiceCount struct {
	Min int
	Max int
}

type ServiceHealth struct {
	Interval int
	Path     string
	Timeout  int
}

type ServicePort struct {
	Port   int
	Scheme string
}

type ServiceScale struct {
	Count  ServiceCount
	Memory int
}

func (s Service) BuildHash() string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("build[path=%q, args=%v] image=%q", s.Build.Path, s.Build.Args, s.Image))))
}

func (s Service) GetName() string {
	return s.Name
}
