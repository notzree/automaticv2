package config

import "os"

type DeployEnvironment string

const (
	EnvProduction DeployEnvironment = "production"
	EnvSandbox    DeployEnvironment = "sandbox"
	EnvStaging    DeployEnvironment = "staging"
	EnvTest       DeployEnvironment = "test"
	EnvDev        DeployEnvironment = "dev"
)

func Env() DeployEnvironment {
	env := GetEnvironmentVariable("ENV")
	if env == "" {
		panic("invalid or missing ENV environment variable")
	}
	return DeployEnvironment(env)
}

func GetEnvironmentVariable(variable string) string {
	return os.Getenv(variable)
}
