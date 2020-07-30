# 課題1 画像変換コマンドを作ろう

## 課題仕様

- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）

## 課題条件

- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
  - 準標準パッケージ：golang.org/x以下のパッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる
- Go Modulesを使ってみる

## どう動かすのか？
```shell script
$ go build -o conv

$ ./conv -h
Usage of ./conv:
  -fdir string
    	変換する画像の保存しているディレクトリ
  -ft string
    	変換する画像の形式 (gif, jpeg, jpg, png, bmp, tiff) (default "jpeg")
  -tdir string
    	変換後の画像を保存するディレクトリ
  -tt string
    	変換後の画像の形式 (gif, jpeg, jpg, png, bmp, tiff) (default "png")

# hogeディレクトリのjpeg形式のファイルをfugaディレクトにpng形式ファイルに変換して出力する。
$ ./conv -ft=jpg -fdir=./hoge -tt=png -tdir=./fuga
Ok!
```

## 悩んだ事
- とりあえず、どんなディレクトリ構成がいいのか？で悩みました。
- もう少しスマートなエラー処理が書けないものかと…。
- やはりGoらしい書き方ってとは？