package models

type Client struct {
	ID      int       `gorm:"primaryKey"`
	Name    string    `gorm:"size:100;not null"`
	Project []Project `gorm:"foreignKey:ClientID;constraint:OnDelete:CASCADE"`
}

type User struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"size:100;not null"`
	Email string `gorm:"size:100;unique;not null"`
}

type Manager struct {
	ID     int    `gorm:"primaryKey"`
	UserID int    `gorm:"not null;index"`
	Title  string `gorm:"size:100"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Investor struct {
	ID      int    `gorm:"primaryKey"`
	UserID  int    `gorm:"not null;index"`
	Company string `gorm:"size:100"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type ProjectInvestor struct {
	ProjectID  int     `gorm:"primaryKey"`
	InvestorID int     `gorm:"primaryKey"`
	Percentage float64 `gorm:"not null"`

	Project  Project  `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
	Investor Investor `gorm:"foreignKey:InvestorID;constraint:OnDelete:CASCADE"`
}

type ProjectManagers struct {
	ProjectID int `gorm:"primaryKey"`
	ManagerID int `gorm:"primaryKey"`

	Project Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
	Manager Manager `gorm:"foreignKey:ManagerID;constraint:OnDelete:CASCADE"`
}

type Project struct {
	ID       int     `gorm:"primaryKey"`
	Name     string  `gorm:"size:100;not null"`
	ClientID int     `gorm:"not null"`
	Client   Client  `gorm:"foreignKey:ClientID;constraint:OnDelete:CASCADE"`
	Fields   []Field `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
}

type Field struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	LeaseType string `gorm:"size:50"`
	ProjectID int    `gorm:"not null"`
	Plots     []Plot `gorm:"foreignKey:FieldID;constraint:OnDelete:CASCADE"`
}

type Plot struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"size:100"`
	Hectares    float64
	PreviusCrop string `gorm:"size:100"`
	CurrentCrop string `gorm:"size:100"`
	Season      string `gorm:"size:50"`
	FieldID     int    `gorm:"not null"`
	Field       Field  `gorm:"foreignKey:FieldID;constraint:OnDelete:CASCADE"`
}
