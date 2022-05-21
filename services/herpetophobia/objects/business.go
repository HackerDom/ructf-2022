package objects

type Map struct {
	Secret string    `json:"secret"`
	Init   [256]byte `json:"init"`
	Flag   string    `json:"flag"`
}

type Ids struct {
	Ids []string `json:"ids"`
}
