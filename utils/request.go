package utils

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
)

func LuyangmaoRequest(host string, path string ,key string,requestType int) []byte{
	//合成完整url
	urls := host + path

	//替换获取没有协议头的host
	host = strings.Replace(host, "https://", "", -1)
	host = strings.Replace(host, "http://", "", -1)

	var data = map[string]interface{}{}
	res := requestType
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
	//multipart/form-data
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
			"p": strings.NewReader(key),
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
					return nil
				}
			} else {
				// Add other fields
				if fw, err = w.CreateFormField(key); err != nil {
					return nil
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
		req.Header.Add("Host", host)
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
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	return body

	//解析返回body的话直接用断言就可以
}