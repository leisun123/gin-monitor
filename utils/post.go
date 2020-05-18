package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func HttpPost(url string, params string) error {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(body)
	fmt.Println(time.Now())
	return nil
}