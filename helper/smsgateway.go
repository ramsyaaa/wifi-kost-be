package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func sendSMS() {

	url := "https://5ymd2y.api.infobip.com/sms/2/text/advanced"
	method := "POST"

	payload := strings.NewReader(`{"messages":[{"destinations":[{"to":"6285155078654"}],"from":"ServiceSMS","text":"Hello,\n\nThis is a test message from Infobip. Have a nice day!"}]}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "App de8a03b0c84fbc5e4d0b324a2a855aea-55ef62bf-ecf0-4356-9d9a-c21a4c457c11")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
