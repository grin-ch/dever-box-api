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

type Server struct {
	Addr string
	Mode string
}

type Token struct {
	Expire int
	Signed string
	Issuer string
}
type config struct {
	Server Server
	Token  Token
}

func initCfg() {
	Config = &config{
		Server: Server{
			Addr: fmt.Sprintf(":%d", viper.GetViper().GetInt("server.port")),
			Mode: "",
		},
		Token: Token{
			Expire: viper.GetInt("token.expire"),
			Signed: viper.GetString("token.signed"),
			Issuer: viper.GetString("token.issuer"),
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
