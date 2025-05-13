package excel

import (
	"context"
	"errors"
	"testing"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/mocks"
	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/usecases/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRepo_SavePerson(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo *mocks.MockPerson_2Repository
	}

	type args struct {
		ctx     context.Context
		persons []domain.Person_2
	}

	tests := []struct {
		name     string
		setup    func(f *fields)
		args     args
		wantErr  bool
		wantResp []string
	}{
		{
			name: "succes save",
			args: args{
				ctx: context.TODO(),
				persons: []domain.Person_2{
					{FirstName: "Luke", LastName: "Brandan"},
				},
			},
			setup: func(f *fields) {
				f.repo.EXPECT().
					SavePerson(gomock.Any(), gomock.Any()).
					Return([]string{"id1"}, nil)
			},
			wantResp: []string{"id1"},
		},
		{
			name: "error while saving",
			args: args{
				ctx: context.TODO(),
				persons: []domain.Person_2{
					{FirstName: "Pedro", LastName: "PicaPiedra"},
				},
			},
			setup: func(f *fields) {
				f.repo.EXPECT().
					SavePerson(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("DB ERROR"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				repo: mocks.NewMockPerson_2Repository(ctrl),
			}
			tt.setup(&f)

			got, err := f.repo.SavePerson(tt.args.ctx, tt.args.persons)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantResp, got)
			}
		})
	}
}

func TestRepo_SaveOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo *mocks.MockOrderRepository
	}

	type args struct {
		ctx    context.Context
		orders []domain.Order
	}

	tests := []struct {
		name     string
		setup    func(f *fields)
		args     args
		wantErr  bool
		wantResp []string
	}{

		{
			name: "save order ok",
			args: args{
				ctx: context.TODO(),
				orders: []domain.Order{
					{OrderID: "001", Customer: "Maria"},
				},
			},
			setup: func(f *fields) {
				f.repo.EXPECT().
					SaveOrder(gomock.Any(), gomock.Any()).
					Return([]string{"orderID_001"}, nil)
			},
			wantResp: []string{"orderID_001"},
		},
		{
			name: "fail saving order",
			args: args{
				ctx: context.TODO(),
				orders: []domain.Order{
					{OrderID: "fail DB"},
				},
			},
			setup: func(f *fields) {
				f.repo.EXPECT().
					SaveOrder(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("connection error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				repo: mocks.NewMockOrderRepository(ctrl),
			}
			tt.setup(&f)

			got, err := f.repo.SaveOrder(tt.args.ctx, tt.args.orders)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantResp, got)
			}
		})
	}

}
