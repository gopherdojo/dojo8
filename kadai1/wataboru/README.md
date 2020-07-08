# 課題 1 【TRY】画像変換コマンドを作ろう

## 次の仕様を満たすコマンドを作って下さい

- ディレクトリを指定する
- 指定したディレクトリ以下の JPG ファイルを PNG に変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）

## 以下を満たすように開発してください

- main パッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
- 準標準パッケージ：golang.org/x 以下のパッケージ
- ユーザ定義型を作ってみる
- GoDoc を生成してみる
- Go Modules を使ってみる

### 動作手順

```
$ go build -o imgconv
$ ./imgconv -h 
Usage of ./imgconv:
  -a string
    	Input extension after conversion. (short) (default "png")
  -after --after=jpg
    	Input extension after conversion.
    	  ex) --after=jpg (default "png")
  -b string
    	Input extension before conversion. (short) (default "jpg")
  -before --before=png
    	Input extension before conversion.
    	  ex) --before=png (default "jpg")
  -d string
    	Input target Directory. (short)
  -dir --dir=./convert_image
    	Input target Directory.
    	  ex) --dir=./convert_image
$ ./imgconv -d ./testdate
  or
$ ./imgconv -d ./testdate -b png -a gif
  or
$ ./imgconv -d ./testdate -b jpeg -a tiff
```

###　感想等

- 前提として、@tenntennさんの公開されている[handsonプロジェクト](https://github.com/tenntenn/gohandson/tree/master/imgconv/ja)と似ている内容でして、以前やったことがありました。
  - そちらに無い部分として、tiff変換やbmp変換の追加などを行いました。
- GoModulesを利用したく、`golang.org/x/image`を導入しました
- 様々なソースをみていると、変数名やコマンド名が自分は冗長かな？と感じています。Go話者の方は短く記載するものなのでしょうか。
