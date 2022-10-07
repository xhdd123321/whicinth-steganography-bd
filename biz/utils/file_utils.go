package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// GetFileHash 计算文件Hash(md5)
func GetFileHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", err
	}
	hash := h.Sum(nil)
	hashValue := hex.EncodeToString(hash)
	return hashValue, nil
}

// GetExtractedFilename 获取被提取文件名
func GetExtractedFilename(origin string) string {
	dir := filepath.Dir(origin)
	ext := filepath.Ext(origin)
	fileNameOnly := strings.TrimSuffix(filepath.Base(origin), ext)
	res := filepath.Join(dir, fileNameOnly+"_extracted"+ext)
	return res
}

// GetMediaAbPath 获取media文件夹绝对路径
func GetMediaAbPath() string {
	return filepath.Join(GetCurrentAbPath(), "media")
}

// GetConfAbPath 获取conf文件夹绝对路径
func GetConfAbPath() string {
	return filepath.Join(GetCurrentAbPath(), "conf")
}

// GetCurrentAbPath 最终方案-全兼容
func GetCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	if strings.Contains(dir, getTmpDir()) {
		dir = getCurrentAbPathByCaller()
	}
	dir = strings.ReplaceAll(dir, "/biz/utils", "")
	dir = strings.ReplaceAll(dir, "\\output", "")
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
