package main

import (
    "exec"
    "fmt"
    "syscall"
    "os"
    //"strconv"
)

func main () {
    cmd, err := exec.LookPath("godoc")
    //fmt.Printf("path=%v, err=%v\n", cmd, err)
    if err != nil {
        println("Please check that godoc is installed\n")
    }
    argv := []string{"", "-http=:8090"}

    //starting server 
    pid, err1 := os.ForkExec(cmd, argv, nil, "", nil)
    fmt.Printf("err=%v, pid=%v\n", err1, pid)

    //starting browser
    cmd2, err1 := exec.LookPath("links")
    if err1 != nil {
        println("In order to run godocs, links needs to be installed")
        println("- Linux: apt-get install links")
        println("- Mac: Fink install links\n")
    }
    argv2 := []string{"", "http://localhost:8090"}
    err2 := syscall.Exec(cmd2, argv2, nil)
    fmt.Printf("err=%v\n", err2)

    /*
    argv3 := []string{"", strconv.Itoa(pid)}
    cmd3, _ := exec.LookPath("kill")
    err3 := syscall.Exec(cmd3, argv3, nil)

    if err3 != nil {
        println("warning: godoc process still running in background !") 
    }
    */

}
