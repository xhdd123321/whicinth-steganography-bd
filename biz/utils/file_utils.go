package utils

import (
	"path/filepath"
	"strings"
)

func GetExtractedFilename(origin string) string {
	dir := filepath.Dir(origin)
	ext := filepath.Ext(origin)
	fileNameOnly := strings.TrimSuffix(filepath.Base(origin), ext)
	res := filepath.Join(dir, fileNameOnly+"_extracted"+ext)
	return res
}
