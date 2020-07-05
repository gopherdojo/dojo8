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

## 回答・動作例

- 自作パッケージとして　 github.com/gopherdojo/dojo8/kadai1/segakazzz/imgconv を作成
- imgconv.RunConverter()でメインの処理を実行
- -d オプションで指定したディレクトリの画像をソースとして使用する
- -d オプションで指定したディレクトリ内に out フォルダが作成され、出力される
- -i オプションで入力画像の  拡張子を指定可能(jpg or png) デフォルトは jpg
- -o オプションで出力画像の拡張子を指定可能(png or jpg) デフォルトは png
- (補足) image フォルダ内の download.zsh 実行して、[The Cat API](https://thecatapi.com/) から jpg ファイルを 10 個ダウンロード可能

### 動作例 main.go

```
$ go build -o kadai1-segakazzz
$ ./kadai1-segakazzz -d [imagedir] -i [jpg|png] -o [png|jpg]
```

```
package main

import "github.com/gopherdojo/dojo8/kadai1/segakazzz/imgconv"

func main() {
    imgconv.RunConverter()
}

```

###　感想等

- Go Modules については、使用するパッケージのバージョン管理に使用されるものと理解しましたが、今回使用したパッケージは標準パッケージのみで、バージョンを指定する必要はなさそうでしたので requires の部分は書いていません。
