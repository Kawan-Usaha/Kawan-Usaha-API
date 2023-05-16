package Model

type Recommender struct {
	Tags       []string   `json:"tags"`
	Categories []Category `json:"categories"`
	Articles   []Article  `json:"articles"`
}
