package util

import (
    "os/exec"
)

func MustRun(name string, args ...string) {
    cmd := exec.Command(name, args...)
    if cmd.Run() != nil {
        panic("Process execution failed")
    }
}

func BinExists(filename string) bool {
    cmd := exec.Command("which", filename)
    err := cmd.Run()
    if err != nil {
        // Return val != 0 ==> not found
        if exitError, ok := err.(*exec.ExitError); ok && !exitError.Success() {
            return false
        } else {
            // Other types of error
            panic("Process execution failed")
        }
    }
    return true
}

