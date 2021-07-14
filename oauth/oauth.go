package oauth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/ngoctutrang/ministore-oauth/errors"
)

const (
	headerXPublic    = "X-Public"
	headerXClientId  = "X-Client-Id"
	headerXCallerId  = "X-User-Id"
	paramAccessToken = "access-token"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 200 * time.Millisecond,
	}
)

type accessToken struct {
	Id       string `json:"id"`
	UserId   int64  `json:"user_id"`
	ClientId int64  `json:"client_id"`
}

type oauthInterface interface {
}

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}

func AuthenticateRequest(request *http.Request) *errors.RestErr {
	if request == nil {
		return nil
	}

	accessToken := strings.TrimSpace(
		request.URL.Query().Get(paramAccessToken),
	)

	if accessToken == "" {
		return nil
	}
	return nil
}

func getAccessToken(accessToken string) (*accessToken, *errors.RestErr) {
	response := oauthRestClient.Get(
		fmt.Sprintf("/oauth/access_token/%s", accessToken),
	)
}
