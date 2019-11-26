package Login

type AccountPayload struct {
	Password string `json:"password"`
	Type     string `json:"type"`
	UserName string `json:"userName"`
}

type Option struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

type Geographic struct {
	City     Option `json:"city"`
	Province Option `json:"province"`
}
