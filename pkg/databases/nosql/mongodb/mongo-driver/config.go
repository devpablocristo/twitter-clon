package pkgmongo

import (
	"fmt"
)

type config struct {
	User         string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

func newConfig(user, password, host, port, databaseName string) *config {
	return &config{
		User:         user,
		Password:     password,
		Host:         host,
		Port:         port,
		DatabaseName: databaseName,
	}
}

func (c *config) GetUser() string {
	return c.User
}

func (c *config) GetPassword() string {
	return c.Password
}

func (c *config) GetHost() string {
	return c.Host
}

func (c *config) GetPort() string {
	return c.Port
}

func (c *config) GetDatabaseName() string {
	return c.DatabaseName
}

func (c *config) DSN() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		c.User, c.Password, c.Host, c.Port, c.DatabaseName)
}

func (c *config) Database() string {
	return c.DatabaseName
}

func (c *config) Validate() error {
	if c.User == "" || c.Password == "" || c.Host == "" || c.Port == "" || c.DatabaseName == "" {
		return fmt.Errorf("incomplete MongoDB configuration")
	}
	return nil
}
