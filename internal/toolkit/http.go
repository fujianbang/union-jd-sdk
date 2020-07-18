package toolkit

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
)

var httpClient = &http.Client{}

//const url = "https://union.jd.com/api/apiDoc/apiSignParam"

func HttpGet(address string, v interface{}) ([]byte, error) {
	urlVal, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	values, _ := query.Values(v)
	urlVal.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", urlVal.String(), nil)
	if err != nil {
		return nil, err
	}
	a := req.URL.RequestURI()
	zap.L().Debug(a)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode != 200 {
		return nil, fmt.Errorf("statusCode: %d", statusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
