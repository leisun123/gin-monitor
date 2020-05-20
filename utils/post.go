package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPost(url string, params string) (string,error) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		fmt.Println(err)
		return "",err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error(),err
	}
	if resp.StatusCode != 200 {
		return "访问失败", nil
	}
	return string(body),nil
}