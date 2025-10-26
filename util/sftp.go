package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SFTPConfig SFTP连接配置
type SFTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

// SFTPClient SFTP客户端封装
type SFTPClient struct {
	config     *SFTPConfig
	sshClient  *ssh.Client
	sftpClient *sftp.Client
}

// NewSFTPClient 创建SFTP客户端
func NewSFTPClient(config *SFTPConfig) (*SFTPClient, error) {
	sshConfig := &ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 生产环境应验证主机密钥
		Timeout:         30 * time.Second,
	}

	// 建立SSH连接
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	sshClient, err := ssh.Dial("tcp", address, sshConfig)
	if err != nil {
		return nil, fmt.Errorf("SSH连接失败: %v", err)
	}

	// 创建SFTP客户端
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		sshClient.Close()
		return nil, fmt.Errorf("SFTP客户端创建失败: %v", err)
	}

	return &SFTPClient{
		config:     config,
		sshClient:  sshClient,
		sftpClient: sftpClient,
	}, nil
}

// Close 关闭连接
func (c *SFTPClient) Close() error {
	if c.sftpClient != nil {
		c.sftpClient.Close()
	}
	if c.sshClient != nil {
		c.sshClient.Close()
	}
	return nil
}

// UploadFile 上传文件到SFTP服务器
func (c *SFTPClient) UploadFile(localPath, remotePath string) error {
	// 打开本地文件
	localFile, err := os.Open(localPath)
	if err != nil {
		return fmt.Errorf("打开本地文件失败: %v", err)
	}
	defer localFile.Close()

	// 创建远程文件
	remoteFile, err := c.sftpClient.Create(remotePath)
	if err != nil {
		return fmt.Errorf("创建远程文件失败: %v", err)
	}
	defer remoteFile.Close()

	// 复制文件内容
	_, err = io.Copy(remoteFile, localFile)
	if err != nil {
		return fmt.Errorf("文件上传失败: %v", err)
	}

	fmt.Printf("文件上传成功: %s -> %s\n", localPath, remotePath)
	return nil
}

// DownloadFile 从SFTP服务器下载文件
func (c *SFTPClient) DownloadFile(remotePath, localPath string) error {
	// 打开远程文件
	remoteFile, err := c.sftpClient.Open(remotePath)
	if err != nil {
		return fmt.Errorf("打开远程文件失败: %v", err)
	}
	defer remoteFile.Close()

	// 创建本地文件
	localFile, err := os.Create(localPath)
	if err != nil {
		return fmt.Errorf("创建本地文件失败: %v", err)
	}
	defer localFile.Close()

	// 复制文件内容
	_, err = io.Copy(localFile, remoteFile)
	if err != nil {
		return fmt.Errorf("文件下载失败: %v", err)
	}

	fmt.Printf("文件下载成功: %s -> %s\n", remotePath, localPath)
	return nil
}

// ListFiles 列出远程目录文件
func (c *SFTPClient) ListFiles(remotePath string) ([]string, error) {
	files, err := c.sftpClient.ReadDir(remotePath)
	if err != nil {
		return nil, fmt.Errorf("读取远程目录失败: %v", err)
	}

	var fileList []string
	for _, file := range files {
		fileType := "文件"
		if file.IsDir() {
			fileType = "目录"
		}
		fileInfo := fmt.Sprintf("%s [%s] %d bytes", file.Name(), fileType, file.Size())
		fileList = append(fileList, fileInfo)
	}

	return fileList, nil
}

// CreateDirectory 创建远程目录
func (c *SFTPClient) CreateDirectory(remotePath string) error {
	err := c.sftpClient.MkdirAll(remotePath)
	if err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}
	fmt.Printf("目录创建成功: %s\n", remotePath)
	return nil
}

// DeleteFile 删除远程文件
func (c *SFTPClient) DeleteFile(remotePath string) error {
	err := c.sftpClient.Remove(remotePath)
	if err != nil {
		return fmt.Errorf("删除文件失败: %v", err)
	}
	fmt.Printf("文件删除成功: %s\n", remotePath)
	return nil
}

// 使用示例
func main() {
	// SFTP配置
	config := &SFTPConfig{
		Host:     "127.0.0.1",  // 替换为你的SFTP服务器地址
		Port:     2222,         // 替换为你的SFTP服务器端口
		Username: "myuser",     // 替换为你的用户名
		Password: "mypassword", // 替换为你的密码
	}

	// 创建SFTP客户端
	client, err := NewSFTPClient(config)
	if err != nil {
		log.Fatalf("创建SFTP客户端失败: %v", err)
	}
	defer client.Close()

	fmt.Println("SFTP连接成功!")

	// 示例操作
	examples := []func(*SFTPClient) error{
		exampleListFiles,
		exampleCreateDirectory,
		exampleUploadFile,
		exampleDownloadFile,
		exampleDeleteFile,
	}

	for i, example := range examples {
		fmt.Printf("\n=== 执行示例 %d ===\n", i+1)
		if err := example(client); err != nil {
			fmt.Printf("示例执行失败: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
}

// 示例函数
func exampleListFiles(client *SFTPClient) error {
	files, err := client.ListFiles(".")
	if err != nil {
		return err
	}

	fmt.Println("远程目录文件列表:")
	for _, file := range files {
		fmt.Printf("  %s\n", file)
	}
	return nil
}

func exampleCreateDirectory(client *SFTPClient) error {
	testDir := "test_directory"
	return client.CreateDirectory(testDir)
}

func exampleUploadFile(client *SFTPClient) error {
	// 创建一个测试文件用于上传
	testContent := "Hello, SFTP! This is a test file.\n "
	localFile := "test_upload.txt"

	if err := os.WriteFile(localFile, []byte(testContent), 0o644); err != nil {
		return fmt.Errorf("创建测试文件失败: %v", err)
	}
	defer os.Remove(localFile) // 清理测试文件

	remoteFile := "test_directory/uploaded_file.txt"
	return client.UploadFile(localFile, remoteFile)
}

func exampleDownloadFile(client *SFTPClient) error {
	remoteFile := "test_directory/uploaded_file.txt"
	localFile := "downloaded_file.txt"

	if err := client.DownloadFile(remoteFile, localFile); err != nil {
		return err
	}
	defer os.Remove(localFile) // 清理下载的文件

	// 读取下载的文件内容验证
	content, err := os.ReadFile(localFile)
	if err != nil {
		return err
	}
	fmt.Printf("下载文件内容: %s", string(content))
	return nil
}

func exampleDeleteFile(client *SFTPClient) error {
	remoteFile := "test_directory/uploaded_file.txt"
	return client.DeleteFile(remoteFile)
}
