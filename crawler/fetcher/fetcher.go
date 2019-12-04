package fetcher

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var rateLimit = time.Tick(100*time.Microsecond)
// Fetch is used to fetch content of html then return *Document and error
func Fetch(url string) (*goquery.Document, error) {
	<-rateLimit
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: ", "code", res.StatusCode)
	}
	return goquery.NewDocumentFromReader(res.Body)
}
