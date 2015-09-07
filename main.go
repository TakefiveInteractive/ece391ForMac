package main

import (
    "fmt"
    "os"
    "github.com/Tedko/ece391ForMac/wget"
    "github.com/Tedko/ece391ForMac/brew"
    "github.com/Tedko/ece391ForMac/util"
)

func credits() {
    fmt.Println(
`
************************************************
Welcome to ece391 setup for OS X/Linux
- Yifei Teng, Yu Wang, Han Yan, Brett Jackson
- Special Thanks for our TA Fei Deng
************************************************
`)
}

func install() {
    credits()
    fmt.Println("Pre-flight checks...")
    out := func (str string) { fmt.Println("\t" + str) }
    if !util.BinExists("unison") {
        out("Unison is not installed")
        brew.Brew("unison")
        // TODO: copy unison config file
    } else {
        out("Found Unison")
    }
}

func main() {
    install()
    wget.WgetCli(os.Args)
    os.Exit(1)
}
