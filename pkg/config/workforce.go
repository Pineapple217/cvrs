package config

type Workforce struct {
	MaxWorkers int `yaml:"maxWorkers"`
}

func (c *Workforce) SetDefault() {
	c.MaxWorkers = 5
}
