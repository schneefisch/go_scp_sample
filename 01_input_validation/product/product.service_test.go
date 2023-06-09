package product

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handleProducts(t *testing.T) {
	type args struct {
		request *http.Request
	}
	tests := []struct {
		name               string
		args               args
		wantStatusCode     int
		expectResponseBody bool
	}{
		{
			name:               "Get all products",
			args:               args{request: httptest.NewRequest(http.MethodGet, "/products", nil)},
			wantStatusCode:     http.StatusOK,
			expectResponseBody: true,
		},
		{
			name:               "Try deleting all products",
			args:               args{request: httptest.NewRequest(http.MethodDelete, "/products", nil)},
			wantStatusCode:     http.StatusNotImplemented,
			expectResponseBody: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := httptest.NewRecorder()

			HandleProducts(response, test.args.request)

			if status := response.Code; status != test.wantStatusCode {
				t.Errorf("Unexpected status-code: got %v, want %v", status, test.wantStatusCode)
			}

			responseBytes, err := io.ReadAll(response.Body)
			if err != nil {
				t.Errorf("Response body cannot be read: %v", err)
			}
			hasResponseBody := len(responseBytes) > 0
			if hasResponseBody != test.expectResponseBody {
				t.Errorf("Expected some response body but got nothing")
			}
		})
	}
}

func Test_handleProduct(t *testing.T) {
	type args struct {
		request *http.Request
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantProduct    *Product
	}{
		{
			name:           "Get product 1",
			args:           args{request: httptest.NewRequest(http.MethodGet, "/products/1", nil)},
			wantStatusCode: http.StatusOK,
			wantProduct: &Product{
				ProductId:      1,
				ProductName:    "Edamer",
				Price:          "7.99",
				QuantityOnHand: 15,
			},
		},
		// todo: add a test with sql-injection
		//{
		//	name:           "Get product with SQL-Injection",
		//	args:           args{request: httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%s", url.PathEscape("1; DELETE FROM products")), nil)},
		//	wantStatusCode: http.StatusOK,
		//},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			HandleProduct(response, test.args.request)

			if status := response.Code; status != test.wantStatusCode {
				t.Errorf("Unexpected return status-code; got %v, want %v", status, test.wantStatusCode)
			}

			_, err := io.ReadAll(response.Body)
			if err != nil {
				t.Errorf("unexpected error reading body; %v", err)
			}

			// verify that there are still some products left
			list, err := getProductList()
			if err != nil {
				t.Errorf("Got some unexpected error when reading all products; %v", err)
			}
			if len(list) < 1 {
				t.Errorf("No more products in the products-database!")
			} else {
				t.Logf("Remaining Products:\n%v", list)
			}
		})
	}
}
