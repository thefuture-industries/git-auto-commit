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
	buffer := make([]byte, 32*1024)

	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			nw, ew := out.Write(buffer[0:n])
			if nw > 0 {
				downloaded += int64(nw)
			}
		}
	}

	_, err = io.Copy(out, resp.Body)
	return err
}
