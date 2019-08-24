package main

import (
	"encoding/json"
	"fmt"
	"gotest/cmd/luyangmao/utils"
)

//定义返回结构
type TaskAdCallback struct {
	Success bool `json:"success"`
	ErrorCode string `json:"error_code"`
	Message string `json:"message"`
	Items TaskAdCallModel `json:"items"`
}

type TaskAdCallModel struct {
	//TODO
	//"Method": string `json:"ceph"`

	Score int `json:"score"`
}

//看热文100片 2500
func main()  {
	requestKey := [] string {
		"Ub06QzjBmGZY=6XAqUcTiR-IV2lvO4s0H7V5XJxff8FnPZr3Xy-Mkw1K4C18STrRvRaRhbEOG1zLGSaNYghOkP1aAIXPG1ZGlXDd1J4khBMXrFINElAU6QBSfq4UAoS9EydoTT4pJGRzagZtU0m_uC99bwfYKxFFgp0TpyMCApZHp_OyxMBSlC4v_dsyUpdsNPZlh9d4tBukIKcO-ijAo_U_0VAM3RS02feK0LVHonINYAwjSOMnHJuB4hSchNkq7wWZApzeF5dds-6pi5pdivvWi9KHa4STAn5JGxEBgoUZaisz0E7VPRM4oTA5e8bDw5iNovC3aW1OszVZsj12jBEfghy1-TGr56v0BXW4Xxzv4xDn3uAdVPwjL2jQ7yRnKMvfAE2U8kK-32L3QeZI7ueWwzo_TXPFpDNM50H1_kadAznpfWzlOvJHT9LaNeeP_XRpEXFfSGjAki1rDwOsco2NqA7gsyuHjzIZ8UwZbFde4DB99ilyRSeWihQLnNjb7SJ8-JW7txmAlRvvNcE_EbT34NUNYTTEEdvC01MIenjbyvN",
	}

	//该接口总共可以获得1125
	allScore := 0
	isFinish := 0
	for  {
		for _,k :=range requestKey {
			body := utils.LuyangmaoRequest("https://ktt.woyaoq.com","/v5/user/task_ad_callback.json?",k,4)

			var r TaskAdCallback
			json.Unmarshal([]byte(body), &r)
			fmt.Println(string(body))
			if(r.Success == true){
				allScore += r.Items.Score
				fmt.Printf("你总共获得 %d 金币",allScore)
			}else{
				isFinish = 1
			}
		}

		if(isFinish == 1){
			break
		}
	}
}

