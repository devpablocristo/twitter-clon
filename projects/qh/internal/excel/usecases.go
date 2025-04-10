package excel

import (
	"context"
	"fmt"
	"mime/multipart"
	"strconv"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/handler/dto"
	"github.com/xuri/excelize/v2"
)

type useCases struct {
	repo Person_2Repository
}

func NewUseCases(r Person_2Repository) useCases {
	return useCases{repo: r}
}

func (uc *useCases) ProccesExcel(ctx context.Context, file multipart.File) error {
	// abre el archivo
	f, err := excelize.OpenReader(file)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo: %w", err)
	}

	//lee la primera hoja
	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("error al leer la hoja: %w", err)
	}

	// valido los encabezados
	expectedHeaders := []string{"FirstName", "LastName", "Age", "Phone"}
	if len(rows) == 0 || len(rows[0]) < len(expectedHeaders) {
		return fmt.Errorf("El archivo no contiene los encabezdos esperados")
	}

	//itero filas
	for i, row := range rows[1:] {
		if len(row) < 4 {
			return fmt.Errorf("fila %d incompleta", i+2) // ---> valido que la fila no este vacia
		}

		age, err := strconv.Atoi(row[2]) // ---> convierto el campo de edad que viene string a int
		if err != nil {
			return fmt.Errorf("edad invalida en fila: %d: %w", i+2, err)
		}

		// convertir a Dto
		excelDto := dto.ExcelPerson{
			FirstName: row[0],
			LastName:  row[1],
			Age:       age,
			Phone:     row[3],
		}

		// convertir a Domain
		person := excelDto.ToDomain()

		if err := uc.repo.CreatePerson(ctx, person); err != nil { // --> lo almaceno en el repositorio
			return fmt.Errorf("error al guardar fila: %d: %w", i+2, err)
		}
	}
	return nil
}
