package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func CallRestAPI(method string, url string, body interface{}) (err error) {

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)
	request, _ := http.NewRequest(method, url, payload)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 || err != nil {
		return err
	}

	return nil
}
