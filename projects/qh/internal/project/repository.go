package project

import (
	"context"

	"github.com/devpablocristo/monorepo/projects/qh/internal/project/usecases/domain"
	"github.com/devpablocristo/monorepo/projects/qh/internal/project/utils"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) FullProjectsRepo {
	return &repository{db: db}
}

func (r *repository) SaveFullProject(ctx context.Context, fps []domain.FullProject) ([]string, error) {
	var createdIDs []string

	for _, fp := range fps {
		models := utils.DomainToModels(fp)

		tx := r.db.WithContext(ctx).Begin()

		// creando client
		if err := tx.Create(&models.Client).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// creando ManagerUser
		if err := tx.Create(&models.ManagerUser).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// creando manager (usando el ID de ManagerUser)
		models.Manager.ID = models.ManagerUser.ID
		if err := tx.Create(&models.Manager).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// creando InvestorUser
		if err := tx.Create(&models.InvestorUser).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// creando investor (usando el ID del InvestorUser)
		models.Investor.UserID = models.InvestorUser.ID
		if err := tx.Create(&models.Investor).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// creando Project (usando el ID de Client)
		models.Project.ClientID = models.Client.ID
		if err := tx.Create(&models.Project).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// Creando  ProjectManager (ProjectID y ManagerID ya disponibles)
		models.ProjectManager.ProjectID = models.Project.ID
		models.ProjectManager.ManagerID = models.Manager.ID
		if err := tx.Create(&models.ProjectManager).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// creando ProjectInvestor
		models.ProjectInvestor.ProjectID = models.Project.ID
		models.ProjectInvestor.InvestorID = models.Investor.ID
		if err := tx.Create(&models.ProjectInvestor).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// creardo campo field
		models.Field.ProjectID = models.Project.ID
		if err := tx.Create(&models.Field).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// creando plot
		models.Plot.FieldID = models.Field.ID
		if err := tx.Create(&models.Plot).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		tx.Commit()

		createdIDs = append(createdIDs, models.Project.Name)
	}

	return createdIDs, nil
}
