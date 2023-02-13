package models

type PersonModel struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Age       int    `db:"age"`
	AddressID int    `db:"address_id"`
}
