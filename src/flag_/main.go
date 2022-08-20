package main

import (
	"flag"
	"log"
	"strings"
	"time"
)

// go run main.go -n 100 -b=true -s "hello world"
// -b true:bool类型不能这样用，要使用-b=true或者-b
// go run main.go -n 100 -s "hello world" -b true =>这样的话，true算其他参数

// bool类型可以接受的值包括：
// 		1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False

// flag 可以将time.Duration支持的字符串转换为time.Duration类型。
// 如：300ms -1.5h 2h45m
// 有效的时间单位包括：ns, us (or µs), ms, s, m, h.

// int,int32,int64: 1234、0664、0x1234等类型，也可以是负数

// flag.Args()  ////返回命令行参数后的其他参数，以[]string类型
// flag.NArg()  //返回命令行参数后的其他参数个数===>未使用参数个数
// flag.NFlag() //返回使用的命令行参数个数

// 方式1
/*
var (
	nFlag *int
	bFlag *bool
	sFlag *string
	dFlag *time.Duration
)

func init() {
	nFlag = flag.Int("n", 1234, "help message for flag n")
	bFlag = flag.Bool("b", false, "help message for bflag b")
	sFlag = flag.String("s", "hello", "help message for flag s")
	dFlag = flag.Duration("d", time.Second, "help message for flag d")
}

func main() {
	flag.Parse()
	fmt.Println(flag.Args(), flag.NArg(), flag.NFlag())
	fmt.Printf("nFlag:%v, bFlag:%v, sFlag:%v, dFlag:%v\n",
		*nFlag, *bFlag, *sFlag, *dFlag)
}
*/

// 方式2
/*
var (
	nFlag int
	bFlag bool
	sFlag string
	dFlag time.Duration
)

func init() {
	flag.IntVar(&nFlag, "n", 1234, "help message for n")
	flag.BoolVar(&bFlag, "b", false, "help message for b")
	flag.StringVar(&sFlag, "s", "hello", "help message for s")
	flag.DurationVar(&dFlag, "d", time.Second, "help message for d")
}

func main() {
	flag.Parse()
	fmt.Printf("nFlag:%v, bFlag:%v, sFlag:%v\n", nFlag, bFlag, sFlag)
}
*/

// 方式3：实现子命令flag
//usage:go run main.go str -n 100 -b=true -s="hellokitty"
/*
func main() {
	foostr := flag.NewFlagSet("str", flag.ExitOnError)
	sFlag := foostr.String("s", "hello", "message for flag s")
	nFlag := foostr.Int("n", 1234, "message for flag n")
	bFlag := foostr.Bool("b", false, "message for flag b")

	if len(os.Args) < 1 {
		fmt.Println("expected str subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "str":
		foostr.Parse(os.Args[2:])
		fmt.Printf("s:%v,b:%v,n:%v\n", *sFlag, *bFlag, *nFlag)
	default:
		fmt.Println("expected str subcommands")
		os.Exit(1)
	}
}
*/
// 方式4-->自定义
// 实现flag.Value接口
//go run main.go -name "gaochuan"
//go run main.go -name "gaochuan" -deltaT "10s,20s,30s"

/*
type testFlag struct {
	Name string
}

func (f *testFlag) String() string {
	return fmt.Sprintf("%s", f.Name)
}

func (f *testFlag) Set(s string) error {
	f.Name = s
	return nil
}

func TestFlag(name string, value string, usage string) *testFlag {
	f := testFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f
}

type interval []time.Duration

func (i *interval) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *interval) Set(s string) error {
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(s, ",") {
		d, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, d)
	}
	return nil
}

var iFlag interval

func init() {
	flag.Var(&iFlag, "deltaT", "help message for interval")
}

func main() {
	var tf = TestFlag("name", "hellokitty", "help message for flag name")
	flag.Parse()
	fmt.Println(tf.Name)
	fmt.Println(iFlag)
}
*/

// 方式5 flag.Func
// 需要注意的是如果某个命令行参数没有提供，相应的flag.Func()函数将不会执行。
// 这意味着不能在flag.Func()中对相应变量设定默认值
/*
flag.Func("pause", "Duration to pause between printing URLs (default 1s)", func(flagValue string) error {
    // 不要这样做！如果参数值为""则不会调用此函数.
    if flagValue == "" {
        pause = time.Second
        return nil
    }
    var err error
    pause, err = time.ParseDuration(flagValue)
    return err
})
*/

// go run main.go --pause=2s --urls="baidu.com jd.com google.com"
var (
	urls  []string
	pause time.Duration = time.Second
)

func main() {
	flag.Func("urls", "List of URLs to print", func(flagValue string) error {
		urls = strings.Fields(flagValue)
		return nil
	})

	flag.Func("pause", "Duration to pause", func(flagValue string) error {
		var err error
		pause, err = time.ParseDuration(flagValue)
		return err
	})
	flag.Parse()

	for _, u := range urls {
		log.Println(u)
		time.Sleep(pause)
	}
}

/*
func main() {
    var (
        environment string = "development"
    )

    enumFlag(&environment, "environment", []string{"development", "staging", "production"}, "Operating environment")

    flag.Parse()

    fmt.Printf("The operating environment is: %s\n", environment)
}

func enumFlag(target *string, name string, safelist []string, usage string) {
    flag.Func(name, usage, func(flagValue string) error {
        for _, allowedValue := range safelist {
            if flagValue == allowedValue {
                *target = flagValue
                return nil
            }
        }

        return fmt.Errorf("must be one of %v", safelist)
    })
}
*/
