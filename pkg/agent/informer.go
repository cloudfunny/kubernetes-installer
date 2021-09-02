package agent

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	apiVersion string = "v1alpha1"
	protocol   string = "http"
)

type Event string

type Informer struct {
	Object   string
	Cli      *client
	locker   sync.RWMutex
	Callback map[Event]func(...interface{})
}

type client struct {
	remoteURL string
	body      io.Reader
	c         *http.Client
}

type RespDecoder struct {
	Opt string
	Obj interface{}
}

func NewInformer(server, obj string) (informer *Informer) {
	newCli := newClient(server, obj)
	informer = &Informer{
		Object:   obj,
		Cli:      newCli,
		Callback: make(map[Event]func(...interface{})),
	}
	return
}

func newClient(server string, obj string) *client {
	objectUrl := protocol + "://" + server + "/" + apiVersion + "/" + obj + "/watch"
	newCli := &client{
		remoteURL: objectUrl,
		body:      nil,
		c:         new(http.Client),
	}
	return newCli
}

func (cli *client) doRequest() *http.Response {
	req, err := http.NewRequest("GET", cli.remoteURL, cli.body)
	if err != nil {
		panic(err)
	}
	resp, err := cli.c.Do(req)
	if err != nil {
		log.Fatalln("http get error: ", err)
		panic(err)
	}
	return resp
}

func (in *Informer) Listen() {
	resp := in.Cli.doRequest()
	defer resp.Body.Close()
	data := make([]byte, 4096)
	target := RespDecoder{}
	for {
		readN, err := resp.Body.Read(data)
		if readN > 0 {
			_ = json.Unmarshal(data[:readN], &target)
			switch target.Opt {
			case "create":
				in.Callback[Event(target.Opt)]()
			}
			fmt.Println(target.Opt, target.Obj)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
}

func (in *Informer) Call() {

}
