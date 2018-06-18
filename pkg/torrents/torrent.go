package torrents

import (
	"fmt"
	"strings"
	"time"
)

// Torrent holds all the information regarding a torrent.
type Torrent struct {
	Title            string       `json:"title"`
	Size             int64        `json:"size"` // In kilobytes
	Seeders          int          `json:"seeders"`
	Leeches          int          `json:"leeches"`
	VerifiedUploader bool         `json:"verified_uploader"`
	VideoQuality     VideoQuality `json:"video_quality"`
	MirrorURL        string       `json:"mirror_url"`
	TorrentURL       string       `json:"torrent_url"`
	Magnet           string       `json:"magnet"`
	UploadTime       time.Time    `json:"upload_time"`
	Uploader         string       `json:"uploader"`
}

// FullURL returns the absolute URL for this torrent, including the mirror it was scraped from.
func (t Torrent) FullURL() string {
	return fmt.Sprintf("%v/%v", strings.Trim(t.MirrorURL, "/"), strings.Trim(t.TorrentURL, "/"))
}

// PeersString returns a string representation of the torrent's connected peers
// in the Seeds/Peers format.
func (t Torrent) PeersString() string {
	return fmt.Sprintf("%v / %v", t.Seeders, t.Seeders+t.Leeches)
}

// SizeString returns a formatted string representation of the torrent's file size.
func (t Torrent) SizeString() string {
	const unit = 1000
	sizeBytes := t.Size * 1000
	if sizeBytes < unit {
		return fmt.Sprintf("%d B", sizeBytes)
	}
	div, exp := int64(unit), 0
	for n := sizeBytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(sizeBytes)/float64(div), "KMGTPE"[exp])
}
