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
	first_path    = "./.cfg" // 优先配置
	default_path  = "./cfg"  // 默认配置
	cfg_file_name = "cfg.yaml"
)

type server struct {
	Addr  string
	Mode  string
	Node  int64
	Debug bool
}

type pprof struct {
	Enable bool
	Port   int
}

type apiDeadline struct {
	Default int
	Appoint map[string]int
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

type oss struct {
	Bucket    string
	AccessKey string
	SecretKey string
	Expire    int
}

func (db database) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.User, db.Passwd, db.Host, db.Port, db.Name) +
		"?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s"
}

type config struct {
	Server      server
	Pprof       pprof
	RateLimiter rateLimit
	ApiDeadline apiDeadline
	Token       token
	Log         log
	DB          database
	OSS         oss
}

func initCfg() {
	Config = &config{
		Server: server{
			Addr:  fmt.Sprintf(":%d", viper.GetInt("server.port")),
			Mode:  viper.GetString("server.mode"),
			Node:  viper.GetInt64("server.node"),
			Debug: viper.GetBool("server.debug"),
		},
		Pprof: pprof{
			Enable: viper.GetBool("pprof.enable"),
			Port:   viper.GetInt("pprof.port"),
		},
		RateLimiter: rateLimit{
			Limit: viper.GetFloat64("rate_limit.limit"),
			Burst: viper.GetInt("rate_limit.burst"),
		},
		ApiDeadline: apiDeadline{
			Default: viper.GetInt("api_deadline.default"),
			Appoint: mapToInt(viper.GetStringMap("api_deadline.appoint")),
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
		OSS: oss{
			Bucket:    viper.GetString("oss.bucket"),
			AccessKey: viper.GetString("oss.access_key"),
			SecretKey: viper.GetString("oss.secret_key"),
			Expire:    viper.GetInt("oss.expire"),
		},
	}
}

func mapToInt(m1 map[string]any) map[string]int {
	m2 := make(map[string]int, len(m1))
	for k, v := range m1 {
		if val, ok := v.(int); ok {
			m2[k] = val
		}
	}
	return m2
}

func defaultCfg() *viper.Viper {
	vp := viper.New()
	vp.AddConfigPath(first_path)
	vp.AddConfigPath(default_path)
	vp.SetConfigName(cfg_file_name)
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		panic(err)
	}
	return vp
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
	viper.AddConfigPath(default_path)
	setEnvConfig()
	vp := defaultCfg()
	for k, v := range vp.AllSettings() {
		viper.Set(k, v)
	}
	setFileConfig()
	initCfg()
}
