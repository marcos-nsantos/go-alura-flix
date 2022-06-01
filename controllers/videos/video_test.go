package videoControllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func videoMock() []models.Video {
	utils.RegisterValidators()
	video := []models.Video{
		{
			Titulo:    "Titulo de Teste",
			Descricao: "Descrição de teste",
			URL:       "https://www.teste.com/video.mp4",
		},
		{
			CategoriaID: 1,
			Titulo:      "Titulo de Teste 2",
			Descricao:   "Descrição de video 2",
			URL:         "https://www.teste.com/video2.mp4",
		},
		{
			Titulo:    "Titulo de Teste 3",
			Descricao: "Descrição de teste 3",
			URL:       "https://www.teste.com/video3.mp4",
		},
	}

	return video
}

func deleteVideo(db *gorm.DB, id uint) {
	db.Delete(&models.Video{}, id)
}

func getLastInsertedID(db *gorm.DB) uint {
	var video models.Video
	db.Last(&video)
	return video.ID
}

func TestCreateVideo(t *testing.T) {
	r := routes.HandleRequests()

	videoMock := videoMock()
	videoJSONMock, _ := json.Marshal(videoMock[0])

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/videos/", bytes.NewBuffer(videoJSONMock))
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

	if video.URL != "https://www.teste.com/video.mp4" {
		t.Errorf("URL expected: https://www.teste.com/teste.mp4, got: %s", video.URL)
	}

	if video.CategoriaID != 1 {
		t.Errorf("CategoriaID expected: 1, got: %d", video.CategoriaID)
	}
}

func TestShowAllVideos(t *testing.T) {
	r := routes.HandleRequests()

	videoMock := videoMock()

	for _, video := range videoMock {
		videoJSONMock, _ := json.Marshal(video)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/videos/", bytes.NewBuffer(videoJSONMock))
		r.ServeHTTP(w, req)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/videos/", nil)
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Status code expected: 200, got: %d", w.Code)
	}

	var videos []models.Video
	json.Unmarshal(w.Body.Bytes(), &videos)

	if len(videos) != 3 {
		t.Errorf("Expected 3 videos, got: %d", len(videos))
	}

	defer func() {
		db, _ := database.Connect()
		for _, video := range videos {
			deleteVideo(db, video.ID)
		}
	}()
}

func TestShowVideo(t *testing.T) {
	r := routes.HandleRequests()

	videoMock := videoMock()
	videoJSONMock, _ := json.Marshal(videoMock[2])

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/videos/", bytes.NewBuffer(videoJSONMock))
	r.ServeHTTP(w, req)

	var video models.Video
	json.Unmarshal(w.Body.Bytes(), &video)

	db, _ := database.Connect()
	lastInsertedID := getLastInsertedID(db)
	defer deleteVideo(db, lastInsertedID)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/videos/%d", video.ID), nil)
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Status code expected: 200, got: %d", w.Code)
	}

	var videoReturned models.Video
	json.Unmarshal(w.Body.Bytes(), &videoReturned)

	if videoReturned.ID != video.ID {
		t.Errorf("ID expected: %d, got: %d", video.ID, videoReturned.ID)
	}

	if videoReturned.Titulo != video.Titulo {
		t.Errorf("Titulo expected: %s, got: %s", video.Titulo, videoReturned.Titulo)
	}

	if videoReturned.Descricao != video.Descricao {
		t.Errorf("Descrição expected: %s, got: %s", video.Descricao, videoReturned.Descricao)
	}

	if videoReturned.URL != video.URL {
		t.Errorf("URL expected: %s, got: %s", video.URL, videoReturned.URL)
	}

	if videoReturned.CategoriaID != video.CategoriaID {
		t.Errorf("CategoriaID expected: %d, got: %d", video.CategoriaID, videoReturned.CategoriaID)
	}
}

func TestUpdateVideo(t *testing.T) {
	r := routes.HandleRequests()

	videoMock := videoMock()
	videoJSONMock, _ := json.Marshal(videoMock[2])

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/videos/", bytes.NewBuffer(videoJSONMock))
	r.ServeHTTP(w, req)

	var video models.Video
	json.Unmarshal(w.Body.Bytes(), &video)

	db, _ := database.Connect()
	lastInsertedID := getLastInsertedID(db)
	defer deleteVideo(db, lastInsertedID)

	video.Titulo = "Titulo de Teste"
	video.Descricao = "Descrição de teste"
	video.URL = "https://www.teste.com/video.mp4"
	video.CategoriaID = 1

	videoJSONMock, _ = json.Marshal(video)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPut, fmt.Sprintf("/videos/%d", lastInsertedID), bytes.NewBuffer(videoJSONMock))
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Status code expected: 200, got: %d", w.Code)
	}

	var videoReturned models.Video
	json.Unmarshal(w.Body.Bytes(), &videoReturned)

	if videoReturned.ID != video.ID {
		t.Errorf("ID expected: %d, got: %d", video.ID, videoReturned.ID)
	}

	if videoReturned.Titulo != video.Titulo {
		t.Errorf("Titulo expected: %s, got: %s", video.Titulo, videoReturned.Titulo)
	}

	if videoReturned.Descricao != video.Descricao {
		t.Errorf("Descrição expected: %s, got: %s", video.Descricao, videoReturned.Descricao)
	}

	if videoReturned.URL != video.URL {
		t.Errorf("URL expected: %s, got: %s", video.URL, videoReturned.URL)
	}

	if videoReturned.CategoriaID != video.CategoriaID {
		t.Errorf("CategoriaID expected: %d, got: %d", video.CategoriaID, videoReturned.CategoriaID)
	}
}

func TestDeleteVideo(t *testing.T) {
	r := routes.HandleRequests()

	videoMock := videoMock()
	videoJSONMock, _ := json.Marshal(videoMock[2])

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/videos/", bytes.NewBuffer(videoJSONMock))
	r.ServeHTTP(w, req)

	var video models.Video
	json.Unmarshal(w.Body.Bytes(), &video)

	db, _ := database.Connect()
	lastInsertedID := getLastInsertedID(db)
	defer deleteVideo(db, lastInsertedID)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodDelete, fmt.Sprintf("/videos/%d", lastInsertedID), nil)
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Status code expected: 200, got: %d", w.Code)
	}
}
