/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package webserver

import (
	"crypto/tls"
	"net/http"

	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
)

func IsAuthenticated(r *http.Request) bool {

	if r.TLS == nil || r.TLS.PeerCertificates == nil {
		return false
	}

	return len(r.TLS.PeerCertificates) >= 1
}

// Loads and returns CA and cert configs for webserver
func LoadTLSConfig(https_crt_file string, https_key_file string) *tls.Config {

	cert, err := tls.LoadX509KeyPair(https_crt_file, https_key_file)
	if err != nil {
		wtCommon.Fatal(err)

	}

	tlsConfig := tls.Config{
		Certificates:             []tls.Certificate{cert},
		MinVersion:               tls.VersionTLS13,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,

			tls.TLS_AES_256_GCM_SHA384,

			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
		},
	}

	return &tlsConfig
}
