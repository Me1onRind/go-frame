package initialize

import (
	"go-frame/internal/core/setting"
	"os"
)

func InitEnvironment() *setting.EnvironmentS {
	environment := &setting.EnvironmentS{}
	// set env
	env := getEnvVariable("env", "ENV")
	switch env {
	case "TEST", "LIVE":
		environment.Env = env
	default:
		environment.Env = "LOCAL"
	}

	return environment
}

func getEnvVariable(variableName ...string) string {
	var value string
	for _, v := range variableName {
		value = os.Getenv(v)
		if len(value) > 0 {
			break
		}
	}
	return value

}
