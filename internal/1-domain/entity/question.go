package entity

// Question は質問エンティティです。
type Question struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
