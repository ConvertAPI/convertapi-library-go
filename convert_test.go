package convertapi

import (
	"net/url"
	"os"
	"path"
	"testing"

	"github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/param"
	"github.com/stretchr/testify/assert"
)

// The tests in this file call the ConvertAPI API. They run only when CONVERTAPI_TOKEN is set, against
// CONVERTAPI_BASE_URL when it is set (CI uses the staging API) and the production API otherwise.
func liveConfig(t *testing.T) *config.Config {
	t.Helper()
	token := os.Getenv("CONVERTAPI_TOKEN")
	if token == "" {
		t.Skip("CONVERTAPI_TOKEN is not set, so the API tests are skipped")
	}
	conf := config.NewDefault(token)
	if base := os.Getenv("CONVERTAPI_BASE_URL"); base != "" {
		baseURL, err := url.ParseRequestURI(base)
		if err != nil {
			t.Fatalf("CONVERTAPI_BASE_URL %q: %v", base, err)
		}
		conf.BaseURL = baseURL
	}
	config.Default = conf
	return conf
}

func TestConvertPath(t *testing.T) {
	liveConfig(t)
	file, errs := convertapi.ConvertPath("assets/test.docx", path.Join(t.TempDir(), "convertapi-test.pdf"))

	assert.Nil(t, errs)
	assert.NotEmpty(t, file.Name())
}

func TestChained(t *testing.T) {
	liveConfig(t)
	jpgRes := convertapi.Convert("docx", "jpg", []param.IParam{
		param.NewPath("file", "assets/test.docx", nil),
	}, nil)

	zipRes := convertapi.Convert("any", "zip", []param.IParam{
		param.NewResult("files", jpgRes, nil),
	}, nil)

	zipCost, err := zipRes.Cost()

	assert.Nil(t, err)
	assert.NotEmpty(t, zipCost)

	files, err := zipRes.Files()

	assert.Nil(t, err)
	if assert.NotEmpty(t, files) {
		assert.Equal(t, "test.zip", files[0].FileName)
	}
}

func TestUserInfo(t *testing.T) {
	conf := liveConfig(t)
	user, err := convertapi.UserInfo(conf)

	assert.Nil(t, err)
	if assert.NotNil(t, user) {
		assert.NotEmpty(t, user.ConversionsTotal)
	}
}
