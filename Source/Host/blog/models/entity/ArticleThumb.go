package entity

import "time"

type ArticleThumb struct {
	ID        int       `json:"id",gorm:"primary_key"`
	ArticleId string    `json:"article_id"`
	ThumbUser int       `json:"thumb_user"`
	ThumbTime time.Time `json:"thumb_time"`
}
