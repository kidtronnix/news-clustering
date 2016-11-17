package article

// Article is the struct that models an article
type Article struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Time        int64  `json:"time"`
	URL         string `json:"url"`
}
