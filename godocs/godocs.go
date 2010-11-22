package main

import (
    "exec"
    "fmt"
    "syscall"
    "os"
    "strconv"
    "time"
    "bufio"
)

var f *os.File
var w *bufio.Writer

func close_server(pid int) {
    argv3 := []string{"", strconv.Itoa(pid)}
    cmd3, _ := exec.LookPath("kill")
    err3 := syscall.Exec(cmd3, argv3, nil)
    fmt.Printf("pid=%v, err3=%v", pid, err3)

    msg :=  "err3=" + strconv.Itoa(err3) + "\n"
    w.WriteString(msg)
    w.WriteString("Done\n")
    w.Flush()

    if err3 != 0 {
        println("warning: godoc process still running in background !")
    }
}


func main () {

    /*
    *TODO:
    *   1) replace the pause by a test to check if the server responding
    *   2) kill server process when the browser exits: Use system calls to monitor the borowser's process  
    */

    f, _ = os.Open("log.txt",  os.O_WRONLY | os.O_CREAT,  0666)
    w = bufio.NewWriter(f)
    w.WriteString("testing log !\n")
    w.Flush()
    cmd, err := exec.LookPath("godoc")
    defer f.Close()

    //fmt.Printf("path=%v, err=%v\n", cmd, err)
    if err != nil {
        println("Please check that godoc is installed\n")
    }
    argv := []string{"", "-http=:8090"}


    //starting server 
    pid, err1 := os.ForkExec(cmd, argv, nil, "", nil)
    fmt.Printf("err1=%v, pid=%v\n", err1, pid)
    time.Sleep(1e9)

    //starting browser

    cmd2, err1 := exec.LookPath("links")
    if err1 != nil {
        println("In order to run godocs, links needs to be installed")
        println("- Linux: apt-get install links")
        println("- Mac: Fink install links\n")
    }

    argv2 := []string{"", "http://localhost:8090"}
    //err2 := syscall.Exec(cmd2, argv2, nil)
     _, err2 := os.ForkExec(cmd2, argv2, nil, "", nil)
    fmt.Printf("err2=%v\n", err2)

    //closing the server.
    //defer close_server(pid)

}

