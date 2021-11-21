package config

import "context"

type ctxKey struct{}

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

func Ctx(ctx context.Context) Config {
	val, ok := ctx.Value(ctxKey{}).(Config)

	if ok {
		return val
	}

	return Config{}
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

func (c Config) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey{}, c)
}
