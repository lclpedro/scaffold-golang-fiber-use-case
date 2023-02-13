package domains

type Address struct {
	ID         int
	StreetName string
	ZipCode    string
	City       string
	State      string
	Country    string
}

func NewAddress(streetName, zipCode, city, state, country string) *Address {
	return &Address{
		StreetName: streetName,
		ZipCode:    zipCode,
		City:       city,
		State:      state,
		Country:    country,
	}
}
