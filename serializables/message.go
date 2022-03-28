package serializables

import "time"

type RestChatMessage struct {
	Id                 string    `json:"id"`
	Type               string    `json:"type"`
	ServerId           string    `json:"serverId"`
	ChannelId          string    `json:"channelId"`
	Content            string    `json:"content"`
	AuthorId           string    `json:"createdBy"`
	ReplyMessageIds    []string  `json:"replyMessageIds"`
	IsPrivate          bool      `json:"isPrivate"`
	CreatedByWebhookId string    `json:"createdByWebhookId"`
	UpdateTime         time.Time `json:"updateTime"`
	CreationTime       time.Time `json:"createdAt"`
}
