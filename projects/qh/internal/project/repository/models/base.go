package models

type Client struct {
	ID      uint      `gorm:"primaryKey"`
	Name    string    `gorm:"size:100;not null"`
	Project []Project `gorm:"foreignKey:ClientID;constraint:OnDelete:CASCADE"`
}

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100;not null"`
	Email string `gorm:"size:100;unique;not null"`
}

type Manager struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"not null;index"`
	Title  string `gorm:"size:100"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Investor struct {
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `gorm:"not null;index"`
	Company string `gorm:"size:100"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type ProjectInvestor struct {
	ProjectID  uint    `gorm:"primaryKey;autoincrement:false"`
	InvestorID uint    `gorm:"primaryKey;autoincrement:false"`
	Percentage float64 `gorm:"not null"`

	Project  Project  `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
	Investor Investor `gorm:"foreignKey:InvestorID;constraint:OnDelete:CASCADE"`
}

type ProjectManager struct {
	ProjectID uint `gorm:"primaryKey;autoincrement:false"`
	ManagerID uint `gorm:"primaryKey;autoincrement:false"`

	Project Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
	Manager Manager `gorm:"foreignKey:ManagerID;constraint:OnDelete:CASCADE"`
}

type Project struct {
	ID       uint    `gorm:"primaryKey"`
	Name     string  `gorm:"size:100;not null"`
	ClientID uint    `gorm:"not null"`
	Client   Client  `gorm:"foreignKey:ClientID;constraint:OnDelete:CASCADE"`
	Fields   []Field `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
}

type Field struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	LeaseType string `gorm:"size:50"`
	ProjectID uint   `gorm:"not null"`
	Plots     []Plot `gorm:"foreignKey:FieldID;constraint:OnDelete:CASCADE"`
}

type Plot struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100"`
	Hectares    float64
	PreviusCrop string `gorm:"size:100"`
	CurrentCrop string `gorm:"size:100"`
	Season      string `gorm:"size:50"`
	FieldID     uint   `gorm:"not null"`
	Field       Field  `gorm:"foreignKey:FieldID;constraint:OnDelete:CASCADE"`
}
