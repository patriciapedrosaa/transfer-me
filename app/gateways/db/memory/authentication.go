package memory

import "time"

type Token struct {
	ID        string
	Name      string
	Subject   string
	Issuer    string
	IssuedAt  time.Time
	ExpiredAt time.Time
}
