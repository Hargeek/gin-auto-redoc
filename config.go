package ginautodoc

import "sync"

type Config struct {
	SwaggerPath string
}

var (
	defaultConfig = Config{
		SwaggerPath: "/api/v1/swagger",
	}
	globalConfig Config
	once         sync.Once
)

func SetConfig(cfg Config) {
	once.Do(func() {
		if cfg.SwaggerPath != "" {
			globalConfig.SwaggerPath = cfg.SwaggerPath
		} else {
			globalConfig.SwaggerPath = defaultConfig.SwaggerPath
		}
	})
}
