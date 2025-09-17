package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestSuccess(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"ApiKey test-key"},
	}

	got, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("error while get api key: %s", err)
	}

	want := "test-key"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestNoPrefix(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"test-key"},
	}

	got, err := GetAPIKey(header)
	if err == nil {
		t.Fatalf("expect error while get api key: %s", header.Get("Authorization"))
	}

	want := "malformed authorization header"
	if !reflect.DeepEqual(want, err.Error()) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestEmptyHeader(t *testing.T) {
	header := http.Header{}

	got, err := GetAPIKey(header)
	if err == nil {
		t.Fatalf("expect error without header: %s", "Authorization")
	}

	want := "no authorization header included"
	if !reflect.DeepEqual(want, err.Error()) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
