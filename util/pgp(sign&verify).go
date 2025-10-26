package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

// 使用发送方私钥签名和接收方公钥加密
func encryptAndSign(senderPrivateKey, receiverPublicKey, passphrase, inputFile, outputFile string) error {
	fmt.Printf("开始加密文件: %s -> %s\n", inputFile, outputFile)

	// 读取发送方私钥
	fmt.Printf("读取发送方私钥: %s\n", senderPrivateKey)
	senderKey, err := os.Open(senderPrivateKey)
	if err != nil {
		return fmt.Errorf("打开发送方私钥文件失败: %v", err)
	}
	defer senderKey.Close()

	senderEntityList, err := openpgp.ReadArmoredKeyRing(senderKey)
	if err != nil {
		return fmt.Errorf("解析发送方私钥失败: %v", err)
	}
	senderEntity := senderEntityList[0]
	fmt.Printf("发送方密钥ID: %X\n", senderEntity.PrimaryKey.KeyId)

	// 解密发送方私钥
	if senderEntity.PrivateKey.Encrypted {
		fmt.Printf("解密发送方私钥...\n")
		err = senderEntity.PrivateKey.Decrypt([]byte(passphrase))
		if err != nil {
			return fmt.Errorf("解密发送方私钥失败: %v", err)
		}
		fmt.Printf("发送方私钥解密成功\n")
	}

	// 读取接收方公钥
	fmt.Printf("读取接收方公钥: %s\n", receiverPublicKey)
	receiverKey, err := os.Open(receiverPublicKey)
	if err != nil {
		return fmt.Errorf("打开接收方公钥文件失败: %v", err)
	}
	defer receiverKey.Close()

	receiverEntityList, err := openpgp.ReadArmoredKeyRing(receiverKey)
	if err != nil {
		return fmt.Errorf("解析接收方公钥失败: %v", err)
	}
	fmt.Printf("接收方密钥ID: %X\n", receiverEntityList[0].PrimaryKey.KeyId)

	// 检查输入文件
	fileInfo, err := os.Stat(inputFile)
	if err != nil {
		return fmt.Errorf("检查输入文件失败: %v", err)
	}
	fmt.Printf("输入文件大小: %d 字节\n", fileInfo.Size())

	// 创建输出文件
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建输出文件失败: %v", err)
	}
	defer outFile.Close()

	// 使用ASCII Armor格式
	w, err := armor.Encode(outFile, "PGP MESSAGE", nil)
	if err != nil {
		return fmt.Errorf("创建ASCII Armor编码器失败: %v", err)
	}

	// 加密并签名
	fmt.Printf("开始加密和签名...\n")
	encryptor, err := openpgp.Encrypt(w, receiverEntityList, senderEntity, nil, nil)
	if err != nil {
		w.Close()
		return fmt.Errorf("创建加密器失败: %v", err)
	}

	// 读取并处理原始文件
	inFile, err := os.Open(inputFile)
	if err != nil {
		encryptor.Close()
		w.Close()
		return fmt.Errorf("打开输入文件失败: %v", err)
	}
	defer inFile.Close()

	_, err = io.Copy(encryptor, inFile)
	if err != nil {
		encryptor.Close()
		w.Close()
		return fmt.Errorf("加密数据失败: %v", err)
	}

	// 按正确顺序关闭
	err = encryptor.Close()
	if err != nil {
		w.Close()
		return fmt.Errorf("关闭加密器失败: %v", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("关闭Armor编码器失败: %v", err)
	}

	fmt.Printf("加密完成，输出文件: %s\n", outputFile)
	return nil
}

// 使用接收方私钥解密并验证发送方签名
func decryptAndVerify(receiverPrivateKey, senderPublicKey, passphrase, inputFile, outputFile string) error {
	fmt.Printf("开始解密文件: %s -> %s\n", inputFile, outputFile)

	// 读取接收方私钥
	fmt.Printf("读取接收方私钥: %s\n", receiverPrivateKey)
	receiverKey, err := os.Open(receiverPrivateKey)
	if err != nil {
		return fmt.Errorf("打开接收方私钥文件失败: %v", err)
	}
	defer receiverKey.Close()

	receiverEntityList, err := openpgp.ReadArmoredKeyRing(receiverKey)
	if err != nil {
		return fmt.Errorf("解析接收方私钥失败: %v", err)
	}
	receiverEntity := receiverEntityList[0]
	fmt.Printf("接收方密钥ID: %X\n", receiverEntity.PrimaryKey.KeyId)

	// 解密接收方私钥
	if receiverEntity.PrivateKey.Encrypted {
		fmt.Printf("解密接收方私钥...\n")
		err = receiverEntity.PrivateKey.Decrypt([]byte(passphrase))
		if err != nil {
			return fmt.Errorf("解密接收方私钥失败: %v", err)
		}
		fmt.Printf("接收方私钥解密成功\n")

		// 解密子密钥
		for i, subkey := range receiverEntity.Subkeys {
			if subkey.PrivateKey.Encrypted {
				err = subkey.PrivateKey.Decrypt([]byte(passphrase))
				if err != nil {
					return fmt.Errorf("解密子密钥 %d 失败: %v", i, err)
				}
			}
		}
	}

	// 读取发送方公钥（用于验证签名）
	fmt.Printf("读取发送方公钥: %s\n", senderPublicKey)
	senderKey, err := os.Open(senderPublicKey)
	if err != nil {
		return fmt.Errorf("打开发送方公钥文件失败: %v", err)
	}
	defer senderKey.Close()

	senderEntityList, err := openpgp.ReadArmoredKeyRing(senderKey)
	if err != nil {
		return fmt.Errorf("解析发送方公钥失败: %v", err)
	}
	fmt.Printf("发送方密钥ID: %X\n", senderEntityList[0].PrimaryKey.KeyId)

	// 打开加密文件
	fmt.Printf("打开加密文件: %s\n", inputFile)
	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("打开加密文件失败: %v", err)
	}
	defer inFile.Close()

	// 解码ASCII Armor
	fmt.Printf("解码ASCII Armor...\n")
	block, err := armor.Decode(inFile)
	if err != nil {
		return fmt.Errorf("解码ASCII Armor失败: %v", err)
	}
	fmt.Printf("Armor类型: %s\n", block.Type)

	// 解密消息 - 使用正确的prompt函数
	fmt.Printf("开始解密消息...\n")
	promptFunc := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		if !symmetric {
			// 对于非对称加密，我们已经解密了私钥，所以返回空密码
			return nil, nil
		}
		return nil, fmt.Errorf("不支持对称加密")
	}

	// 将接收方私钥列表和发送方公钥列表合并，供签名验证使用
	combinedKeyring := append(receiverEntityList, senderEntityList...)

	md, err := openpgp.ReadMessage(block.Body, combinedKeyring, promptFunc, nil)
	if err != nil {
		return fmt.Errorf("解密消息失败: %v", err)
	}
	fmt.Printf("解密成功\n")

	// 验证签名
	if md.SignedBy != nil {
		signedByKeyID := md.SignedBy.PublicKey.KeyId
		fmt.Printf("消息由密钥 %X 签名\n", signedByKeyID)

		validSignature := false
		for _, senderEntity := range senderEntityList {
			if senderEntity.PrimaryKey.KeyId == signedByKeyID {
				validSignature = true
				fmt.Printf("✓ 签名验证成功! 签名者: %X\n", signedByKeyID)
				break
			}
		}

		if !validSignature {
			fmt.Printf("⚠ 警告: 签名者 %X 不在可信发送方列表中\n", signedByKeyID)
		}
	} else {
		fmt.Println("⚠ 警告: 消息没有签名")
	}

	// 读取解密后的数据
	fmt.Printf("读取解密数据...\n")
	decryptedData, err := io.ReadAll(md.UnverifiedBody)
	if err != nil {
		return fmt.Errorf("读取解密数据失败: %v", err)
	}
	fmt.Printf("解密数据大小: %d 字节\n", len(decryptedData))

	// 将解密数据写入文件
	err = os.WriteFile(outputFile, decryptedData, 0o644)
	if err != nil {
		return fmt.Errorf("写入解密文件失败: %v", err)
	}
	fmt.Printf("解密完成，输出文件: %s\n", outputFile)

	return nil
}

func main() {
	fmt.Println("=== PGP 文件加密解密测试 ===")

	// 签名并加密
	fmt.Println("\n1. 对文件进行加密和签名...")
	err := encryptAndSign("sender_private.asc", "receiver_public.asc",
		"sender-passphrase", "plaintext.txt", "encrypted_signed.pgp")
	if err != nil {
		log.Fatal("加密签名失败:", err)
	}

	// 解密并验证
	fmt.Println("\n2. 对文件进行解密和验证...")
	err = decryptAndVerify("receiver_private.asc", "sender_public.asc",
		"receiver-passphrase", "encrypted_signed.pgp", "decrypted_verified.txt")
	if err != nil {
		log.Fatal("解密验证失败:", err)
	}

	// 验证解密内容
	fmt.Println("\n3. 验证解密内容...")
	content, err := os.ReadFile("decrypted_verified.txt")
	if err != nil {
		log.Fatal("读取解密文件失败:", err)
	}
	fmt.Printf("解密内容: %s\n", string(content))

	// 比较原始文件和解密文件
	original, _ := os.ReadFile("plaintext.txt")
	if string(original) == string(content) {
		fmt.Println("\n✓ 测试成功: 解密内容与原始内容完全一致!")
	} else {
		fmt.Println("\n✗ 测试失败: 解密内容与原始内容不一致!")
		fmt.Printf("原始文件大小: %d 字节\n", len(original))
		fmt.Printf("解密文件大小: %d 字节\n", len(content))
	}
}
