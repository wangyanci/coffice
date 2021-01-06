package setting

import (
	"fmt"

	"github.com/astaxie/beego/config"
	"github.com/wangyanci/coffice/utils"
)

var Conf config.Configer

var (
	DataBaseDriver string
	DataBaseAddr   string
	DataBasePort   string
	DataBaseName   string
	DataBaseUser   string
	DataBasePasswd string

	SigningKey string

	LogAdapter string
	LogLevel   int
	LogCache   int64
	LogPath    string
)

func init() {
	//	conf, err := config.NewConfig("ini", "E:\\workspace\\workspace_go\\src\\com.609.huangsewangzhan\\x\\conf\\app.conf")
	var err error
	Conf, err = config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Println(err)
		return
	}

	InitDBConfig()
	InitAuthConfig()
	InitLogConfig()
}

func InitDBConfig() {
	DataBaseAddr = Conf.String("database::db_addr")
	DataBasePort = Conf.String("database::db_port")
	DataBaseName = Conf.String("database::db_name")
	DataBaseUser = Conf.String("database::db_user")
	DataBaseDriver = Conf.String("database::db_driver")
	DataBasePasswd = utils.SimpDecrypt(Conf.String("database::db_passwd"))
}

func InitAuthConfig() {
	SigningKey = utils.SimpDecrypt(Conf.String("auth::sig_key"))
}

func InitLogConfig() {
	LogAdapter = Conf.String("log::log_adapter")

	if logLevel, err := Conf.Int("log::log_level"); err != nil {
		LogLevel = 7
	} else {
		LogLevel = logLevel
	}
	if logCache, err := Conf.Int64("log::log_cache"); err != nil {
		LogCache = 10000
	} else {
		LogCache = logCache
	}
	LogPath = Conf.String("log::log_path")
}
