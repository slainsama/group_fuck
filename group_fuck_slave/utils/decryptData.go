package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"group_fuck_slave/globals"
)

func DecryptData(data string) (string, error) {
	var aesKey = globals.GlobalConfig.Secure.AesKey

	innerBase64Data, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	encryptedData, err := base64.StdEncoding.DecodeString(string(innerBase64Data))
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return "", err
	}

	// 验证加密数据长度
	if len(encryptedData) < 12+aes.BlockSize { // 至少需要12字节nonce和16字节的GCM块
		err := errors.New("加密数据长度不足")
		return "", err
	}

	// 提取 nonce 和实际加密数据（包含 tag 和 cipherText）
	nonce := encryptedData[:12]
	cipherText := encryptedData[12:]

	// 使用 GCM 模式解密
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 解密数据，cipherText 中已包含认证标签，无需单独处理
	decryptedData, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedData), nil
}
