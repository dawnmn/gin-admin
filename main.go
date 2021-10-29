package main

import (
	"gin-admin/app"
)

func main() {
	var err error
	err = app.Router.Run(app.Config.App.Host + ":" + app.Config.App.Port)
	if err != nil {
		app.Log.Fatal(err)
	}
}
