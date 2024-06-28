package httprequest

import (
	"net/http"
)

func HttpRequest(url string) (int, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
		return resp.StatusCode, err
	}
	return 0, err
}
