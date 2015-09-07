package brew

import (
    "github.com/Tedko/ece391ForMac/util"
)

var ensureBrew = (func () func() {
    calc := false
    return func() {
        if !calc {
            if util.BinExists("brew") {
                return
            }
            util.MustRun("ruby", "-e",
                `"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`)
            calc = true
        }
    }
})()

func Brew(name string) {
    ensureBrew()
    util.MustRun("brew", "install", name)
}
