# 【TRY】タイピングゲームを作ろう
- 標準出力に英単語を出す（出すものは自由）
- 標準入力から1行受け取る
-制限時間内に何問解けたか表示する

## 回答
### 実行方法
~~~
cd cmd
go build -o kadai3-1
 ./kadai3-1 -json ../testdata/words.json
~~~

### オプション
- -json 入力データをjsonfileで指定
- -sec  制限秒数を指定

~~~
$ ./kadai3-1 -h
Usage of ./kadai3-1:
  -json string
        Path to source json file to import (default "../testdata/words.json")
  -sec int
        Seconds to timeout (default 10)
~~~

### 実行出力例
~~~
$ ./kadai3-1 -json ../testdata/words.json
[1] drive >>> drive
[2] part >>> part
[3] available >>> available
[4] yellow >>> yellow
[5] hollow >>> 
--------------------------------------------------------------------------------
Timeout!
#          Your Input           Answer               Correct?
================================================================================
1          drive                drive                ⭕    
2          part                 part                 ⭕    
3          available            available            ⭕    
4          yellow               yellow               ⭕    
[Summary]
Num of Questions     Num of Correct ANS   Match Ratio[%]       Timeout Duration[sec]
================================================================================
4                    4                    100.00               10s  
~~~

### jsonのデータ形式
- オンラインでランダムワードから生成できるツールを用いて、10000個のワードを作成 https://onlinerandomtools.com/generate-random-json
- 実行時には、ランダムソートを行うので、10000個の中からランダムに出題できる
~~~
[
  "recall",
  "work",
  "cold",
  "star",
  ...
 ]
~~~

## 感想及び課題
- テストコード生成をする。コード生成と同時にテストを書く癖をつけたい。
