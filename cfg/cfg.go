package cfg

var (
	Config config
)

type config struct {
	Token struct {
		Expire int    `yaml:"token_expire"`
		Signed string `yaml:"token_signed"`
		Issuer string `yaml:"token_issuer"`
	}
}

func init() {
	initTestCfg()
}

func initTestCfg() {
	Config.Token.Expire = 15
	Config.Token.Signed = "dever-box"
	Config.Token.Issuer = "grin"
}
