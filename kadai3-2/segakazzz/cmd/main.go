package main

import "github.com/gopherdojo/dojo8/kadai3-2/segakazzz/download"

func main() {

	download.Download("http://ipv4.download.thinkbroadband.com/1GB.zip", "./testdata/", 10)
	download.Download("http://ipv4.download.thinkbroadband.com/1GB.zip", "./testdata/", 50)
	download.Download("http://ipv4.download.thinkbroadband.com/1GB.zip", "./testdata/", 100)
	download.Download("http://ipv4.download.thinkbroadband.com/1GB.zip", "./testdata/", 1000)


	//download.Download("https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4", "./", 1)
	//download.Download("https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4", "./", 2)
	//download.Download("https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4", "./", 50)
	//download.Download("https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4", "./", 100)


	//download.Download("https://d2qguwbxlx1sbt.cloudfront.net/TextInMotion-VideoSample-1080p.mp4", "./", 1)
	//download.Download("https://d2qguwbxlx1sbt.cloudfront.net/TextInMotion-VideoSample-1080p.mp4", "./", 2)
	//download.Download("https://d2qguwbxlx1sbt.cloudfront.net/TextInMotion-VideoSample-1080p.mp4", "./", 50)
	//download.Download("https://d2qguwbxlx1sbt.cloudfront.net/TextInMotion-VideoSample-1080p.mp4", "./", 100)

}

