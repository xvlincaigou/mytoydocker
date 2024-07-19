package main

import (
    "os"
    "os/exec"
    "fmt"
    "syscall"
    "bufio"
)

const configFile = ".mytoydockerrc"

//docker         run image <cmd> <params> This is equal to:
//go run main.go run       <cmd> <params>
func main() {
    if len(os.Args) < 3 {
        panic("Man! What can I say? Bad command!")
    }
    switch os.Args[1] {
        case "run":
            parent()
        case "child":
            child()
        case "pull":
            if len(os.Args) < 3 {
                panic("You need to specify a distribution to pull!")
            }
            pull(os.Args[2])
        case "activate":
            if len(os.Args) < 3 {
                panic("You need to specify a distribution to activate!")
            }
            activate(os.Args[2])
        default:
            panic("Bad command!")
    }
}

func parent() {
    cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
    cmd.SysProcAttr = &syscall.SysProcAttr {
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
    }
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err:= cmd.Run(); err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }
}

func child() {
    rootfsPath := "./" + os.Args[2]
    must(syscall.Mount(rootfsPath, rootfsPath, "", syscall.MS_BIND, ""))
    must(os.MkdirAll(rootfsPath+"/oldrootfs", 0700))
    must(syscall.PivotRoot(rootfsPath, rootfsPath+"/oldrootfs"))
    must(os.Chdir("/"))

    if err := syscall.Sethostname([]byte(os.Args[2])); err != nil {
        fmt.Errorf("Sethostname: %v", err)
    }

    cmd := exec.Command(os.Args[3], os.Args[4:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }
}

func pull(distribution string) {
    if inConfigFile(distribution) {
        fmt.Println("Distribution already exists.")
        return
    }

    path := "./" + distribution
    if err:= os.Mkdir(path, 0755); err != nil {
        panic(fmt.Sprintf("Failed to create directory: %v\n", err))
    }

    fmt.Println("Pulling distribution...")

    cmd := exec.Command("sudo", "debootstrap", distribution, path)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout

    if err := cmd.Run(); err != nil {
        fmt.Println("ERROR", err)
        os.RemoveAll(path)
        os.Exit(1)
    }

    appendToConfigFile(distribution)
}

func activate(distribution string) {
    if !inConfigFile(distribution) {
        fmt.Println("Distribution not found. Please pull it first.")
        return
    }

    cmd := exec.Command("/proc/self/exe", append([]string{"run", distribution, "/bin/bash"})...)
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
    }
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }
}

func inConfigFile(distribution string) bool {
    file, err := os.Open(configFile)
    if err != nil {
        return false
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if scanner.Text() == distribution {
            return true
        }
    }
    return false
}

func appendToConfigFile(distribution string) {
    file, err := os.OpenFile(configFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(fmt.Sprintf("Failed to open config file: %v", err))
    }
    defer file.Close()

    if _, err := file.WriteString(distribution + "\n"); err != nil {
        panic(fmt.Sprintf("Failed to write to config file: %v", err))
    }
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
