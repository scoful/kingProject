# 1. 功能
1. 一次性的任务提醒，基于server酱的微信推送和基于钉钉群的自定义机器人
2. 接入github的webhook，处理github项目的事件，比如有人star项目，根据不同事件再推送通知，基于server酱的微信推送和基于钉钉群的自定义机器人

# 2. 准备工作
## 2.1 获取server酱的微信推送链接
[参考官方文档](https://sc.ftqq.com/?c=wechat&a=bind)
## 2.2 获取钉钉群的自定义机器人链接
[参考官方文档](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq/26eaddd5)
## 2.3 补充config配置
把2.1和2.2获取的链接填入对应的config里

# 3. 如何打包
# 3.1 打win包
命令行里使用以下命令：

    set GOOS=windows
    set GOARCH=amd64
    go build main.go
会打出一个exe文件，不需要安装go环境就可以直接运行，切记要把config目录也带到运行路径下。
# 3.2 打Linux包
命令行里使用以下命令：

    set GOOS=linux
    set GOARCH=amd64
    go build main.go
会打出一个二进制文件，不需要安装go环境就可以直接运行，切记要把config目录也带到运行路径下。

# 4. 项目介绍
直接基于goframe一把梭的练手玩具，以上是完全没看过go语法出品的。

[goframe官网](https://goframe.org/index)

