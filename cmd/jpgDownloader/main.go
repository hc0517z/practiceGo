package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func ensureDir(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, s := os.Stat(dirName); s != nil {
		m := os.MkdirAll(dirName, os.ModePerm)
		if m != nil {
			panic(m)
		}
	}
}

func download(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	filename, err := urlToFilename(url)
	if err != nil {
		return "", err
	}

	filename = "output/" + filename
	ensureDir(filename)

	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	_, err = io.Copy(f, resp.Body)
	return filename, err
}

func urlToFilename(rawUrl string) (string, error) {
	parse, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	return filepath.Base(parse.Path), nil
}

func writeZip(outFilename string, filenames []string) error {

	ensureDir(outFilename)

	var files []*os.File

	defer func() {
		for _, file := range files {
			err := file.Close()
			if err != nil {
				return
			}
		}
	}()

	outFile, err := os.Create(outFilename)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(outFile)
	for _, filename := range filenames {
		w, err := zw.Create(filename)
		if err != nil {
			return err
		}
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		files = append(files, f)

		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}
	return zw.Close()
}

func main() {
	runtime.GOMAXPROCS(4)
	var wait sync.WaitGroup
	var urls = []string{
		"http://xkxqjlzvieat874751.gcdn.ntruss.com/2/2019/d265/2d2651001bb575d64812b398661b39589500a9084c38a772f4b409035f74bf4e5_o_st.jpg",
		"https://file.mk.co.kr/meet/neds/2019/06/image_readtop_2019_441884_15610734753796599.jpg",
		"https://t1.daumcdn.net/news/201806/15/seouleconomy/20180615151617982vtvw.jpg",
	}

	for _, u := range urls {
		wait.Add(1)
		go func(url string) {
			defer wait.Done()
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(u)
	}
	wait.Wait()

	filenames, err := filepath.Glob("output/*.jpg")
	if err != nil {
		log.Fatal(err)
	}
	err = writeZip("output/kei_img.zip", filenames)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE!")
}
