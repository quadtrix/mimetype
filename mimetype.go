// package mimetype Identifies the mime type of a file
package mimetype

import (
	"errors"
	"os"
	"strings"
)

type MIME struct {
	mtype     string
	subtype   string
	filename  string
	extension string
}

func getfile(file string, limit int32) ([]byte, error) {
	fileinfo, err := os.Stat(file)
	if err != nil {
		return []byte(""), err
	}
	if fileinfo.IsDir() {
		return []byte(""), errors.New("not a file")
	}
	fh, err := os.Open(file)
	if err != nil {
		return []byte(""), err
	}
	defer fh.Close()
	if limit > 50 {
		limit = 50
	}
	byteslice := make([]byte, limit)
	_, err = fh.Read(byteslice)
	if err != nil {
		return []byte(""), err
	}
	return byteslice, err
}

func gethint(file string) (m MIME) {
	parts := strings.Split(file, ".")
	ext := parts[len(parts)-1]
	m.filename = file
	m.extension = ext
	switch ext {
	case "js":
		m.mtype = "application"
		m.subtype = "javascript"
	case "css":
		m.mtype = "text"
		m.subtype = "css"
	case "htm", "html":
		m.mtype = "text"
		m.subtype = "html"
	case "jpg", "jpeg":
		m.mtype = "image"
		m.subtype = "jpeg"
	case "png":
		m.mtype = "image"
		m.subtype = "png"
	case "gif":
		m.mtype = "image"
		m.subtype = "gif"
	case "jar":
		m.mtype = "application"
		m.subtype = "jar"
	case "pdf":
		m.mtype = "application"
		m.subtype = "pdf"
	case "webp":
		m.mtype = "image"
		m.subtype = "webp"
	case "tif", "tiff":
		m.mtype = "image"
		m.subtype = "tiff"
	case "bmp":
		m.mtype = "image"
		m.subtype = "bmp"
	case "ico":
		m.mtype = "image"
		m.subtype = "x-icon"
	case "gz":
		m.mtype = "application"
		m.subtype = "gzip"
	case "tar":
		m.mtype = "application"
		m.subtype = "tar"
	case "zip":
		m.mtype = "application"
		m.subtype = "zip"
	case "txt":
		m.mtype = "text"
		m.subtype = "plain"
	case "rss":
		m.mtype = "application"
		m.subtype = "rss+xml"
	case "atom":
		m.mtype = "application"
		m.subtype = "atom+xml"
	case "xml":
		m.mtype = "text"
		m.subtype = "xml"
	case "json":
		m.mtype = "application"
		m.subtype = "json"
	case "war":
		m.mtype = "application"
		m.subtype = "octet-stream"
	case "svg":
		m.mtype = "image"
		m.subtype = "svg+xml"
	default:
		m.mtype = "application"
		m.subtype = "octet-stream"
	}
	return m
}

func analyze(fragment []byte) (m MIME) {
	fragment = []byte("unused")
	m.mtype = "application"
	m.subtype = "octet-stream"
	return m
}

// Detect detects a file's MIME type
func Detect(file string) (m MIME) {
	hint := gethint(file)
	if m.mtype == "application" && m.subtype == "octet-stream" {
		fragment, err := getfile(file, 50)
		if err != nil {
			return hint
		}
		return analyze(fragment)
	} else {
		return hint
	}
}

// String returns the MIME type and subtype as a string
func (m MIME) String() string {
	return m.mtype + "/" + m.subtype
}

// Extension returns the MIME extension
func (m MIME) Extension() string {
	return m.extension
}
