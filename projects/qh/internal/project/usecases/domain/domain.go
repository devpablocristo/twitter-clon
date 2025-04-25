package domain

type Client struct {
	ID      int
	Name    string
	Project []Project
}

type User struct {
	ID    int
	Name  string
	Email string
}

type Manager struct {
	ID     int
	UserID int
	Title  string
	User   User
}

type Investor struct {
	ID      int
	UserID  int
	Company string
	User    User
}

type ProjectInvestor struct {
	ProjectID  int
	InvestorID int
	Percentage float64
	Project    Project
	Investor   Investor
}

type ProjectManagers struct {
	ProjectID int
	ManagerID int
	Project   Project
	Manager   Manager
}

type Project struct {
	ID       int
	Name     string
	ClientID int
	Client   Client
	Fields   []Field
}

type Field struct {
	ID        int
	Name      string
	LeaseType string
	ProjectID int
	Plots     []Plot
}

type Plot struct {
	ID          int
	Name        string
	Hectares    float64
	PreviusCrop string
	CurrentCrop string
	Season      string
	FieldID     int
	Field       Field
}
