package main

import (
    "fmt"
)

const (
    spanish = "Spanish"
    french = "French"
    engHello = "Hello, "
    spanHello = "Hola, "
    freHello = "Bonjour, "
)

func hello(name string, lang string) string {
    
    if name == "" {
        name = "World"
    }


    return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) string {
    
    prefix := engHello

    switch lang {
        case spanish:
            prefix = spanHello
        case french:
            prefix = freHello
        default:
            prefix = engHello
    }
    return prefix
}

func main() {
    fmt.Println(hello("Kobe", ""))
}
