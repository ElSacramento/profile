package models

type User struct {
	ID        uint64 `pg:"id"`
	FirstName string `pg:"first_name"`
	LastName  string `pg:"last_name"`
	Nickname  string `pg:"nickname"`
	Password  string `pg:"password"`
	Email     string `pg:"email"`
	Country   string `pg:"country"` // todo: enum
}
