package smuggle

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func GetBytes(input string) []byte {
	var bytes []byte
	var errBytes error

	bytes, errBytes = os.ReadFile(input)

	if errBytes != nil {
		color.Red(fmt.Sprintf("[!]%s", errBytes.Error()))
		os.Exit(1)
	}

	return bytes

}

func (s *Smuggler) AESEncrypt() (string, error) {
	var err error

	if s.Key == nil {
		s.NewAESKey()
	}

	s.EncBuf, err = AESCBCEncrypt(s.Key, s.Buf)
	if err != nil {
		return "", err
	}

	encstr := base64.StdEncoding.EncodeToString(s.EncBuf)
	s.KeyStr = base64.StdEncoding.EncodeToString(s.Key)

	return encstr, nil

}
