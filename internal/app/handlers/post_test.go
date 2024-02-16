package handlers

import (
	"github.com/Tarasusrus/url-shortener/internal/app/configs"
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlePost(t *testing.T) {
	config, _ := configs.NewFlagConfig()
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
			name: "Test case 1 - Check Content - Type",
			args: args{
				w:     httptest.NewRecorder(),
				r:     httptest.NewRequest("POST", "/ajdbkl", nil),
				store: stores.NewStore(),
			},
			want: want{
				body:       "",
				statusCode: 400,
			},
		},
	}
	for _, tt := range tests {

		req := tt.args.r
		req.Header.Set("Content-Type", "text/plain")

		t.Run(tt.name, func(t *testing.T) {
			HandlePost(tt.args.w, tt.args.r, tt.args.store, config)
			assert.Equal(t, tt.want.statusCode, tt.args.w.Code)
		})
	}
}
