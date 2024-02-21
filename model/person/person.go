package person

type Persons struct {
	PersonID  int    `bun:"personid"`
	LastName  string `bun:"lastname"`
	FirstName string `bun:"firstname"`
	Address   string `bun:"address"`
	City      string `bun:"city,"`
}

type PersonsOBJ struct {
	PersonID  int    `json:"personid"`
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
	Address   string `json:"address"`
	City      string `json:"city,"`
}
