package config

import (
	"testing"

	gc "gopkg.in/check.v1"
)

var _ = gc.Suite(&ConfigSuite{})

type ConfigSuite struct{}

type MockLoader struct{}

func (f *MockLoader) Load(cfg *Config) {
	return
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gc.TestingT(t) }

func (s *ConfigSuite) TestNewConfigShouldReturnErrorIfNoLoaderProvided(c *gc.C) {
	cfg, err := NewConfig([]Loader{})

	c.Check(err, gc.NotNil)
	c.Check(cfg, gc.IsNil)
}

func (s *ConfigSuite) TestNewConfigShouldHaveDefaultValues(c *gc.C) {
	cfg, err := NewConfig([]Loader{&MockLoader{}})

	c.Check(err, gc.IsNil)

	c.Check(cfg.Address, gc.Equals, ":8000")
	c.Check(cfg.OperationMode, gc.Equals, "debug")
	c.Check(cfg.PreferredAPI, gc.Equals, "viacep")
	c.Check(cfg.CepAbertoToken, gc.Equals, "")
}
