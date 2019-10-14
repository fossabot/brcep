package flag

import (
	"os"
	"testing"

	gc "gopkg.in/check.v1"

	"github.com/leogregianin/brcep/config"
)

var _ = gc.Suite(&FlagLoaderSuite{})

type FlagLoaderSuite struct{}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gc.TestingT(t) }

func (s *FlagLoaderSuite) TestNewFlagLoaderShouldLoadValuesIntoConfig(c *gc.C) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{
		"brcep",
		"-address=:8080",
		"-log-level=test",
		"-preferred-api=cep-aberto",
		"-via-cep-url=http://localhost:8000/",
		"-cep-aberto-url=http://localhost:8010/",
		"-cep-aberto-token=token-sample",
		"-correios-url=http://localhost:8001/",
	}

	var (
		cfg    = &config.Config{}
		loader = NewFlagLoader()
	)

	loader.Load(cfg)
	c.Check(cfg.Address, gc.Equals, ":8080")
	c.Check(cfg.LogLevel, gc.Equals, "test")
	c.Check(cfg.PreferredAPI, gc.Equals, "cep-aberto")
	c.Check(cfg.ViaCepURL, gc.Equals, "http://localhost:8000/")
	c.Check(cfg.CepAbertoURL, gc.Equals, "http://localhost:8010/")
	c.Check(cfg.CepAbertoToken, gc.Equals, "token-sample")
	c.Check(cfg.CorreiosURL, gc.Equals, "http://localhost:8001/")
}
