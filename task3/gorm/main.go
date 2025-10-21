package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

/*
进阶gorm
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。

题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。

题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func main() {
	db, err := gorm.Open(sqlite.Open("gorm1.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Migrator().DropTable(&User{}, &Post{}, &Comment{})
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	gorm1(db)

	queryUserPostsWithComments(db)

	queryMostComments(db)

	testHooks(db)
}

func gorm1(db *gorm.DB) {
	u := &User{
		UserName: "张三",
	}
	db.Create(u)

	p1 := &Post{
		Title:  "张三的标题1",
		UserId: u.ID,
	}

	p2 := &Post{
		Title:  "张三的标题2",
		UserId: u.ID,
	}

	db.Create(p1)
	db.Create(p2)

	c1 := &Comment{
		Content: "回复1",
		PostId:  p1.ID,
		UserId:  u.ID,
	}

	c2 := &Comment{
		Content: "回复2",
		PostId:  p1.ID,
		UserId:  u.ID,
	}

	c3 := &Comment{
		Content: "回复3",
		PostId:  p2.ID,
		UserId:  u.ID,
	}
	c4 := &Comment{
		Content: "回复4",
		PostId:  p2.ID,
		UserId:  u.ID,
	}
	c5 := &Comment{
		Content: "回复5",
		PostId:  p2.ID,
		UserId:  u.ID,
	}
	db.Create(c1)
	db.Create(c2)
	db.Create(c3)
	db.Create(c4)
	db.Create(c5)

}

func queryUserPostsWithComments(db *gorm.DB) {
	var user User
	result := db.Preload("Posts.Comments").Find(&user, 1)
	if result.Error != nil {
		fmt.Println("查询失败", result.Error)
	} else {
		log.Printf("用户：%s\n", user.UserName)
		for _, p := range user.Posts {
			log.Printf("  文章《%s》 共有 %d 条评论", p.Title, len(p.Comments))
			for _, c := range p.Comments {
				log.Printf("    评论：%s", c.Content)
			}
		}
	}
}

func queryMostComments(db *gorm.DB) {
	var id uint
	// 找出评论最多的文章
	db.Model(&Comment{}).
		Select("post_id").
		Group("post_id").
		Order("count(*)  desc").
		Limit(1).
		Pluck("post_id", &id)

	var post Post
	db.Preload("Comments").Where("id = ?", id).First(&post)

	fmt.Printf("评论最多的文章标题 : %s \n", post.Title)

	for _, comment := range post.Comments {
		fmt.Printf("评论最多的文章的回复 : %s \n", comment.Content)
	}

}

func testHooks(db *gorm.DB) {
	// 测试用户文章数钩子
	var user User
	db.First(&user, 1)
	fmt.Printf("用户 %s 的文章数为 %d \n", user.UserName, user.PostCount)

	// 先查询出要删除的评论
	var comments []Comment
	db.Where("post_id = ?", 1).Find(&comments)

	// 逐个删除，这样钩子才能正确获取 PostId
	for _, comment := range comments {
		db.Delete(&comment)
	}

	var post Post
	db.First(&post, 1)
	fmt.Printf("文章 %s 的评论状态: %s \n", post.Title, post.CommentStatus)
}
