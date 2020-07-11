# 画像変換コマンドを作ろう

## 概要

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

# Usage

## build

```
cd kynefuk
go build -o imgconverter
```

## command

```
imgconverter -d [directory] -f [from format] -t [to format]
```
