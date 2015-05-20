package target

type Target_data struct {
	Id int `json:"id"`
	Destination string `json:"destination"`
	Status bool `json:"status"`
}
