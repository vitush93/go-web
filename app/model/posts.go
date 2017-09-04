package model

import "time"

// CreatePost will create a new blog post.
func CreatePost(title string, text string) {
	now := time.Now()

	DB.Create(&Post{Title: title, Text: text, Posted: now})
}

// FindAllPosts retrieves all posts from database ordered by date posted descending.
func FindAllPosts() []Post {
	var posts []Post
	DB.Order("posted desc").Find(&posts)

	return posts
}
