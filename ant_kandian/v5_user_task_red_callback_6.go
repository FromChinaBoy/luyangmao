package main

import (
	"encoding/json"
	"fmt"
	"gotest/cmd/luyangmao/utils"
)

//定义返回结构
type Red6Start struct {
	Success bool `json:"success"`
	ErrorCode string `json:"error_code"`
	Message string `json:"message"`
	Items Red6Model `json:"items"`
}

type Red6Model struct {
	//TODO
	//"Method": string `json:"ceph"`

	Score int `json:"score"`
}

//一分钟阅读红包 总1800
func main()  {
	requestKey := [] string {
		"LI0ueNrUge-I=WWRRHWUvVxwpW9vgNsd5g_Uofah1gcPiXEmnVH33c8onoShgDKjPXN2stIBYnIud2cdAvRqnseO8knf09dTZn7bna99TTCjGh80DS4y_3bqm-g7PkNWwGFqhvssW5Xek7cSj1fbRxSkB20jVs9rz6hDzyEEgDSsNgaXHoUwJzndABhjlt_u-5SVU0ViPTKEP1s_UpFpwIoyZ57_lwOksTc2PdmFN9__8lZCAv3IE7vxF_2tQSIodoOddlsrkZ7JAOP0yTDCQFH5rvQcvYSXSzyHKdJkP70yWserSvlAdTIHd83pjW1T2kWInB1CAhY7z5dpuIh6e6TjoXVj0p8RG9Jzf0JL-DZtcM_Wh2-c2h6aWYyHEVyH4GWTYaaJxO3g9DwU_6OnLhTlQSlIYl6ALhMfjUOMSmjRJCoqbzXU3NmwmYArZ56MudjMzM_tgpTtGM1RVj-Dxi07kDs8ChfTT61F5z-pX-G-2QfvfU4y1xiLCBD5L0F3vUF5WyNzKuAadneR0J01XIHw-Rr8EA25ZTb8XgJ0u-jtf",
	}
	allScore := 0
	for _,v :=range requestKey {
		for i:=0;i<20 ;i++  {
			body := utils.LuyangmaoRequest("https://ktt.woyaoq.com","/v5/user/task_ad_callback.json?",v,4)
			fmt.Println(string(body))
			var r Red6Start
			json.Unmarshal([]byte(body), &r)
			if(r.Success == true){
				allScore += r.Items.Score
				fmt.Printf("你总共获得 %d 金币",allScore)
			}else {
				break
			}
		}
	}
}


