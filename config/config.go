package config

const APPSECRET = "dunedain"

type DeploymentEnvironment string
const DEBUG DeploymentEnvironment = "debug"
const PRODUCTION DeploymentEnvironment = "production"

const LOGPATH string = "/home/sid/Desktop/Workspace/go/src/shiriff/shiriff.logs"
const Namespace string = "Shiriff"

type Config struct {
	Namespace string
	Environment DeploymentEnvironment
	LogFilePath string
}
var configVal Config
func init() {
	configVal = Config {
		Namespace: "Shiriff",
		Environment: PRODUCTION,
		LogFilePath: LOGPATH,
	}
}

func GetConfig() Config {
	return configVal
}