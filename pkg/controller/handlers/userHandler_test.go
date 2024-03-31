package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/gin-gonic/gin"
)

func mockRegisterUserForTesting(userData *domain.User) error {
	return nil
}

func TestRegisterUserHandler(t *testing.T) {
	Register_user = mockRegisterUserForTesting

	jsonPayload := `{"email": "test@example.com", "password": "password123"}`
	req, err := http.NewRequest("POST", "/register", strings.NewReader(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	RegisterUser(ctx)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"Error":"nil","Message":"Redirect: http://localhost:8000/user/register/validate","Success":"true"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}
