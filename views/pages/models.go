package pages

type BaseModel struct {
	Title string
}

type HomeModel struct {
	BaseModel
	Post string
	Toc  string
}
