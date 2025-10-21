package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string

	// 文章数量统计
	PostCount uint
	Posts     []Post
}

type Post struct {
	gorm.Model
	Title string

	UserId uint // 外键
	User   User

	// 评论状态
	CommentStatus string
	Comments      []Comment
}

type Comment struct {
	Content string
	gorm.Model

	PostId uint
	Post   Post

	UserId uint
	User   User
}

// 创建之后自动调用的钩子
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", p.UserId).
		UpdateColumn("post_count", gorm.Expr("post_count +?", 1)).
		Error
}

// 删除后触发自动调用的钩子
func (c *Comment) AfterDelete(tx *gorm.DB) error {

	// 统计文章的回复
	var count int64
	tx.Debug().Model(&Comment{}).Where("post_id =?", c.PostId).Count(&count)

	// 如果没有回复的话,就把文章改成无评论
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostId).Update("comment_status", "无评论").Error
	}
	return nil
}
