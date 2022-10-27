package urls

import (
	"net/url"
	"path"
)

func ComposeFullURL(baseURL, partialURL string, appendTrailingSlash bool, queryParams *map[string]string, queryString *url.Values) (*url.URL, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = path.Join(URL.Path, partialURL)

	if appendTrailingSlash == true {
		URL.Path += "/"
	}

	query := URL.Query()

	if queryString != nil {
		for k, v := range *queryString {
			if v[0] == []string{""}[0] {
				queryString.Del(k)
			}
		}
		URL.RawQuery = queryString.Encode()
		return URL, nil
	}

	for k, v := range *queryParams {
		if v != "" {
			query.Set(k, v)
		}
	}

	URL.RawQuery = query.Encode()

	return URL, nil
}
