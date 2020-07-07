# imgconv

imgconv is an image converter.

# Install

Use go get to install this package:

```
$ go get github.com/gopherdojo/dojo8/kadai1/sadah/imgconv
```

# Build


```
$ git clone git@github.com:gopherdojo/dojo8.git
$ cd dojo8/kadai1/sadah/cmd/imgconv
$ go build -o imgconv
```

# Usage

```
./imgconv [options...] <path>
```

To check the all available options,

```
$ imgconv -h
imgconv is an image converter
Usage: imgconv [options...] [path]
Use "imgconv --help" for more information about a command.

Supported formats: [.jpg .jpeg .JPG .JPEG .png .PNG .gif .GIF .bmp .BMP .tiff .TIFF]

  -s string
        Set a source extension (default "jpg")
  -source string
        Set a source extension (default "jpg")
  -t string
        Set a target extension (default "png")
  -target string
        Set a target extension (default "png")
  -v    Show version
  -verbose
        Print verbose messages
  -version
        Show version
```

# Spec

## Required Spec

次の仕様を満たすコマンドを作って下さい

* [x] ディレクトリを指定する
* [x] 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
* [x] ディレクトリ以下は再帰的に処理する
* [x] 変換前と変換後の画像形式を指定できる（オプション）

以下を満たすように開発してください
* [x] mainパッケージと分離する
* [x] 自作パッケージと標準パッケージと準標準パッケージのみ使う
    * 準標準パッケージ：golang.org/x以下のパッケージ
* [x] ユーザ定義型を作ってみる
* [x] GoDocを生成してみる
* [x] Go Modulesを使ってみる

## Featured Spec

* コマンドライン引数で PATH を指定。指定がない場合はカレントディレクトリが対象となる
* ディレクトリ以下は再帰的に処理する
* デフォルトでは指定したディレクトリ以下のJPGファイルをPNGに変換
* `-s` で変換前、 `-t` 変換後の画像形式を指定できる
    * jpg, png, gif, bmp, tiffに対応
* main パッケージと imgconv パッケージを分離
* 自作パッケージと標準パッケージと準標準パッケージのみ使う
* 拡張子リストに `exts` というユーザ定義型を利用
* GoDocを生成した
* Go Modulesを使った

# 感想

`type` をどういったときに使うとよいのか、まだいまひとつわかっていない。


struct の初期化で、flagから受け取った値を使いたかったが、よい使い方が思いつかなかった。

愚直に代入すればできるけど、ポインタをうまく使う方法がわからなかった。以下の `&verbose` を struct に定義したかった…。

```go
flag.BoolVar(&verbose, "verbose", false, optVerboseText)
```

# Auther

@sadah