package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	ServerAddr = 9123
)

var (
	mtx   sync.Mutex
	GData []string
)

func init() {
	/**
	如果初始化时有需求，可以在这里导入数据
	t := []string{}
	GData = append(GData, t...)
	*/
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		defer mtx.Unlock()
		mtx.Lock()

		if len(GData) != 0 {
			var sb strings.Builder
			for i := range GData {
				sb.WriteString(GData[i])
				sb.WriteString("\n")
			}

			GData = GData[:0]
			fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " data queried\n")
			fmt.Fprintf(w, sb.String())
		} else {
			fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " no data\n")
			fmt.Fprintf(w, "")
		}
	})

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("read http contents err:%v\n", err)
			fmt.Fprintf(w, "error:"+fmt.Sprint(err))
		}

		s := string(bs)
		{
			mtx.Lock()
			GData = append(GData, s)
			mtx.Unlock()
		}

		fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " data rcved\n")
		fmt.Fprintf(w, "data add done\n")
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", ServerAddr),
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("start server error:%v\n", err)
	}
}
