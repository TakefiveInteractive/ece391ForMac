package brew

import (
    "github.com/Tedko/ece391ForMac/util"
    "runtime"
)

func getBrewDownloadString() string {
    if runtime.GOOS == "linux" {
        return `"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/linuxbrew/go/install)"`
    } else {
        return `"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`
    }
}

var ensureBrew = (func () func() {
    calc := false
    return func() {
        if !calc {
            if util.BinExists("brew") {
                return
            }
            util.MustRun("ruby", "-e", getBrewDownloadString())
            calc = true
        }
    }
})()

func Brew(name string) {
    ensureBrew()
    util.MustRun("brew", "install", name)
}
