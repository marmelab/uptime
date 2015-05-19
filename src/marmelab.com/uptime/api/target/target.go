package target

type Ip struct {
	Id int `json:"id"`
	Destination string `json:"destination"`
	Status bool `json:"status"`
}
