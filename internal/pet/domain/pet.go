package domain

type Pet struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Especie string `json:"especie"`
}
