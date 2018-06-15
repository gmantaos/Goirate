package main

import (
	"encoding/json"
	"git.gmantaos.com/haath/Goirate/pkg/torrents"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strconv"
	"testing"
)

func TestSearchExecute(t *testing.T) {

	var cmd SearchCommand
	cmd.Args.Query = "avengers"
	Options.JSON = true

	output := CaptureCommand(func() { cmd.Execute([]string{}) })

	var mirrors []torrents.Mirror
	json.Unmarshal([]byte(output), &mirrors)

	cmd.MagnetLink = true

	err := cmd.Execute([]string{})

	if err == nil {
		t.Errorf("Expected error")
	}

	Options.JSON = false

	cmd.SourceURL = "http://localhost"

	err = cmd.Execute([]string{})

	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestGetTorrentsTable(t *testing.T) {
	var table = []struct {
		in  []torrents.Torrent
		out string
	}{
		{[]torrents.Torrent{}, " Title  Size  Seeds/Peers \n--------------------------\n"},
	}

	for _, tt := range table {
		s := getTorrentsTable(tt.in)
		if s != tt.out {
			t.Errorf("\ngot : %v\nwant: %v", s, tt.out)
		}
	}
}

func TestFilterTorrentList(t *testing.T) {

	file, err := os.Open("../../samples/piratebay_search.html")

	if err != nil {
		t.Error(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Error(err)
	}

	scraper := torrents.NewScraper("localhost")

	torrentList := scraper.ParseSearchPage(doc)

	var table = []struct {
		in  func() SearchCommand
		out int
	}{
		{func() SearchCommand { return SearchCommand{} }, 30},
		{func() SearchCommand {
			cmd := SearchCommand{}
			cmd.VerifiedUploader = true
			return cmd
		}, 21},
		{func() SearchCommand {
			cmd := SearchCommand{}
			cmd.Count = 1
			return cmd
		}, 1},
	}

	for _, tt := range table {
		t.Run(strconv.Itoa(tt.out), func(t *testing.T) {
			filt := tt.in()
			s := filt.filterTorrentList(torrentList)
			if len(s) != tt.out {
				t.Errorf("\ngot: %v\nwant: %v", len(s), tt.out)
			}
		})
	}
}
