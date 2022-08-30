package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 开启http服务的三种方式

/*
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/bye", sayBye)
	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe(":1210", nil))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye, this is v1 httpServer"))
}
*/

/*
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	log.Println("Starting v2 httpServer")
	log.Fatal(http.ListenAndServe(":1210", mux))
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 2"))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye, this is v2 httpServer"))
}
*/

/*
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	server := &http.Server{
		Addr:         ":1210",
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	log.Println("Starting v3 httpServer")
	log.Fatal(server.ListenAndServe())
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version3"))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(4 * time.Second)
	w.Write([]byte("bye bye,this is v3 httpServer"))
}
*/

/*
var server *http.Server

var mquit = make(chan int, 1)

	func main() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)

		mux := http.NewServeMux()
		mux.Handle("/", &myHandler{})
		mux.HandleFunc("/bye", sayBye)

		server = &http.Server{
			Addr:         ":1210",
			WriteTimeout: time.Second * 4,
			Handler:      mux,
		}
		go func() {
			<-quit
			err := server.Shutdown(nil)
			//if err := server.Close(); err != nil {
			//	log.Fatal("Close server:", err)
			//}
			if err != nil {
				fmt.Println("err:", err)
			}
			fmt.Println("quit.............................")
		}()

		log.Println("Starting v3 httpServer")
		err := server.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				//log.Fatal("Server closed under request")
			} else {
				//log.Fatal("Server closed unexpected:", err)
			}
		}
		//log.Fatal("Server exited")
		fmt.Println("hhhhhhhhhhhhhhhhh")
	}

type myHandler struct{}

	func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is version3"))
	}

	func sayBye(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bye bye, shutdown the server")) //没有输出----->为什么呢？不是说shutdown是优雅关闭？
		time.Sleep(time.Second * 3)
		err := server.Shutdown(nil)
		fmt.Println("sayBye err:", err)
		if err != nil {
			//	log.Fatal([]byte("shutdown the server err:"), err)
		}
		fmt.Println("KKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK")
		//mquit <- 0
	}
*/

func main() {
	log.Println("test-----------------")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		fmt.Fprintln(w, "Hello world!")
	})
	server := &http.Server{
		Addr:    ":1210",
		Handler: mux,
	}
	go server.ListenAndServe()

	listenSignal(context.Background(), server)
}

func listenSignal(ctx context.Context, httpSrv *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-sigs:
		fmt.Println("notify sigs")
		httpSrv.Shutdown(ctx)
		fmt.Println("http shutdown")
	}
}

// 下面方式好点
// ************1.信号捕捉写在新goroutine中,主goroutine方便捕获http错误主动退出以方便defer函数调用****************

// 作者原话:
// ListenAndServe()在goroutine中的话，错误处理大概率是log.Fatal(err)这样的操作，
// 如果服务并不是主动退出的（比如启动时立马遇到端口占用的错误），主函数main()中的defer是不会执行的。
// 我这里用了一些额外的复杂度让安全退出的逻辑更圆满了一些。

/*
type GracefulServer struct {
	Server           *http.Server
	shutdownFinished chan struct{}
}

func (s *GracefulServer) ListenAndServe() (err error) {
	if s.shutdownFinished == nil {
		s.shutdownFinished = make(chan struct{})
	}

	err = s.Server.ListenAndServe()
	if err == http.ErrServerClosed {
		// expected error after calling Server.Shutdown().
		err = nil
	} else if err != nil {
		err = fmt.Errorf("unexpected error from ListenAndServe: %w", err)
		return
	}

	log.Println("waiting for shutdown finishing...")
	<-s.shutdownFinished
	log.Println("shutdown finished")

	return
}

func (s *GracefulServer) WaitForExitingSignal(timeout time.Duration) {
	var waiter = make(chan os.Signal, 1) // buffered channel
	signal.Notify(waiter, syscall.SIGTERM, syscall.SIGINT)

	// blocks here until there's a signal
	<-waiter

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := s.Server.Shutdown(ctx)
	if err != nil {
		log.Println("shutting down: " + err.Error())
	} else {
		log.Println("shutdown processed successfully")
		close(s.shutdownFinished)
	}
}

func main() {
	flag.Parse()

	var err error
	defer func() {
		if err != nil {
			log.Println("exited with error: " + err.Error())
		}
	}()

	// 各种各样的初始化和依赖注入...
	//
	// defer tearDown()
	// defer beforeClose() // 注册各种各样的主程序退出时的清理工作

	server := &GracefulServer{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: myAwesomeHandler,
		},
	}

	go server.WaitForExitingSignal(10 * time.Second)

	log.Printf("listening on port %d...", port)
	err = server.ListenAndServe()
	if err != nil {
		err = fmt.Errorf("unexpected error from ListenAndServe: %w", err)
	}
	log.Println("main goroutine exited.")
}
*/
