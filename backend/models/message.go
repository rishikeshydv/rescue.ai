package models

import "time"

type Message struct {
	MessageDepartment string    `json:"messageDepartment"`
	MessageType       string    `json:"messageType"`
	DeliveredTime     time.Time `json:"deliveredTime"`
	MessageContent    string    `json:"messageContent"`
	SenderUser        string    `json:"senderUser"`
}
