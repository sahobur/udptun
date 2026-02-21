package cmd

import (
	"github.com/Russiancold/eviper"
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

type preauthKey string

type Config struct {
	Port    int
	Listen  string
	TunAddr string `mapstructure:"tun_addr"`
	//Clients map[preauthKey]Client
	ServerKey string   `mapstructure:"server_key"`
	CfgFile   string   `mapstructure:"config_file"`
	Clients   []Client `mapstructure:"clients"`
}

type Client struct {
	Name string `mapstructure:"name"`
	Key  string `mapstructure:"key"`
	IP   string `mapstructure:"ip"`
}

func (c *Config) Load() error {
	v := eviper.New(viper.New())
	v.AddConfigPath("./config")
	v.AddConfigPath("/etc/udptun")
	v.AddConfigPath("../")
	v.AddConfigPath("./")
	v.SetConfigType("yaml")
	v.SetConfigFile("config.yaml")
	v.AutomaticEnv()
	err := v.Unmarshal(c)
	if err != nil {
		return err
	}
	c.CfgFile = v.ConfigFileUsed()
	return nil
}

func (c *Config) Save() error {
	var configMap map[string]interface{}
	err := mapstructure.Decode(c, &configMap)
	if err != nil {
		return err
	}
	if err = viper.MergeConfigMap(configMap); err != nil {
		return err
	}
	viper.SetConfigType("yaml")
	return viper.WriteConfigAs(c.CfgFile)
}

func (c *Config) AddClient(client Client) error {
	c.Clients = append(c.Clients, client)
	return c.Save()
}

func (c *Config) DeleteClient(name string) error {
	return nil
}
