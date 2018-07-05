package utils

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// HTTPClient holds parameters for performing HTTP requests in the library.
type HTTPClient struct {
	// Timeout is the time limit for requests.
	Timeout time.Duration
}

// Get fetches an HTTP url and returns a goquery.Document.
// It will also set the appropriate headers to make sure the pages are returned in English.
func (c HTTPClient) Get(url string) (*goquery.Document, error) {

	client := http.Client{
		Timeout: c.Timeout,
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept-Language", "en-US,en;q=0.8,gd;q=0.6")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Set("X-FORWARDED-FOR", "165.234.102.177")

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http: %v -> %v", url, res.StatusCode)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return doc, err
	}

	return doc, err
}

// HTTPGet fetches an HTTP url and returns a goquery.Document.
// It will also set the appropriate headers to make sure the pages are returned in English.
func HTTPGet(url string) (*goquery.Document, error) {
	var client HTTPClient
	return client.Get(url)
}

// GetFileDocument opens an HTML file and returns a GoQuery document from that file.
func GetFileDocument(filePath string) (*goquery.Document, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(file)

	if err != nil {
		return nil, err
	}

	return doc, nil
}
