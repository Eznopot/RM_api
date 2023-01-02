package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"time"

	logger "github.com/Eznopot/RM_api/src/Logger"
)

func ZipFiles(month int, files []string) (string, bool) {
	now := time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	filename := fmt.Sprintf("CRA_RMS_%s.zip", now.Month())
	zipFile, err := os.Create(filename)
	if err != nil {
		logger.Error(err.Error())
		return "", false
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		zipfile, err := os.Open(file)
		if err != nil {
			logger.Error(err.Error())
			return "", false
		}
		defer zipfile.Close()
		info, err := zipfile.Stat()
		if err != nil {
			logger.Error(err.Error())
			return "", false
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			logger.Error(err.Error())
			return "", false
		}
		header.Method = zip.Store

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			logger.Error(err.Error())
			return "", false
		}
		_, err = io.Copy(writer, zipfile)
		if err != nil {
			logger.Error(err.Error())
			return "", false
		}
	}
	return filename, true
}
