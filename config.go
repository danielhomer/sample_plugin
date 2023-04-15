package sample_plugin

type Config struct {
	keySample string `mapstructure:"key"`
}

func (c *Config) InitDefaults() {
	if c.keySample == "" {
		c.keySample = "some_value"
	}
}
