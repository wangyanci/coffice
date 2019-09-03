package main

import (
	"os"

	"vueApp/app"

	_ "github.com/lib/pq"
	_ "vueApp/model"
)

func main() {
	app.NewApp().Run(os.Args)
}
