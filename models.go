package main


// Model of a Author
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Model of a book
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Author *Author `json:"author"`
	Title  string  `json:"title"`
}

type Error struct {
	Code string `json:"error_code"`
	Message string `json:"message"`
}
