package main

import (
    "exec"
    "os"
    "strconv"
    "time"
    "log"
)

var f *os.File
var w *log.Logger

func kill_process(pid int) {
    argv3 := []string{"", strconv.Itoa(pid)}
    cmd3, _ := exec.LookPath("kill")
    err3 := os.Exec(cmd3, argv3, nil)
    w.Printf("pid=%v, err3=%v", pid, err3.String())

    msg :=  "err3=" + err3.String() + "\n"
    w.Println(msg)
    w.Println("Done !\n")
    if err3 != nil {
        println("warning: godoc process still running in background !")
    }

}


func main () {

    /*
    *TODO:
    *   1) replace the pause by a test to check if the server responding
    */

    var e os.Error
    f, e = os.Open("log.txt",  os.O_WRONLY | os.O_CREAT,  0666)
    w = log.New(f,">>>",0)
    w.Printf("testing log !\n")

    if e != nil {
        println("Open log file failed !")
    }

    defer f.Close()

    /*
    //---------------
    // Test for os.ForkExec
    //---------------
    streams := []*os.File{ os.Stdin, os.Stdout, os.Stderr}

    //---------------
    // test
    //---------------
    argv0 := []string{"ls", "-lAF",  "/"}
    cmd0, _ := exec.LookPath("ls")
    pid, err0 := os.ForkExec(cmd0, argv0, nil, "", streams)
    fmt.Printf("pid = %d\n", pid)
    if err0 != nil {
        println("test with ls Failed !")
    }
    */

    //---------------
    // Starting server 
    //---------------
    cmd, err := exec.LookPath("godoc")

    //fmt.Printf("path=%v, err=%v\n", cmd, err)
    if err != nil {
        println("Please check that godoc is installed\n")
    }
    argv := []string{"godoc", "-http=:8090"}
    pid1, err1 := os.ForkExec(cmd, argv, nil, "", nil)
    w.Printf("err1=%v, pid=%v\n", err1, pid1)

    //pause in the parent process to allow the server to start
    time.Sleep(1e9)

    // Waiting for the server to complete
    os.Wait(pid1, os.WNOHANG)


    //---------------
    // Starting browser
    //---------------
    cmd2, err1 := exec.LookPath("links")
    if err1 != nil {
        println("In order to run godocs, links needs to be installed")
        println("- Linux: apt-get install links")
        println("- Mac: Fink install links\n")
    }

    argv2 := []string{"links", "http://localhost:8090/pkg"}
    //pid2, err2 := os.ForkExec(cmd2, argv2, nil, "", nil)
    err2 := os.Exec(cmd2, argv2, nil)
    w.Printf("err2=%v\n", err2)

    // Waiting for the browser to complete
    //os.Wait(pid2, os.WNOHANG)
    // Waiting for the server to complete
    os.Wait(pid1, os.WNOHANG)

    // deferred closing the server.
    defer kill_process(pid1)
    // Closing the browser.
    //defer kill_process(pid2)
}

