package redirect_test

import (
	"net/http/httptest"
	"testing"
	"url_shortener/internal/http-server/handlers/redirect"
	"url_shortener/internal/http-server/handlers/redirect/mocks"
	"url_shortener/internal/lib/api"
	"url_shortener/internal/lib/logger/handlers/slogDiscard"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestSaveHandler(t *testing.T) {
	cases := []struct {
		name      string
		alias     string
		url       string
		respError string
		mockError error
	}{
		{
			name:  "Success",
			alias: "test_alias",
			url:   "https://www.google.com/",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			urlGetterMock := mocks.NewUrlGetter(t)

			if tc.respError == "" || tc.mockError != nil {
				urlGetterMock.On("GetUrl", tc.alias).
					Return(tc.url, tc.mockError).Once()
			}

			r := chi.NewRouter()
			r.Get("/{alias}", redirect.New(slogDiscard.NewDiscardLogger(), urlGetterMock))

			ts := httptest.NewServer(r)
			defer ts.Close()

			redirectedToURL, err := api.GetRedirect(ts.URL + "/" + tc.alias)
			require.NoError(t, err)

			// Check the final URL after redirection.
			assert.Equal(t, tc.url, redirectedToURL)
		})
	}
}
