package main

import (
    "os"
    "os/exec"
    "syscall"
    "fmt"
)

//docker         run image <cmd> <params> This is equal to:
//go run main.go run       <cmd> <params>
func main() {
    switch os.Args[1] {
        case "run":
            parent()
        case "child":
            child()
        default:
            panic("You are shabby! Bad command!")
    }
}

func parent() {
    cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err:= cmd.Run(); err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }
}

func child() {
    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
