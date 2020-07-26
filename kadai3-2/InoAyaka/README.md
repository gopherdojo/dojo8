# ＜課題3-2＞【TRY】分割ダウンローダを作ろう
## 概要

- Rangeアクセスを用いる
- いくつかのゴルーチンでダウンロードしてマージする
- エラー処理を工夫する
-- golang.org/x/sync/errgourpパッケージなどを使ってみる
- キャンセルが発生した場合の実装を行う



## 使い方
```
$ go build -o downloader
$
$ ./downloader -h
Usage of ./downloader:
  -URL string
        URL to download
  -dir string
        Download directory (default "./")
  -div uint
        Specifying the number of divisions (default 5)
$
$
```

実行例
```
$ ./downloader -dir testdata -URL https://golang.org/doc/gopher/bumper640x360.png
checkDir(testdata): OK
checkURL(https://golang.org/doc/gopher/bumper640x360.png): OK
--------------------------------------------------
finished download [05]　: bytes  33608 - 42013
finished download [02]　: bytes   8402 - 16803
finished download [04]　: bytes  25206 - 33607
finished download [01]　: bytes      0 -  8401
finished download [03]　: bytes  16804 - 25205
--------------------------------------------------
download complete

```

### 感想

- 個人的には今までの課題の中で難易度は高いように感じた
- 分割ダウンローダーというものをどういう構成でプログラムすれば実現できるのか、という点から非常に時間がかかってしまった
- [Code-Hex/pget](https://github.com/Code-Hex/pget) を参考に作成した
- 参考にするために"Code-Hex/pget"のコードを読み、読む力についても、力をつける必要があると感じた
