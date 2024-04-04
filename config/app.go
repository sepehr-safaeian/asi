package config

type AutoInstaller struct {
	Version string
}

func AppVersion(AI *AutoInstaller) {
	AI.Version = "1.0.0"
}
