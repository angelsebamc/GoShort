package link_dto

type LinkDTO_Post struct {
	OriginalUrl string `json:"original_url"`
}

type LinkDTO_Info struct {
	ShortUrl    string `json:"short_url"`
	OriginalUrl string `json:"original_url"`
	UserID      string `json:"user_id"`
}

type LinkDTO_Get struct {
	ID          string `json:"id"`
	ShortUrl    string `json:"short_url"`
	OriginalUrl string `json:"original_url"`
	UserID      string `json:"user_id"`
	Clicks      int    `json:"clicks"`
}
