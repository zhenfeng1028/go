package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	mchID             string
	mchSerialNo       string
	mchPrivateKeyPath string
	mchAPIv3Key       string

	wechatPayCertificatePath string
	outputPath               string
)

const (
	errCodeParamError = 1
	errCodeRunError   = 2
)

func init() {
	flag.StringVar(&mchID, "m", "", "【必传】`商户号`")
	flag.StringVar(&mchSerialNo, "s", "", "【必传】`商户证书序列号`")
	flag.StringVar(&mchPrivateKeyPath, "p", "", "【必传】`商户私钥路径`")
	flag.StringVar(&mchAPIv3Key, "k", "", "【必传】`商户APIv3密钥`")

	flag.StringVar(&wechatPayCertificatePath, "c", "", "【可选】`微信支付平台证书路径`，用于验签。省略则跳过验签")
	flag.StringVar(&outputPath, "o", "./", "【可选】`证书下载保存目录`")
}

func main() {
	printUsageAndExit()
}

func printUsageAndExit() {
	_, _ = fmt.Fprintf(os.Stderr, "usage of wechatpay_download_certs:\n")
	flag.PrintDefaults()
	os.Exit(errCodeParamError)
}
