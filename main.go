package main

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/Tedko/ece391ForMac/wget"
    "github.com/Tedko/ece391ForMac/brew"
    "github.com/Tedko/ece391ForMac/util"
    "github.com/mitchellh/go-homedir"
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

func install() error {
    credits()

    var homePath string
    var err error
    if homePath, err = homedir.Dir(); err != nil { return err }
    basePath := filepath.Join(homePath, "ece391ForMac")
    if os.Getenv("ECE391_INSTALL_PATH") != "" {
        if result, err := homedir.Expand(os.Getenv("ECE391_INSTALL_PATH")); err != nil {
            return err
        } else {
            if basePath, err = filepath.Abs(result); err != nil { return err }
        }
    }
    fmt.Println("Installing to " + basePath)
    fmt.Println("Pre-flight checks...")
    out := func (str string) { fmt.Println("\t" + str) }
    if !util.BinExists("unison") {
        out("Unison is not installed")
        brew.Brew("unison")
        // TODO: copy unison config file
    } else {
        out("Found Unison")
    }
    if _, err := os.Stat(""); os.IsNotExist(err) {

    }
    return nil
}

func main() {
    install()
    wget.WgetCli(os.Args)
}
