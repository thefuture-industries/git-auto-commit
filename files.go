package main

import (
	"io"
	"net/http"
	"os"
)

func DownloadBinAutoCommit(url, destPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	size := resp.ContentLength
	if size <= 0 {
		_, err := io.Copy(out, resp.Body)
		return err
	}

	const progressBarWidth = 70
	var downloaded int64 = 0
	buffer := make([]byte, 321025)

	_, err = io.Copy(out, resp.Body)
	return err
}
