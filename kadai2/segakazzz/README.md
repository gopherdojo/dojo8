# Try 1: io.Reader と io.Writer について調べてみよう

## 標準パッケージでどのように使われているか

### io.Reader/io.Writer とは

- golang.org の定義によると、それぞれメソッド Read／Write を持つインターフェイス
- io.Reader -> バイト列を読み、Read 関数のバイトスライスへ格納
- io.Writer -> Write 関数の引数バイトスライスを書き出し
- インターフェースで宣言されているメソッドがすべて実装されている構造体ならばどんな型でも対象のインターフェイスとして扱うことができる

参照）https://golang.org/pkg/io/

```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

```
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### 標準パッケージでの使用

リンクの通り多くのパッケージで使用されている

- io.Reader https://golang.org/search?q=Read#Global
- io.Writer https://golang.org/search?q=Write#Global

#### 代表的なもの

- [package bufio](https://golang.org/search?q=Read#Global_pkg/bufio)
- [package csv](https://golang.org/search?q=Read#Global_pkg/csv)
- os.Stdin
- os.Stdout
- os.File

## io.Reader と io.Writer があることでどういう利点があるのか具体例を挙げて考えてみる

### 1. どこからデータを読み込み、どこへ書き出すかについて自由に実装ができる

os.Stdin、os.File からの読み込み、os.Stdout, os.File への書き出しなど、Read/Write 関数の引数・戻値が同じのため、呼び出す元の型を変えるだけで、異なる読み出し・書き出し先を実装できる

### 2. シンプルな構造で、カスタマイズが簡単に実装できる

引数１つ、戻値が２つしかないシンプルな関数を書くだけでよい。下のコードはアルファベット以外の文字を読み込まないようにカスタマイズした、Reader の例。

```golang
type alphaReader struct {
	src string
	cur int
}

func newAlphaReader(src string) *alphaReader {
	return &alphaReader{src: src}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	x := len(a.src) - a.cur
	n, bound := 0, 0
	if x >= len(p) {
		bound = len(p)
	} else if x <= len(p) {
		bound = x
	}

	buf := make([]byte, bound)
	for n < bound {
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
		}
		n++
		a.cur++
	}
	copy(p, buf)
	return n, nil
}

func main() {
	reader := newAlphaReader("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}
```

### 3. 既存の Reader/Writer からの拡張が簡単に実装できる

下は io.Reader を持つ構造体の例。下の main 関数では strings.Reader に 2.での例と同様、アルファベットのみ読み込む機能を拡張しているが、この部分は alphaReader を変更することなく、main 関数内で os.File、bufio.Reader などに変更するだけで実現が可能になる。

```golang
type alphaReader struct {
	reader io.Reader
}

func newAlphaReader(reader io.Reader) *alphaReader {
	return &alphaReader{reader: reader}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil
}

func main() {
	// use an io.Reader as source for alphaReader
	reader := newAlphaReader(strings.NewReader("Hello! It's 9am, where is the sun?"))
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}

```

#### 参考文献

- https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185

- https://qiita.com/ktnyt/items/8ede94469ba8b1399b12

# Try 2: テストを書いてみよう

## 1 回目の課題のテストを作ってみてください。

- テストのしやすさを考えてリファクタリングしてみる
- テストのカバレッジを取ってみる
- テーブル駆動テストを行う
- テストヘルパーを作ってみる

### 参考文献

- カバレッジとは　https://www.techmatrix.co.jp/t/quality/coverage.html
- テーブル駆動テストとは　https://github.com/golang/go/wiki/TableDrivenTests
- テストヘルパーとは　https://qiita.com/atotto/items/f6b8c773264a3183a53c

```
testComputeにt.Helper()の１行を追加します。すると、testCompute内で発生したエラーは呼び元のTestComputeのどの行で失敗したのかを表示するようになります。
```
