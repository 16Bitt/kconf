package dockercompose

// Version is the docker-compose version we're targetting
const Version = "3"

// DockerCompose is the root object in a docker-compose file
type DockerCompose struct {
	Version  string
	Services map[string]Service
}

// Service represents an image within a docker-compose application
type Service struct {
	Name        string            `yaml:",omitempty"`
	Image       string            `yaml:",omitempty"`
	Build       string            `yaml:",omitempty"`
	Ports       []string          `yaml:",omitempty"`
	Volumes     []string          `yaml:",omitempty"`
	Environment map[string]string `yaml:",omitempty"`
}

// New returns a fresh docker-compose struct
func New() DockerCompose {
	return DockerCompose{
		Version:  Version,
		Services: make(map[string]Service),
	}
}
