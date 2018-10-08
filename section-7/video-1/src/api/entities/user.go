package entities

import "time"

type User struct {
	ID        uint32    `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
	Added     time.Time `json:"added,omitempty"`
	Account   uint32    `json:"account"`
	Password  string    `json:"password"`
}

type UserAccountType int

const (
	NormalAccount UserAccountType = iota
	SilverAccount
	GoldAccount
	DiamondAccount
)

func (u *User) GetAccountType() UserAccountType {
	if u.Account >= 5000 {
		return DiamondAccount
	} else if u.Account >= 3500 {
		return GoldAccount
	} else if u.Account >= 1000 {
		return SilverAccount
	} else {
		return NormalAccount
	}
}
