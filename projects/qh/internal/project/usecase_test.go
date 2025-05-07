package project

import (
	"context"
	"errors"
	"testing"

	"github.com/devpablocristo/monorepo/projects/qh/internal/project/handler/dto"
	"github.com/devpablocristo/monorepo/projects/qh/internal/project/mocks"
	"github.com/devpablocristo/monorepo/projects/qh/internal/project/usecases/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRepo_SaveFullProject(t *testing.T) {

	type fields struct {
		repo *mocks.MockFullProjectsRepo
	}

	type args struct {
		ctx      context.Context
		projects []domain.FullProject
	}

	tests := []struct {
		name    string
		setup   func(f *fields)
		args    args
		wantErr bool
		wantRes []string
	}{
		{
			name: "Succes salving full project",
			setup: func(f *fields) {
				f.repo.EXPECT().
					SaveFullProject(gomock.Any(), gomock.Any()).
					Return([]string{"fp1"}, nil)
			},
			args: args{
				ctx:      context.Background(),
				projects: []domain.FullProject{{}},
			},
			wantRes: []string{"fp1"},
		},
		{
			name: "Failed salving full project",
			setup: func(f *fields) {
				f.repo.EXPECT().
					SaveFullProject(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("Error DB"))
			},
			args: args{
				ctx:      context.Background(),
				projects: []domain.FullProject{{}},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mocks.NewMockFullProjectsRepo(ctrl),
			}
			tt.setup(&f)

			got, err := f.repo.SaveFullProject(tt.args.ctx, tt.args.projects)
			if tt.wantErr {
				require.Error(t, err, "expected an error but got none")
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, got, "response mismatch")
			}
		})
	}
}

func TestProcessFullProjects(t *testing.T) {
	type fields struct {
		usecase *mocks.MockUseCases
	}

	type args struct {
		ctx     context.Context
		project []dto.FullProjectDTO
	}

	tests := []struct {
		name    string
		setup   func(f *fields)
		args    args
		wantErr bool
		wantRes []string
	}{
		{
			name: "sucess process",
			setup: func(f *fields) {
				f.usecase.EXPECT().
					ProcessFullProjects(gomock.Any(), gomock.Any()).
					Return([]string{"id1"}, nil)
			},
			args: args{
				ctx:     context.Background(),
				project: []dto.FullProjectDTO{{}},
			},
			wantRes: []string{"id1"},
		},
		{
			name: "failed process",
			setup: func(f *fields) {
				f.usecase.EXPECT().
					ProcessFullProjects(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("file cannot be empty"))
			},
			args: args{
				ctx:     context.Background(),
				project: []dto.FullProjectDTO{{}},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				usecase: mocks.NewMockUseCases(ctrl),
			}
			tt.setup(&f)

			got, err := f.usecase.ProcessFullProjects(tt.args.ctx, tt.args.project)
			if tt.wantErr {
				require.Error(t, err, "expected an error but got none")
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantRes, got, "response mismatch")
			}
		})
	}
}
