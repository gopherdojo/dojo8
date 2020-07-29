#【TRY】おみくじAPI
## おみくじAPIを作ってみよう
- JSON形式でおみくじの結果を返す
- 正月（1/1-1/3）だけ大吉にする
- ハンドラのテストを書いてみる

## 回答
### 実行方法
~~~
$ cd cmd
$ go build -o omkj main.go
$ ./omkj
Server is running with port 8080👍
~~~
上記実行後任意のブラウザでアクセスすると以下のような結果が帰ってくる
~~~
{
    "time": "2020-07-26T22:54:31.541464+09:00",
    "dice": 1,
    "result": "凶"
}
~~~

### 処理説明補足
- 以前作成した[Try]おみくじプログラムを作ろうを応用しました。
- テストパッケージをomikuji_testとして作成しました。（課題２は同じパッケージ内で行った）
- テストでは勉強のため、一部のテストをStandard Library提供の関数をMockして実行しました。

~~~
$ cd omikuji
$ go test  --cover --coverprofile=coverprofile.out 
Server is starting with port 8000 👍
PASS
coverage: 97.6% of statements
ok      github.com/gopherdojo/dojo8/kadai4/segakazzz/omikuji    0.016s
~~~

## 感想
- APIの作成自体はすんなり終わりましたが、テストコードの作成にかなり苦戦しました