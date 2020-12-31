package main

import (
	"os"

	"github.com/wangyanci/coffice/app"

	_ "github.com/lib/pq"
)

func main() {
	app.NewApp().Run(os.Args)
}
