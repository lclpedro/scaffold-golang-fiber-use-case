package models

type AddressModel struct {
	ID         int    `db:"id"`
	StreetName string `db:"street_name"`
	ZipCode    string `db:"zip_code"`
	City       string `db:"city"`
	State      string `db:"state"`
	Country    string `db:"country"`
}
