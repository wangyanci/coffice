package filter

import (
	"github.com/wangyanci/coffice/filter/globalfilter"
	"github.com/wangyanci/coffice/filter/identfyfilter"

	"github.com/astaxie/beego"
)

func InitFilter() {
	beego.InsertFilter("*", beego.BeforeRouter, globalfilter.PreDeal, true)
	//beego.InsertFilter("/*", beego.BeforeRouter, identfyfilter.Identfy, true)
	beego.InsertFilter("/*", beego.BeforeExec, identfyfilter.Identfy, true)
}
