package email

type Email struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type ReqEmail struct {
	Email string `json:"email"`
}
