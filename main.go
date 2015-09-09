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
    out := func (str string) { fmt.Println("\t> " + str) }
    if !util.BinExists("unison") {
        out("Unison is not installed")
        brew.Brew("unison")
        // TODO: copy unison config file
    } else {
        out("Found Unison")
    }
    // Create ece391 folder
    baseImagePath := filepath.Join(basePath, "assets", "ece391.tar.gz")
    if _, err := os.Stat(baseImagePath); os.IsNotExist(err) {
        out("Downloading ece391 qemu base image")
        os.MkdirAll(filepath.Dir(baseImagePath), 0777)
        err, _ := wget.WgetCli([]string {
            os.Args[0],
            "-O", filepath.Base(baseImagePath),
            "-P", filepath.Dir(baseImagePath),
            "http://duke8253.net/ece391/ece391.tar.gz" })
        if err != nil { return err }
    }
    if _, err := os.Stat(filepath.Join(baseImagePath, "ece391")); os.IsNotExist(err) {
        out("Extracting ece391 folder")
        if err := util.Untar(baseImagePath, basePath); err != nil { return err }
    } else {
        out("ece391 folder exists")
    }
    if util.Run("smbd", "--version") != nil {
        out("Smbd is not good")
        brew.Brew("smbd")
    } else {
        out("Smbd detected")
    }
    return nil
}

func main() {
    if err := install(); err != nil {
        fmt.Printf("Error: %+v. \n", err)
        os.Exit(1)
    }
}
