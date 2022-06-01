package videoControllers_test

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

func videoMock() models.Video {
	utils.RegisterValidators()
	video := models.Video{
		Model:       gorm.Model{},
		CategoriaID: 1,
		Titulo:      "Titulo de Teste",
		Descricao:   "Descrição de teste",
		URL:         "https://www.teste.com/teste.mp4",
	}

	return video
}

func deleteVideo(db *gorm.DB, id uint) {
	db.Delete(&models.Video{}, id)
}

func TestCreateVideo(t *testing.T) {
	r := routes.HandleRequests()

	videoMock := videoMock()
	videoJSONMock, _ := json.Marshal(videoMock)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/videos/", bytes.NewBuffer(videoJSONMock))
	r.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Errorf("Status code expected: 201, got: %d", w.Code)
	}

	var video models.Video
	json.Unmarshal(w.Body.Bytes(), &video)

	db, _ := database.Connect()
	defer deleteVideo(db, video.ID)

	if video.Titulo != "Titulo de Teste" {
		t.Errorf("Titulo expected: Titulo de Teste, got: %s", video.Titulo)
	}

	if video.Descricao != "Descrição de teste" {
		t.Errorf("Descrição expected: Descrição de teste, got: %s", video.Descricao)
	}

	if video.URL != "https://www.teste.com/teste.mp4" {
		t.Errorf("URL expected: https://www.teste.com/teste.mp4, got: %s", video.URL)
	}

	if video.CategoriaID != 1 {
		t.Errorf("CategoriaID expected: 1, got: %d", video.CategoriaID)
	}
}
