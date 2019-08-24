package main

import(
	"encoding/json"
	"fmt"
	"gotest/cmd/luyangmao/utils"
	"time"
)

//定义返回结构
type Red4Start struct {
	Success bool `json:"success"`
	ErrorCode string `json:"error_code"`
	Message string `json:"message"`
	Items Red4Model `json:"items"`
}

type Red4Model struct {
	//TODO
	//"Method": string `json:"ceph"`

	Score int `json:"score"`
	UserReadNum int `json:"user_read_num"`
}

//一分钟阅读红包 总1800
func main()  {
	requestKey := [] string {
		"2G3OyHQhK6SQ=PBkS9uRm5YZ39RAiNIt9Hhmig_UG_wJniuA3aJhpE1sSCNng06g_8MauiwxecFw8Q2dv56Z_2VmGJkU7TiZ4xjimTpr8MRhz9slOoNm4keOttWKzCZNCkMT1yV2KU_agyNqX_iPjZsHjaPrmm7tQ3IuXljUUPeTuxYFgn78qNGCRdumnGGu4mTIwvPJWUPwxzrEnsdL5FfiOyCY7lLOKNLlNs9PZaGFo9k8gKx2PHOZWfjfk05m9MEPM747galc5tWWO615QCwwHkBsnHD975je8ydqnceyJn6XqhVcMboDXvq7ZGL2V78STb8dxFPVIwnWPlwGkpMn-Xk8SVwvOuHhs9BxuAcIq3bMKwUXXTDn46R5J9muKuiZBqITVTVw1FGeQYo6QV0x5snRHDUkHZ0WUPbbb42zyNYKBL6MCdaD09k5erSs4Jz36QKaorSw13_VCNOUzhNFydaqhBuKoXZWIMGg3Tu0mB0IhP_cBNxWI4RjsBkU3QZ8O8pmmokBlOdJHnPst0IiMg3i_lmN-SAvWs2yfns4nJuqRFHleJ5k=",
	}

	allScore := 0
	isFinish := 0
	for  {
		for _,k :=range requestKey {
			body := utils.LuyangmaoRequest("https://ktt.woyaoq.com","/v5/user/task_red_callback.json?",k,4)

			var r Red4Start
			json.Unmarshal([]byte(body), &r)
			fmt.Println(string(body))

			if(r.Success == true){
				allScore += r.Items.Score
				fmt.Printf("你总共获得 %d 金币",allScore)
				if(r.Items.UserReadNum == 20) {
					isFinish = 1
				}
			}
			time.Sleep(time.Duration(60)*time.Second)

		}

		if(1 == isFinish){
			break
		}
	}
}

