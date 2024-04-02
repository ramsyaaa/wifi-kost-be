package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendSMS(msisdn string, guestHouseName string, expiryDays int) error {
	url := "https://5ymd2y.api.infobip.com/sms/2/text/advanced"
	method := "POST"

	message := fmt.Sprintf("Paket Internet %s WIFI Anda untuk %d Hari telah aktif, berikut password untuk akses Anda: password\nAnda dapat login melalui link http://192.168.1.1", guestHouseName, expiryDays)

	payloadData := map[string]interface{}{
		"messages": []map[string]interface{}{
			{
				"destinations": []map[string]interface{}{
					{"to": msisdn},
				},
				"from": "ServiceSMS",
				"text": message,
			},
		},
	}

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		return err
	}

	payload := bytes.NewReader(payloadBytes)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "App de8a03b0c84fbc5e4d0b324a2a855aea-55ef62bf-ecf0-4356-9d9a-c21a4c457c11")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
