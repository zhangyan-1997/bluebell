package mysql

import (
	"bluebell/models"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) (err error) {
	sqlStr := `insert into post(
				post_id, title, content, author_id, community_id)
				values(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, post.PostID, post.Title, post.Content, post.AuthorId, post.CommunityID)
	return err
}

// GetPostById 根据id查询单个帖子数据
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select
			post_id, title, content, author_id, community_id, create_time
			from post
			where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}

// GetPostList 查询帖子列表
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select
			post_id, title, content, author_id, community_id, create_time
			from post
			ORDER BY create_time
			DESC
			limit ?, ?
			`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
