/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package webserver

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net/http"

	wtCommon "github.com/diligencewatchtower-client/common"
)

var isAlive = false
var publicKey = "INVALID"
var lastAlive = ""
var currentlyWatchingL2 = ""

var https = true
var hostname = ""
var webserverStarted = false

func OkResponse() string {

	return fmt.Sprintf("{\"isAlive\":%t,\"publicKey\":\"%s\",\"lastAlive\":\"%s\",\"currentlyWatchingL2\":\"%s\"}", isAlive, publicKey, lastAlive, currentlyWatchingL2)
}

func Start(server *http.Server, webServerConfig *wtCommon.WebServerConfig, config_chan chan bool) {

	now := time.Now()
	lastAlive = now.Format(time.ANSIC)

	isAlive = true
	hostname = webServerConfig.HostName
	publicKey = webServerConfig.PublicKeyAddressHex
	currentlyWatchingL2 = webServerConfig.CurrentlyWatchingL2

	if !webserverStarted {

		// get ip : curl https://checkip.amazonaws.com

		if hostname == "" {

			res, err := http.Get("https://checkip.amazonaws.com")

			if err != nil {
				https = false
			}

			if (res != nil) && (res.StatusCode != 200) {
				https = false
			}

			if https {

				body, err := ioutil.ReadAll(res.Body)

				if err != nil {
					https = false
					wtCommon.Warning("Your Public IP could not be determined")
				} else {

					ip := strings.TrimSpace(string(body))

					wtCommon.Info("Your Public IP is")
					wtCommon.Success(ip)

					// get hostname : nslookup ip

					looked_up_hostname, err := net.LookupAddr(ip)

					if err != nil {
						https = false
						wtCommon.Warning("Could not lookup your hostname")
					} else {
						if len(looked_up_hostname) > 0 {
							hostname = looked_up_hostname[0]
						} else {
							https = false
							wtCommon.Error("looked_up_hostname is empty!")
						}
					}
				}
			}
		}

		https_crt := "https-fullchain.pem"
		https_key := "https-privkey.pem"

		cert_exists := true
		cert_key_exists := true

		c, err := os.Stat(https_crt)
		if errors.Is(err, os.ErrNotExist) {
			cert_exists = false
		}

		k, err := os.Stat(https_key)
		if errors.Is(err, os.ErrNotExist) {
			cert_key_exists = false
		}

		_ = c
		_ = k

		if cert_exists == false || cert_key_exists == false {
			https = false
		}

		if cert_exists == true && cert_key_exists == true {
			https = true
		}

		if hostname == "" && https == true {

			pemData, err := ioutil.ReadFile(https_crt)

			if err == nil {

				block, err := pem.Decode(pemData)

				if err == nil {

					cert, err := x509.ParseCertificate(block.Bytes)

					if err == nil {

						wtCommon.Info("Getting hostname from `" + https_crt + "` ...")
						wtCommon.Success(cert.Subject.CommonName)

						hostname = cert.Subject.CommonName
					}
				}
			}
		}

		wtCommon.Info("Starting webserver @ " + hostname + ":8443 ...")

		server = &http.Server{
			Addr:              ":8443",
			IdleTimeout:       5 * time.Second,
			ReadTimeout:       2 * time.Second,
			WriteTimeout:      2 * time.Second,
			ReadHeaderTimeout: 2 * time.Second,
		}

		http.HandleFunc("/ok", Ok)

		wtCommon.Success("Webserver started!")

		webserverStarted = true

		if https {

			server.TLSConfig = LoadTLSConfig(https_crt, https_key)

			err := server.ListenAndServeTLS(
				https_crt, https_key,
			)

			if errors.Is(err, http.ErrServerClosed) {
				wtCommon.Fatal(err)
			} else {
				webserverStarted = true
			}

		} else {

			wtCommon.Warning("Not running with HTTPS")

			err := server.ListenAndServe()

			if err != http.ErrServerClosed {
				webserverStarted = true
			} else {
				wtCommon.Fatal(err)
				return
			}
		}
	}
}

func Stop(server *http.Server) {

	isAlive = false
}

func SetHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "OPTIONS, GET")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Request-Method", "*")
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
}

func Ok(w http.ResponseWriter, r *http.Request) {

	SetHeaders(w)

	now := time.Now()
	lastAlive = now.Format(time.ANSIC)

	fmt.Fprintf(w, OkResponse())
}
