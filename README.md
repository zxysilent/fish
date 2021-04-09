
## fish
fish 是一个方便开发go程序的工具。

监视文件修改，然后编译go程序并自动运行。


## 安装

```bash
go get github.com/zxysilent/fish
```

确保 `fish.exe` 所在路径在你的环境变量里面

如果配置了 `GOPATH` 环境变量 `fish.exe` 默认在 `GOPATH/bin`

请确保 `GOPATH/bin` 在你的环境变量里

## 升级

```bash
go get -u github.com/zxysilent/fish
```

## 基本命令

```
    version     打印版本信息

    run         热重载，监听 .go 文件变化，自动编译并运行

```

### fish version

打印版本信息

```bash
$ fish version
   ____  _         __
  / __/ (_) ___   / /
 / _/  / / (_-<  / _ \
/_/   /_/ /___/ /_//_/ v0.3.3

├── Go      : go1.16.3
├── GOOS    : windows
├── GOARCH  : amd64
├── NumCPU  : 8
├── GOPATH  : D:\App\Go
├── GOROOT  : D:\Program Files\Go
└── Date    : 2021-04-09 13:39:00
```


### fish run

热重载，监听 .go 文件变化，自动编译并运行

```bash
$ fish run
2020/02/27 15:19:34 INFO  ▶ 00001 Using 'blog' as app name
2020/02/27 15:19:34 INFO  ▶ 00002 Loading watcher...
2020/02/27 15:19:34 INFO  ▶ 00003 Watching: D:\App\Go\src\blog\conf
2020/02/27 15:19:34 INFO  ▶ 00004 Watching: D:\App\Go\src\blog\control
2020/02/27 15:19:34 INFO  ▶ 00005 Watching: D:\App\Go\src\blog
2020/02/27 15:19:34 INFO  ▶ 00006 Watching: D:\App\Go\src\blog\model
2020/02/27 15:19:34 INFO  ▶ 00007 Watching: D:\App\Go\src\blog\router
2020/02/27 15:19:35 SUCC  ▶ 00008 Built successfully
2020/02/27 15:19:35 INFO  ▶ 00009 Starting 'blog.exe'
2020/02/27 15:19:35 SUCC  ▶ 00010 './blog.exe' is running
⇨ http server started on [::]:88
```

