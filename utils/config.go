package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Unknwon/goconfig"
)

type Config struct {
	File *goconfig.ConfigFile
}

func NewConfig() *Config {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir = strings.Replace(dir, "\\", "/", -1)
	config, err := goconfig.LoadConfigFile(dir + "/config/config.ini")
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Config{
		File: config,
	}
}

func (c *Config) GetServerPath() (string, error) {
	path, err := c.File.GetValue("proxy_server", "path")
	if err != nil {
		return "", err
	}
	return path, nil
}

func (c *Config) GetServerPort() (string, error) {
	path, err := c.File.GetValue("proxy_server", "port")
	if err != nil {
		return "", err
	}
	return path, nil
}

func (c *Config) GetUsernameAndPassword() (string, string, error) {
	username, err := c.File.GetValue("proxy_server", "username")
	if err != nil {
		return "", "", err
	}
	password, err := c.File.GetValue("proxy_server", "password")
	if err != nil {
		return "", "", err
	}
	return username, password, nil
}

func (c *Config) GetServicesFilePath() (string, error) {
	serviceFile, err := c.File.GetValue("proxy_server", "services")
	if err != nil {
		return "", err
	}
	return dir + serviceFile, nil
}

func (c *Config) UpdateAutoStart(autoStart string) (isSuccess bool) {
	c.File.DeleteKey("config", "auto_start")
	isSuccess = c.File.SetValue("config", "auto_start", autoStart)
	goconfig.SaveConfigFile(c.File, dir+"/config/config.ini")
	return
}

func (c *Config) UpdateProxy(proxy string) (isSuccess bool) {
	c.File.DeleteKey("config", "proxy")
	isSuccess = c.File.SetValue("config", "proxy", proxy)
	goconfig.SaveConfigFile(c.File, dir+"/config/config.ini")
	return
}

func (c *Config) GetAutoStart() (autoStart bool) {
	autoStart = c.File.MustBool("config", "auto_start")
	return
}

func (c *Config) GetProxySetting() (proxy bool) {
	proxy = c.File.MustBool("config", "proxy")
	return
}
