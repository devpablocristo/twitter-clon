package utils

import (
	"github.com/devpablocristo/monorepo/projects/qh/internal/project/repository/models"
	"github.com/devpablocristo/monorepo/projects/qh/internal/project/usecases/domain"
)

type FullProjectModels struct {
	Client          models.Client
	ManagerUser     models.User
	Manager         models.Manager
	InvestorUser    models.User
	Investor        models.Investor
	Project         models.Project
	ProjectInvestor models.ProjectInvestor
	ProjectManager  models.ProjectManager
	Field           models.Field
	Plot            models.Plot
}

// esta funcion mapea los tipos domain a tipo modelos de GORM
func DomainToModels(fp domain.FullProject) *FullProjectModels {
	//users
	managerUser := models.User{
		Name:  fp.Manager.User.Name,
		Email: fp.Manager.User.Email,
	}
	investorUser := models.User{
		Name:  fp.Investor.User.Name,
		Email: fp.Investor.User.Email,
	}

	// plot  <--- solo estoy usando el primer plot, luego voy a agregar por si hay mas de uno
	var plot models.Plot
	if len(fp.Field.Plots) > 0 {
		p := fp.Field.Plots[0]
		plot = models.Plot{
			Name:        p.Name,
			Hectares:    p.Hectares,
			PreviusCrop: p.PreviusCrop,
			CurrentCrop: p.CurrentCrop,
			Season:      p.Season,
		}
	}

	return &FullProjectModels{
		Client: models.Client{
			Name: fp.Client.Name,
		},
		ManagerUser: managerUser,
		Manager: models.Manager{
			Title: fp.Manager.Title,
			// UserID se asigna despues de guardar el managerUser
		},
		InvestorUser: investorUser,
		Investor: models.Investor{
			Company: fp.Investor.Company,
			// UserID se asigna despues de guardar el investorUser
		},
		Project: models.Project{
			Name: fp.Project.Name,
			// ClientID se asigna despues de guardar Client
		},
		ProjectManager: models.ProjectManager{
			// ProjectID y ManagerID se asigna despues de crear Project y Manager
		},

		ProjectInvestor: models.ProjectInvestor{
			Percentage: fp.Investor.Percentage,
			// ProjectID y InvestorID se asigna despues
		},
		Field: models.Field{
			Name:      fp.Field.Name,
			LeaseType: fp.Field.LeaseType,
			// ProjectID se asigna despues de crear Project
		},
		Plot: plot,
	}
}
