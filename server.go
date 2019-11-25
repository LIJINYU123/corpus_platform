package main

import (
	"ecovacs.com/corpus_platform/Login"
	"ecovacs.com/corpus_platform/utils"
	"fmt"
	"github.com/kataras/iris/v12"
	"os"
)

func init()  {
	utils.LogInit(os.Stderr, os.Stdout, os.Stdout, os.Stderr)
	if len(os.Args) > 1 {
		err := utils.LoadConfigFromFile(os.Args[1])
		if err != nil {
			utils.LogError.Printf(err.Error())
			os.Exit(-1)
		}
	}
}

func setRoute(app *iris.Application)  {
	modules := []interface{}{
		Login.ModuleInfo,
	}

	for _, module := range modules {
		info := module.(utils.ModuleInfo)
		for _, entrypoint := range info.EntryPoints {
			entryPath := fmt.Sprintf("/api/%s/%s", info.ModuleName, entrypoint.EntryPath)
			if app.Handle(entrypoint.AllowMethod, entryPath, entrypoint.Callback) == nil {
				utils.LogInfo.Printf("Add route for %s (%s) fail", entryPath, entrypoint.AllowMethod)
			} else {
				utils.LogInfo.Printf("Add route for %s (%s) success", entryPath, entrypoint.AllowMethod)
			}
		}
	}
}

func main() {
	app := iris.New()

	setRoute(app)

	if ok, port:= utils.GetEnv("SERVER_PORT"); ok {
		_ = app.Run(iris.Addr(":" + port))
	} else {
		_ = app.Run(iris.Addr(":8080"))
	}
}


