package utils

import (
	"github.com/kataras/iris/v12"
)

type EntryPoint struct {
	AllowMethod string
	EntryPath   string
	Callback    func(ctx iris.Context)
}

type ModuleInfo struct {
	ModuleName string
	EntryPoints []EntryPoint
}

func NewEntryPint(method string, path string, callback func(ctx iris.Context)) EntryPoint {
	entrypoint := EntryPoint{}
	entrypoint.AllowMethod = method
	entrypoint.EntryPath = path
	entrypoint.Callback = callback
	return entrypoint
}

