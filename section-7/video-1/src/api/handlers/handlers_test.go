package handlers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/utils/appErrors"
	"github.com/gorilla/mux"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/usecases"
)

func Test_GetUserByUserNameHandler_NoUsername(t *testing.T) {
	h := &Handlers{
		GetUserUsecase:    &MockGetUserUsecase{},
		UpdateUserUsecase: &MockUpdateUserUsecase{},
	}

	r := mux.NewRouter()
	r.HandleFunc("/user/{username}", http.HandlerFunc(h.GetUserByUsername))

	ts := httptest.NewServer(r)
	defer ts.Close()

	url := ts.URL + "/user/ "
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("Test failed. Unexpected error: %s", err.Error())
	}

	if resp.StatusCode != 400 {
		t.Errorf("Test failed. Expected 400 Status. Got: %s", resp.Status)
	}
}

func Test_GetUserByUserNameHandler_EverythingRight(t *testing.T) {
	h := &Handlers{
		GetUserUsecase:    &MockGetUserUsecase{},
		UpdateUserUsecase: &MockUpdateUserUsecase{},
	}

	r := mux.NewRouter()
	r.HandleFunc("/user/{username}", http.HandlerFunc(h.GetUserByUsername))

	ts := httptest.NewServer(r)
	defer ts.Close()

	url := ts.URL + "/user/pepe"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("Test failed. Unexpected error: %s", err.Error())
	}

	if resp.StatusCode != 200 {
		t.Errorf("Test failed. Expected 200 Status. Got: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Test failed. Unexpected error: %s", err.Error())
	}
	if string(body) != `{"id":500,"username":"pepe","first_name":"Josua","last_name":"Smith","email":"pepe@example.com","birth_date":"1992-02-25T00:00:00Z","added":"2018-07-10T00:00:00Z","account":800,"password":"HASHED","AccountType":0}` {
		t.Errorf("Test failed. Wrong body.")
	}
}

func Test_GetUserByUserNameHandler_NotFoundOnUsecase(t *testing.T) {
	h := &Handlers{
		GetUserUsecase:    &MockGetUserUsecase{},
		UpdateUserUsecase: &MockUpdateUserUsecase{},
	}

	r := mux.NewRouter()
	r.HandleFunc("/user/{username}", http.HandlerFunc(h.GetUserByUsername))

	ts := httptest.NewServer(r)
	defer ts.Close()

	url := ts.URL + "/user/test2"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("Test failed. Unexpected error: %s", err.Error())
	}

	if resp.StatusCode != 404 {
		t.Errorf("Test failed. Expected 404 Status. Got: %s", resp.Status)
	}
}

func Test_GetUserByUserNameHandler_OtherErrorOnUsecase(t *testing.T) {
	h := &Handlers{
		GetUserUsecase:    &MockGetUserUsecase{},
		UpdateUserUsecase: &MockUpdateUserUsecase{},
	}

	r := mux.NewRouter()
	r.HandleFunc("/user/{username}", http.HandlerFunc(h.GetUserByUsername))

	ts := httptest.NewServer(r)
	defer ts.Close()

	url := ts.URL + "/user/test3"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("Test failed. Unexpected error: %s", err.Error())
	}

	if resp.StatusCode != 500 {
		t.Errorf("Test failed. Expected 500 Status. Got: %s", resp.Status)
	}
}

type MockGetUserUsecase struct{}

func (m *MockGetUserUsecase) GetUser(username string) (*usecases.UserDTO, error) {
	udto := &usecases.UserDTO{}
	udto.Account = 800
	udto.AccountType = udto.GetAccountType()
	udto.Added = time.Date(2018, 07, 10, 0, 0, 0, 0, time.UTC)
	udto.BirthDate = time.Date(1992, 02, 25, 0, 0, 0, 0, time.UTC)
	udto.Username = "pepe"
	udto.Email = "pepe@example.com"
	udto.FirstName = "Josua"
	udto.LastName = "Smith"
	udto.ID = 500
	udto.Password = "HASHED"

	if username == "pepe" {
		return udto, nil
	} else if username == "test2" {
		return nil, appErrors.ErrorNotFound
	} else if username == "test3" {
		return nil, errors.New("Unknown error.")
	}
	return nil, nil
}

type MockUpdateUserUsecase struct{}

func (m *MockUpdateUserUsecase) UpdateUser(user *entities.User) error {
	//We are not testing this usecase
	return nil
}
