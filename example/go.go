package example

import "time"

// Article Object for article
type Article struct {
	Id         int64     `json:"id" db:"id"`                   // id
	UserId     string    `json:"user_id" db:"user_id"`         // 用户id
	Content    string    `json:"content" db:"content"`         // 内容
	CreateTime time.Time `json:"create_time" db:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time" db:"update_time"` // 更新时间
}
