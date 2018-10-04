package utils

import (
	"reflect"
	"testing"
)

func TestHTTPGet(t *testing.T) {
	_, err := HTTPGet("https://www.google.com/")

	if err != nil {
		t.Errorf("HTTPGet status code: %s", err.Error())
	}
}

func TestHTTPGetError(t *testing.T) {
	_, err := HTTPGet("1.2.3.4")

	if err == nil {
		t.Errorf("HTTPGet status code: %s", err.Error())
	}
}

func TestHTTPGetStatus(t *testing.T) {

	res, err := HTTPGet("https://www.reddit.com/asdf")

	if err == nil {
		t.Errorf("Expected 404, got: %v", res)
	}
}

func TestHTTPPost(t *testing.T) {
	obj := []struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
		UserID int    `json:"userId"`
		ID     int    `json:"id"`
	}{
		{
			"some title",
			"super special body",
			1234567,
			1234567,
		},
		{},
	}

	err := HTTPPost("https://jsonplaceholder.typicode.com/posts", obj[0], &obj[1])

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(obj[0], obj[1]) {
		t.Fatalf("got %v, want %v", obj[1], obj[0])
	}

	HTTPGetJSON("https://jsonplaceholder.typicode.com/posts/102", &obj[1])
}

func TestGetFileDocument(t *testing.T) {

	expected := "Cast Away (2000) - IMDbTryIMDbProFree"

	doc, err := GetFileDocument("../../test_samples/imdb.html")

	if err != nil {
		t.Error(err)
	}

	title := doc.Find("title").Text()

	if title != expected {
		t.Errorf("got %v want %v", title, expected)
	}
}
