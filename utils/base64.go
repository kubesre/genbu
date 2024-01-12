/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

package utils

import (
	"encoding/base64"
)

// base64解密

func DecodeBase64(source string) (dest string, err error) {
	destByte, err := base64.StdEncoding.DecodeString(source)
	if err != nil {
		return "", err
	}
	dest = string(destByte)
	return dest, nil
}
