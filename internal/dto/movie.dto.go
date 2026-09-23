package dto

type Movie struct {
	Title  string `form:"title"`
	Genres string `form:"genre"`
}
