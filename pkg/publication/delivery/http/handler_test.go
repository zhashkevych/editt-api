package http

import (
	"bytes"
	"edittapi/pkg/models"
	"edittapi/pkg/publication/upload"
	"edittapi/pkg/publication/usecase"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	//"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type publishTest struct {
	input      *publishInput
	StatusCode int
}

func TestPublish(t *testing.T) {
	r := gin.Default()
	group := r.Group("/api")

	uc := new(usecase.PublicationUseCaseMock)
	uploader := new(upload.UploaderMock)

	RegisterHTTPEndpoints(group, uc, uploader)

	tests := []*publishTest{
		{ // Ok
			input: &publishInput{
				Author:    "Maksim Zhashkevych",
				Tags:      []string{"мінімалізм", "лайфстайл"},
				Title:     "Title",
				Body:      "крута стаття",
				ImageLink: "https://link-to.image",
			},
			StatusCode: 200,
		},
		{ // Author name too long
			input: &publishInput{
				Author:    "Maksim Zhashkevych asfaskfkasjfkassaf sf",
				Tags:      []string{"мінімалізм", "лайфстайл"},
				Body:      "крута стаття",
				Title:     "Title",
				ImageLink: "https://link-to.image",
			},
			StatusCode: 400,
		},
		{ // Author name too short
			input: &publishInput{
				Author:    "Ma",
				Tags:      []string{"мінімалізм", "лайфстайл"},
				Body:      "крута стаття",
				Title:     "Title",
				ImageLink: "https://link-to.image",
			},
			StatusCode: 400,
		},
		{ // No tags
			input: &publishInput{
				Author:    "Maksim",
				Tags:      []string{},
				Body:      "крута стаття",
				Title:     "Title",
				ImageLink: "https://link-to.image",
			},
			StatusCode: 400,
		},
		{ // Too much tags
			input: &publishInput{
				Author:    "Maksim",
				Tags:      []string{"1", "2", "3", "4"},
				Body:      "крута стаття",
				Title:     "Title",
				ImageLink: "https://link-to.image",
			},
			StatusCode: 400,
		},
	}

	for _, test := range tests {
		body, err := json.Marshal(test.input)
		assert.NoError(t, err)

		p := toPublicationModel(test.input)

		uc.On("Publish", p).Return(nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/publications", bytes.NewBuffer(body))
		r.ServeHTTP(w, req)

		assert.Equal(t, test.StatusCode, w.Code)
	}
}

func TestGetPublications(t *testing.T) {
	r := gin.Default()
	group := r.Group("/api")

	uc := new(usecase.PublicationUseCaseMock)
	uploader := new(upload.UploaderMock)

	RegisterHTTPEndpoints(group, uc, uploader)

	// Type = popular
	results := []*models.Publication{
		{
			Author:      "Maksim",
			Tags:        []string{"1", "2", "3", "4"},
			Body:        "крута стаття",
			Title:       "Title",
			ImageLink:   "https://link-to.image",
			Views:       25,
			Reactions:   3,
			ReadingTime: 4,
		},
		{
			Author:      "Roman",
			Tags:        []string{"1", "2", "3", "4"},
			Body:        "крута стаття",
			Title:       "Title",
			ImageLink:   "https://link-to.image",
			Views:       56,
			Reactions:   8,
			ReadingTime: 6,
		},
	}
	uc.On("GetPopularPublications", int64(2)).Return(results, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/publications?limit=2&type=popular", nil)
	r.ServeHTTP(w, req)

	expectedResponse, _ := json.Marshal(&getPublicationsResponse{
		results,
	})

	expectedResponse = append(expectedResponse, []byte("\n")...)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedResponse), w.Body.String())

	// Type = latest
	results = []*models.Publication{
		{
			Author:      "Roman",
			Tags:        []string{"1", "2", "3", "4"},
			Body:        "крута стаття",
			Title:       "Title",
			ImageLink:   "https://link-to.image",
			Views:       56,
			Reactions:   8,
			ReadingTime: 6,
		},
		{
			Author:      "Maksim",
			Tags:        []string{"1", "2", "3", "4"},
			Body:        "крута стаття",
			Title:       "Title",
			ImageLink:   "https://link-to.image",
			Views:       25,
			Reactions:   3,
			ReadingTime: 4,
		},
		{
			Author:      "Oleg",
			Tags:        []string{"1", "2", "3", "4"},
			Body:        "крута стаття",
			Title:       "Title",
			ImageLink:   "https://link-to.image",
			Views:       25,
			Reactions:   3,
			ReadingTime: 4,
		},
	}
	uc.On("GetLatestPublications", int64(3)).Return(results, nil)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/publications?limit=3&type=latest", nil)
	r.ServeHTTP(w, req)

	expectedResponse, _ = json.Marshal(&getPublicationsResponse{
		results,
	})

	expectedResponse = append(expectedResponse, []byte("\n")...)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedResponse), w.Body.String())

	// Type = unknown
	results = []*models.Publication{
		{
			Author:      "Roman",
			Tags:        []string{"1", "2", "3", "4"},
			Body:        "крута стаття",
			Title:       "Title",
			ImageLink:   "https://link-to.image",
			Views:       56,
			Reactions:   8,
			ReadingTime: 6,
		},
	}
	uc.On("GetPopularPublications", int64(1)).Return(results, nil)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/publications?limit=1&type=jsafjasfj", nil)
	r.ServeHTTP(w, req)

	expectedResponse, _ = json.Marshal(&getPublicationsResponse{
		results,
	})

	expectedResponse = append(expectedResponse, []byte("\n")...)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedResponse), w.Body.String())
}

func TestGetById(t *testing.T) {
	r := gin.Default()
	group := r.Group("/api")

	uc := new(usecase.PublicationUseCaseMock)
	uploader := new(upload.UploaderMock)

	RegisterHTTPEndpoints(group, uc, uploader)

	p := &models.Publication{
		ID:          "5e6a03309ea43ef775bd247e",
		Author:      "Roman",
		Tags:        []string{"1", "2", "3", "4"},
		Body:        "крута стаття",
		Title:       "Title",
		ImageLink:   "https://link-to.image",
		Views:       56,
		Reactions:   8,
		ReadingTime: 6,
	}

	uc.On("GetById", "5e6a03309ea43ef775bd247e").Return(p, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/publications/5e6a03309ea43ef775bd247e", nil)
	r.ServeHTTP(w, req)

	expectedResponse, _ := json.Marshal(p)
	
	expectedResponse = append(expectedResponse, []byte("\n")...)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedResponse), w.Body.String())
}

func TestUploadInvalidFile(t *testing.T) {
	r := gin.Default()
	group := r.Group("/api")

	uc := new(usecase.PublicationUseCaseMock)
	uploader := new(upload.UploaderMock)

	RegisterHTTPEndpoints(group, uc, uploader)

	file, err := ioutil.TempFile("", "testUpload.*.txt")
	assert.NoError(t, err)

	fileContents, err := ioutil.ReadAll(file)
	assert.NoError(t, err)

	fi, err := file.Stat()
	assert.NoError(t, err)

	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fi.Name())
	assert.NoError(t, err)

	_, err = part.Write(fileContents)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/upload", body)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
