# 課題2
## 【TRY】io.Readerとio.Writer
### io.Readerとio.Writerについて調べてみよう
#### 標準パッケージでどのように使われているか
* データの入出力処理を抽象化するために `io.Reader`, `io.Writer` は使われている
* `*os.File`, `*bytes.Buffer`, `net.Conn` などが `io.Reader`, `io.Writer` を実装している
* 利用例. `func json.NewEncoder(w io.Writer) *Encoder`
  * この関数は json エンコーダを返す
  * `Encoder` の `Encode` を実行すると、内部で `w.Write` を実行し、エンコード結果を `w` に出力する

#### io.Readerとio.Writerがあることで、どういう利点があるのか具体例を挙げて考えてみる
* 上記利用例のように `func json.NewEncoder(w io.Writer) *Encoder` の `w` には `*os.File`, `*bytes.Buffer` など様々なものを入れることができる
* `w` は `io.Writer` を実装さえしていればよく、`Encoder` が `w` の内部実装まで考える必要がない
* 新しく型を追加してもそれが、`io.Writer` を実装していれば `Encode` を実行できる。`Encoder` 側の修正は必要ない

## 【TRY】テストを書いてみよう
### 1回目の課題のテストを作ってみて下さい
* テストのしやすさを考えてリファクタリングしてみる
* テストのカバレッジを取ってみる
* テーブル駆動テストを行う
* テストヘルパーを作ってみる

### テストのカバレッジ
```
$ go test -coverprofile=profile github.com/gopherdojo/dojo8/kadai2/hiroygo/imgconv
ok  	github.com/gopherdojo/dojo8/kadai2/hiroygo/imgconv	0.497s	coverage: 88.1% of statements
```
