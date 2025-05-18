package models

type NoteCreateEntry struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	IsPublic *bool  `json:"is_public" binding:"required"`
}

type NoteUpdateEntry struct {
	Title    *string `json:"title"`
	Content  *string `json:"content"`
	IsPublic *bool   `json:"is_public"`
}
