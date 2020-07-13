#  課題 2

## 【TRY】io.Readerとio.Writer

### io.Readerとio.Writerについて調べてみよう

- 標準パッケージでどのように使われているか
- io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

> - 標準パッケージでどのように使われているか

- 実装している標準パッケージ
  - bufio
  - bytes
  - zip
  - strings
 
- どの様に使われているか
  - strings
  
```
func (w *appendSliceWriter) Write(p []byte) (int, error) {
	*w = append(*w, p...)
	return len(p), nil
}
```

  - bufio
  
```
func (b *Writer) Write(p []byte) (nn int, err error) {
	for len(p) > b.Available() && b.err == nil {
		var n int
		if b.Buffered() == 0 {
			// Large write, empty buffer.
			// Write directly from p to avoid copy.
			n, b.err = b.wr.Write(p)
		} else {
			n = copy(b.buf[b.n:], p)
			b.n += n
			b.Flush()
		}
		nn += n
		p = p[n:]
	}
	if b.err != nil {
		return nn, b.err
	}
	n := copy(b.buf[b.n:], p)
	b.n += n
	nn += n
	return nn, nil
}
```

  - zip
  
```
func (w *fileWriter) Write(p []byte) (int, error) {
	if w.closed {
		return 0, errors.New("zip: write to closed file")
	}
	w.crc32.Write(p)
	return w.rawCount.Write(p)
}
```

> - io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

1. 各パッケージで扱っている書き込み先は全て違うが、同じインターフェースにて抽象化されていること中の実装は隠蔽され、利用者としてはそれぞれ全て等しく `Write` や `Read` で利用することができる。
2. 上記の通り隠蔽されることで、書き込み先や読み込み先という仕様が変更になったとしても、利用するパッケージを変更するだけで実現が可能になる。（変更容易性が上がる）

## 【TRY】テストを書いてみよう

### 1回目の課題のテストを作ってみて下さい
- テストのしやすさを考えてリファクタリングしてみる
- テストのカバレッジを取ってみる
- テーブル駆動テストを行う
- テストヘルパーを作ってみる

## 動作手順

```
$ go test ./imageconverter
$ go test ./commands/imageconverter
```

## カバレッジ

- imageconverter
`./imageconverter_coveage.html`

- main
`./imageconverter_coveage.html`

##　感想等

### io.Readerとio.Writerについて調べてみよう

- Interfaceを利用した

### 【TRY】テストを書いてみよう

- 可能な限り、テスト対象の関数内で利用しているパッケージをMockして、依存を減らす様にしました。（Mockのパッケージを使わないで実現するのに相当苦労しました。）
- 画像のエンコードを、ストラテジーパターンを利用してリファクタリングしました。
- テスタブルなコードを目指すあまり、実装として読みにくくなっていないか気になっています。

