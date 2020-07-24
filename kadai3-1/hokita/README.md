# 課題3-1 タイピングゲームを作ろう

## 課題内容
- 標準出力に英単語を出す（出すものは自由）
- 標準入力から1行受け取る
- 制限時間内に何問解けたか表示する

### ヒント
- 制限時間にはtime.After関数を用いる
  - context.WithTimeoutでもよい
- select構文を用いる
  - 制限時間と入力を同時に待つ

## 動作
```shell
$ go build

$ go build -o typing

$ ./typing
ddd
>ddd
Good!

aaa
>aab
Bad..

aaa
>aaa
Good!

bbb
>
------
Finish!!
Result: 2 points
```

## カバレッジ
```shell
$ go test -coverprofile=profile github.com/gopherdojo/dojo8/kadai3-1/hokita/typing
ok      github.com/gopherdojo/dojo8/kadai3-1/hokita/typing      0.046s  coverage: 46.7% of statements
```
（低い。。）

## 工夫したこと
- 入力は「【TRY】チャンネルを使ってみよう」を参考に作成
- 過去のgopherdojoのレビューコメントを確認してscanner.Scan()のエラー処理を書いた。

## わからなかったこと、むずかしかったこと
- `run()`関数の良いテスト方法が思いつかなかった。
  - `input()`のようにテストしやすい単位でメソッドとして外だしをしたかったがこちらも思いつかなったか。
- エラーハンドリングが合っているか自信がない。
