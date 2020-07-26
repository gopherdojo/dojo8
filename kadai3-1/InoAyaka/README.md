# ＜課題3-1＞【TRY】タイピングゲームを作ろう
## 概要

- 標準出力に英単語を出す（出すものは自由）
- 標準入力から1行受け取る
- 制限時間内に何問解けたか表示する


## 使い方
```
$ go build -o typing
$
$ ./typing -h
Usage of ./typing:
  -t int
        time limit(unit:seconds) (default 30)
$
$
```

実行例
```
$ ./typing -t 15
society
society
ability
ablity
bless
bless
anybody
anybody
title
title
understand
understand
commercial
commercial
表示                   入力                   (一致)
---------------------------------------------------------------------------
society              society              (◯)
ability              ablity               (×)
bless                bless                (◯)
anybody              anybody              (◯)
title                title                (◯)
understand           understand           (◯)
commercial           commercial           (◯)
---------------------------------------------------------------------------
7問中 6問 (85.71 %)
---------------------------------------------------------------------------
```

### 感想

- 前回の課題にあったテストのしやすさを考慮したコードをいかに書くかに悩みました
- 課題3-2に比べると、難しい課題ではないように感じたが、いざ書いてみると読みやすく書くにはどうしたら良いか悩んだ部分も多かった
