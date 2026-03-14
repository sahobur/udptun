package cmd

import (
	"fmt"
	"log/slog"

	"github.com/Russiancold/eviper"
	"github.com/go-viper/mapstructure/v2"
	"github.com/sahobur/udptun/internal/model"
	"github.com/spf13/viper"
)

type preauthKey string

type Config struct {
	Port    int
	Listen  string
	TunAddr string `mapstructure:"tun_addr"`
	//Clients map[preauthKey]Client
	ServerKey string         `mapstructure:"server_key"`
	CfgFile   string         `mapstructure:"config_file"`
	Clients   []model.Client `mapstructure:"clients"`
}

// todo vaidate
func (c *Config) Load() error {
	v := eviper.New(viper.New())
	v.AddConfigPath("/etc/udptun")
	v.AddConfigPath("../")
	v.AddConfigPath("../../")
	v.AddConfigPath(".")
	v.AddConfigPath("/app")

	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AutomaticEnv()
	err := v.Unmarshal(c)
	if err != nil {
		return err
	}
	c.CfgFile = v.ConfigFileUsed()
	slog.Info("config", slog.String("file:", c.CfgFile))
	return c.Validate()
}

func (c *Config) Validate() error {
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("invalid port: %d", c.Port)
	}
	if c.Listen == "" {
		return fmt.Errorf("listen address is required")
	}
	if c.TunAddr == "" {
		return fmt.Errorf("tun address is required")
	}
	if len(c.Clients) == 0 {
		return fmt.Errorf("at least one client is required")
	}
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

func (c *Config) AddClient(client model.Client) error {
	c.Clients = append(c.Clients, client)
	return c.Save()
}

func (c *Config) DeleteClient(name string) error {
	return nil
}
