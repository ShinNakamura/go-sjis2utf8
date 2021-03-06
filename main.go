// 参考: https://dev.classmethod.jp/articles/golang-iconv/
// 今のところほぼそのまま
package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	infGlob := os.Args[1] // glob が渡されたら展開して最初のファイルだけ採用
	outf := os.Args[2] // glob 不可

	infs, err := filepath.Glob(infGlob)
	if err != nil {
		log.Fatal(err)
	}

	// ShiftJISファイルを開く
	inf := infs[0]
	sjisFile, err := os.Open(inf)
	defer sjisFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// ShiftJISのデコーダーを噛ませたReaderを作成する
	reader := transform.NewReader(sjisFile, japanese.ShiftJIS.NewDecoder())

	// 書き込み先ファイルを用意
	utf8File, err := os.Create(outf)
	defer utf8File.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 書き込み
	tee := io.TeeReader(reader, utf8File)
	s := bufio.NewScanner(tee)
	for s.Scan() {
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	// log.Println("done")
}
