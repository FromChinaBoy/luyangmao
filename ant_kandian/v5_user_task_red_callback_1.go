package main

import (
	"encoding/json"
	"fmt"
	"gotest/cmd/luyangmao/utils"
	"time"
)

//定义返回结构
type TaskRedStart struct {
	Success bool `json:"success"`
	ErrorCode string `json:"error_code"`
	Message string `json:"message"`
	Items StartModel `json:"items"`
}

type StartModel struct {
	//TODO
	//"Method": string `json:"ceph"`

	Score int `json:"score"`
	MaxTime   string `json:"max_time"`
	TotalTime   string `json:"max_time"`
	OverrideUrlInterval   string `json:"override_url_interval"`
	Title int `json:"title"`
	ChargingField int `json:"charging_field"`
	DefaultHint int `json:"default_hint"`
	OverrideUrlHint int `json:"override_url_hint"`
	UserReadNum int `json:"user_read_num"`
}

//30秒阅读红包 总1125
func main()  {
	requestKey := [] string {
		"WBlmVd_t-K9Y=mnG9Xa9rH5mYAwS_kj9yNOWBIglt0kCHPM9PqJiq9E9RCUrZh28urmI_4ju-k4BwpaWeryNiNOJxwyiIHOHQ6nD-ZRzK-Rfn-zJL7UdtsHGXaSpIYLcMjRTijWGup0lKV5Jf1VjoJYTJFbdlKUTJXOiiSs79TKPvOZMmZPnH4UJui8FpmPD_2JJs4zJw9TcoUPnU-oxHmk5ZXG2syUSdRASKCIZnhucf7l2qZoUqv-qdOvszmzhefSy94ONsA4xb_LTDRBcHW8iyXNEBi57LLzx_0RfzIrtARIZnbbV6KkJO5M9ysNJDKhwiIL7xUMxJnFidj6VUY6TF2nb1MUoU70mbBxE9VkGYtI9aVJ71_KeavyZqqTQTIc82FZWbWO72CycuKlKpQeY5Pq0Eg8AKQhYGOt_W6mEMRWjFEFbSZAZvTq5Hfkrx09Erpy1fNT5yvV3wRkJ-spDCLKKpt4E8i_JUh7YeKkJDxaKlU2dmyZi9JU1KBIFMAfEBJwPZ3Wz3adRXlaToCbTjCQx-gEBjJL9Gjd7JkWY-OYCBqJuOy3Y=s",
	}

		//该接口总共可以获得1125
	allScore := 0
	for  {
		for _,k :=range requestKey {
			body := utils.LuyangmaoRequest("https://ktt.woyaoq.com","/v5/user/task_red_callback.json?",k,4)

			var r TaskRedStart
			json.Unmarshal([]byte(body), &r)
			fmt.Println(string(body))
			time.Sleep(time.Duration(30)*time.Second)
			if(r.Success == true){
				allScore += r.Items.Score
				fmt.Printf("你总共获得 %d 金币",allScore)


			}else{
				continue
			}
		}

		if(allScore >= 1125){
			break
		}
	}
}

