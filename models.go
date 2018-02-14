package dcmd

type Config struct {
	PATH       string      `json:"dockerPath"`
	CONTAINERS []Container `json:"containers"`
}

type Container struct {
	Name    string   `json:"name"`
	Service []string `json:"service"`
}

func (c *Config) GetService(s string) *Container {
	for _, container := range c.CONTAINERS {
		if container.Name == s {
			return &container
		}
	}
	return nil
}

type DockerYAML struct {
	Version  string
	Services map[string]interface{} `yaml:"services"`
	Networks map[string]interface{} `yaml:"networks"`
	Volumes  map[string]interface{} `yaml:"volumes"`
}

//type services map[string]string
