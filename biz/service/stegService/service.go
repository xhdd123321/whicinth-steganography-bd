package stegService

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/stegify"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"
)

func EncodeImage(carrierFile string, dataFile string, outFile string) error {
	fByte, err := os.ReadFile(carrierFile)
	if err != nil {
		fmt.Printf("read carrierFile [%s] failed, err: %v\n", carrierFile, err)
		return err
	}
	carrierReader := bytes.NewReader(fByte)
	cByte, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Printf("read dataFile [%s] failed, err: %v\n", dataFile, err)
		return err
	}
	dataReader := bytes.NewReader(cByte)
	outWriter, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("Create failed, err: %v\n", err)
		return err
	}

	// 加密
	err = stegify.Encode(carrierReader, dataReader, outWriter)
	if err != nil {
		fmt.Printf("encode failed, err: %v\n", err)
		return err
	}
	return nil
}

func DecodeImage(decodeFile string) (string, error) {
	fByte, err := os.ReadFile(decodeFile)
	if err != nil {
		fmt.Printf("read file [%s] failed, err: %v\n", decodeFile, err)
		return "", err
	}
	decodeReader := bytes.NewReader(fByte)

	resFile := utils.GetExtractedFilename(decodeFile)
	outWriter, err := os.Create(resFile)
	if err != nil {
		fmt.Printf("Create failed, err: %v\n", err)
		return "", err
	}
	// 解密
	err = stegify.Decode(decodeReader, outWriter)
	if err != nil {
		fmt.Printf("decode failed, err: %v\n", err)
		return "", err
	}
	return resFile, err
}

func EncodeDoc(carrierFile string, data string, outFile string) error {
	fByte, err := os.ReadFile(carrierFile)
	if err != nil {
		fmt.Printf("read file [%s] failed, err: %v\n", carrierFile, err)
		return err
	}
	carrierReader := bytes.NewReader(fByte)
	dataReader := strings.NewReader(data)
	outWriter, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("Create failed, err: %v\n", err)
		return err
	}

	// 加密
	err = stegify.Encode(carrierReader, dataReader, outWriter)
	if err != nil {
		fmt.Printf("encode failed, err: %v\n", err)
		return err
	}
	return nil
}

func DecodeDoc(decodeFile string) (string, error) {
	var res string
	fByte, err := os.ReadFile(decodeFile)
	if err != nil {
		fmt.Printf("read file [%s] failed, err: %v\n", decodeFile, err)
		return res, err
	}
	decodeReader := bytes.NewReader(fByte)

	originData := bytes.NewBufferString("")
	// 解密
	err = stegify.Decode(decodeReader, originData)
	if err != nil {
		fmt.Printf("decode failed, err: %v\n", err)
		return res, err
	}
	res = originData.String()
	return res, err
}
