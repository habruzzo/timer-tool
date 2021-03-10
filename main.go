package main

import (
	"fmt"
	"net/http"
	"time"
)

type Timer struct {
	name string
	start time.Time
	stop time.Time
}

var timeMap = map[string]*Timer{}

func NewTimer(name string) *Timer {
	return &Timer{
		name: name,
		start: time.Time{},
		stop: time.Time{},
	}
}

func (t *Timer) timerStart() {
	t.start = time.Now()
}

func (t *Timer) timerStop() {
	t.stop = time.Now()
}

func (t *Timer) timerReport() string {
	return fmt.Sprintf("timer name= %s, timer start= %s, timer stop = %s, total time = %v", t.name, t.start, t.stop, t.stop.Sub(t.start))
}

func (t *Timer) timerSub() time.Duration {
	return t.stop.Sub(t.start)
}

func register(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query()["name"][0]
	fmt.Printf("creating timer %s\n", name)
	t := NewTimer(name)
	timeMap[name] = t
	fmt.Fprint(w, "register\n")
}

func start(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "start\n")
	name := req.URL.Query()["name"][0]
	t := timeMap[name]
	if t == nil {
		fmt.Fprint(w, fmt.Sprintf("no timer with name %s\n", name))
		return
	}
	fmt.Printf("starting timer %s\n", name)
	t.timerStart()
}

func stop(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "stop\n")
	name := req.URL.Query()["name"][0]
	t := timeMap[name]
	if t == nil {
		fmt.Fprint(w, fmt.Sprintf("no timer with name %s\n", name))
		return
	}
	fmt.Printf("stopping timer %s\n", name)
	t.timerStop()
}

func report(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "report\n")
	name := req.URL.Query()["name"][0]
	t := timeMap[name]
	if t == nil {
		fmt.Fprint(w, fmt.Sprintf("no timer with name %s\n", name))
		return
	}
	fmt.Printf("reporting timer %s\n", name)
	fmt.Fprint(w, t.timerReport())
	fmt.Printf("total time: %v", t.timerSub())
}

func exportData(w http.ResponseWriter, req *http.Request) {
	//var average int64
	//var standardDeviation int64
	//
	//for k, v := range timeMap {
	//	sub := v.timerSub()
	//
	//}
}

func main() {
	timeMap = make(map[string]*Timer)
	http.HandleFunc("/register", register)
	http.HandleFunc("/start", start)
	http.HandleFunc("/stop", stop)
	http.HandleFunc("/report", report)
	http.HandleFunc("/export", exportData)

	http.ListenAndServe(":8090", nil)
}