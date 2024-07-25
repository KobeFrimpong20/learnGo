package main

import (
    "os"
    "time"

    "github.com/KobeFrimpong20/learnGo/maths/clockface"
)

func main() {
    t := time.Now()
    clockface.SVGWriter(os.Stdout, t)
}
