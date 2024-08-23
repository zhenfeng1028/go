package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"log"
)

const (
	// 私钥
	PRIVATE_KEY = `
-----BEGIN PRIVATE KEY-----
MIIEowIBAAKCAQEArk8Pt3qCZ/TXiCz9O1Q0xuCfTnudSM14AoejFYWdnZqIStuV
sEhlgL88VkxGnbuSmyKZn70CjyFgcDQwdbFsHxHsP5ph63i+oPawEQ3aZdhnRqe7
LOt8Iej9h/Qvd7F3d76ftHiF4/vBT4gdsc7IjdW7gJ7kyBcC6a2hVthpOK9CTN/7
jneQroXAkoEWZyeAzf8OwL0v+dh8EbuHzcQL+aZ6VAf72rYinNrwXfGsVgoUCMtF
/kKGCIlu35lxzAOPHWDGrfHPMDdcQkk3MblVrdqlf3U17kiD6QlYtksdxjkTgI3w
s011RiLZE8gDfqJqwhDuSfkshbcPbBRxPsB16QIDAQABAoIBAHoIS96OXGHfXk2u
AZPZviFF/QH7NQBVKSHMR/a3VsOR/r05wjBqXvWl7N4WBJJFUbxemuLkdrgyC8xC
HXMoRaPy7k0RDmDGZd9XFai65B3NRXCH2Vj3fC5ZabYW1GydyCADStgUUNht1saR
awii1nHGIsFHfLBAu/9RWdTtKa1OiRpDK/uRvOD2MHqYxOupqyxM3ECYI82cj/VK
8H2Gl3qNG08xrQHx34PhtE/w0ojNGYxThdQsM9FCBVIkTraluhlBGrOYQtwL08v0
29HtYo9Y4laU4PdB+AtkwnV0fWo/KZXcA7NXBAWWiYFOYU7a6kklDo1YtMTAmr/X
q65IxWkCgYEA/TqKLAyxKfIQdOkEmJ2Ir8xvSYgvN/HOuTE130QqH7DC8BYNBzds
7XCyEqYQvNQfgS5h1sNnh27awWkuKP7kPUWNmDrk0vBZdgrwCTAligxXGqWiULJ6
Y1ZlYtit5+u70bWmSqc8D6Ot2bsaSZ74MQMSlPFThMANXvgXEDSSeQMCgYEAsDdq
TBf9GbwYWthbyW60SCPxRnIk8AVExdLnLbfiDK8rZNuKbTD3t1Jtq/khaPQhno0S
1t2kxdCbrn/cmZCxjhXZ92Xa98i9cD5rZ1kiwF4JwRCHpwVmSPBOHVWGKXwxJzS9
WNo0jTZg3x6yYD+WRE1dUHYfT5ARwu4UJW2mI6MCgYAEu7hjdGVHXG/0T3q3g5JP
funBnRsM16c+jO3t229SIlKyfCcIMIAHvCiH8rSzYt+n4X8QrZxcyrSqEtY7C/Cm
OIuaAdfPHZD+hNvRwCWje/GQ3PG4wm1JOA8zFsvBXwWmuNd/47parHuOa1HleGQ5
Pk51nEEsZQ6a7NY5hlIKkwKBgA/d3X9ehZaknhjH9HPNuQPp4Ja0PmkyQEADvnca
YIs91o1tSyiLyTJDhMaAVybJHKHXGAQqzuRU5T+wGZ/mLGabxxWEoVPASnZp2Gfk
2cEydnRg1aYPUWdt48q2Ya8olJBwKkBu3V59r/lHHU6XSyN5R9av3B/g3AiOQVgh
aF07AoGBALYSh3aGyrR0r8F6cjbzc+8zrx//Vsl+R5aJGVYCCF0GzWaLVP51UASa
NCuqXKoTcf+6L9oEH3tZqgnSqNKL76HIUc4puVyL4sfsByKRWXUH806jPBYvKLEm
cmKerBCl6jRhzLEmkp0RnLokNL4GkO1O/EZlYuAMU75wErMHFBfA
-----END PRIVATE KEY-----
`
	// 公钥
	PUBLIC_KEY = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArk8Pt3qCZ/TXiCz9O1Q0
xuCfTnudSM14AoejFYWdnZqIStuVsEhlgL88VkxGnbuSmyKZn70CjyFgcDQwdbFs
HxHsP5ph63i+oPawEQ3aZdhnRqe7LOt8Iej9h/Qvd7F3d76ftHiF4/vBT4gdsc7I
jdW7gJ7kyBcC6a2hVthpOK9CTN/7jneQroXAkoEWZyeAzf8OwL0v+dh8EbuHzcQL
+aZ6VAf72rYinNrwXfGsVgoUCMtF/kKGCIlu35lxzAOPHWDGrfHPMDdcQkk3MblV
rdqlf3U17kiD6QlYtksdxjkTgI3ws011RiLZE8gDfqJqwhDuSfkshbcPbBRxPsB1
6QIDAQAB
-----END PUBLIC KEY-----
`
)

func main() {
	// 原内容
	str := "F78691A5-2963-42D2-A0C7-3A4F31CF5EB6"
	// 生成签名
	sig := Rsa2PriSign(str, PRIVATE_KEY, crypto.SHA256)
	log.Println(sig)
	// 验证原内容与签名是否一致
	res := Rsa2PubCheckSign(str, sig, PUBLIC_KEY, crypto.SHA256)
	log.Println(res)
}

// RSA2私钥签名
func Rsa2PriSign(signContent string, privateKey string, hash crypto.Hash) string {
	shaNew := hash.New()
	shaNew.Write([]byte(signContent))
	hashed := shaNew.Sum(nil)
	priKey, err := ParsePrivateKey(privateKey)
	if err != nil {
		return ""
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, hash, hashed)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(signature)
}

// 解析私钥
func ParsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priKey, nil
}

// RSA2公钥验证签名
func Rsa2PubCheckSign(signContent, sign, publicKey string, hash crypto.Hash) bool {
	hashed := sha256.Sum256([]byte(signContent))
	pubKey, err := ParsePublicKey(publicKey)
	if err != nil {
		return false
	}
	sig, _ := base64.StdEncoding.DecodeString(sign)
	err = rsa.VerifyPKCS1v15(pubKey, hash, hashed[:], sig)
	return err == nil
}

// 解析公钥
func ParsePublicKey(publicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("公钥信息错误！")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKey.(*rsa.PublicKey), nil
}
