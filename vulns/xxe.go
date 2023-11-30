package vulns

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	ID      int      `xml:"id"`
	Name    string   `xml:"name"`
}

func parseXML(data []byte) (*User, error) {
	var user User
	err := xml.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Xxe(w http.ResponseWriter, r *http.Request) {
	xmlData := r.URL.Query().Get("xmlData")

	if xmlData == "" {
		http.Error(w, "Missing xmlData parameter", http.StatusBadRequest)
		return
	}

	user, err := parseXML([]byte(xmlData))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing XML: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Parsed user: %+v\n", user)
}