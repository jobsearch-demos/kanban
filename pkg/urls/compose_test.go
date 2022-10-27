package urls

import (
	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestComposeFullURLWithTrailingSlash(t *testing.T) {
	baseUrl := "https://google.com"
	partialUrl := "testing"
	queries := map[string]string{
		"name":    "name",
		"surname": "surname",
	}

	// with trailing slash
	expectedUrl := "https://google.com/testing/?name=name&surname=surname"
	fullUrl, err := ComposeFullURL(baseUrl, partialUrl, true, &queries, nil)
	require.NoError(t, err)
	require.Equal(t, expectedUrl, fullUrl.String())
}

func TestComposeFullURLWithoutTrailingSlash(t *testing.T) {
	baseUrl := "https://google.com"
	partialUrl := "testing"
	queries := map[string]string{
		"name":    "name",
		"surname": "surname",
	}

	// with trailing slash
	expectedUrl := "https://google.com/testing?name=name&surname=surname"
	fullUrl, err := ComposeFullURL(baseUrl, partialUrl, false, &queries, nil)
	require.NoError(t, err)
	require.Equal(t, expectedUrl, fullUrl.String())
}

func TestComposeFullURLWithCustomQS(t *testing.T) {
	baseUrl := "https://google.com"
	partialUrl := "testing"

	type QSStruct struct {
		Name    string `url:"name"`
		Surname string `url:"surname"`
	}

	qString, err := query.Values(QSStruct{Name: "name", Surname: "surname"})

	require.NoError(t, err)
	require.NotEmpty(t, qString)

	// with trailing slash
	expectedUrl := "https://google.com/testing?name=name&surname=surname"
	fullUrl, err := ComposeFullURL(baseUrl, partialUrl, false, nil, &qString)

	require.NoError(t, err)
	require.Equal(t, expectedUrl, fullUrl.String())

}

func TestComposeFullURLWithCustomQSAndTrailingSlash(t *testing.T) {
	baseUrl := "https://google.com"
	partialUrl := "testing"

	type QSStruct struct {
		Name    string `url:"name"`
		Surname string `url:"surname"`
	}

	qString, err := query.Values(QSStruct{Name: "name", Surname: "surname"})

	require.NoError(t, err)
	require.NotEmpty(t, qString)

	// with trailing slash
	expectedUrl := "https://google.com/testing/?name=name&surname=surname"
	fullUrl, err := ComposeFullURL(baseUrl, partialUrl, true, nil, &qString)

	require.NoError(t, err)
	require.Equal(t, expectedUrl, fullUrl.String())
}

func TestComposeFullURLWithEmptyQueryStringStructParam(t *testing.T) {
	baseUrl := "https://google.com"
	partialUrl := "testing"

	type QSStruct struct {
		Name    string `url:"name"`
		Surname string `url:"surname"`
	}

	qString, err := query.Values(QSStruct{Name: "", Surname: "surname"})

	require.NoError(t, err)
	require.NotEmpty(t, qString)

	// with trailing slash
	expectedUrl := "https://google.com/testing?surname=surname"
	fullUrl, err := ComposeFullURL(baseUrl, partialUrl, false, nil, &qString)

	require.NoError(t, err)
	require.Equal(t, expectedUrl, fullUrl.String())
}

func TestComposeFullURLWithEmptyQueryStringMapParam(t *testing.T) {
	baseUrl := "https://google.com"
	partialUrl := "testing"

	// with trailing slash
	expectedUrl := "https://google.com/testing?surname=surname"
	fullUrl, err := ComposeFullURL(baseUrl, partialUrl, false, &map[string]string{
		"name":    "",
		"surname": "surname",
	}, nil)

	require.NoError(t, err)
	require.Equal(t, expectedUrl, fullUrl.String())
}
