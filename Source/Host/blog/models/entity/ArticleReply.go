package entity

import "time"

type ArticleReply struct {
	ID        int       `json:"id",gorm:"primary_key"`
	ReplyUUID string    `json:"reply_uuid"`
	ReplyTo   string    `json:"reply_to"`
	ReplyType int       `json:"reply_type"`
	Value     string    `json:"value"`
	ReplyTime time.Time `json:"reply_time"`
}

func (ArticleReply) TableName() string {
	return "article_reply"
}
