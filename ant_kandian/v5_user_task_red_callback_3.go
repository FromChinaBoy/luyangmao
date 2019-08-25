package main

import (
	"encoding/json"
	"fmt"
	"gotest/cmd/luyangmao/utils"
	"time"
)

//定义返回结构
type TaskPlusStart struct {
	Success bool `json:"success"`
	ErrorCode string `json:"error_code"`
	Message string `json:"message"`
	Items TaskPlusModel `json:"items"`
}

type TaskPlusModel struct {
	//TODO
	//"Method": string `json:"ceph"`

	Score int `json:"score"`
	UserReadNum int `json:"user_read_num"`
}

//一分钟阅读红包 总1800
func main()  {
	requestKey := [] string {
		"nG3OyHQhK6SQ=PBkS9uRm5Ya8GXc-vyca_98-LWhYkGLk8nJnNTbt-QIJKu-X50Zi5Ln0aSr528E4clYfWXl43tyb-ZDBOXCckpMU-Ju0IC0muSrt-No6KeAVJ4CaZ0Y0l9thc3CKEtT8LCmiY4Y4ORo-p3b0k6nK2uHdhTnfMfetY3UIbzkGIH27WPKukxC6ws4Lgj5LSdHnJSFjiNE5LAW5wjNplaRuV6i2guVZkiIagOzvv73i_LImAYCTnPz23KILwuSRLbvADJSr4fUEFCx354JLtJT8p-seoTouj3uFoUugd0Itcy42aIASaSfHCjWLM3NR6bO-_MUja93zxHGWDrc7lHeU77Xo6vUZrnnc_n73f9xmHK3Pf4lDlY5WDIL95pOZMQfY80B3ViE5yscIIDoiDrxHxHvOJ65KF0O0hLB_Ow-PmP9g67GR3W259a4UPNdc7QnRPuNkyOZZaYvJ2jl64CIBvuICFCJur_67xtOJMWC-wWTsWb5_eaaNzeOUbxSdrdtWkMwRTb1C49ZsGSCrbpuok-f6BLJkuigmj4EzpH_BJGA=",
	}

	allScore := 0
	isFinish := 0
	for  {
		for _,k :=range requestKey {
			body := utils.LuyangmaoRequest("https://ktt.woyaoq.com","/v5/user/task_red_callback.json?",k,4)

			var r TaskPlusStart
			json.Unmarshal([]byte(body), &r)
			fmt.Println(string(body))

			if(r.Success == true){
				allScore += r.Items.Score
				fmt.Printf("你总共获得 %d 金币",allScore)
				if(r.Items.UserReadNum == 20) {
					isFinish = 1
				}
			}
			time.Sleep(time.Duration(65)*time.Second)

		}

		if(1 == isFinish){
			break
		}
	}
}

