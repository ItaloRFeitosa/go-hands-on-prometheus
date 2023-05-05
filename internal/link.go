package internal

type Link struct {
	ID   int64  `json:"-"`
	URL  string `json:"url" validate:"required,url"`
	Slug string `json:"slug"`
}
