package main

import (
	"io"
	"log"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func main() {
	err := encryptFile("public.asc", "plaintext.txt", "encrypted.pgp")
	if err != nil {
		log.Fatal(err)
	}
	println("File encrypted successfully to encrypted.pgp")

	err = decryptFile("private.asc", "your-passphrase", "encrypted.pgp", "decrypted.txt")
	if err != nil {
		log.Fatal(err)
	}
	println("File decrypted successfully to decrypted.txt")
}

func encryptFile(publicKeyPath, inputFile, outputFile string) error {
	// 读取公钥
	pubKey, err := os.Open(publicKeyPath)
	if err != nil {
		return err
	}
	defer pubKey.Close()

	entityList, err := openpgp.ReadArmoredKeyRing(pubKey)
	if err != nil {
		return err
	}

	// 创建输出文件
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// 使用ASCII Armor格式创建加密流
	w, err := armor.Encode(outFile, "PGP MESSAGE", nil)
	if err != nil {
		return err
	}
	defer w.Close()

	// 创建加密器
	encryptor, err := openpgp.Encrypt(w, entityList, nil, nil, nil)
	if err != nil {
		return err
	}
	defer encryptor.Close()

	// 读取并加密原始文件
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	_, err = io.Copy(encryptor, inFile)
	return err
}

func decryptFile(privateKeyPath, passphrase, inputFile, outputFile string) error {
	// 读取私钥
	privKey, err := os.Open(privateKeyPath)
	if err != nil {
		return err
	}
	defer privKey.Close()

	entityList, err := openpgp.ReadArmoredKeyRing(privKey)
	if err != nil {
		return err
	}

	// 如果私钥有密码，需要先解密
	entity := entityList[0]
	passphraseByte := []byte(passphrase)
	err = entity.PrivateKey.Decrypt(passphraseByte)
	if err != nil {
		return err
	}
	for _, subkey := range entity.Subkeys {
		subkey.PrivateKey.Decrypt(passphraseByte)
	}

	// 打开被加密的文件
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// 解码ASCII Armor
	block, err := armor.Decode(inFile)
	if err != nil {
		return err
	}
	if block.Type != "PGP MESSAGE" {
		log.Fatal("Invalid message type")
	}

	// 使用密钥环解密消息
	md, err := openpgp.ReadMessage(block.Body, entityList, nil, nil)
	if err != nil {
		return err
	}

	// 读取解密后的数据
	decryptedData, err := io.ReadAll(md.UnverifiedBody)
	if err != nil {
		return err
	}

	// 将解密数据写入文件
	return os.WriteFile(outputFile, decryptedData, 0o644)
}
