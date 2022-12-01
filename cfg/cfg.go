package cfg

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var (
	Config *config
)

const (
	server_name = "DEVER_BOX"
)

var (
	cfg_paths     = []string{"./", "./cfg"}
	cfg_file_name = "cfg.yaml"
)

type server struct {
	Addr  string
	Mode  string
	Node  int64
	Debug bool
}

type rateLimit struct {
	Limit float64
	Burst int
}
type token struct {
	Expire int
	Signed string
	Issuer string
}

type log struct {
	Path      string
	Level     int
	MaxAge    int
	HasCollor bool
	HasCaller bool
}

type database struct {
	Port   int
	Host   string
	Name   string
	User   string
	Passwd string
}

func (db database) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.User, db.Passwd, db.Host, db.Port, db.Name) +
		"?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s"
}

type config struct {
	Server      server
	RateLimiter rateLimit
	Token       token
	Log         log
	DB          database
}

func initCfg() {
	Config = &config{
		Server: server{
			Addr:  fmt.Sprintf(":%d", viper.GetInt("server.port")),
			Mode:  viper.GetString("server.mode"),
			Node:  viper.GetInt64("server.node"),
			Debug: viper.GetBool("server.debug"),
		},
		RateLimiter: rateLimit{
			Limit: viper.GetFloat64("rate_limit.limit"),
			Burst: viper.GetInt("rate_limit.burst"),
		},
		Token: token{
			Expire: viper.GetInt("token.expire"),
			Signed: viper.GetString("token.signed"),
			Issuer: viper.GetString("token.issuer"),
		},
		Log: log{
			Path:      viper.GetString("log.path"),
			Level:     viper.GetInt("log.level"),
			MaxAge:    viper.GetInt("log.max_age"),
			HasCollor: viper.GetBool("log.has_collor"),
			HasCaller: viper.GetBool("log.has_caller"),
		},
		DB: database{
			Port:   viper.GetInt("database.port"),
			Host:   viper.GetString("database.host"),
			Name:   viper.GetString("database.name"),
			User:   viper.GetString("database.user"),
			Passwd: viper.GetString("database.passwd"),
		},
	}
}

func setFileConfig() {
	viper.SetConfigName(cfg_file_name)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func setEnvConfig() {
	viper.SetEnvPrefix(server_name)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func init() {
	for _, path := range cfg_paths {
		viper.AddConfigPath(path)
	}
	setEnvConfig()
	setFileConfig()
	initCfg()
}
