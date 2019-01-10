package types

type User struct {
	Id string	`json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Relationship struct {
	UserId string `json:"user_id"`
	State string `json:"state"`
	Type string `json:"type"`
}
