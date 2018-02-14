package dcmd

type AppConfig struct {
	Name        string
	Usage       string
	Version     string
	ConfigName  string
	CheckSymbol string
	CrossSymbol string
	EditSymbol  string
	up          string
	down        string
	Detached    string
	yamlName    string
	StorePath   string
}

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
