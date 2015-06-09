package target

type Target_data struct {
	Id int `json:"id"`
	Destination string `json:"destination"`
	Status bool `json:"status"`
}

type Load_fixtures struct {

}

func addTarget(destination string) {
	db, _ := sql.Open("postgres", "host=db user=postgres dbname=uptime sslmode=disable")
	var expectedTarget target.Target_data{Destination=destination}
	_,_ = db.Exec("INSERT testDestination (destination) VALUES($1)", expectedTarget.Destination)
}
