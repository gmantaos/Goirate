package movies

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"

	"gitlab.com/haath/goirate/pkg/torrents"
	"gitlab.com/haath/goirate/pkg/utils"
)

// MovieID holds the defining properties of an IMDb movie as they appear in search results.
type MovieID struct {
	IMDbID   string `json:"imdb_id"`
	Title    string `json:"title"`
	Year     uint   `json:"year"`
	AltTitle string `json:"alt_title"`
}

// Movie holds all the information regarding a movie on IMDb.
type Movie struct {
	MovieID
	Duration  int     `json:"duration"`
	Rating    float32 `json:"rating"`
	PosterURL string  `json:"poster_url"`
}

// GetURL formats the IMDbID of the movie object and returns the full
// URL to the movie's page on IMDb.
func (m MovieID) GetURL() (*url.URL, error) {
	formattedID, err := FormatIMDbID(m.IMDbID)

	if err != nil {
		return nil, err
	}

	urlString := fmt.Sprintf(BaseURL+"/title/tt%v/", formattedID)

	return url.Parse(urlString)
}

// FormattedDuration returns the duration of the movie in human-readable format.
func (m Movie) FormattedDuration() string {
	hours := m.Duration / 60
	minutes := m.Duration % 60

	var buf bytes.Buffer

	if hours > 0 {
		buf.WriteString(fmt.Sprintf("%vh ", hours))
	}

	if minutes > 0 {
		buf.WriteString(fmt.Sprintf("%vmin", minutes))
	}

	return strings.TrimSpace(buf.String())
}

// GetTorrent will search The Pirate Bay and return the best torrent that complies with the given filters.
func (m Movie) GetTorrent(scraper torrents.PirateBayScaper, filters torrents.SearchFilters) (*torrents.Torrent, error) {

	filteredTorrents, err := m.GetTorrents(scraper, filters)

	if err != nil {
		return nil, err
	}

	return torrents.PickVideoTorrent(filteredTorrents, filters)
}

// GetTorrents will search The Pirate Bay for torrents of this movie that comply with the given filters.
// It will return one torrent for each video quality.
func (m Movie) GetTorrents(scraper torrents.PirateBayScaper, filters torrents.SearchFilters) ([]torrents.Torrent, error) {

	trnts, err := getTorrents(scraper, filters, m.Title, m.Year)

	if err != nil {
		return nil, err
	}

	if trnts == nil && m.AltTitle != "" {

		trnts, err = getTorrents(scraper, filters, m.AltTitle, m.Year)

		if err != nil {
			return nil, err
		}

	}

	return trnts, nil
}

// SearchQuery returns the normalized title of the movie, as it will be used when searching for torrents.
func (m Movie) SearchQuery() string {
	return utils.NormalizeQuery(m.Title)
}

func getTorrents(scraper torrents.PirateBayScaper, filters torrents.SearchFilters, title string, year uint) ([]torrents.Torrent, error) {

	if year == 0 {

		return scraper.SearchVideoTorrents(title, filters, title)

	}

	return scraper.SearchVideoTorrents(title, filters, title, fmt.Sprint(year))
}
