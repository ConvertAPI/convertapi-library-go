package convertapi

import (
	"github.com/ConvertAPI/convertapi-go/config"
	"github.com/ConvertAPI/convertapi-go/param"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestSetup(t *testing.T) {
	config.Default.Secret = os.Getenv("CONVERTAPI_SECRET")
	assert.Equal(t, config.Default.Secret, os.Getenv("CONVERTAPI_SECRET"))
}

func TestConvertPath(t *testing.T) {
	file, errs := ConvertPath("assets/test.docx", path.Join(os.TempDir(), "convertapi-test.pdf"))

	assert.Nil(t, errs)
	assert.NotEmpty(t, file.Name())
}

func TestChained(t *testing.T) {
	jpgRes := Convert("docx", "jpg", []param.IParam{
		param.NewPath("file", "assets/test.docx", nil),
	}, nil)

	zipRes := Convert("jpg", "zip", []param.IParam{
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
	user, err := UserInfo(nil)

	assert.Nil(t, err)
	assert.NotEmpty(t, user.SecondsLeft)
}
