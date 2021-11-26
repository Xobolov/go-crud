package models

import (
	"go-simple-crud/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title, Content string
	Description    string
}

func (Post) Migrate() {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.CheckDbCon(err)

	db.AutoMigrate(&Post{})
}

func (Post) GetAll() []Post {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.CheckDbCon(err)

	var posts []Post

	db.Find(&posts)
	return posts
}

func (post Post) Get(id string) Post {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.CheckDbCon(err)

	db.First(&post, id)

	return post
}

func (Post) Add(title, content, description string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.CheckDbCon(err)

	db.Create(&Post{Title: title, Content: content, Description: description})
}

func (post Post) Update(title, content, description string, id uint)  {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.CheckDbCon(err)

	post.ID = id
	post.Title = title
	post.Content = content
	post.Description = description

	db.Save(&post)
}

func (Post) Delete(id string)  {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.CheckDbCon(err)

	db.Delete(&Post{}, id)
}


