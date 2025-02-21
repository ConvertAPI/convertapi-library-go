package convertapi

import (
	"github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/param"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

var authCred = os.Getenv("CONVERTAPI_SECRET")

func TestSetup(t *testing.T) {
	config.Default = config.NewDefault(authCred)
	assert.Equal(t, config.Default.CaTransport.AuthCred, authCred)
}

func TestConvertPath(t *testing.T) {
	config.Default = config.NewDefault(authCred)
	file, errs := convertapi.ConvertPath("assets/test.docx", path.Join(os.TempDir(), "convertapi-test.pdf"))

	assert.Nil(t, errs)
	assert.NotEmpty(t, file.Name())
}

func TestChained(t *testing.T) {
	config.Default = config.NewDefault(authCred)
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
	assert.Equal(t, "test.zip", files[0].FileName)
}

func TestUserInfo(t *testing.T) {
	config.Default = config.NewDefault(authCred)
	user, err := convertapi.UserInfo(nil)

	assert.Nil(t, err)
	assert.NotEmpty(t, user.ConversionsTotal)
	assert.NotEmpty(t, user.ConversionsConsumed)
}
