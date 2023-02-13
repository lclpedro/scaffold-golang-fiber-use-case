package domains

type Person struct {
	ID        int64
	FistName  string
	LastName  string
	Age       int
	AddressID int64
}

func NewPerson(firstName, lastName string, age int, addressID int64) *Person {
	return &Person{
		FistName:  firstName,
		LastName:  lastName,
		Age:       age,
		AddressID: addressID,
	}
}
