package main

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
	"time"

	"flag"
)

// Config 用于解析配置文件
type Config struct {
	Path   string `yaml:"path"`
	RowKey string `yaml:"RowKey"`
	IV     string `yaml:"IV"`
}

// 读取配置文件
func LoadConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// EncryptFile 用于加密文件
func EncryptFile(ctx context.Context, inputFile, keyHex, ivHex string) error {
	startTime := time.Now()
	defer func() {
		dur := time.Now().Sub(startTime).Milliseconds()
		fmt.Println("consume time: ", dur)
	}()

	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return err
	}
	iv, err := hex.DecodeString(ivHex)
	if err != nil {
		return err
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer func() {
		_ = inFile.Close()
	}()

	outFile, err := os.Create(inputFile + ".tmp")
	if err != nil {
		return err
	}
	defer func() {
		_ = outFile.Close()
	}()

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	stream := cipher.NewCBCEncrypter(block, iv)
	buf := make([]byte, aes.BlockSize*1024) // Use a larger buffer to reduce padding frequency

	for {
		n, err := inFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if n%aes.BlockSize != 0 {
			buf = pad(buf[:n], aes.BlockSize)
			n = len(buf)
		} else {
			buf = buf[:n]
		}

		encrypted := make([]byte, n)
		stream.CryptBlocks(encrypted, buf)
		if _, err := outFile.Write(encrypted); err != nil {
			return err
		}
		if err == io.EOF {
			break
		}
	}

	// Close files to flush all buffers
	_ = outFile.Close()
	_ = inFile.Close()

	// Rename original file to .bak
	//if err = os.Rename(inputFile, inputFile+".bak"); err != nil {
	//	return err
	//}

	// Rename temp file to original file name
	if err = os.Rename(inputFile+".tmp", inputFile); err != nil {
		return err
	}

	return nil
}

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func unpad(src []byte) ([]byte, error) {
	length := len(src)
	padding := int(src[length-1])
	if padding < 1 || padding > aes.BlockSize {
		return nil, errors.New("invalid padding size")
	}
	return src[:length-padding], nil
}

// 遍历目录并加密所有文件
func EncryptFilesInDir(config *Config) error {
	// 遍历目录
	err := filepath.Walk(config.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() { // 只处理文件
			fmt.Printf("Encrypting file: %s\n", path)
			err := EncryptFile(context.Background(), path, config.RowKey, config.IV)
			if err != nil {
				return fmt.Errorf("failed to encrypt %s: %v", path, err)
			}
		}
		return nil
	})
	return err
}

func main() {
	// 命令行参数
	configFile := flag.String("config", "config.yaml", "Path to the config.yaml file")
	flag.Parse()

	// 加载配置
	config, err := LoadConfig(*configFile)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// 加密目录中的所有文件
	err = EncryptFilesInDir(config)
	if err != nil {
		fmt.Println("Error encrypting files:", err)
	}
}
