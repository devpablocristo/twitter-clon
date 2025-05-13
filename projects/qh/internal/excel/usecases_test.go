package excel

import (
	"context"
	"errors"
	"testing"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/handler/dto"
	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_ProccesPerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		useCase *mocks.MockUseCases
	}

	type args struct {
		ctx    context.Context
		person []dto.ExcelPerson
	}

	tests := []struct {
		name     string
		setup    func(f *fields)
		args     args
		wantErr  bool
		WantResp []string
	}{
		{
			name: "succes",
			args: args{
				ctx: context.TODO(),
				person: []dto.ExcelPerson{
					{FirstName: "Luke", LastName: "Brandan", Age: 20, Phone: "3858384321"},
				},
			},
			setup: func(f *fields) {
				f.useCase.EXPECT().
					ProccesPerson(gomock.Any(), gomock.Any()).
					Return([]string{"id1"}, nil)
			},
			WantResp: []string{"id1"},
		},
		{
			name: "error usecase",
			args: args{
				ctx: context.TODO(),
				person: []dto.ExcelPerson{
					{FirstName: "john", LastName: "Smith"},
				},
			},
			setup: func(f *fields) {
				f.useCase.EXPECT().
					ProccesPerson(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("procces error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				useCase: mocks.NewMockUseCases(ctrl),
			}
			tt.setup(&f)

			got, err := f.useCase.ProccesPerson(tt.args.ctx, tt.args.person)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.WantResp, got)
			}
		})
	}
}

func TestUsecase_ProccesOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		usecase *mocks.MockUseCases
	}

	type args struct {
		ctx    context.Context
		orders []dto.OrderDto
	}

	tests := []struct {
		name     string
		setup    func(f *fields)
		args     args
		wantErr  bool
		wantResp []string
	}{
		{
			name: "succes",
			args: args{
				ctx: context.TODO(),
				orders: []dto.OrderDto{
					{OrderID: "001", Customer: "Felipe", TotalAmount: 100.0, Status: "Paid", Date: "2024-01-01"},
				},
			},
			setup: func(f *fields) {
				f.usecase.EXPECT().
					ProccesOrder(gomock.Any(), gomock.Any()).
					Return([]string{"order001"}, nil)
			},
			wantResp: []string{"order001"},
		},
		{
			name: "error from usecase",
			args: args{
				ctx: context.TODO(),
				orders: []dto.OrderDto{
					{OrderID: "fail"},
				},
			},
			setup: func(f *fields) {
				f.usecase.EXPECT().
					ProccesOrder(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("procces error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				usecase: mocks.NewMockUseCases(ctrl),
			}
			tt.setup(&f)

			got, err := f.usecase.ProccesOrder(tt.args.ctx, tt.args.orders)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantResp, got)
			}
		})
	}

}
