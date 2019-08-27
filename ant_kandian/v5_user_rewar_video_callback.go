package main

import (
	"encoding/json"
	"fmt"
	"gotest/cmd/luyangmao/utils"
	"time"
)

//定义返回结构
type VideoCallback struct {
	Success bool `json:"success"`
	ErrorCode string `json:"error_code"`
	Message string `json:"message"`
	Items VideoItemModel `json:"items"`
}

type VideoItemModel struct {
	//TODO
	//"Method": string `json:"ceph"`

	Score int `json:"score"`
	PlayNum int `json:"play_num"`
}

//看视频 总1200
func main()  {
	requestKey := [] string {
		"PG3OyHQhK6SQ=lg1-m4fcElCQbwXN3gI8yh380vLq7kdmnlreLhIeGUQh-U8qa0f4-dYk8tjI1fIeHKf_hFzkh9gYmtXgfrho32Yrj-EGmQWGtEtclpai6vhh8Za2TflhetbNetOTqB7JSkEVJ7bxaS-asSh9awVJiqg2azbNX50vm2fnDYU4a80-d_J1-d58vblWso2tZgHZxRPn97kCGlMcBEY8qOenBjkOyKz8STAN-Q4Ig5yPVd1ctMP9Higlevqt_DOMCdq4IAY3GfNFpOfUaH0HdNzaDNvuHtfpsoLP9UMLnsKnNwZtycOFakhubHM861V1TnpnLjo7ScCsjBsgFQTDq6FZgpMm4gRvQsbkTO039ab2oSujq9e_O43Y-yG3Cf-vh385MCfRnCt4e2ddc3wMZdTicSd72wgfkD-rJp4eW46bN4SgK7tOPebFj5zKAe3xEGmjMfN5bMeHV0zzcA_YeL6iDuTXhHuMFS6jShv3xk0aYGSbewLZfewGsB6p6pXW30YANr_2IYWdEV2f0v1MS3QsNMEhxSsNHkfb6iTcDMLNEBeI3VRfNqM9vDrVFCQ7ryi2qPpD8qJKHiBBGAWDIuiM9Z1qSklBxpV2xXycnghv7gXbWdSJA_4HhFKxSdYdMqaD",
	}

	//该接口总共可以获得1125
	allScore := 0
	isFinish := 0
	for  {
		for _,k :=range requestKey {
			body := utils.LuyangmaoRequest("https://ktt.woyaoq.com","/v5/user/rewar_video_callback.json?",k,4)

			var r VideoCallback
			json.Unmarshal([]byte(body), &r)
			fmt.Println(string(body))

			if(r.Success == true){
				allScore += r.Items.Score
				fmt.Printf("你总共获得 %d 金币",allScore)
			}

			if(r.Items.PlayNum == 40){
				isFinish = 1
			}
			time.Sleep(time.Duration(65)*time.Second)
		}

		if(isFinish == 1){
			break
		}
	}
}

