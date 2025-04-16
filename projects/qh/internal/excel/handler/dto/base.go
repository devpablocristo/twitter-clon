package dto

type ExcelPerson struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
}

/*func (E *ExcelPerson) ToDomain() domain.Person_2 {
	return domain.Person_2{
		FirstName: E.FirstName,
		LastName:  E.LastName,
		Age:       E.Age,
		Phone:     E.Phone,
	}
}*/
