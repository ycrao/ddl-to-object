package entity
// generated by ddl-to-object <https://github.com/ycrao/ddl-to-object>

import (
	"time"
)

// Article article
type Article struct {  
	Id uint64 `json:"id" db:"id"`  // id 
	UserId int64 `json:"user_id" db:"user_id"`  // 用户id 
	Content string `json:"content" db:"content"`  // 正文 
	CreateTime time.Time `json:"create_time" db:"create_time"`  // 创建时间 
	UpdateTime time.Time `json:"update_time" db:"update_time"`  // 更新时间  
}
