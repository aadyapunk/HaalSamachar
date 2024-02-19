package models

type User struct {
	ID       string
	Username string
	Email    string
	// Add more fields as needed
}

type BlogPost struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type Comment struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	BlogID    int    `json:"blog_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
