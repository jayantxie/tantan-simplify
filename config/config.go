package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// global config
type Config struct {
	Host       string     `toml:"host"`       // service host
	Port       int        `toml:"port"`       // service tcp port
	LogDir     string     `toml:"log_dir"`    // service log directory
	LogLevel   string     `toml:"log_level"`  // service log level
	PostgreSQL PostgreSQL `toml:"postgresql"` // postgreSQL config
}

// 操作 MySQL 对应的 MySQL 配置
type PostgreSQL struct {
	User     string `toml:"user"`     // MySQL user name
	Password string `toml:"password"` // MySQL user password
	Addr     string `toml:"addr"`     // MySQL host name
	DBName   string `toml:"db_name"`  // MySQL db name
}

func MustLoad(path string) (*Config, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	var conf Config
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
