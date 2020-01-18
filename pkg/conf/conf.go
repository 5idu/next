package conf

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Cf conf instance
var Cf *viper.Viper

func init() {
	Cf = viper.New()
	Cf.SetConfigName("config")                // name of config file (without extension)
	Cf.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	Cf.AddConfigPath(".")                     // optionally look for config in the working directory
	if err := Cf.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(errors.WithMessage(err, "fatal error config file"))
	}
}
