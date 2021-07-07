package run

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	path "path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/zxysilent/fish/internal/cmds"
	. "github.com/zxysilent/fish/logger"
)

var CmdRun = &cmds.Command{
	UsageLine: "run [-name=name] [-all=false] [-doc=false]",
	Short:     "watch your .go files and restart your go application",
	Long: `
Run command will monitor any changes to the application file and recompile/restart it.`,
	Run: runRun,
}

var (
	buildname  string                   // 编译名称😎
	runname    string                   // 运行名称😝
	watchall   bool                     // 监听所有包括静态文件✔
	gendoc     bool                     // 是否生成文档✔
	always     chan struct{}            // 保持一直运行😋
	cmd        *exec.Cmd                // 命令
	locker     sync.Mutex               // 锁🔒
	modTimes   = make(map[string]int64) // 修改时间⏰
	aimExts    = []string{".go"}        // 监听目标后缀
	staticExts = []string{".html"}      // 静态资源后缀➰
	// 临时文件🚫
	ignoreRegexps = []*regexp.Regexp{
		regexp.MustCompile(`(\w+).go~$`),
		regexp.MustCompile(`(\w+).tmp$`),
		regexp.MustCompile(`.#(\w+).go$`),
		regexp.MustCompile(`.(\w+).go.swp$`),
	}
	// 一定要排除的目录
	excludes = []string{"docs", "node_modules", "dist", "vendor", "upload"}
)

func init() {
	CmdRun.Flag.StringVar(&buildname, "name", "", "Set the app name.")
	CmdRun.Flag.BoolVar(&watchall, "all", false, "Enable watch all files. eg: .html")
	CmdRun.Flag.BoolVar(&gendoc, "doc", false, "Enable generate swagger")
	always = make(chan struct{})
	cmds.Regcmd(CmdRun)
}

// runRun
func runRun(cmd *cmds.Command, args []string) {
	// Getwd 返回一个对应当前工作目录的根路径。
	apppath, _ := os.Getwd()
	// 监听所有包含静态资源
	if watchall {
		aimExts = append(aimExts, staticExts...)
	}
	if buildname == "" {
		buildname = path.Base(apppath)
	}
	Flog.Infof("Using '%s' as app name", buildname)
	if runtime.GOOS == "windows" {
		buildname += ".exe"
	}
	// 运行名称
	runname = "./" + buildname
	var dirs []string
	// 加载监听目录
	loadWatchDirs(apppath, &dirs)
	// 新建监听
	newWatcher(dirs)
	// 初始构建
	buildApp()
	<-always
}

// newWatcher 监听文件变动
func newWatcher(dirs []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		Flog.Fatalf("Failed to create watcher: %s", err)
	}
	go func() {
		build := true
		for {
			select {
			case evts := <-watcher.Events:
				build = true
				// 忽略的文件
				if isIgnoreFile(evts.Name) {
					continue
				}
				// 不是目标文件类型
				if !isAimFiles(evts.Name) {
					continue
				}
				fmu := fileModUnix(evts.Name)
				if mt := modTimes[evts.Name]; fmu == mt {
					build = false
				}
				modTimes[evts.Name] = fmu
				if build {
					go func() {
						time.Sleep(time.Millisecond * 800)
						buildApp()
					}()
				}
			case errs := <-watcher.Errors:
				Flog.Errorf("Watcher error: %s", errs.Error())
			}
		}
	}()
	Flog.Info("Loading watcher...")
	for _, dir := range dirs {
		Flog.Infof("Watching: %s", dir)
		err = watcher.Add(dir)
		if err != nil {
			Flog.Fatalf("Failed to watch directory: %s", err)
		}
	}
}

// genDoc 生成 swagger
func genDoc() {
	Flog.Info("Generating swagger...")
	args := []string{"init"}
	build := exec.Command("swag", args...)
	stderr := bytes.Buffer{}
	build.Stderr = &stderr
	if err := build.Run(); err != nil {
		Flog.Errorf("Failed to generate: %s", stderr.String())
		return
	}
	Flog.Succ("Generated successfully")
}

// buildApp 编译APP
func buildApp() {
	locker.Lock()
	defer locker.Unlock()
	if gendoc {
		genDoc()
	}
	killApp()
	args := []string{"build"}
	args = append(args, "-o", buildname)
	build := exec.Command("go", args...)
	build.Env = append(os.Environ(), "GOGC=off")
	stderr := bytes.Buffer{}
	build.Stderr = &stderr
	if err := build.Run(); err != nil {
		Flog.Errorf("Failed to build: %s", stderr.String())
		return
	}
	Flog.Succ("Built successfully")
	runApp()
}

// runApp 启动命令
func runApp() {
	Flog.Infof("Starting '%s'", runname)
	cmd = exec.Command(runname)
	// Stdin指定进程的标准输入，如为nil，进程会从空设备读取（os.DevNull）
	// Stdout和Stderr指定进程的标准输出和标准错误输出。
	// 如果任一个为nil，Run方法会将对应的文件描述符关联到空设备（os.DevNull）
	// 如果两个字段相同，同一时间最多有一个线程可以写入。
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Args = []string{runname}
	// Env指定进程的环境，如为nil，则是在当前进程的环境下执行。
	cmd.Env = os.Environ()
	// Run 执行c包含的命令，并阻塞直到完成。
	// runApp 执行c包含的命令，但并不会等待该命令完成即返回
	if err := cmd.Start(); err != nil {
		Flog.Warnf("Starting '%s' error : %s", runname, err.Error())
		return
	}
	go cmd.Wait()
	Flog.Succf("'%s' is running", runname)
}

// killApp 结束
func killApp() {
	defer func() {
		if msg := recover(); msg != nil {
			Flog.Errorf("Kill recover: %s", msg)
		}
	}()
	if cmd != nil && cmd.Process != nil {
		if runtime.GOOS == "windows" {
			// Signal方法向进程发送一个信号
			// 在windows中向进程发送Interrupt信号尚未实现
			// os.Kill 强制进程退出 🐌
			cmd.Process.Signal(os.Kill)
		} else {
			// os.Interrupt 向进程发送中断
			cmd.Process.Signal(os.Interrupt)
		}
		wait := make(chan struct{})
		go func() {
			// Wait会阻塞直到该命令执行完成
			// 该命令必须是被Start方法开始执行的。
			// Wait方法会在命令返回后释放相关的资源。
			cmd.Wait()
			wait <- struct{}{}
		}()
		select {
		case <-wait:
			Flog.Info("Kill running process")
			return
		case <-time.After(time.Second * 5):
			Flog.Warn("Timeout. Force kill process")
			// Kill 让进程立刻退出。❌
			if err := cmd.Process.Kill(); err != nil {
				Flog.Errorf("Error while killing process: %s", err)
			}
			return
		}
	}
}

// loadWatchDirs 监听目录
func loadWatchDirs(dir string, dirs *[]string) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	flag := false
	for _, info := range infos {
		name := info.Name()
		// 一定要排除的目录
		if isExclude(name) {
			continue
		}
		if info.IsDir() && name[0] != '.' {
			loadWatchDirs(dir+"/"+name, dirs)
			continue
		}
		// 当前目录已经添加
		if flag {
			continue
		}
		if strings.HasSuffix(name, ".go") || (isStatic(name) && watchall) {
			*dirs = append(*dirs, path.Clean(dir))
			flag = true
		}
	}
}

// isExclude 一定要排除的目录
func isExclude(name string) bool {
	for _, exc := range excludes {
		if strings.HasPrefix(name, exc) {
			return true
		}
	}
	return false
}

// isStatic 是静态资源
func isStatic(name string) bool {
	for _, ext := range staticExts {
		if strings.HasSuffix(name, ext) {
			return true
		}
	}
	return false
}

// isIgnoreFile 是临时文件
func isIgnoreFile(name string) bool {
	for _, re := range ignoreRegexps {
		if re.MatchString(name) {
			return true
		}
	}
	return false
}

// isAimFiles 是目标文件
func isAimFiles(name string) bool {
	for _, s := range aimExts {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}

// fileModUnix 获取修改时间
func fileModUnix(path string) int64 {
	path = strings.Replace(path, "\\", "/", -1)
	fi, err := os.Stat(path)
	if err != nil {
		Flog.Warnf("Failed to open file on '%s': %s", path, err)
		return time.Now().Unix()
	}
	return fi.ModTime().Unix()
}
