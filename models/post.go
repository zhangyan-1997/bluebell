package models

import "time"

//注意内存对齐，相同类型的字段放一块
type Post struct {
	PostID      int64     `json:"post_id" db:"post_id"`
	AuthorId    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content"`
	CreateTime  time.Time `json:"-" db:"create_time"`
}

// ApiPostDetail 帖子详情接口的结构
type ApiPostDetail struct {
	AuthorName       string `json:"author_name"`
	VoteNum          int64
	*Post                               //嵌入帖子结构体
	*CommunityDetail `json:"community"` //嵌入社区信息
}
