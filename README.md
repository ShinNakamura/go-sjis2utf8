# SJIS to UTF-8

## 参考

Go言語(golang)でShiftJISのファイルをutf-8に変換する
https://dev.classmethod.jp/articles/golang-iconv/


## 準備

```sh
go mod init github.com/ShinNakamura/go-sjis2utf8
go get -u golang.org/x/text
mkdir bin
```


## usage

```sh
$0 /path/to/sjisFile /path/to/utf8File
```

`/path/to/sjisFile` は存在しないとエラー

`/path/to/utf8File` はファイルは新規作成／上書き。事前にディレクトリを作成しておくこと
