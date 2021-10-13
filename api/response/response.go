package response

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GitResponse struct {
	Status string
	Repos  []Repo
}
