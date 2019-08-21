package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//定义返回结构
type Response struct {
	Success bool `json:"success"`
	ErrorCode string `json:"error_code"`
	Message string `json:"message"`
	Items ItemModel `json:"items"`
}

type ItemModel struct {
	//TODO
	//"Method": string `json:"ceph"`
	Desc   string `json:"desc"`
	Score int `json:"score"`
	IsShow int `json:"is_show"`
	Complete int `json:"complete"`
	IsProgress int `json:"is_progress"`
}


func main()  {
	requestKey := [] string {
		"3O7oFQGUcudM=r5BAb-UsVAtXd6jfl7gJ_OqTgJn8-0x1H1m-goyWKASfHTPXjMvdm92LyYdD6isXszdv3OkNQ-N_RA-T99ytB7TFvyqLnzR9wKOURN6CMdBVYHBATNTYqVfyJkxcveuOYmFJ0gBMBXox6IF3uNRIDhLjfVDlUk-jOa4s86PKUFOxZCwBt7trrLuTVHG2ETMEcCXx8aEaYIrgW5F4siKnnJzOlPlyMwU33GHeaoJTd-Xs6NIoJb7LO77OLJb0nZCMEAxisSoy3yonsU5yhihE5CaII_auhgqS57GOjhOqYaEBwGJxBUbl0hhsUbyFgiaHVHSR-5TzxnoAewqAjAx2sSewausbsC-_Wh1w1zDX1RaoTwCtvcFF9_gHEpdMUn7kF9XeCRU4xQbN7_Dxu3WQNwK4RihglQc01J9INZoNlqeuKFJY2EpOpvgPxK5HwBPE8yjwsQknsIJ3VhYAyI6azfBUwu00wF-dOp7KdSdse7ajtFHszi4007vcfdrCcEgate8okoHmsEqpXZKU_Ei3PBZCMsgt-Gu8K",
		"q308dO6RZsFg=Yn3ZUPMh-_a06Iyt3PNnJSpaHmkhm3PrUA0ToSqvi9R-kG8OmnjDCFq7G06ZxE4ToxQ551OhNqIAqfPsg0-iO2JmWr4JGHt7HabTV39KEO2qBkYPyzgHqQw7JbXcnebRVoy9uNw57SnDUeTLdHqAXpQixIvzqHqOIRNnF2JEQJ0_2JZVz_0z2YwsY0ROqpxP6rVV46GEjBk4fiDV82dqz7ZPF1L85YOsI-MZ3KtXCAHzLpzXgBOnSWxucamVe8FOPhjlKtvM0JqlOQlXC9qv0Naj0ozDIc2j6ivL4EL-UZ4V2b-pNDy5SMa_o9JGfpYwD8qEoxjQYG-polrhkiIIQJL-ubtO5srrR5PEZ63jw4fUgcOnSlAmSK7EOkxxM6QAtu9pjRH6ar5p8SxmvKLx0LxXr034znXuyiKL9CUDoc-KaSqNL-RWSeeujG1_WIS3yVkUwbxQG8zDx13XDYSPf55EUS-_7_xQRF6muQnnfT7IZNGc1bEAat0JSl6fPm50QwYsS6drcEX7Ahv2cQcSo_MHhv9EuFlM",
		"J1FyL9-SGPvc=j1YAMFD8qU3VsT9ea42LdbSPrnlMZ2i0mpL--bitZ_67-DI-R_9WrCOcwU_lA5pXOxtB0etl3Bl9t-XCtLTUXkMEy7DADzIDF0BIRUJhKHP-_Y1mK3M3SbNEuUHnzrjttqeNx-6-R7f0g9a-Mi--M2d3XANPepDTRXfVr3ABRrjMHSSgi5fE6z22B7PgQADcUtiPt_LGsGxULdX7E5TmZEEWM_W40zMCGsKJu9dSLyeCv-ByMtn63lePZO5R1z6S9lhr2clKKLa-pPpMh_x1Wr6taGDRiayDtJBRAeYovMfzcgqq3dEwoYBR1vtUJfXBYHggE6b6p6elunRWm9WXwCWWx97FaYB4gaFz8ffT29JUwdboAcrsrYvjwrDy7y2lk7mQn5GPSjcZXKZpcxNuOKgfT7ooziFLH4nXXsc3j_MPKAc37EDXbuKqEa9DaOYGf9JProJkFEwHzNij7B-6M3I99UqE7T290oY0B9CpYewBcAQCrJSzEwaIE1DcgL9HSIWpTVNkO05qT6GvbIEkwUsPp7zWc548w",
		"GO7oFQGUcudM=r5BAb-UsVAvZXl_FRyovpgPMmnJp39BhBWjuAY36v2_q3xEMi6NoihARbuX-GhHHCPAj6W86RkyCoOl3z1Ib-S9IzDNF-YczTjeaVyrzw98ZIcwHTouGV9gp19M1gB4x4e4tnzf4o6RJo77ErvUyYMZdY4epmWgQYg_VL12Y8KzQD0Ztd0ObuB4Mvllt_R9KBagqdHHRavEICowBhE-doAVHOpQYMr5ZBwJ2t8y0LL_aGadfM1v2Pal1X5jUEpK0L7TQDODK3E9mmVLsrqY3TAFWNWC1--Jf4w4j54sN7b2FLk7NoRidnfMIv8i2opHEqx57jvBdR-Ih16cc9eetcq8OoE38dcF5RLnXw9HXQ6EBbNw9zO15elbjT6F6fe6KT6ylfAytMuoJV2wkreRVPJ3Ylrj55Bd4b7BbcvaFXQcc2WlsoHyvd-MwDjjOGn2gQHO1oU1wtC8LKotdtpAY92cTsRQnpN3s6-MpJz_7vkX3KycgepVlGBAllYnXbFgcJ3WcM1xP-dTCjoUW7ZHSqnZwL-EG4Rldw",
		"XX_uqcajdUDhI=SYmy1thwDFyEC1Y04UpYIC0J18h6WQ4dq_dIb6E07YbFDATjLisk-1PUL6EpAuVRcO_LUbqDWMnNKDR8rlm-OXL5XMpae8qVXbE7GsPhECnl9sQpQlTua6loUW60aa_d7QzixQ2vRgiQPIJWmSapaEE6rmi8d4EejxoUSrdgkodlpSyL9sGH32ASYYxgHuHPW2N_0m_Cx0lKuz9_Z6dueT-zyG3ACzzpfVczaON980ykOBLHgM3OihkysYroRh8SPHlh2xJvS1aSpsCVizOVIAFOSdnYV7tQA7-FpeymMeDLUXvynkDyhV4n36sibe5NJdgPL42iSJ2YtxojpXzd7oUdZ5cWnHJ1oT3o_ORgXz48K2qrrffc0C5z9awWalT63LgiwFthMBdbl3XhaHNpMwApx8B7hWP2FyfISVEx4nwUbt-VIvXihNif9RRJO8_PiOwJA4fuc-z5DrcNFddtmOym0o0B4qg2Wjrr9C4_RyIUTs6SxQV3VfO34xIobon0ShHg_6XgbRxHKY8Yg_pH2dnRg3g_ELP-Tz",
		"h1FyL9-SGPvc=j1YAMFD8qU0Bb4GwcwA_4At0X4Z59vd76tniiLUOyB4X342xiPqgN2kmJ9N_HFlrr1xj3IiZHhfSpfqCUKM1tJSkN7Fpued_0RRypVmqpNyYU2GbCEl6aG-GHb-WcVF2xnAvMUZjMYrwg2Iexz6qrP2wT6Zladahz4qpuY9UfYQxl69DmpZM69LQQxWnvWH9l0gM7y3SYP9Ec-fzUNdExA9sCigN-C-fubt0ex9byqLo2PAkkJM2HUSI4v9fdg8vjqQK-80By8Qly2Ia87HUZmfs4elGFq3Gf60ySQMfJtpr-wkZu-Kakl7tdxM8aax5BBiRQ7wfp39RvIiJUfK9UPTDYTsiodHCrdsbR-fFtX2J3sXC5xOceJmgR3tbF5C-hfJ75yjQ0JqWyiczz0WOjy1gO-rTvweLp3ri9C6tD40fCKeHgCzN1vmrp-7eHolp8LCOdrIiqyDvWCd-H1K4r3UwJHwclecBpJKLlcgVDn4bImvHGzpWeLN32DcAldhDIHstGylQjU72_LJ7BaBCtjiz05LEOnorm",
		"FG3OyHQhK6SQ=l5_qV9eaQimyW-lrMxMMtiZ3FW5XyewHNTR-ccCjg5UrAH13v7u3QlZJSA9d90Di_EluMXY832D2QTAY6JAZpA02VtS_uRSFUMxTkBjueD6ElLM5XqujJY_BRd02zV2ox8VPdWvbI8QTZjFEWsCkJUl7gnUFwRVS1jQHxFhgwSPUXPlMKmH2xyKHkNJF5oHk_jdKM4SBsOFPXmHl502YgQfevlSMnpA9b0PVKjPbI1TpHVQSA5o91cLTc_9EgayW6m7uXB-Zlq18Rf3hmvsplpuUrkxqMMR1vtdzkx2Eoi_ipNjdS7zWl24jU6lg8Zj8GdXSPp1Ih6ATrQbXZJeLPVKHSUkLjvVzjXysGO7bcBKHQ7IIuBxhiFiHEInuHDisAg93qo9RmT3_zKZU4yLLNqs0FLOvNJNK_4biRlOS_g0z5xK5Y2AEQNzXQw9ha0IiziWpw7yc81A5UfY4CUW572r0jJzUVr0xKsrFLahTO10dvc9Nvkwcs5auSUIKTnrMVbOP4OEp9Y94rVWutH5OTbx40coxqWo5",
		"oO7oFQGUcudM=r5BAb-UsVAvZXl_FRyovpgPMmnJp39BhBWjuAY36v2_q3xEMi6NoihARbuX-GhHHCPAj6W86RkyCoOl3z1Ib-S9IzDNF-YczTjeaVyrzw98ZIcwHTouGV9gp19M1gB4x4e4tnzf4o6RJo77ErvUyYMZdY4epmWgQYg_VL12Y8KzQD0Ztd0ObuB4Mvllt_R9KBagqdHHRavEICowBhE-doAVHOpQYMr5ZBwJ2t8y0LL_aGadfM1v2Pal1X5jUEpK0L7TQDODK3E9mmVLsrqY3TAFWNWC1--Jf4w4j54sN7b2FLk7NoRidnfMIv8i2opHEqx57jvBdR-Ih16cc9eetcq8OoE38dcF5RLnXw9HXQ6EBbNw9zO15elbjT6F6fe6KT6ylfAytMuoJV2wkreRVPJ3Ylrj55Bd4b7BbcvaFXQcc2WlsoHyvd-MwDjjOGn2gQHO1oU1wtC8LKotdtpAY92qczRaKQD2k5fL8y-asDeil3CtHetz4-KYflZjpkW4VuwAV8dG_eJ40FU5nrFBZf1MdsP33i1l-D",
		"w_AKV1AEqa-c=ZIafDtZioW6P4TNuzTH4mu1Q409IEF47Xg1-cLhH4lKq7O922uE3R2b35IhvMHLrelMQBKRRhBzjViqYTDPIFXSmzYDLdw2AaltQ8EVpYAsE6Lec2TZ0h-PMl-4m2ZB2RJb7rwuKADwo2Syb16LG1iYabdx5NJqY5ckDhlRMNSo9qDGsCYx57ANG1SX9YgxPp8q1VwGSDBhMspMAOrIBYLE4YXf3JWJwJ9KZdN8dRhR5EAfm_zLB4SJE9L1Ya6G7NOY4PcMO6VjJ5TnAlNIp0frC0D-E3HjqR_o4fDOfP8P73L8G6EdNltPT7g-K1Tnj8o_hInFTduPOp5whIrsGlQeQZ6cTbNB02x5kGyF0oxnXchUQgy1e1UJly4Uhn2s3_edA87vAHAyFzli29lmp3-Benl5wkTt_fHycACgPIV__Tt8eWDtEcA3KEHdQnTUrezk0pc6VkmHtWUCfsrvleF5celqwXwD2U0R6_mOVlKyQipUoHUoGG0v4ZBL3VVoC0dEelj08fMq7BkdX2P7zNd8eXIPeZBrF",
		"h1FyL9-SGPvc=j1YAMFD8qU1D6a4-EhUhWugnboBkSP24PeEVNjVMPF_U6I7lqIjaoI2tWD6DCEVn0dE4rTACzl8gpOPLSYM6QiemnMFT6rCq-Fm0UonJMl6B3FWr6hoFc_JtUL0A3vRvPwOG8iM9rESdnbuAIXx6MdT0MWpByC7f04vWkfXVHPyYABHhFtf8SWZMFaedq7wISsf3rA7HAFKRpjUBAHEgdjpL6CMFGbvU8rIy2FqnNTtNpri8H9mZenFfaxSNIf_Fhd3m8nCWfPcesTu-OVFynypzAR4GvCsOQr1tuWYeOFubo3Cnu1DLpUXpecVsRzFGbz3XjiLxg_SBi8VRmALkbLz7AKYD3rwmjBt-8tJ_YaKGvhyrqNEWj5_zzMGCNqSC62g8YVTtjqctwb9dG45YRPvjBT0nmxIZdic-zYvNAWYZ-jhZDt_ImLnIrSbDRYwg4YwQhEan-GVgCnPZ5ioXc6jp0OC7Ps6BjJk1YORdInrPW3_9TnTWGSk0mqu6u0RPmFEtMVc6fWDHxX8mviyppF1OyfNYVlvQN",
	}

	//一篇文章请求key可以循环六次
	allScore := 0
	for i:=0;i < 6 ; i++ {
		for _,k :=range requestKey {
			score := luyangmao("https://ktt.woyaoq.com/v5/article/complete_article.json",k)
			allScore += score
			fmt.Printf("你总共获得 %d 金币",allScore)
			time.Sleep(time.Duration(30)*time.Second)
		}
	}
}

func luyangmao(urls string,cookies string) int{

	var data = map[string]interface{}{}
	res :=4
	var resp *http.Response
	var err error


	//-------------------------发起请求------------------------------------
	switch  res{
		case 1:
			// get请求
			resp, err = http.Get(urls)
			break
		case 2:
			// post发送json
			datas, _ := json.Marshal(data)
			resp, err = http.Post(urls, "application/json", bytes.NewBuffer(datas))
		case 3:
			// post发送from表单
			datas := make(url.Values)
			for key, value := range data {
				datas[string(key)] = []string{value.(string)}
			}
			resp, err = http.PostForm(urls, datas)
		case 4:
			// 创建请求设置参数
			// unicode
			//var clusterinfo = url.Values{}
			////var clusterinfo = map[string]string{}
			//clusterinfo.Add("userName","sss" )
			//clusterinfo.Add("password", "2121")
			//clusterinfo.Add("cloudName", "2121")
			//clusterinfo.Add("masterIp", "2121")
			//clusterinfo.Add("cacrt", string("2121"))


			var b bytes.Buffer
			w := multipart.NewWriter(&b)
			values := map[string]io.Reader{
				//"file":  mustOpen("main.go"), // lets assume its this file
				"p": strings.NewReader(cookies),
				"application_id": strings.NewReader("com.ldzs.zhangxin"),
			}

			//拼接body体 multipart/form-data
			for key, r := range values {
				var fw io.Writer
				if x, ok := r.(io.Closer); ok {
					defer x.Close()
				}
				// Add an image file
				if x, ok := r.(*os.File); ok {
					if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
						return 0
					}
				} else {
					// Add other fields
					if fw, err = w.CreateFormField(key); err != nil {
						return 0
					}
				}
				if _, err = io.Copy(fw, r); err != nil {

				}

			}
			// Don't forget to close the multipart writer.
			// If you don't close it, your request will be missing the terminating boundary.
			w.Close()

			//实例请求
			req, _ := http.NewRequest("POST",urls, &b)

			//添加header头
			req.Header.Add("POST", "/v5/article/complete_article.json? HTTP/1.1")
			req.Header.Add("iid", "0")
			req.Header.Add("sm_device_id", "2019081220510120358e7508e0422050d596224920381401a35b5ae7309211")
			req.Header.Add("uuid", "868241031584825")
			req.Header.Add("access", "WIFI")
			req.Header.Add("app-version", "5.6.1")
			req.Header.Add("device-platform", "android")
			req.Header.Add("os-version", "ZQL1711-vince-build-20180504184901")
			req.Header.Add("os-api", "25")
			req.Header.Add("device-model", "Redmi+5+Plus")
			req.Header.Add("sign-tag", "084ae9243762e750")
			req.Header.Add("openudid", "f5d4c76537a06a57")
			req.Header.Add("phone-sim", "1")
			req.Header.Add("carrier", "%E4%B8%AD%E5%9B%BD%E7%A7%BB%E5%8A%A8")
			req.Header.Add("Content-Type",  w.FormDataContentType())
			req.Header.Add("Content-Length", "843")
			req.Header.Add("Host", "ktt.woyaoq.com")
			req.Header.Add("Connection", "Keep-Alive")
			req.Header.Add("Accept-Encoding", "gzip")
			req.Header.Add("User-Agent", "okhttp/3.11.0")

			//发起请求
			client := &http.Client{}
			resp, err = client.Do(req)

	}
	//-------------------------发起请求成功------------------------------------

	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))
	var r Response
	json.Unmarshal([]byte(body), &r)
	if(r.Success == true){
		return r.Items.Score
	}else{
		return 0
	}

	//解析返回body的话直接用断言就可以
}

func print_json(m map[string]interface{}) {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float", int64(vv))
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case nil:
			fmt.Println(k, "is nil", "null")
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			print_json(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
		}
	}
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}