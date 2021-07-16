package memory

import "time"

type Account struct {
	id        string
	name      string
	cpf       string
	secret    string
	balance   int
	createdAt time.Time
}
