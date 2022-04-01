package config

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Mysql *MysqlConfig
}

type MysqlConfig struct {
	Host            string
	Port            string
	User            string
	Passwd          string
	DB              string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
	Loc             string
}

func NewConfig() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("sentinel")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return nil, errors.Wrap(err, "viper: can't bind pflags")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "can't read config file")
	}

	return &Config{
		Mysql: &MysqlConfig{
			Host:            viper.GetString("mysql.host"),
			Port:            viper.GetString("mysql.port"),
			User:            viper.GetString("mysql.user"),
			Passwd:          viper.GetString("mysql.passwd"),
			DB:              viper.GetString("mysql.db"),
			MaxOpenConns:    viper.GetInt("mysql.maxOpenConns"),
			MaxIdleConns:    viper.GetInt("mysql.maxIdleConns"),
			ConnMaxLifetime: viper.GetInt("mysql.connMaxLifetimeMinutes"),
			Loc:             viper.GetString("mysql.loc"),
		},
	}, nil
}
