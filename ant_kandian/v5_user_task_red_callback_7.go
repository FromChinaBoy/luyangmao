package main

import (
	"encoding/json"
	"fmt"
	"gotest/cmd/luyangmao/utils"
)

//定义返回结构
type Red7Start struct {
	Success bool `json:"success"`
	ErrorCode string `json:"error_code"`
	Message string `json:"message"`
	Items Red7Model `json:"items"`
}

type Red7Model struct {
	//TODO
	//"Method": string `json:"ceph"`

	Score int `json:"score"`
}

//一分钟阅读红包 总1800
func main()  {
	requestKey := [] string {
		"2G3OyHQhK6SQ=cV9ejyp5jcUumb0bcGKN8lyo6-1sgAXFnaYhX0e4foKLTzZ-0EkMKTtRrU3WdnAit9_LpmrP9gSDZpj7gNYDPYAottuN2w7kUErX2nxjgZAlqlIA8wVSQi07plCVYCrH9LMnyS6sy0uf-RqBoH_WMTkuf7S7JwTnrcAs6tyyJO85ypHCBO4d70TioPw0lezwdMY8KCuP5oRxClIRlwneTzfdJmvYlMNfQOQQB-6Wdfs1eWjPzvVyYE8timqMlUDWzM-ydkkD4csxgf__tQRUK5KRwI7IIoXnJTzihd5T5yayyj6Dj5QP_NfdXC3TNJCztdiETYZfyfgs8YqDP5qSpMxDT2VznHxpm9G8IJcBPmaFfmLn-9eh6dhQSYBsR-E85az63b6kVxfp68WxcMNGVw6Uv1-dd2Wczh8QHlbyOEfh7bOAKiZasqMf4gfa4MjgLdw4_D9r0CMdtQjEmlpfNAIYR1RWy2_LIFhvuCxLFlAeHle0a45jXbofzS1HMOoUE3pGE7gf0bkH-HOodJlumQ==",
		"ib06QzjBmGZY=0w9kqvJV3Gwpg1O9QncFCQ4PSqO3CTwzrgkvepAlZAcMNZJlqEyXNZMEVTeHvHigncO6kQlE21bSEpEJtBhMQ57A3k_YTt7jkP8p4wOoMsgXF3ie81ClZJiYhQMwHBfArm2PA0ZC3v2FhutJNpeiqd5K4bvfkeNBFfQ6RAmpMGCiVkdVNTDr-SHuqvzbWzoBspID-3ZiaeyOyusVAp-j24XjsB0eLvhBY34psj5TundcR_6hnjyEO4CusHnSpaMFsJnjTr1iHC31XV4fBw_70ZY7R39Sxi2_hcBxcDBovMbYmWOShEMx_O-NXt-0fjuBazt2msbi1MY1MVsaXUTOFGkNrYtdOAOj8SVudwsc_T_peIMWv4R0DcGr8CrsTPngEPP7CLdCoJJbinhrkx2NXf7MVenHB_EITBw1nqTDGahPnUiuxUbJPOY40x8rcto4ZRZTzmIr3fSTgmKwUGAxUk__4t8uDpVl5tFFCLV47pbBy5qBcAF1MSK3sqVS4F8D7WhCW2GqbhiI7WiWBLvDRw==JU",
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
					var r Red7Start
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

