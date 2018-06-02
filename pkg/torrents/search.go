package torrents

import (
	"errors"
)

// SearchTorrentList
func SearchTorrentList(torrents []Torrent, filters SearchFilters) (*Torrent, error) {

	maxSize, err := filters.MaxSizeKB()

	if err != nil {
		return nil, err
	}

	minSize, err := filters.MinSizeKB()

	if err != nil {
		return nil, err
	}

	for _, t := range torrents {

		if filters.VerifiedUploader && !t.VerifiedUploader {
			continue
		}

		if t.Size > maxSize ||
			t.Size < minSize {
			continue
		}

		return &t, nil
	}

	return nil, errors.New("No torrent found with the specified filters")
}