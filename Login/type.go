package Login

type AccountPayload struct {
	Password     string    `json:"password"`
	Type         string	   `json:"type"`
	UserName     string    `json:"userName"`
}
