package leveltravel

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	apiURL     = "https://api.level.travel"
	baseURL    = "https://level.travel"
	apiVersion = "3"
)

// API is the struct for Level.Travel API
type API struct {
	apiKey    string
	affiliate string
}

// NewAPI returns new API instance
func NewAPI(apiKey, affiliate string) *API {
	return &API{
		apiKey:    apiKey,
		affiliate: affiliate,
	}
}

func (a *API) get(path string, args map[string]string, v interface{}) error {
	parseURL, err := url.Parse(apiURL + path)
	if err != nil {
		return err
	}
	params := url.Values{}
	params.Add("api_version", apiVersion)
	for k, v := range args {
		params.Add(k, v)
	}

	parseURL.RawQuery = params.Encode()

	resp, err := fetch(parseURL.String(), a.getHeaders())
	if err != nil {
		return err
	}
	defer resp.Close()

	return json.NewDecoder(resp).Decode(&v)
}

func (a *API) getHeaders() map[string]string {
	return map[string]string{
		"Accept":        "application/vnd.leveltravel.v3",
		"Authorization": fmt.Sprintf("Token token=\"%s\"", a.apiKey),
	}
}

// GetURL returns full URL with affiliate mark
func (a *API) GetURL(link string) string {
	return fmt.Sprintf("%s%s?aflt=%s", baseURL, link, a.affiliate)
}

func fetch(url string, headers map[string]string) (resp io.ReadCloser, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return resp, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	httpResp, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	resp = httpResp.Body
	return resp, err
}

func convertToMap(m interface{}) map[string]string {
	var i map[string]interface{}
	j, _ := json.Marshal(m)
	_ = json.Unmarshal(j, &i)

	res := make(map[string]string)
	for k, v := range i {
		res[k] = fmt.Sprintf("%v", v)
	}

	return res
}
