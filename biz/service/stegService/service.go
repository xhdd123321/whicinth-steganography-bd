package stegService

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/stegify"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"
)

// EncodeImage 图片中加密图片
func EncodeImage(carrierFile string, dataFile string, outFile string) error {
	fByte, err := os.ReadFile(carrierFile)
	if err != nil {
		return fmt.Errorf("read carrierFile [%s] failed, err: %v", carrierFile, err)
	}
	carrierReader := bytes.NewReader(fByte)
	cByte, err := os.ReadFile(dataFile)
	if err != nil {
		return fmt.Errorf("read dataFile [%s] failed, err: %v", dataFile, err)
	}
	dataReader := bytes.NewReader(cByte)
	outWriter, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("create failed, err: %v", err)
	}
	// 加密
	err = stegify.Encode(carrierReader, dataReader, outWriter)
	if err != nil {
		return fmt.Errorf("encode failed, err: %v", err)
	}
	return nil
}

// DecodeImage 图片中解密图片
func DecodeImage(decodeFile string) (string, error) {
	fByte, err := os.ReadFile(decodeFile)
	if err != nil {
		return "", fmt.Errorf("read file [%s] failed, err: %v", decodeFile, err)
	}
	decodeReader := bytes.NewReader(fByte)

	resFile := utils.GetExtractedFilename(decodeFile)
	outWriter, err := os.Create(resFile)
	if err != nil {
		return "", fmt.Errorf("create failed, err: %v", err)
	}
	// 解密
	err = stegify.Decode(decodeReader, outWriter)
	if err != nil {
		return "", fmt.Errorf("decode failed, err: %v", err)
	}
	return resFile, nil
}

// EncodeDoc 图片中加密文字信息
func EncodeDoc(carrierFile string, data string, outFile string) error {
	fByte, err := os.ReadFile(carrierFile)
	if err != nil {
		return fmt.Errorf("read file [%s] failed, err: %v", carrierFile, err)
	}
	carrierReader := bytes.NewReader(fByte)
	dataReader := strings.NewReader(data)
	outWriter, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("create failed, err: %v", err)
	}

	// 加密
	err = stegify.Encode(carrierReader, dataReader, outWriter)
	if err != nil {
		return fmt.Errorf("encode failed, err: %v", err)
	}
	return nil
}

// DecodeDoc 图片中解密文字信息
func DecodeDoc(decodeFile string) (string, error) {
	var res string
	fByte, err := os.ReadFile(decodeFile)
	if err != nil {
		return res, fmt.Errorf("read file [%s] failed, err: %v\n", decodeFile, err)
	}
	decodeReader := bytes.NewReader(fByte)

	originData := bytes.NewBufferString("")
	// 解密
	err = stegify.Decode(decodeReader, originData)
	if err != nil {
		return res, fmt.Errorf("decode failed, err: %v\n", err)
	}
	res = originData.String()
	return res, nil
}
