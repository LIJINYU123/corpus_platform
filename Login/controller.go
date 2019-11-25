package Login

import (
	"ecovacs.com/corpus_platform/utils"
	"github.com/kataras/iris/v12"
)

var (
	ModuleInfo utils.ModuleInfo
)

func init() {
	ModuleInfo = utils.ModuleInfo{
		ModuleName: "login",
		EntryPoints: []utils.EntryPoint{
			utils.EntryPoint{"POST", "account", handleAccount},
		},
	}
}

func handleAccount(ctx iris.Context) {

	body := AccountPayload{}
	err := ctx.ReadJSON(&body)

	type Response struct {
		CurrentAuthority  string    `json:"currentAuthority"`
		Status            string    `json:"status"`
		Type              string    `json:"type"`
	}

	var response Response

	if err != nil {
		utils.LogInfo.Printf("Bad request when loading from input: %s", err.Error())
		response = Response{"", "error", ""}
		ctx.JSON(response)
		return
	}

	response = Response{"admin", "ok", "account"}
	ctx.JSON(response)
}
