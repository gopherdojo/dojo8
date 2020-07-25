# 【TRY】分割ダウンローダを作ろう
## 分割ダウンロードを行う

- Rangeアクセスを用いる
- いくつかのゴルーチンでダウンロードしてマージする
- エラー処理を工夫する
- golang.org/x/sync/errgourpパッケージなどを使ってみる
- キャンセルが発生した場合の実装を行う


## 回答
### 実行方法
~~~
cd cmd
go build -o kadai3-2
./kadai3-2 -u https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4 -d ../testdata/ -s 10 -n 10
~~~

###　オプション
- -d 出力先のディレクトリを指定
- -n ダウンロード分割数を指定
- -s タイムアウト秒を指定
- -u ダウンロードするファイルのURLを指定
~~~
$ ./kadai3-2 -h
Usage of ./kadai3-2:
  -d string
        Directory to save file (default "./testdata/")
  -n int
        Number of parallel process (default 10)
  -s int
        Seconds to timeout (default 10)
  -u string
        Target URL to download (default "https://d2qguwbxlx1sbt.cloudfront.net/TextInMotion-VideoSample-1080p.mp4")
~~~

### 実行出力例
~~~
$ ./kadai3-2 -u https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4 -d ../testdata/ -s 10 -n 10
[4]...Downloaded. Start: 7135936, End: 8919919, Size:1783983
[5]...Downloaded. Start: 8919920, End: 10703903, Size:1783983
[0]...Downloaded. Start: 0, End: 1783983, Size:1783983
[3]...Downloaded. Start: 5351952, End: 7135935, Size:1783983
[6]...Downloaded. Start: 10703904, End: 12487887, Size:1783983
[8]...Downloaded. Start: 14271872, End: 16055855, Size:1783983
[9]...Downloaded. Start: 16055856, End: 17839844, Size:1783988
[1]...Downloaded. Start: 1783984, End: 3567967, Size:1783983
[7]...Downloaded. Start: 12487888, End: 14271871, Size:1783983
[2]...Downloaded. Start: 3567968, End: 5351951, Size:1783983
====================================================================================================
Download Completed!
[Summary]
----------------------------------------------------------------------------------------------------
URL                            https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4
Output File                    ../testdata/file_example_MP4_1920_18MG.mp4
Split Count                    10                            
Remote Size (Bytes)            17839845                      
Local Size (Bytes)             17839845                      
Elapsed                        3.520312025s                  
====================================================================================================
~~~

### 処理の説明
- http.Clientで対象ファイルのヘッダー情報から、ファイルサイズを取得
- サイズをn分割して、ダウンロードするバイト位置を計算
- Go Routineを用い、各々ダウンロードを行う
- 一つのGo Routineでエラーが発生した場合は処理をキャンセルする 
- 指定のタイムアウト秒数を過ぎたら、親処理をキャンセルし、contextを用い子処理（分割ダウンローダー）を全てキャンセルする

## 感想及び課題
- テストコードをしっかり書きたい。Go Routine仕様時のテストの書き方など、調査が必要。
- contextでの処理はもっと色々なことができそうなので、勉強したい。 