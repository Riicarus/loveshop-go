package conf

import (
	"os"

	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Gorm Gorm

	Cache Cache

	Jwt Jwt

	Server Server
}

type Gorm struct {
	Mysql Mysql
	Pool  Pool
}

type Mysql struct {
	Dsn string
}

type Pool struct {
	MaxIdle     int
	MaxOpen     int
	MaxLifeTime int
}

type Cache struct {
	Redis Redis
}

type Redis struct {
	Addr        string
	Password    string
	DB          int
	PoolSize    int
	DialTimeout int
}

type Server struct {
	Port int
}

type Jwt struct {
	Issuer string
	Expire int
	Secret string
}

var ServiceConf *ServiceConfig

func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/conf")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	ServiceConf = &ServiceConfig{
		Gorm: Gorm{
			Mysql: Mysql{
				Dsn: viper.GetString("Gorm.Mysql.Dsn"),
			},
			Pool: Pool{
				MaxIdle:     viper.GetInt("Gorm.Pool.MaxIdle"),
				MaxOpen:     viper.GetInt("Gorm.Pool.MaxOpen"),
				MaxLifeTime: viper.GetInt("Gorm.Pool.MaxLifeTime"),
			},
		},
		Cache: Cache{
			Redis: Redis{
				Addr:        viper.GetString("Cache.Redis.Addr"),
				Password:    viper.GetString("Cache.Redis.Password"),
				DB:          viper.GetInt("Cache.Redis.DB"),
				PoolSize:    viper.GetInt("Cache.Redis.PoolSize"),
				DialTimeout: viper.GetInt("Cache.Redis.DialTimeout"),
			},
		},
		Jwt: Jwt{
			Issuer: viper.GetString("Jwt.Issuer"),
			Expire: viper.GetInt("Jwt.Expire"),
			Secret: viper.GetString("Jwt.Secret"),
		},
		Server: Server{
			Port: viper.GetInt("Server.Port"),
		},
	}
}
