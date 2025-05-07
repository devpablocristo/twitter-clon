package integrationtest_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/devpablocristo/monorepo/projects/qh/internal/project"
	"github.com/devpablocristo/monorepo/projects/qh/internal/project/adapter"
	"github.com/devpablocristo/monorepo/projects/qh/internal/project/usecases/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type memoryRepo struct {
	saved []domain.FullProject
}

func (m *memoryRepo) SaveFullProject(ctx context.Context, fps []domain.FullProject) ([]string, error) {
	m.saved = append(m.saved, fps...)
	ids := make([]string, len(fps))
	for i := range fps {
		ids[i] = fmt.Sprintf("mem-%d", i+1)
	}

	return ids, nil
}

func TestIntegration_UploadFullProjectExcel(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//repositorio de memoria
	repo := &memoryRepo{}

	//contruyendo dependencias reales
	usecase := project.NewUseCases(repo)
	adapter := adapter.NewExcelAdapter()
	server := newFakeServer("v1")

	handler := project.NewHandler(server, usecase, nil, adapter)
	handler.Routes()

	// preparando peticiones con el archivo test
	testFile := filepath.Join("testdata", "full_projects.xlsx")
	// verifico que el archivo exista
	require.FileExists(t, testFile, "Archivo Excel no encontrado")

	req := newMultipartRequest(t, "file", "full_projects.xlsx", testFile)
	w := httptest.NewRecorder()
	server.GetRouter().ServeHTTP(w, req)

	// verifico las respuestas http
	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), `"projects_count":`)

	// verifico que se guardaron los pryectos
	require.Len(t, repo.saved, 3)
	require.Equal(t, "ACME Corp", repo.saved[0].Client.Name)
	require.Equal(t, "BioFarms", repo.saved[1].Client.Name)
	require.Equal(t, "AgroAndes", repo.saved[2].Client.Name)

	// Valido el JSON devuelto
	type response struct {
		Message       string   `json:"message"`
		ProjectIDs    []string `json:"project_ids"`
		ProjectsCount int      `json:"projects_count"`
	}

	var res response
	err := json.Unmarshal(w.Body.Bytes(), &res)
	require.NoError(t, err)

	require.Equal(t, 3, res.ProjectsCount)
	require.Equal(t, []string{"mem-1", "mem-2", "mem-3"}, res.ProjectIDs)
	require.Equal(t, "projects processed successfully", res.Message)
}

// newMultipartRequest crea una peticion HTTP simulada con un archivo adjunto
func newMultipartRequest(t *testing.T, field, filename, path string) *http.Request {
	t.Helper()
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	fileWriter, err := writer.CreateFormFile(field, filename)
	require.NoError(t, err)

	file, err := os.Open(path)
	require.NoError(t, err)
	defer file.Close()

	_, err = io.Copy(fileWriter, file)
	require.NoError(t, err)

	require.NoError(t, writer.Close())

	req := httptest.NewRequest(http.MethodPost, "/api/v1/upload/fullProject", buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req
}
