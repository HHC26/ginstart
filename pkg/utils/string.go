package utils

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

// BatchSearchToStringSlice 批量搜索对字符串的分割
func BatchSearchToStringSlice(text string) []string {
	var nilSlice []string
	if text == "" {
		return nilSlice
	}
	res := strings.Split(text, ",")
	if len(res) > 1 {
		return res
	}

	res = strings.Split(text, " ")
	if len(res) > 1 {
		return res
	}

	return []string{text}
}

func NewMkdir(path string) string {
	floderName := time.Now().Format(time.DateOnly)
	floderPath := filepath.Join(path, floderName)
	err := os.MkdirAll(floderPath, os.ModePerm)
	if err != nil {
		return ""
	}
	return floderPath
}
