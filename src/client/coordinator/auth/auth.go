package auth

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/publicsuffix"

	wtCommon "github.com/diligencewatchtower-client/common"
	coordCfg "github.com/diligencewatchtower-client/watchtower/coordinator/configuration"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type CoordinatorClient struct {
	client http.Client
	config coordCfg.CoordinatorConfiguration
}

const (
	PRELOGIN string = "tracer/v1/watchtower/pre-login"
	LOGIN    string = "tracer/v1/watchtower/login"
)

var BASEURL string

func SignCoordinatorMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}

func (cc *CoordinatorClient) Initialize(coordinatorUrl string, privateKey *ecdsa.PrivateKey, chainID string) error {
	BASEURL = coordinatorUrl
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}

	cc.client = http.Client{
		Jar: jar,
	}

	cc.config.PrivateKey = privateKey
	cc.config.WatchingChainID = chainID
	addr, _ := wtCommon.GetPublicKeyAddressFromPrivateKey(privateKey)
	cc.config.PublicKeyHex = addr.Hex()
	return nil
}

func (cc *CoordinatorClient) AuthenticateToWitnesschain() (bool, error) {
	message, err := cc.doPrelogin()
	if err != nil {
		return false, err
	}
	statusCode, err := cc.doLogin(message)
	if err != nil {
		return false, err
	}

	if statusCode != 200 {
		wtCommon.Error(fmt.Sprintf("Unable to Login - Err[%d]", statusCode))
		return false, nil
	}
	return true, nil
}

func (cc CoordinatorClient) GetHeaders() http.Header {
	url, _ := url.Parse(BASEURL)
	cookies := cc.client.Jar.Cookies(url)
	var header http.Header = http.Header{}
	if len(cookies) > 1 {
		header.Add("Content-Type", "application/json")
		header.Add("Cookie", "__Secure-session="+(*cookies[0]).Value+"; __Secure-session.hash="+(*cookies[1]).Value)
	}
	return header
}

func (cc CoordinatorClient) doPrelogin() (string, error) {
	body, _ := json.Marshal(map[string]string{
		"publicKey":         cc.config.PublicKeyHex,
		"keyType":           "ethereum", //ToDo: Get this dynamically
		"currentlyWatching": cc.config.WatchingChainID,
	})
	requestBody := bytes.NewBuffer(body)
	resp, err := cc.client.Post(
		BASEURL+PRELOGIN,
		"application/json",
		requestBody,
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	apiPreloginResponse := preloginResponse{}
	err = json.NewDecoder(resp.Body).Decode(&apiPreloginResponse)
	if err != nil {
		return "", err
	}
	return apiPreloginResponse.Result.Message, nil
}

func (cc CoordinatorClient) doLogin(message string) (int, error) {
	signedMessage, err := SignCoordinatorMessage(message, cc.config.PrivateKey)
	if err != nil {
		return -1, err
	}
	body, _ := json.Marshal(map[string]string{
		"message":   message,
		"signature": signedMessage,
	})
	requestBody := bytes.NewBuffer(body)
	resp, err := cc.client.Post(
		BASEURL+LOGIN,
		"application/json",
		requestBody,
	)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

type preloginResponse struct {
	Result struct {
		Message string
	}
}
