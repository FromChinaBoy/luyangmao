package main

import (
	"encoding/json"
	"fmt"
	"gotest/cmd/luyangmao/utils"
)

//定义返回结构
type Red5Start struct {
	Success bool `json:"success"`
	ErrorCode string `json:"error_code"`
	Message string `json:"message"`
	Items Red5Model `json:"items"`
}

type Red5Model struct {
	//TODO
	//"Method": string `json:"ceph"`

	Score int `json:"score"`
}

//一分钟阅读红包 总1800
func main()  {
	requestKey := [] string {
		"LI0ueNrUge-I=WWRRHWUvVxwpW9vgNsd5g_Uofah1gcPiXEmnVH33c8onoShgDKjPXN2stIBYnIud2cdAvRqnseO8knf09dTZn7bna99TTCjGh80DS4y_3bqm-g7PkNWwGFqhvssW5Xek7cSj1fbRxSkB20jVs9rz6hDzyEEgDSsNgaXHoUwJzndABhjlt_u-5SVU0ViPTKEP1s_UpFpwIoyZ57_lwOksTc2PdmFN9__8lZCAv3IE7vxF_2tQSIodoOddlsrkZ7JAOP0yTDCQFH5rvQcvYSXSzyHKdJkP70yWserSvlAdTIHd83pjW1T2kWInB1CAhY7z5dpuIh6e6TjoXVj0p8RG9Jzf0JL-DZtcM_Wh2-c2h6aWYyHEVyH4GWTYaaJxO3g9DwU_6OnLhTlQSlIYl6ALhMfjUOMSmjRJCoqbzXU3NmwmYArZ56MudjMzM_tgpTtGM1RVj-Dxi07kDs8ChfTT61F5z-pX-G-2QfvfU4y1xiLCBD5L0F3vUF5WyNzKuAadneR0J01XIHw-Rr8EA25ZTb8XgJ0u-jtf",
		"jI0ueNrUge-I=WWRRHWUvVxxWvm9h0xi1BqRSfBylyTfoLzz7A-Sj5fQl34S5s9vG9UpOEMy_wNn4QBtJ0vpFT1mBljSR-XECukX1rAl8GQmApxLvSE6NEh4dFDShh8KJWWTdRIOiJg3edOi2VFkr5Cq6Nu0Qsyqa-F3UKRPbVBpxi958yumveUiPl05DSbrFrgW-xcrNmgK0y5W8nOIvyg_jSFnCXJaUjAsDUuLH9BK4LidBOYotGGJz6vioYhLiClb3YJj3Q31cNyXy9GxNcajI96kreTmXQuHF8ylcqsAXV4Ol8pNQbBApnDXnfduBI1DM9pnJ6DHVT9QbT93yQ_Yf0uXhqCfxD6A2UgBnHrfkTElAAboE9Tm2j4TOxz2pdGAjwYAMZTzl5gorPzc7nyp-EEbglh9QKxjzDWB8IrXC6yV6xVwbtXseLcO3L1MBABchIzeU8X2-plDdQEFjfjxdXnDlNNv8tfnyZDlAtjGFZeQqo7ByVkgY0kvc9NJFbLTMVZVeyEH8d2jLLbjLZi57V_-RUPHLPm62rPsAjP9d",
	}

	allScore := 0
	isFinish := 0
	for  {
		for k,v :=range requestKey {
			switch k {
				case 0:
					for i:=0;i<2 ;i++  {
						body := utils.LuyangmaoRequest("https://ktt.woyaoq.com","/v5/user/task_ad_status.json?",v,4)
						fmt.Println(string(body))
					}
					break
				case 1:
					body := utils.LuyangmaoRequest("https://ktt.woyaoq.com","/v5/user/task_ad_callback.json?",v,4)
					fmt.Println(string(body))
					var r Red5Start
					json.Unmarshal([]byte(body), &r)
					if(r.Success == true){
						allScore += r.Items.Score
						fmt.Printf("你总共获得 %d 金币",allScore)

					}else {
						isFinish = 1
					}
					break
			}
		}
		if(1 == isFinish){
			break
		}

	}
}

