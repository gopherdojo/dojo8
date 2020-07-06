# ＜課題1＞【TRY】画像変換コマンドを作ろう
## 概要
### 次の仕様を満たすコマンドを作って下さい

- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）


### 以下を満たすように開発してください

- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
- 準標準パッケージ：golang.org/x以下のパッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる
- Go Modulesを使ってみる


## 使い方
```
$ go build -o imageConv
$
$ ./imageConv -h
Usage of ./imageConv:
  -af string
        Image format after conversion (default "png")
  -bf string
        Image format before conversion (default "jpg")
  -dir string
        Directory containing images to be converted
$
$
```

実行例
```
$ ./imageConv -bf=png -af=gif -dir=./image
```

### 感想

-  mainパッケージと自作パッケージの分離をどこにするのか悩みました。（画像変換部分のみを切り出して自作パッケージにしましたが、ディレクトリを渡して検索〜変換までの一連の処理をパッケージに纏めた方が良かったのかもしれないと考えています。）
- errの扱いについても、errを返してmainで一貫して処理を行うのか、それぞれで行うべきなのか悩んだ点になります。
- 画像変換を行うImageConverter()での、errチェックを行う部分が適切かも少し不安な点ではあります。
