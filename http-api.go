package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"sync/atomic"
	"strconv"
	"os/exec"
	"fmt"
	"io"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		now := time.Now()
		f, err := os.Open(`imgs/id_card.jpeg`)
		panicIfError(err)
		resp, err := http.Post(`http://127.0.0.1:8080/image_to_text`, `binary/octet-stream`, f)
		panicIfError(err)
		defer resp.Body.Close()

		_, _ = io.Copy(os.Stdout, resp.Body)
		fmt.Println(time.Since(now))
	} else {
		runServer()
	}
}

func runServer() {
	{
		err := os.MkdirAll(`upload`, 0777)
		panicIfError(err)
	}

	var fileId int64

	http.HandleFunc(`/image_to_text`, func(writer http.ResponseWriter, request *http.Request) {
		content, err := ioutil.ReadAll(request.Body)
		panicIfError(err)

		id := atomic.AddInt64(&fileId, 1)
		fileName := `upload/` + strconv.Itoa(int(id))
		err = ioutil.WriteFile(fileName, content, 0777)
		panicIfError(err)

		defer os.Remove(fileName)
		
		out, err := exec.Command(`python`, `image_to_text.py`, fileName).CombinedOutput()
		panicIfError(err)

		_, _ = writer.Write(out)
	})
	{
		fmt.Println(`start listen :8080`)
		err := http.ListenAndServe(":8080", nil)
		panicIfError(err)
	}
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
