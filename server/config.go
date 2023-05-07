package server

import "fmt"

type Config struct {
	AppName string
	Host    string
	Port    string
}

func NewConfig(appName, host, port string) *Config {
	return &Config{
		AppName: appName,
		Host:    host,
		Port:    port,
	}
}

func (c *Config) GetAppName() string {
	return c.AppName
}

func (c *Config) SetAppName(appName string) {
	c.AppName = appName
}

func (c *Config) GetHost() string {
	return c.Host
}

func (c *Config) SetHost(host string) {
	c.Host = host
}

func (c *Config) GetPort() string {
	return c.Port
}

func (c *Config) SetPort(port string) {
	c.Port = port
}

func (c *Config) ToString() string {
	return fmt.Sprintf("Config:{AppName:  %s, Host: %s, Port: %s}", c.GetAppName(), c.GetHost(), c.GetPort())
}
