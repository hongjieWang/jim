package main

import "fmt"
import "github.com/jim/message/feishu"

func main() {
	app := feishu.Feishu{AppSecret: "TFzUlKx3FJrDFhZdrwpxahXrtpIU31xP", AppId: "cli_a208d1b17ef8900d"}
	token := app.GetTenantAccessToken()
	//info := app.GetUserInfo("ou_16e983082ce312de36a654414c6ef2be")
	dept := app.GetUsersByDept("0")
	fmt.Println(dept)
	fmt.Println(token)
	//app.SendAlertMessage("text", "ou_16e983082ce312de36a654414c6ef2be")
	//fmt.Println(info)
}
