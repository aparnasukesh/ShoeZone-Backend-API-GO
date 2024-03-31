package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/gin-gonic/gin"
)

// Register User
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

// User Login
func mockUserLoginForTesting(userData *domain.User) (error, *domain.User) {
	return nil, &domain.User{}
}

func mockGenerateJWTForTesting(userData domain.User) (string, error) {
	return "", nil
}

func TestUserLoginHandler(t *testing.T) {
	User_login = mockUserLoginForTesting
	Generate_JWt = mockGenerateJWTForTesting

	jsonPayload := `{"email": "test@example.com", "password": "password123"}`
	req, err := http.NewRequest("POST", "/register", strings.NewReader(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	Login(ctx)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"Error":null,"Message":"Succesfully Login","Success":true,"Token":""}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v\n want %v\n", w.Body.String(), expected)
	}
}

// User -Products

func mockGetProductsForTest(limit, offset int) ([]domain.Product, error) {
	return nil, nil
}
func TestGetProductsHandler(t *testing.T) {
	Get_Products = mockGetProductsForTest
	req, err := http.NewRequest("GET", "/products?page=0&limit=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	GetProducts(ctx)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"Error":null,"Message":"Product Details","Products":null,"Success":true}`

	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v\n want %v\n", w.Body.String(), expected)
	}
}

// User- Porduct By Id
func mockGetProductByIDForTest(id int) (*domain.Product, error) {
	return nil, nil
}

func TestGetProductByIDHandler(t *testing.T) {
	Get_ProductByID = mockGetProductByIDForTest
	req, err := http.NewRequest("GET", "/product/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	GetProductByID(ctx)

	if status := w.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	expected := `{"Error":"strconv.Atoi: parsing \"\": invalid syntax","Message":"No product found","Success":false}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v\n want %v\n", w.Body.String(), expected)
	}
}

// User - Products By Brand ID
func mockGetProductByBrandIDForTest(limit, offset int) ([]domain.Product, error) {
	return nil, nil
}
func TestGetProductByBrandIDHandler(t *testing.T) {
	Get_Products = mockGetProductByBrandIDForTest
	req, err := http.NewRequest("GET", "/products?page=0&limit=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	GetProducts(ctx)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"Error":null,"Message":"Product Details","Products":null,"Success":true}`

	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v\n want %v\n", w.Body.String(), expected)
	}
}
