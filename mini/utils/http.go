package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

//post请求
func Post(url string, data interface{}, contentType map[string]string) (string, error) {

	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest(`POST`, url, bytes.NewBuffer(jsonStr))
	for k, v := range contentType {
		req.Header.Add(k, v)
	}
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}
//get
func Get(url string, headers map[string]string) (string, error) {
	req, err := http.NewRequest(`GET`, url, nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	if err != nil {
		panic(err)
	}

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}
func PostFormData(url string,files,data map[string]string) (*http.Response, error) {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	for key, val := range files {
		fw, err := body_writer.CreateFormFile(key, val)
		if err != nil {
			fmt.Println("error writing to buffer")
			return nil, err
		}
		fb,_:=ioutil.ReadFile(val)
		_,_=fw.Write(fb)
	}
	for key, val := range data {
		err := body_writer.WriteField(key, val)
		if err != nil {
			fmt.Println("error writing to buffer")
			return nil, err
		}
	}
	_=body_writer.Close()

	//发送
	req, err := http.NewRequest("POST", url, body_buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type",body_writer.FormDataContentType())

	return http.DefaultClient.Do(req)
}