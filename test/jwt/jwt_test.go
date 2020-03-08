package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCreate_jwt(t *testing.T) {
	header:=NewHeader()
	payload:=Payload{
		Iss:      "liu",
		Exp:       strconv.FormatInt(time.Now().Add(3*time.Hour).Unix(), 10),
		Iat:      strconv.FormatInt(time.Now().Unix(), 10),
		Username: "hahah",
	}
	h, _ := json.Marshal(header)
	p, _ := json.Marshal(payload)
	headerBase64 := base64.StdEncoding.EncodeToString(h)
	payloadBase64 := base64.StdEncoding.EncodeToString(p)
	str1 := strings.Join([]string{headerBase64, payloadBase64}, ".")
	key:="liu"
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str1))
	s := mac.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(s)
	token := str1 + "." + signature
	fmt.Println(token)
}
