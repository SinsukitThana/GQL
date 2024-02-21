package person

type Persons struct {
	PersonID  int    `bun:"personid"`
	LastName  string `bun:"lastname"`
	FirstName string `bun:"firstname"`
	Address   string `bun:"address"`
	City      string `bun:"city,"`
}
