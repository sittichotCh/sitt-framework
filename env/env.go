package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

func LoadConfig(prefix string, i interface{}) error {
	if err := LoadEnv(); err != nil {
		return err
	}

	if err := envconfig.Process(prefix, i); err != nil {
		return err
	}

	return nil
}

func LoadEnv() error {
	env := os.Getenv("ENV")
	fName := fmt.Sprintf(".env")
	if env != "" {
		fName = fmt.Sprintf("%s.%s", fName, env)
	}
	return godotenv.Load(fName)
}
