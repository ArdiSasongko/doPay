package config

import "time"

type Config struct {
	App      AppConfig      `json:"app" mapstructure:"app"`
	Http     HttpConfig     `json:"http" mapstructure:"http"`
	Redis    RedisConfig    `json:"redis" mapstructure:"redis"`
	Database DatabaseConfig `json:"database" mapstructure:"database"`
}

type AppConfig struct {
	Address string `json:"address" mapstructure:"address"`
	Label   string `json:"label" mapstructure:"label"`
}

type HttpConfig struct {
	Timeout time.Duration `json:"timeout" mapstructure:"timeout"`
}

type RedisConfig struct {
	Address  string `json:"address" mapstructure:"address"`
	PoolSize int    `json:"poolsize" mapstructure:"poolsize"`
}

type DatabaseConfig struct {
	Host            string `json:"host" mapstructure:"host"`
	Port            string `json:"port" mapstructure:"port"`
	User            string `json:"user" mapstructure:"user"`
	Password        string `json:"password" mapstructure:"password"`
	DBName          string `json:"dbname" mapstructure:"dbname"`
	SSLMode         string `json:"sslmode" mapstructure:"sslmode"`
	MaxOpenConns    int    `json:"maxopenconns" mapstructure:"maxOpenConns"`
	MaxIdleConns    int    `json:"maxidleconns" mapstructure:"maxIdleConns"`
	ConnMaxIdleTime int    `json:"connmaxidletime" mapstructure:"connMaxIdletime"`
	ConnMaxLifeTime int    `json:"connmaxlifetime" mapstructure:"connMaxLifetime"`
}
