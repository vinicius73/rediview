package config

type Config struct {
	version   string
	commit    string
	buildDate string
	appStage  string
}

func Build() Config {
	return Config{
		version:   version,
		commit:    commit,
		buildDate: buildDate,
		appStage:  appStage,
	}
}

func (c Config) Version() string {
	return c.version
}

func (c Config) Commit() string {
	return c.commit
}

func (c Config) BuildDate() string {
	return c.buildDate
}

func (c Config) AppStage() string {
	return c.appStage
}
