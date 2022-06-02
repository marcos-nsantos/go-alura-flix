package categoriaControllers_test

import (
	"bytes"
	"encoding/json"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/routes"
	"github.com/marcos-nsantos/alura-flix/utils"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	database.Migrate()
}

func categoriaMock() []models.Categoria {
	utils.RegisterValidators()
	categorias := []models.Categoria{
		{
			Titulo: "Livre",
			Cor:    "#993399",
		},
		{
			Titulo: "Programação",
			Cor:    "#008000",
		},
		{
			Titulo: "Front-end",
			Cor:    "#0000ff",
		},
		{
			Titulo: "Data Science",
			Cor:    "#32cd32",
		},
		{
			Titulo: "DevOps",
			Cor:    "#ff0000",
		},
		{
			Titulo: "UX",
			Cor:    "#ffa500",
		},
		{
			Titulo: "Mobile",
			Cor:    "#ffd700",
		},
	}
	return categorias
}

func deleteCategoria(db *gorm.DB, id uint) {
	db.Delete(&models.Categoria{}, id)
}

func getLastInsertedID(db *gorm.DB) uint {
	var categoria models.Categoria
	db.Last(&categoria)
	return categoria.ID
}

func TestCreateCategoria(t *testing.T) {
	r := routes.HandleRequests()

	categorias := categoriaMock()
	categoriaJSONMock, _ := json.Marshal(categorias[0])

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/categorias/", bytes.NewBuffer(categoriaJSONMock))
	r.ServeHTTP(w, req)

	defer func() {
		db, _ := database.Connect()
		defer deleteCategoria(db, getLastInsertedID(db))
	}()

	if w.Code != http.StatusCreated {
		t.Errorf("Status code should be 201, was: %d", w.Code)
	}

	var categoria models.Categoria
	json.NewDecoder(w.Body).Decode(&categoria)

	if categoria.Titulo != categorias[0].Titulo {
		t.Errorf("Titulo should be %s, was: %s", categorias[0].Titulo, categoria.Titulo)
	}

	if categoria.Cor != categorias[0].Cor {
		t.Errorf("Cor should be %s, was: %s", categorias[0].Cor, categoria.Cor)
	}
}

func TestShowAllCategorias(t *testing.T) {
	r := routes.HandleRequests()

	categoriasMock := categoriaMock()

	for _, categoria := range categoriasMock {
		categoriaJSONMock, _ := json.Marshal(categoria)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/categorias/", bytes.NewBuffer(categoriaJSONMock))
		r.ServeHTTP(w, req)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/categorias/", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be 200, was: %d", w.Code)
	}

	var categorias []models.Categoria
	json.NewDecoder(w.Body).Decode(&categorias)

	defer func() {
		db, _ := database.Connect()
		for _, categoria := range categorias {
			deleteCategoria(db, categoria.ID)
		}
	}()

	if len(categorias) != len(categoriasMock) {
		t.Errorf("Should be %d, was: %d", len(categoriasMock), len(categorias))
	}

	for index, categoria := range categorias {
		if categoria.Titulo != categoriasMock[index].Titulo {
			t.Errorf("Titulo should be %s, was: %s", categoriasMock[index].Titulo, categoria.Titulo)
		}
		if categoria.Cor != categoriasMock[index].Cor {
			t.Errorf("Cor should be %s, was: %s", categoriasMock[index].Cor, categoria.Cor)
		}
	}
}