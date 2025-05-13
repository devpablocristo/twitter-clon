package dto

type ExcelPerson struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
}

type OrderDto struct {
	OrderID     string  `json:"order_id"`
	Customer    string  `json:"customer"`
	TotalAmount float64 `json:"total_amount"`
	Date        string  `json:"date"`
	Status      string  `json:"status"`
}

/*func (E *ExcelPerson) ToDomain() domain.Person_2 {
	return domain.Person_2{
		FirstName: E.FirstName,
		LastName:  E.LastName,
		Age:       E.Age,
		Phone:     E.Phone,
	}
}*/
