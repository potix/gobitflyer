package api

import (
        "os"
	"strconv"
        "strings"
	"time"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
        "io/ioutil"
        "github.com/pkg/errors"
)

type Authenticator interface {
	SetAuthHeaders(headers map[string]string, now time.Time, method string, path string, body []byte)
}

type authenticator struct {
	apiKeyFile string
	apiKey     string
	apiSecret  string
}

func (a *authenticator) SetAuthHeaders(headers map[string]string, now time.Time, method string, path string, body []byte) {
	timestamp := strconv.FormatInt(now.UnixNano(), 10)
	mac := hmac.New(sha256.New, []byte(a.apiSecret))
	mac.Write([]byte(timestamp))
	mac.Write([]byte(method))
	mac.Write([]byte(path))
	if len(body) != 0 {
		mac.Write(body)
	}
	sign := hex.EncodeToString(mac.Sum(nil))
	headers["ACCESS-KEY"] = a.apiKey
	headers["ACCESS-TIMESTAMP"] = timestamp
	headers["ACCESS-SIGN"] = sign
}

func (a *authenticator) LoadAPIKey() (error) {
        fileInfo, err := os.Stat(a.apiKeyFile)
        if err != nil {
                return errors.Wrapf(err, "not exists api key file (%v)", a.apiKeyFile)
        }
        if fileInfo.Mode().Perm() != 0600 {
                return errors.Errorf("api key file have insecure permission (e.g. !=  0600) (%v)", a.apiKeyFile)
        }
        apiKeyPair, err := ioutil.ReadFile(a.apiKeyFile)
        if err != nil {
                return errors.Wrapf(err, "can not read api key file (%v)", a.apiKeyFile)
        }
        s := strings.SplitN(string(apiKeyPair), "\n", 4)
        if len(s) < 2 {
                return errors.Wrapf(err, "can not parse api key file (%v)", a.apiKeyFile)
        }
	a.apiKey = strings.TrimSpace(s[0])
	a.apiSecret = strings.TrimSpace(s[1])
        return nil
}

func NewAuthenticator(apiKeyFile string) (Authenticator, error){
	a := &authenticator {
		apiKeyFile: apiKeyFile,
	}
	err := a.LoadAPIKey()
	if err != nil {
		return nil, errors.Wrapf(err, "can not load api key from file (%v)", a.apiKeyFile)
	}
	return a, nil
}
