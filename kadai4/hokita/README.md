# 課題4
## おみくじAPIを作ってみよう
- JSON形式でおみくじの結果を返す
- 正月（1/1-1/3）だけ大吉にする
- ハンドラのテストを書いてみる

## 動作
```shell
$ go build -o omkj
$ ./omkj
```

別ターミナルで
```shell
$ curl localhost:8080/omikuji/
{"result":"末吉"}

$ curl localhost:8080/omikuji/
{"result":"吉"}

$ curl localhost:8080/omikuji/
{"result":"中吉"}

$ curl localhost:8080/omikuji/
{"result":"凶"}
```

## 参考
- https://github.com/quii/learn-go-with-tests/blob/main/http-server.md
- https://github.com/quii/learn-go-with-tests/blob/main/json.md
