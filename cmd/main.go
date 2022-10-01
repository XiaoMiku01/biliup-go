package main

import (
	"github.com/XiaoMiku01/biliup-go/login"
	"github.com/XiaoMiku01/biliup-go/upload"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var loginCmd = kingpin.Command("login", "登录")
	var uploadCmd = kingpin.Command("upload", "上传视频")
	var cookie = uploadCmd.Flag("c", "cookie文件路径").Default("./cookie.json").String()
	var videosPath = uploadCmd.Flag("v", "视频路径").Required().String()
	var coverPath = uploadCmd.Flag("cover", "封面路径").Default("").String()
	var videoTitle = uploadCmd.Flag("title", "视频标题").Default("").String()
	var videoDesc = uploadCmd.Flag("desc", "视频简介").Default("").String()
	var upType = uploadCmd.Flag("t", "上传类型,1:原创 2:转载").Default("1").Int64()
	var tid = uploadCmd.Flag("tid", "分区id").Default("47").Int64()
	var tag = uploadCmd.Flag("tags", "标签").Default("").String()
	var source = uploadCmd.Flag("source", "来源 类型为转载时填写").Default("").String()
	kingpin.Parse()
	switch kingpin.Parse() {
	case loginCmd.FullCommand():
		login.LoginBili()
	case uploadCmd.FullCommand():
		upload.NewUp(*cookie).
			SetVideos(*tid, *upType, *videosPath, *coverPath, *videoTitle, *videoDesc, *tag, *source).
			Up()
	}
}
