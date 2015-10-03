package config

import (
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"

	"github.com/ilkka/seita/cli"
)

// Config is the unmarshaled form of the configuration.
type Config struct {
	repo string
}

// Make sure a basic configuration exists.
func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.seita")

	err := viper.ReadInConfig()
	if err != nil {
		initializeConfig()
		write()
	}
}

// GetRuntimeConfig returns runtime configuration.
func GetRuntimeConfig() (cfg Config) {
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Error unmarshaling config: %s", err)
	}
	return cfg
}

// write writes configuration to user config file.
func write() {
	cfg := GetRuntimeConfig()

	cfgdir := path.Join(os.ExpandEnv("$HOME"), ".seita")
	err := os.MkdirAll(cfgdir, 0755)
	if err != nil {
		log.Fatalf("Could not create directory ~/.seita: %s", err)
	}

	cfgfilename := path.Join(cfgdir, "config.toml")
	f, err := os.Create(cfgfilename)
	if err != nil {
		log.Fatalf("Could not create file %s", cfgfilename)
	}

	defer f.Close()
	enc := toml.NewEncoder(f)
	enc.Encode(cfg)
}

// Take care of asking user to provide initial config
func initializeConfig() {
	cli.Printf("Initial configuration:\n")

	cli.Ask("What is the location of your seita repo? ", func(val string) error {
		if len(val) > 0 {
			viper.Set("repo", val)
			return nil
		}
		return cli.Errorf("This value is required")
	})
}
