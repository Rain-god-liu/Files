package jwt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Iss      string `json:"iss"`
	Exp      string `json:"exp"`
	Iat      string `json:"iat"`
	Username string `json:"username"`
	Uid      int
}

type Jwt struct {
	header Header
	payload Payload
	signature string
}

func NewHeader() Header {
	return Header{
		Alg: "HS256",
		Typ: "JWT",
	}
}

func Create_jwt(usernaem string)(string){
	header:=NewHeader()
	payload:=Payload{
		Iss:      "liu",
		Exp:       strconv.FormatInt(time.Now().Add(3*time.Hour).Unix(), 10),
		Iat:      strconv.FormatInt(time.Now().Unix(), 10),
		Username: usernaem,
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
	return token
}

func Makesure_token(token string)( username string, err error){
	arr:= strings.Split(token, ".")
	if len(arr)!=3 {
		err = errors.New("token error")
	}
	_, err = base64.StdEncoding.DecodeString(arr[0])
	if err != nil {
		err = errors.New("token error")
	}
	pay, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		err = errors.New("token error")
	}
	sign, err := base64.StdEncoding.DecodeString(arr[2])
	if err != nil {
		err = errors.New("token error")
	}
	str1 := arr[0] + "." + arr[1]
	key := []byte("liu")
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(str1))
	s := mac.Sum(nil)
	fmt.Println(sign,"这是一次测试打印")
	fmt.Println(s,"这是测试是否成功的测试打印")
	if res := bytes.Compare(sign, s); res != 0 {
		fmt.Println("如果出现了这个消息，说明你的token出现了错误")
		return
	}
	var payload Payload
	json.Unmarshal(pay,&payload)
	username =payload.Username
	return
}
