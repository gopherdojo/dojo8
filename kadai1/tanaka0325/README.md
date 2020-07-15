# Image Converter

## Spec

```
## 次の仕様を満たすコマンドを作って下さい

- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）

## 以下を満たすように開発してください

- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
- 準標準パッケージ：golang.org/x以下のパッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる
- Go Modulesを使ってみる
```

## Usage

```zsh
# build
$ go build -o imgconv ./cmd/imgconv

# display help
$ ./imgconv -h
Usage of ./imgconv:
  -f string
        file extention before convert (default "jpg")
  -n    dry run
  -t string
        file extention after convert (default "png")

# single directory
$ ./imgconv testdata/images

# multi directories
$ ./imgconv testdata/images testdata/images2

# customize ext
$ ./imgconv -f png -t gif testdata/images

# dry run
$ ./imgconv -n testdata/images testdata/images2
testdata/images/sample1.jpg => testdata/images/sample1.png
testdata/images2/img/sample3.jpg => testdata/images2/img/sample3.png
testdata/images2/sample2.jpg => testdata/images2/sample2.png

```

## 感想

- long option(?) はどうやってやれば良いのだろうか
- そもそも作りとしてこれで良いのだろうか・・・めっちゃ悩みました

