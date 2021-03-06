package helpers

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/Spazzy757/paul/pkg/config"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTAuth token issued by Github in response to signed JWT Token
type JWTAuth struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

/*
GetEnv looks up an env key or returns a default
*/
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//GetAccessToken returns a Github OAuth Token
func GetAccessToken(config config.Config, installationID int64) (string, error) {
	token := os.Getenv("PERSONAL_ACCESS_TOKEN")
	if len(token) == 0 {
		installationToken, tokenErr := makeAccessTokenForInstallation(
			config.ApplicationID,
			installationID,
			config.PrivateKey,
		)
		if tokenErr != nil {
			return "", tokenErr
		}
		token = installationToken
	}
	return token, nil
}

// MakeAccessTokenForInstallation makes an access token for an installation / private key
func makeAccessTokenForInstallation(appID string, installation int64, privateKey string) (string, error) {
	signed, err := getSignedJwtToken(appID, privateKey)

	if err != nil {
		msg := fmt.Sprintf("can't run GetSignedJwtToken for app_id: %s and installation_id: %d, error: %v", appID, installation, err)

		fmt.Printf("Error %s\n", msg)
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("https://api.github.com/app/installations/%d/access_tokens", installation), nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", signed))
	req.Header.Add("Accept", "application/vnd.github.machine-man-preview+json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		msg := fmt.Sprintf("can't get access_token for app_id: %s and installation_id: %d error: %v", appID, installation, err)
		fmt.Printf("Error: %s\n", msg)
		return "", fmt.Errorf("%s", msg)
	}

	defer res.Body.Close()

	bytesOut, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return "", readErr
	}

	jwtAuth := JWTAuth{}
	jsonErr := json.Unmarshal(bytesOut, &jwtAuth)
	if jsonErr != nil {
		return "", jsonErr
	}
	return jwtAuth.Token, nil
}

// GetSignedJwtToken get a tokens signed with private key
func getSignedJwtToken(appID string, privateKey string) (string, error) {

	keyBytes := []byte(privateKey)

	key, keyErr := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if keyErr != nil {
		return "", keyErr
	}

	now := time.Now()
	claims := jwt.StandardClaims{
		Issuer:    appID,
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Minute * 9).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedVal, signErr := token.SignedString(key)
	if signErr != nil {
		return "", signErr
	}

	return string(signedVal), nil
}

//MockHTTPClient A helper to Mock out Http Servers for testing
func MockHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewTLSServer(handler)
	cli := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
		},
	}

	return cli, s.Close
}
