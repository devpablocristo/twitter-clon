package integrationtest

import (
	"context"
	"testing"

	"github.com/devpablocristo/monorepo/projects/qh/internal/project"
	"github.com/devpablocristo/monorepo/projects/qh/internal/project/repository/models"
	"github.com/devpablocristo/monorepo/projects/qh/internal/project/usecases/domain"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Dialector{
		DSN:        "file::memory:?cache=shared",
		DriverName: "sqlite",
	}, &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&models.Client{}, &models.Project{}, &models.User{}, &models.Investor{}, &models.ProjectManager{}, &models.ProjectInvestor{}, &models.Field{}, &models.Plot{})
	require.NoError(t, err)

	return db
}

func TestSaveFullProject(t *testing.T) {
	db := setupTestDB(t)
	repo := project.NewRepository(db)

	fullProject := domain.FullProject{
		Client:  domain.Client{Name: "Test Client"},
		Project: domain.Project{Name: "Test Project"},
		Manager: domain.Manager{
			Title: "Supervisor",
			User: domain.User{
				Name:  "Alice",
				Email: "alice@example.com",
			},
		},
		Investor: domain.Investor{
			Company: "InvestCo",
			User: domain.User{
				Name:  "Bob",
				Email: "bob@example.com",
			},
			Percentage: 45.0,
		},
		Field: domain.Field{
			Name:      "North Field",
			LeaseType: "Arrendamiento",
			Plots: []domain.Plot{
				{
					Name:        "Lote A",
					Hectares:    5.0,
					PreviusCrop: "Soja",
					CurrentCrop: "Maíz",
					Season:      "2025",
				},
			},
		},
	}

	ids, err := repo.SaveFullProject(context.TODO(), []domain.FullProject{fullProject})
	require.NoError(t, err)
	require.Len(t, ids, 1)

	var p models.Project
	err = db.Preload("Client").First(&p).Error
	require.Equal(t, "Test Client", p.Client.Name)

	var pm models.ProjectManager
	require.NoError(t, db.Preload("Manager").First(&pm).Error)

	require.NoError(t, db.Preload("User").First(&pm.Manager).Error)

	require.Equal(t, "Alice", pm.Manager.User.Name)

	var pi models.ProjectInvestor
	require.NoError(t, db.Preload("Investor").First(&pi).Error)
	require.NoError(t, db.Preload("User").First(&pi.Investor).Error)

	require.Equal(t, "Bob", pi.Investor.User.Name)

	var f models.Field
	require.NoError(t, db.Where("project_id = ?", p.ID).First(&f).Error)
	require.Equal(t, "North Field", f.Name)

	var plot models.Plot
	require.NoError(t, db.Where("field_id = ?", f.ID).First(&plot).Error)
	require.Equal(t, "Lote A", plot.Name)
	require.Equal(t, "Maíz", plot.CurrentCrop)
}
