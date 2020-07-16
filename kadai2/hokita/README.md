# 課題 1 画像変換コマンドを作ろう

## 課題内容
### 次の仕様を満たすコマンドを作って下さい

- ディレクトリを指定する
- 指定したディレクトリ以下の JPG ファイルを PNG に変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）

### 以下を満たすように開発してください

- main パッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
- 準標準パッケージ：golang.org/x 以下のパッケージ
- ユーザ定義型を作ってみる
- GoDoc を生成してみる
- Go Modules を使ってみる

## 対応したこと
- 画像を変換
  - 現状はjpg, pngのみ
  - jpg, png以外はエラー表示
  - 画像出力先は対象画像と同じディレクトリ
- 指定したディレクトリが無いとエラーを表示

## 動作
```shell
$ go build -o test_imgconv

$ ./test_imgconv -h
Usage of ./test_imgconv:
  -from string
        Conversion source extension. (default "jpg")
  -to string
        Conversion target extension. (default "png")

# testdata内のすべてのjpgファイルをpngに変換する
$ ./test_imgconv testdata
Conversion finished!

# testdata内のすべてのpngファイルをjpgに変換する
$ ./test_imgconv -from png -to jpg testdata
Conversion finished!

# ディレクトリの指定が無い場合はエラー
$ ./test_imgconv
Please specify a directory.

# 存在しないディレクトリの場合はエラー
$ ./test_imgconv non_exist_dir
Cannot find directory.

# 対応していない拡張子の場合はエラー
$ ./test_imgconv -from txt -to jpg testdata
Selected extension is not supported.
```

## 工夫したこと
- png, jpg以外にも拡張子が増えそうなので、`image_type`というinterfaceを作ってみた。
- 拡張子の微妙な違い（jpg, jpeg, JPGなど）にも対応できるようにした。

## わからなかったこと、むずかしかったこと
- go mod initで指定するmodule名に命名規則があるのか。
- 普段オブジェクト指向（その上動的型付け言語）で書いているので、それがgoらしいコードになっているのか不安。
  - なんでもかんでも構造体メソッドにしたい願望がでてくる
