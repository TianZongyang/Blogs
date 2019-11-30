package facade

import "time"

type Article struct {
	ArticleId   string    `json:"uuid"`         //文章编号ID
	Title       string    `json:"title"`        //文章标题
	AuthorId    int       `json:"author_id"`    //作者ID编号
	Content     string    `json:"content"`      //文章内容
	Attaches    []string  `json:"attaches"`     //文章附件地址
	PublishTime time.Time `json:"publish_time"` //发布时间
}
