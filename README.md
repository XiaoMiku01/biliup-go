# biliup-go

B 站命令行投稿工具 Golang 实现,支持 **扫码登录**, 并将登录后返回的 cookie 和 token 保存在 `cookie.json` 中，可用于其他项目。

对 [biliup-rs](https://github.com/ForgQi/biliup-rs) 的 Golang 实现。

# 使用命令行  

二进制文件 [下载](https://github.com/XiaoMiku01/biliup-go/releases)

## 登录

```bash
./biliup-go login
```

## 上传视频

```bash
./biliup-go upload --v <视频路径> --cover <封面路径> --title <视频标题> --desc <视频简介> --t <投稿类型,1:原创 2:转载> --tid <视频分区> --tags <视频标签 多个用英文逗号隔开> --source <视频来源 类型为转载时填写>
```
eg:
```bash
./biliup-go upload --v test.mp4 --cover cover.jpg --title "[test]标题" --desc "test简介" --t 2 --tid 47 --tags "动画,音乐" --source "抖音"
```

* 查看完整用法命令行输入 `./biliup-go upload --help`
* 分区 tid : [https://biliup.github.io/tid-ref.html](https://biliup.github.io/tid-ref.html)

# 其他 go 程序使用

``` bash
go get -u github.com/XiaoMiku01/biliup-go
```

``` go
package main

import (
	// "github.com/XiaoMiku01/biliup-go/login"
	"github.com/XiaoMiku01/biliup-go/upload"
)

func main() {
	// login.LoginBili() // 登录 获取cookie
	cookieFile := "cookie.json" // cookie文件
	title := "测试"               // 视频标题
	desc := "测试"                // 视频简介
	upType := 1                 // 1:原创 2:转载
	videoPath := "test.mp4"     // 视频路径
	coverPath := "cover.jpg"    // 封面路径
	tid := 1                    // 分区id
	tag := "测试"                 // 标签 , 分割
	source := "测试"              // 来源 upType 为 2 时必填
	threadNum := 10              // 上传线程数
	// 上传视频
	upload.NewUp(cookieFile,threadNum).SetVideos(int64(tid), int64(upType), videoPath, coverPath, title, desc, tag, source).Up()
}
```    