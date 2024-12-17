package config

type Config struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Command string            `yaml:"command"`
	Args    []string          `yaml:"args,omitempty"`
	Env     map[string]string `yaml:"env,omitempty"`
}
