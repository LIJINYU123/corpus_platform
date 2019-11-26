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
			utils.EntryPoint{"POST", "login/account", handleAccount},
			utils.EntryPoint{"GET", "currentUser", handleCurrentUser},
		},
	}
}

func handleAccount(ctx iris.Context) {

	body := AccountPayload{}
	err := ctx.ReadJSON(&body)

	type Response struct {
		CurrentAuthority string `json:"currentAuthority"`
		Status           string `json:"status"`
		Type             string `json:"type"`
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

func handleCurrentUser(ctx iris.Context) {

	type Response struct {
		Address     string     `json:"address"`
		Avatar      string     `json:"avatar"`
		Country     string     `json:"country"`
		Email       string     `json:"email"`
		Geographic  Geographic `json:"geographic"`
		Group       string     `json:"group"`
		Name        string     `json:"name"`
		NotifyCount int        `json:"notifyCount"`
		Phone       string     `json:"phone"`
		Signature   string     `json:"signature"`
		Tags        []Option   `json:"tags"`
		Title       string     `json:"title"`
		UnreadCount int        `json:"unreadCount"`
		Userid      string     `json:"userid"`
	}

	var response Response

	response = Response{
		"苏州市友翔路18号",
		"https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		"China",
		"jinyu.li@ecovacs.com",
		Geographic{Option{"320000", "江苏省"}, Option{"320500", "苏州市"}},
		"科沃斯商用机器人-研发部",
		"李锦宇",
		12,
		"0512-4008078999",
		"海纳百川，有容乃大",
		[]Option{{"0", "很有想法的"}, {"1", "专注设计"}},
		"交互专家",
		11,
		"00000001",
	}

	ctx.JSON(response)
}
