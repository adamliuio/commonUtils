package commonUtils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Adam's notifications
func NotifyMaster(strToSend, webhookURL string) {
	values := map[string]string{"text": strToSend}
	jsonValue, _ := json.Marshal(values)

	res, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Println(err)
	}
	if Mode == "testing" {
		log.Printf("res Status: %s\n", res.Status)
		log.Printf("res: %+v\n", res)
	}
}

func ServiceStatusUpdate(serviceName string, arr []map[string]string, webhookURL string) {
	values := map[string]string{"blocks": blocksFormatter(serviceName, arr)}
	jsonValue, _ := json.Marshal(values)

	res, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Println(err)
	}
	if Mode == "testing" {
		log.Printf("res Status: %s\n", res.Status)
		log.Printf("res: %+v\n", res)
	}
}

func blocksFormatter(serviceName string, arr []map[string]string) (msg string) {
	fieldFormat := `{"type": "mrkdwn", "text": "*%s:*\n%s"},`
	for _, a := range arr {
		msg += fmt.Sprintf(fieldFormat, a["title"], a["value"])
	}
	msg = fmt.Sprintf(blockString, serviceName, string(msg[:len(msg)-1]))
	return
}

var blockString string = `[
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": "New Service Update:\n💡 *%s*"
		}
	},
	{
		"type": "section",
		"fields": [
			%s
		]
	}
]`
