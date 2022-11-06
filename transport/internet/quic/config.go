package quic

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"

	"golang.org/x/crypto/chacha20poly1305"

	"github.com/robovpn/github.com/robovpn/v2fly_core/common"
	"github.com/robovpn/github.com/robovpn/v2fly_core/common/protocol"
	"github.com/robovpn/github.com/robovpn/v2fly_core/common/serial"
	"github.com/robovpn/github.com/robovpn/v2fly_core/transport/internet"
)

func getAuth(config *Config) (cipher.AEAD, error) {
	security := config.Security.GetSecurityType()
	if security == protocol.SecurityType_NONE {
		return nil, nil
	}

	salted := []byte(config.Key + "v2fly-quic-salt")
	key := sha256.Sum256(salted)

	if security == protocol.SecurityType_AES128_GCM {
		block, err := aes.NewCipher(key[:16])
		common.Must(err)
		return cipher.NewGCM(block)
	}

	if security == protocol.SecurityType_CHACHA20_POLY1305 {
		return chacha20poly1305.New(key[:])
	}

	return nil, newError("unsupported security type")
}

func getHeader(config *Config) (internet.PacketHeader, error) {
	if config.Header == nil {
		return nil, nil
	}

	msg, err := serial.GetInstanceOf(config.Header)
	if err != nil {
		return nil, err
	}

	return internet.CreatePacketHeader(msg)
}
