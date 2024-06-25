package common

import "io"

func ParseResponseBody(responseBody io.ReadCloser) string {
	bodyBytes, err := io.ReadAll(responseBody)
	if err != nil {
		Warning(err)
		return ""
	}
	return string(bodyBytes)
}
