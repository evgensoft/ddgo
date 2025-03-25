package ddgo

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// DefaultUserAgent defines a default value for user-agent header.
const DefaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"

// Result holds the returned query data
type Result struct {
	Title string
	Info  string
	URL   string
}

// Requests the query and puts the results into an array
func Query(query string, maxResult int) ([]Result, error) {
	results := []Result{}
	queryUrl := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", url.QueryEscape(query))

	client := &http.Client{}
	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return results, fmt.Errorf("new request error: %w", err)
	}
	req.Header.Set("User-Agent", DefaultUserAgent)

	response, err := client.Do(req)
	if err != nil {
		return results, fmt.Errorf("get %v error: %w", queryUrl, err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusAccepted {
		return results, fmt.Errorf("status code error: %d %s", response.StatusCode, response.Status)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return results, fmt.Errorf("goquery.NewDocument error: %w", err)
	}

	sel := doc.Find(".web-result")

	for i := range sel.Nodes {
		// Break loop once required amount of results are add
		if maxResult == len(results) {
			break
		}
		node := sel.Eq(i)
		titleNode := node.Find(".result__a")

		info := node.Find(".result__snippet").Text()
		title := titleNode.Text()
		ref := ""

		if len(titleNode.Nodes) > 0 && len(titleNode.Nodes[0].Attr) > 2 {
			ref = getDDGUrl(titleNode.Nodes[0].Attr[2].Val)
		}

		results = append(results[:], Result{title, info, ref})

	}

	return results, nil
}

func getDDGUrl(urlStr string) string {
	trimmed := strings.TrimPrefix(urlStr, "//duckduckgo.com/l/?uddg=")
	if idx := strings.Index(trimmed, "&rut="); idx != -1 {
		decodedStr, err := url.PathUnescape(trimmed[:idx])
		if err != nil {
			return ""
		}

		return decodedStr
	}

	return ""
}
