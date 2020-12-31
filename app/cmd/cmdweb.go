package cmd

import (
	"fmt"
	"time"

	"github.com/wangyanci/coffice/filter"
	"github.com/wangyanci/coffice/router"
	"github.com/wangyanci/coffice/setting"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gopkg.in/urfave/cli.v2"
)

var WebCmd *cli.Command

func init() {
	WebCmd = &cli.Command{
		Name:    "web",
		Aliases: []string{},
		Usage:   "run the web app!",
		Action:  RunWebApp,
	}
}

//to sure execute InitOrm after register all model, then: register db `default`, sql: unknown driver "postgres" (forgotten import?)
func InitOrm() {
	DRIVER := map[string]orm.DriverType{
		"mysql":    orm.DRMySQL,
		"sqlite3":  orm.DRSqlite,
		"postgres": orm.DRPostgres,
	}[setting.DataBaseDriver]

	DNS := map[string]string{
		"sqlite3": fmt.Sprintf("%s", setting.DataBaseAddr),
		"mysql": fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8",
			setting.DataBaseUser, setting.DataBasePasswd, setting.DataBaseAddr, setting.DataBasePort, setting.DataBaseName),
		"postgres": fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			setting.DataBaseUser, setting.DataBasePasswd, setting.DataBaseAddr, setting.DataBasePort, setting.DataBaseName),
	}[setting.DataBaseDriver]

	fmt.Println("DNS: ", DNS)

	orm.Debug = true
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver(setting.DataBaseDriver, DRIVER)
	orm.RegisterDataBase("default", setting.DataBaseDriver, DNS)
	//orm.RegisterDataBase("default", setting.DataBaseDriver, "user=postgres password=postgres dbname=postgres host=192.168.76.3 port=5432 sslmode=disable")
	//to sure execute RunSyncdb after RegisterDriver and RegisterDataBase, rhen: must have one register DataBase alias named `default
	orm.RunSyncdb("default", false, true)
}

func RunWebApp(c *cli.Context) error {
	beego.BConfig.Listen.ServerTimeOut = 0
	beego.BConfig.Listen.EnableHTTP = true
	beego.BConfig.Listen.HTTPPort = 9076
	beego.BConfig.WebConfig.TemplateLeft = "<<<"
	beego.BConfig.WebConfig.TemplateRight = ">>>"
	//beego.BConfig.WebConfig.ViewsPath
	//	beego.SetViewsPath("E:\\workspace\\workspace_go\\src\\vueApp\\mviews")
	//	beego.SetViewsPath("mviews")
	beego.BConfig.WebConfig.AutoRender = true
	//路由匹配信息是否输出到控制台
	beego.BConfig.Log.AccessLogs = true
	beego.BConfig.Log.Outputs = map[string]string{"console": ""}
	//InitOrm()
	filter.InitFilter()
	router.InitRouter()
	beego.Run()
	return nil
}
