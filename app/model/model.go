package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Post represents single blog post.
type Post struct {
	gorm.Model

	Title  string
	Posted time.Time
	Text   string
}
