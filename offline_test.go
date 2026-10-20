package convertapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/param"
	"github.com/stretchr/testify/assert"
)

// fakeAPI answers like the ConvertAPI API: upload, convert, download and user. Every request must carry
// the token and the library's User-Agent.
func fakeAPI(t *testing.T) *httptest.Server {
	t.Helper()
	var srv *httptest.Server
	srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		assert.True(t, strings.HasPrefix(r.Header.Get("User-Agent"), "ConvertAPI-Go/"+config.Version+" ("),
			"User-Agent %q", r.Header.Get("User-Agent"))

		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/upload":
			assert.Equal(t, "test.docx", r.URL.Query().Get("filename"))
			io.WriteString(w, "upload-id")
		case r.Method == http.MethodPost && r.URL.Path == "/convert/docx/to/pdf":
			assert.Equal(t, "true", r.URL.Query().Get("storefile"))
			assert.NoError(t, r.ParseForm())
			assert.Equal(t, "upload-id", r.PostForm.Get("file"))
			assert.Equal(t, "2", r.PostForm.Get("pagerange"))
			json.NewEncoder(w).Encode(map[string]any{
				"ConversionCost": 1,
				"Files": []map[string]any{
					{"FileName": "test.pdf", "FileSize": 3, "Url": srv.URL + "/d/test.pdf"},
				},
			})
		case r.Method == http.MethodGet && r.URL.Path == "/d/test.pdf":
			io.WriteString(w, "PDF")
		case r.Method == http.MethodGet && r.URL.Path == "/user":
			io.WriteString(w, `{"Active":true,"ConversionsTotal":250,"ConversionsConsumed":7}`)
		default:
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"Code":4013,"Message":"Unknown endpoint"}`)
		}
	}))
	t.Cleanup(srv.Close)
	return srv
}

func fakeConfig(t *testing.T, srv *httptest.Server) *config.Config {
	t.Helper()
	baseURL, err := url.Parse(srv.URL)
	assert.NoError(t, err)
	conf := config.New("test-token", baseURL, nil)
	config.Default = conf
	return conf
}

func TestOfflineConvert(t *testing.T) {
	conf := fakeConfig(t, fakeAPI(t))
	assert.Equal(t, "test-token", conf.CaTransport.AuthCred)

	res := convertapi.ConvDef("docx", "pdf",
		param.NewPath("file", "assets/test.docx", nil),
		param.NewInt("pagerange", 2),
	)

	cost, err := res.Cost()
	assert.NoError(t, err)
	assert.Equal(t, 1, cost)

	out := path.Join(t.TempDir(), "result.pdf")
	files, errs := res.ToPath(out)
	assert.Nil(t, errs)
	if assert.Len(t, files, 1) {
		body, err := os.ReadFile(files[0].Name())
		assert.NoError(t, err)
		assert.Equal(t, "PDF", string(body))
	}
}

func TestOfflineUserInfo(t *testing.T) {
	conf := fakeConfig(t, fakeAPI(t))
	user, err := convertapi.UserInfo(conf)

	assert.NoError(t, err)
	if assert.NotNil(t, user) {
		assert.Equal(t, 250, user.ConversionsTotal)
		assert.Equal(t, 7, user.ConversionsConsumed)
	}
}

func TestOfflineError(t *testing.T) {
	fakeConfig(t, fakeAPI(t))
	_, err := convertapi.ConvDef("pptx", "pdf", param.NewString("file", "https://example.com/a.pptx")).Files()

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "4013")
	}
}

func TestNewWithoutDefault(t *testing.T) {
	config.Default = nil
	conf := config.New("test-token", nil, nil)

	assert.Equal(t, "https://v2.convertapi.com", conf.BaseURL.String())
	assert.Equal(t, "test-token", conf.CaTransport.AuthCred)
}
