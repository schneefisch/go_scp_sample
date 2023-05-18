package product

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
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

			handleProducts(response, test.args.request)

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
			} else if hasResponseBody {
				products := make([]Product, 0)
				err = json.Unmarshal(responseBytes, &products)
				if err != nil {
					t.Errorf("Response body cannot be parsed: %v", err)
				}

				if len(products) == 0 {
					t.Errorf("Products response is empty")
				}
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

	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			handleProduct(response, test.args.request)

			if status := response.Code; status != test.wantStatusCode {
				t.Errorf("Unexpected return status-code; got %v, want %v", status, test.wantStatusCode)
			}

			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				t.Errorf("unexpected error reading body; %v", err)
			}
			var product Product
			err = json.Unmarshal(bodyBytes, &product)
			if err != nil {
				t.Errorf("Unexpected error parsing product response; %v", err)
			}
			if test.wantProduct != nil && !reflect.DeepEqual(product, *test.wantProduct) {
				t.Errorf("Unexpected product response;\ngot: %v\nwant: %v", product, *test.wantProduct)
			}

		})
	}
}
