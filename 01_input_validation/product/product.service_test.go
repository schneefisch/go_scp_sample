package product

import (
	"encoding/json"
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
		name           string
		args           args
		wantStatusCode int
	}{
		{
			name:           "Get all products",
			args:           args{request: httptest.NewRequest(http.MethodGet, "/products", nil)},
			wantStatusCode: http.StatusOK,
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
			products := make([]Product, 0)
			err = json.Unmarshal(responseBytes, &products)
			if err != nil {
				t.Errorf("Response body cannot be parsed: %v", err)
			}

			if len(products) == 0 {
				t.Errorf("Products response is empty")
			}
		})
	}
}
