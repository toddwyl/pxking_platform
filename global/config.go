package global

import (
	"github.com/go-eagle/eagle/infrastructure/config"
	"github.com/go-eagle/eagle/infrastructure/logger"
	"github.com/go-eagle/eagle/infrastructure/storage/orm"
	"path/filepath"
	"time"
)

var (
	// Conf global app var
	Conf *Config
)

// Config global config
// nolint
type Config struct {
	Name              string
	Version           string
	Mode              string
	PprofPort         string
	URL               string
	JwtSecret         string
	JwtTimeout        int
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
	EnableTrace       bool
	EnablePprof       bool
	HTTP              ServerConfig
	GRPC              ServerConfig
	Logger            logger.Config
	Mysql             orm.Config
}

// ServerConfig server config.
type ServerConfig struct {
	Network      string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func InitConfig(cfgDir string, env string, filename string) *Config {
	c := config.New(filepath.Join(cfgDir, env))
	var cfg = Config{}
	if err := c.Load(filename, &cfg); err != nil {
		panic(err)
	}
	// set global
	Conf = &cfg
	return Conf
}
