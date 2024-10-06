package pages

import "chipa.me/pkg/services"

type BaseModel struct {
	Title string
}

type BlogModel struct {
	BaseModel
	Posts []*services.Post
}

type BlogEntryModel struct {
	BaseModel
	Post string
	Toc  string
}

type HomeModel struct {
	BaseModel
}

type EditorModel struct {
	BaseModel
}
