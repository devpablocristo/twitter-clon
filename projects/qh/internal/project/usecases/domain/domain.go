package domain

type FullProject struct {
	Client   Client
	Project  Project
	Manager  Manager
	Investor Investor
	Field    Field
	Plot     Plot
}

type Client struct {
	ID   uint
	Name string
}

type User struct {
	ID    uint
	Name  string
	Email string
}

type Manager struct {
	ID     uint
	UserID uint
	Title  string
	User   User
}

type Investor struct {
	ID         uint
	UserID     uint
	Company    string
	User       User
	Percentage float64 // este campo no está en DB, pero lo usamos para el ProjectInvestor
}

type Project struct {
	ID       uint
	Name     string
	ClientID uint
	Client   Client
	Fields   Field
}

type ProjectManagers struct {
	ProjectID uint
	ManagerID uint
	Project   Project
	Manager   Manager
}

type ProjectInvestor struct {
	ProjectID  uint
	InvestorID uint
	Percentage float64
	Project    Project
	Investor   Investor
}

type Field struct {
	ID        uint
	Name      string
	LeaseType string
	ProjectID uint
	Plots     []Plot
}

type Plot struct {
	ID          uint
	Name        string
	Hectares    float64
	PreviusCrop string
	CurrentCrop string
	Season      string
	FieldID     uint
}
