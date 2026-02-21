package model

type Client struct {
	Name string `mapstructure:"name"`
	Key  string `mapstructure:"key"`
	IP   string `mapstructure:"ip"`
}
