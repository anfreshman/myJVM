package main
import "flag"
import "fmt"
import "os"

type Cmd struct{
	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args []string
}

func parseCmd() *Cmd{
	cmd := &Cmd{}

	// 显示函数
	flag.Usage = printUsage
	// （被绑定变量，“参数解析标志”，默认值，描述语句）
	flag.BoolVar(&cmd.helpFlag,"help",false,"help指令的绑定参数")
	flag.BoolVar(&cmd.helpFlag,"?",false,"help指令的绑定参数")
	flag.BoolVar(&cmd.versionFlag,"version",false,"输出版本号")
	flag.StringVar(&cmd.cpOption,"classpath","","输出classpath")
	flag.StringVar(&cmd.cpOption,"cp","","输出classpath")
	// 解析函数
	flag.Parse()

	// 获取命令行参数
	args := flag.Args()
	// 有输入参数时
	if len(args) > 0{
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage(){
	fmt.Printf("Usage: %s [-Option] class [args...]\n",os.Args[0])
}

func main() {
	cmd := parseCmd()
	if cmd.versionFlag{
		fmt.Println("version 0.0.1")
	}else if cmd.helpFlag || cmd.class == ""{
		printUsage()
	}else{
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd){
	fmt.Println("JVM虚拟机已启动")
}