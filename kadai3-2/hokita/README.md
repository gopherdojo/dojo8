# 課題3-2
## 分割ダウンローダを作ろう
- 分割ダウンロードを行う
  - Rangeアクセスを用いる
  - いくつかのゴルーチンでダウンロードしてマージする
  - エラー処理を工夫する
  - golang.org/x/sync/errgourpパッケージなどを使ってみる
  - キャンセルが発生した場合の実装を行う

### 動作
```shell
$ go build -o pdl cmd/pdl/main.go

$ ./pdl -proc 10 https://blog.golang.org/gopher/header.jpg
start download worker: 1
start download worker: 10
start download worker: 4
start download worker: 2
start download worker: 5
start download worker: 7
start download worker: 9
start download worker: 6
start download worker: 3
start download worker: 8
finish download worker: 4
finish download worker: 8
finish download worker: 5
finish download worker: 3
finish download worker: 7
finish download worker: 9
finish download worker: 1
finish download worker: 10
finish download worker: 2
finish download worker: 6
finished

$ ls testdata/header.jpg
testdata/header.jpg
```

## わからなかったこと、むずかしかったこと
- そもそもurlからダウンロードをどう実現するのかを考えるのに時間がかかった。
  - `pget`を参考にした。
    - cf. https://github.com/Code-Hex/pget
  - 結局はurlでアクセスして読み込んだ情報(`io.Reader`)をファイルに書き込む(`io.writer`)だけだった。
