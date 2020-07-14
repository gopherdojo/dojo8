# 課題2 tsuchinaga

## 課題
スライドより転載

* io.Readerとio.Writerについて調べてみよう
    * 標準パッケージでどのように使われているか
    * io.Readerとio.Writerがあることで  
      どういう利点があるのか具体例を挙げて考えてみる
* 1回目の課題のテストを作ってみて下さい
    * テストのしやすさを考えてリファクタリングしてみる
    * テストのカバレッジを取ってみる
    * テーブル駆動テストを行う
    * テストヘルパーを作ってみる


### io.Readerとio.Writerについて調べてみよう

`io.Reader` と `io.Writer` を実装している代表的なのはやはり `os.File` だと思う。

`os.File` には `os.Stdin` 、 `os.Stdout` 、 `os.Stderr` がある。  
これらはそれぞれ、標準入力、標準出力、標準エラーにあたり、誰もが入出力先にしたことがあると思う。

例えば、「出力先を標準出力からログファイルにかえたいなー」というときに、
`io.Writer` がなければ新たに出力用の処理を用意して上げる必要があったり、
標準ライブラリが `io.Writer` を出力先に指定していなければ、標準ライブラリと同じことをするのに自前で用意する必要があったりする。  
さらに、「標準出力とログファイルの両方に出したいなー」みたいなときに、それぞれを同じように扱えるようにしていないと処理を作ることも難しい。

読み込みの代表例として、 `bufio.Scanner` をあげる。

競技プログラミングで大きな入力を受け取るときや、サイズの大きいファイルを読み込むときにお世話になる。  
※ 他にも1単語ずつじゃなくて、1行ずつ読むときとかにも重宝する

この `Scanner` を生成する関数である `NewScanner` は下記のようになっている。
```go
// NewScanner returns a new Scanner to read from r.
// The split function defaults to ScanLines.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r:            r,
		split:        ScanLines,
		maxTokenSize: MaxScanTokenSize,
	}
}
```
引数が `io.Reader` となっている。  
これは標準入力だろうが、ファイルだろうが、HTTPリクエストのBodyだろうが、なんでもかんでも同じように扱える。

同じような処理をいっぱい作らなくていい以外に、モックを作りやすいというメリットもある。(interfaceのメリットかも)

テストコードを書いていると、ここなんか適当なもので置き換えたいなーというのは頻繁にある。  
そんな時に構造体に依存していたりすると置き換えることができず、密結合が起きてしまい、テストしづらくなる。  
そこをinterfaceに依存するようにしていると、置き換えることが容易になる。


### 1回目の課題のテストを作ってみて下さい

リファクタリングの域をこえてるきもするけど、いつもやってるようにTODO出してテストを作っていく  
課題1の時よりも責務を意識してTODOを分ける

* [x] ディレクトリが指定でき、その配下のディレクトリとファイルを再帰的に探索し、変換前として指定されたフォーマットの画像ファイルを変換後として指定されたフォーマットの画像ファイルに変換する
    * [x] コマンドのパラメータを指定する
        * [x] ベースのディレクトリを指定する
            * [x] 必須
        * [x] 変換前のファイルフォーマットを指定する
            * [x] 任意
            * [x] デフォルトjpeg
            * [x] 許容されるのはjpeg/png
        * [x] 変換後のファイルフォーマットを指定する
            * [x] 任意
            * [x] デフォルトpng
            * [x] 許容されるのはjpeg/png
            * [x] 変換前と同じフォーマットの指定は不可
    * [x] 指定されたディレクトリの配下のディレクトリとファイルの一覧を取得する
        * [x] 配下にディレクトリがあればさらに読み込めるようにする
        * [x] 指定されたパスがディレクトリでなければ中止
            * [x] ディレクトリかを判断する
    * [x] ファイルの画像フォーマットを取得する
        * [x] ファイルが画像なら画像フォーマットを取得する
        * [x] ファイルが画像でなければフォーマットは取得できないので中止
    * [x] ファイルフォーマットを変換する


#### カバレッジ
```bash
$ go tool cover -func=cover.out
github.com/gopherdojo/dojo8/kadai2/tsuchinaga/imgconv/imgconv.go:27:            NewIMGConverter 100.0%
github.com/gopherdojo/dojo8/kadai2/tsuchinaga/imgconv/imgconv.go:42:            Do              93.3%
github.com/gopherdojo/dojo8/kadai2/tsuchinaga/imgconv/imgconv.go:74:            NewConverter    100.0%
github.com/gopherdojo/dojo8/kadai2/tsuchinaga/imgconv/imgconv.go:90:            IsDir           100.0%
github.com/gopherdojo/dojo8/kadai2/tsuchinaga/imgconv/imgconv.go:99:            GetIMGType      90.0%
github.com/gopherdojo/dojo8/kadai2/tsuchinaga/imgconv/imgconv.go:118:           DirFileList     100.0%
github.com/gopherdojo/dojo8/kadai2/tsuchinaga/imgconv/imgconv.go:137:           Convert         69.2%
github.com/gopherdojo/dojo8/kadai2/tsuchinaga/validation/validation.go:4:       NewValidator    100.0%
github.com/gopherdojo/dojo8/kadai2/tsuchinaga/validation/validation.go:19:      IsValidDir      100.0%
github.com/gopherdojo/dojo8/kadai2/tsuchinaga/validation/validation.go:26:      IsValidFileType 100.0%
github.com/gopherdojo/dojo8/kadai2/tsuchinaga/validation/validation.go:31:      IsValidSrc      100.0%
github.com/gopherdojo/dojo8/kadai2/tsuchinaga/validation/validation.go:36:      IsValidDest     100.0%
total:                                                                          (statements)    86.1%
```
