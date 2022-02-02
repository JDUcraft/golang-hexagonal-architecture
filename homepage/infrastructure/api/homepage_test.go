package api

import (
	"io"
	"main/homepage/infrastructure/repository"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHomePage(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	old := repository.Timer
	repository.Timer = func() time.Time {
		return time.Date(2022, 01, 21, 15, 54, 28, 0, time.Local)
	}
	defer func() { repository.Timer = old }()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "hello", args: args{w: w, r: req}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HomePage(tt.args.w, tt.args.r)
			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)
			if resp.StatusCode != 200 {
				t.Errorf("Expected: 200 \t Received: %d", resp.StatusCode)
			}

			if string(body) != "The time is: 15:54:28" {
				t.Errorf("Expected: The time is: 15:54:28 \t Received: %s", string(body))
			}
		})
	}
}
