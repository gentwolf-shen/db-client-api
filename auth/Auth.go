package auth

import (
	"github.com/gentwolf-shen/gohelper/convert"
	"github.com/gentwolf-shen/gohelper/cryptohelper/aes"
	"github.com/gentwolf-shen/gohelper/timehelper"
	"github.com/gentwolf-shen/gohelper/util"
)

func GetToken(appKey, appSecret string) (string, error) {
	str := util.Uuid() + "|" + appKey + "|" + convert.ToStr(timehelper.Timestamp())
	crypto := aes.New(aes.CBC, []byte(appSecret[0:16]), []byte(appSecret[16:]))
	token, err := crypto.EncryptToString([]byte(str))
	if err != nil {
		return "", err
	}
	return token, nil
}
