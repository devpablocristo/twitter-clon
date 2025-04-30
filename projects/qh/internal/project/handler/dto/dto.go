package dto

import "github.com/devpablocristo/monorepo/projects/qh/internal/project/usecases/domain"

type FullProjectDTO struct {
	// Client y Project
	ClientName  string `json:"client_name"`
	ProjectName string `json:"project_name"`

	// Manager
	ManagerName  string `json:"manager_name"`
	ManagerEmail string `json:"manager_email"`
	ManagerTitle string `json:"manager_title"`

	// Investor
	InvestorName    string  `json:"investor_name"`
	InvestorEmail   string  `json:"investor_email"`
	InvestorCompany string  `json:"investor_company"`
	InvestmentPct   float64 `json:"investment_pct"`

	// Field y Plot
	FieldName    string  `json:"field_name"`
	LeaseType    string  `json:"lease_type"`
	PlotName     string  `json:"plot_name"`
	Hectares     float64 `json:"hectares"`
	PreviousCrop string  `json:"previous_crop"`
	CurrentCrop  string  `json:"current_crop"`
	Season       string  `json:"season"`
}

// convierte a tipos DTO en tipo Domain
func DtoToDomain(dto FullProjectDTO) domain.FullProject {
	return domain.FullProject{
		Client: domain.Client{
			Name: dto.ClientName,
		},
		Project: domain.Project{
			Name: dto.ProjectName,
		},
		Manager: domain.Manager{
			Title: dto.ManagerTitle,
			User: domain.User{
				Name:  dto.ManagerName,
				Email: dto.ManagerEmail,
			},
		},
		Investor: domain.Investor{
			Company: dto.InvestorCompany,
			User: domain.User{
				Name:  dto.InvestorName,
				Email: dto.InvestorEmail,
			},
			Percentage: dto.InvestmentPct,
		},
		Field: domain.Field{
			Name:      dto.FieldName,
			LeaseType: dto.LeaseType,
			Plots: []domain.Plot{
				{
					Name:        dto.PlotName,
					Hectares:    dto.Hectares,
					PreviusCrop: dto.PreviousCrop,
					CurrentCrop: dto.CurrentCrop,
					Season:      dto.Season,
				},
			},
		},
	}
}
