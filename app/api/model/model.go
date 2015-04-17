package api 

type Ip struct{
	Destination string   `json:"destination"`
}

type Ips []Ip
