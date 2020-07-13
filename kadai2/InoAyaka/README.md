# ＜課題2＞【TRY】io.Readerとio.Writer
## 概要
### io.Readerとio.Writerについて調べてみよう

- 標準パッケージでどのように使われているか
- io.Readerとio.Writerがあることで、どういう利点があるのか具体例を挙げて考えてみる

### 調べてみた結果

<details>
<summary>長くなってしまったため、折りたたみ</summary>

#### io.Readerとは
なんらかのバイト列の入力を抽象化するための基本的なインタフェース
メソッドには、Read()が定義されている

#### io.Writerとは
なんらかのバイト列の出力を抽象化するための基本的なインタフェース
メソッドには、Write()が定義されている

#### io.Reader、io.Writerを実装している標準パッケージ
方法）godoc -analysis=type で確認
※goModuleモードは対応していないため、一時GOPATHモードにして実施>

os.File
pipe
bytes.Buffer などなど

#### 標準パッケージでどのように使っているか
os.File はReadとWriteを実装している
func (f *File) Read(b []byte) (n int, err error)
func (f *File) Write(b []byte) (n int, err error)
→ io.Readerまたはio.Writerを引数に指定している関数にos.Fileを渡すことができる


#### どういう利点があるのか
標準入力、標準出力、ファイル、ネットワーク、その他さまざまな入力処理、出力処理において
io.Reader、io.Writerという共通の型で扱うことができる　→　共通化が実現できる

例えば、標準入出力、ファイル、ネットワークなんに関しても同じような処理を行う場合には、
引数にio.Readerもしくはio.Writer を指定してあげると良い


また、io.Writer、io.Readerがさまざまな入出力を共通化しているので、
テスト等により出力先を一時的に変えたいといった場合にも、以下のように変更することができる

``` 例
//この変数outをtestファイルで上書きする
var out io.Writer = os.Stdout
fmt.Fprintln(out,string.Join(args,sep))
```  
</details>



# ＜課題2＞【TRY】テストを書いてみよう
## 概要
### 1回目の課題のテストを作ってみて下さい

- テストのしやすさを考えてリファクタリングしてみる
- テストのカバレッジを取ってみる
- テーブル駆動テストを行う
- テストヘルパーを作ってみる

### 動作手順
```
$ go test ./...

# HTMLにカバレッジを保存する場合
$ go test -coverprofile=cover.out ./...
$ go tool cover -html=cover.out -o cover.html
```

### カバレッジ
[cover.html](https://github.com/gopherdojo/dojo8/blob/kadai2-InoAyaka/kadai2/InoAyaka/cover.html) 参照


##### 補足事項
（testdataディレクトリ配下に入っているテスト用イメージについて）
オリジナルのThe Go gopher（Gopherくん）は、Renée Frenchによってデザインされました。


### 感想
- カバレッジをあげるために、どのように工夫すればいいのか試行錯誤といった感じでした
- テストヘルパーについても、正直もう少しスッキリとした形にできるのではないか...といった感じになってしまいました
- テスト全般において、他の方のコードも参考にもっと勉強が必要だなと強く感じました

