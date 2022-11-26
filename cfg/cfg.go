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

type config struct {
	Server server
	Token  token
	Log    log
}

func initCfg() {
	Config = &config{
		Server: server{
			Addr:  fmt.Sprintf(":%d", viper.GetInt("server.port")),
			Mode:  viper.GetString("server.mode"),
			Node:  viper.GetInt64("server.node"),
			Debug: viper.GetBool("server.debug"),
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
