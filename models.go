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
