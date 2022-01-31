package v1

import (
	validator "gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"

	metav1 "github.com/marmotedu/goserver/pkg/meta/v1"
	"github.com/marmotedu/goserver/pkg/util/id"
)

// Post represents a blog for a user.
type Post struct {
	BaseModel

	Username string `json:"username,omitempty" gorm:"column:username;not null"`

	PostID string `json:"postID,omitempty" gorm:"column:postID;not null"`

	Title string `json:"title" gorm:"column:title;not null" binding:"required" validate:"min=1,max=256"`

	Content string `json:"content" gorm:"column:content" binding:"required"`
}

// TableName maps to mysql table name.
func (p *Post) TableName() string {
	return "post"
}

// BeforeCreate run before create database record.
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	p.PostID = "post-" + id.GenShortId()

	return nil
}

// PostList is the whole list of all posts which have been stored in stroage.
type PostList struct {
	metav1.ListMeta `json:",inline"`

	Items []*Post `json:"items"`
}

// Validate the fields.
func (p *Post) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
