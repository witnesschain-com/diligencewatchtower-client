package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/publicsuffix"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	coordCfg "github.com/witnesschain-com/diligencewatchtower-client/coordinator/configuration"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
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

func SignCoordinatorMessage(message string, watchtowerAddress common.Address, vault *keystore.Vault) (string, error) {
	signature, err := vault.SignData([]byte(message))
	if err != nil {
		wtCommon.Error(err)
		return "", err
	}
	return hexutil.Encode(signature), nil
}

func (cc *CoordinatorClient) Initialize(coordinatorUrl string, watchtowerAddress common.Address, chainID string, config *wtCommon.SimplifiedConfig) error {
	BASEURL = coordinatorUrl
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}

	cc.client = http.Client{
		Jar: jar,
	}

	cc.config.WatchtowerAddress = watchtowerAddress

	cc.config.Vault, err = keystore.SetupVault(config)
	if err != nil {
		wtCommon.Error(err)
		return err
	}

	cc.config.WatchingChainID = chainID
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
		wtCommon.Error(fmt.Sprintf("Unable to Login - Err[%d] %v", statusCode, err))
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
		"publicKey": cc.config.WatchtowerAddress.Hex(),
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
	signedMessage, err := SignCoordinatorMessage(message, cc.config.WatchtowerAddress, cc.config.Vault)
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
