package toolkit

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
)

//const url = "https://union.jd.com/api/apiDoc/apiSignParam"
// 0aa97446a384a7f4fc97296fca506d59
// 12ad4702be644255b1f7d1d5f71f6208
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

	c := &http.Client{}
	resp, err := c.Do(req)
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
