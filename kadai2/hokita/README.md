# 課題2
## io.Readerとio.Writer
### io.Readerとio.Writerについて調べてみよう
- io.Readerとio.Writerについて調べてみよう
  - 標準パッケージでどのように使われているか
  - io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

### 解答
#### 標準パッケージでどのように使われているか
- strings
  - 文字列操作
- bufio
  - バッファリングしながら読み書きする
- bytes
  - バイトスライスを操作する

#### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
IOが統一されているため、標準入出力、ファイル、ネットワーク通信どのような場合でも、入出力処理をする側、それを使う側がお互いの処理内容を知らなくてすみ、付け替えも可能。

例えばファイルを読み込んで処理Aをするアプリがあり、今回新しくHTTPレスポンスからA処理をする機能を追加したい場合、io.Readerを返すHTTPレスポンス読込処理を用意すれば、処理A自体を変更する必要はない。また処理Aのテストは入力を意識する必要もない。

## テストを書いてみよう
### 1回目の課題のテストを作ってみて下さい
  - テストのしやすさを考えてリファクタリングしてみる
  - テストのカバレッジを取ってみる
  - テーブル駆動テストを行う
  - テストヘルパーを作ってみる

### 対応
#### テストのしやすさを考えてリファクタリングしてみる
- `io.Readerとio.Writer`の課題を通して`(c *Converter) Execute(in io.Reader, out io.Writer)`を作成した。

#### テストのカバレッジを取ってみる
```shell
$ go test -coverprofile=profile github.com/gopherdojo/dojo8/kadai2/hokita/imgconv/
ok      github.com/gopherdojo/dojo8/kadai2/hokita/imgconv       0.438s  coverage: 79.5% o
f statements

$ go test -coverprofile=profile github.com/gopherdojo/dojo8/kadai2/hokita/imgconv/imgconv
ok      github.com/gopherdojo/dojo8/kadai2/hokita/imgconv/imgconv       0.476s  coverage:
 95.7% of statements
```

#### テーブル駆動テストを行う
- `map[string]struct`で作ってみた。

#### テストヘルパーを作ってみる
- convertで作成されたファイル確認、削除処理をする`checkAndDeleteFile`を作成

## わからなかったこと、むずかしかったこと
- モックをしたいがためにinterfaceを無理やり作ることはあるのだろうか。
- だからと言って抽象化してモックしないとunitテストではなくE2Eテストになってしまいそう。
