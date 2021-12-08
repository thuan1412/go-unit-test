package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_hello_2(t *testing.T) {
	t.Run("test 2", func(t *testing.T) {
		if got := hello2(""); got != "asd" {
			t.Errorf("hello() = %v, want %v", got, "asd")
		}
	})
}

func Test_hello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{

			"1",
			args{
				"",
			},
			"Hello, what's your name?",
		}, {
			"2",
			args{
				"thuan",
			},
			"Hello thuan",
		},
		{
			"2",
			args{
				"thuan",
			},
			"Hello thuan",
		},
		{
			"2",
			args{
				"thuan",
			},
			"Hello thuan",
		},
		{
			"2",
			args{
				"thuan",
			},
			"Hello thuan",
		},
		{
			"2",
			args{
				"thuan",
			},
			"Hello thuan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello(tt.args.name); got != tt.want {
				t.Errorf("hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_welcome(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", welcome)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestNew(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"normal string",
			args{
				"hello",
			},
			"hello",
		}, {
			"empty string",
			args{
				"",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.a); got != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
