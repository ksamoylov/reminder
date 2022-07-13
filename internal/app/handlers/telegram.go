package handlers

import (
	"io/ioutil"
	"net/http"
)

const MessageList = "list"

func HandleMessage(message string) []byte {
	var repliedMessage []byte

	switch message {
	case MessageList:
		repliedMessage = handleList()
	}

	return repliedMessage
}

func handleList() []byte {
	resp, err := http.Get("http://localhost:8080/note/list/")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	jsonData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return jsonData
}
