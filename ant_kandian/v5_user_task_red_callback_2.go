package main

import (
	"encoding/json"
	"fmt"
	"gotest/cmd/luyangmao/utils"
	"time"
)

//定义返回结构
type TaskRed2Start struct {
	Success bool `json:"success"`
	ErrorCode string `json:"error_code"`
	Message string `json:"message"`
	Items Red2Model `json:"items"`
}

type Red2Model struct {
	//TODO
	//"Method": string `json:"ceph"`

	Score int `json:"score"`
	UserReadNum int `json:"user_read_num"`
}

//30秒浏览红包 总1125
func main()  {
	requestKey := [] string {
		"WBlmVd_t-K9Y=mnG9Xa9rH5mGwUv0qMIr3pcqUJmvOjdsOvwrdCldzKgcltj8mKH-9_oc9gFcZQS7og_Tup-06UF9EvBEQjwCWigNE0tVmi7bCB6tHRljn3jz11wu1mFNxbZ0c1bYhLLbXlZilEYfGKmT9RWx1s-5XSlp_iIIn3dCjkNeZWbvwvnlJBOLTz8r2qKstmVnRClvaACN4uOg9r9X_-V61CDuYEH8hEHbtD4I5-V2wbBQAnSzMORp8y1m2ZEMA7l38H54x5ADWFRcv6Bb7PQyMrFWCfl8lmUaSCZ0JadSLG2Lx9ED4nuvNA2lM_U6okp3l0onLeJUnhGa3B1IpH64oD__Z-AcsWzPpE8m36dEpZU-GYoyt9KoayiKypXAqZ0LHE2NBSQK_yoogBCD_CKABOHxsT5XCN8XLCb-46MLH0yjp7YUZf27tzVNYzisfPBm4tvtpvgWY93FJO89GHxvUrPpBR1IkVcTZjSFucIp707Uf5noJcrsV_qGjjR7-GVymklyfD2sCLbMoSBInnLMZb2AhJHAFz7KnCkf-kCp44drokc=u",
	}

		//该接口总共可以获得1800
	allScore := 0
	for  {
		for _,k :=range requestKey {
			body := utils.LuyangmaoRequest("https://ktt.woyaoq.com","/v5/user/task_red_callback.json?",k,4)

			var r TaskRed2Start
			json.Unmarshal([]byte(body), &r)
			fmt.Println(string(body))

			if(r.Success == true){
				allScore += r.Items.Score
				fmt.Printf("你总共获得 %d 金币",allScore)
			}
			time.Sleep(time.Duration(30)*time.Second)

		}

		if(allScore == 1125){
			break
		}
	}
}

