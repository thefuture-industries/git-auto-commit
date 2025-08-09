package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

	if errProgress := DownloadProgress(resp, out); errProgress != nil {
		return errProgress
	}

	_, err = io.Copy(out, resp.Body)
	return err
}

func DownloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func DownloadProgress(resp *http.Response, file *os.File) error {
	size := resp.ContentLength
	if size <= 0 {
		_, err := io.Copy(file, resp.Body)
		return err
	}

	const progressBarWidth = 70
	var downloaded int64 = 0
	buffer := make([]byte, 32*1024)

	for {
		n, errRead := resp.Body.Read(buffer)
		if n > 0 {
			nw, ew := file.Write(buffer[0:n])
			if nw > 0 {
				downloaded += int64(nw)
				progress := float64(downloaded) * 100 / float64(size)
				hashes := int(progress * progressBarWidth)

				fmt.Printf("\r[%-*s] %3.0f%%", progressBarWidth, strings.Repeat("#", hashes), progress*100)
			}

			if ew != nil {
				return ew
			}

			if n != nw {
				return io.ErrShortWrite
			}
		}

		if errRead != nil {
			if errRead == io.EOF {
				break
			}

			return errRead
		}
	}

	fmt.Println()
	return nil
}
