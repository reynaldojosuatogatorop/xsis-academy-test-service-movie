package helper

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func SaveImageToLocalDrive(image multipart.FileHeader, parentPath string, subpath string) (string, error) {
	// Membuka file gambar
	src, err := image.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Membuat path penyimpanan
	fileName := strings.ToLower(filepath.Base(image.Filename))
	dstPath := filepath.Join(parentPath, subpath, fileName)
	savePathDB := filepath.Join(subpath, fileName)

	// Membuka file tujuan untuk penulisan
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Menyalin isi file gambar ke file tujuan
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	// Mengembalikan path gambar yang disimpan
	return savePathDB, nil
}
