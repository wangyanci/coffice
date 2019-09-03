package filter

import (
	"vueApp/filter/identfyfilter"
	"vueApp/filter/globalfilter"

	"github.com/astaxie/beego"
)

func InitFilter() {
	beego.InsertFilter("*", beego.BeforeRouter, globalfilter.PreDeal, true)
	beego.InsertFilter("/*", beego.BeforeRouter, identfyfilter.Identfy, true)
}