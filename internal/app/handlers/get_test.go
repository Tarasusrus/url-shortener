package handlers

import (
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGet(t *testing.T) {
	type args struct {
		w     *httptest.ResponseRecorder
		r     *http.Request
		store *stores.Store
	}

	type want struct {
		body        string
		contentType string
		ID          string
		statusCode  int
	}

	tests := []struct {
		name string
		want want
		args args
	}{
		{
			name: "Test case 1 - URL not found for ID",
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("GET", "/ajdbkl", nil),
				store: stores.NewStore(),
			},
			want: want{
				body:       "",
				statusCode: 400,
			},
		},

		{
			name: "Test case 2 - Empty ID received",
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("GET", "/", nil),
				store: stores.NewStore(),
			},
			want: want{
				body:       "",
				statusCode: 400,
			},
		},

		{
			name: "Test case 3 - URL found for ID",
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("GET", "/test-id", nil),
				store: &stores.Store{Urls: map[string]string{"test-id": "http://example.com"}},
			},
			want: want{
				body:       "",
				statusCode: 307,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleGet(tt.args.w, tt.args.r, tt.args.store)
			assert.Equal(t, tt.want.statusCode, tt.args.w.Code)
		})
	}
}
