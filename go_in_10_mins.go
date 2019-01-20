package main

import (
    "os"
    "fmt"
    "time"
    "strconv"
)

func show() {}

func ss(x int) {}

func kk(x int) int {
    return x + 10
}

func ww(s1 string, s2 string) string {
    return s1 + s2
}

func gg() (int, string) {
    return 123, "abc"
}

const PI = 1000

type vt struct {
    x, y int
}

func (d vt) add() int {
    return d.x + d.y
}

func (d vt) String() string {
    return strconv.Itoa(d.x) + "" + strconv.Itoa(d.y)
}

func (d vt) div() int {
    return d.x / d.y
}

func dv(str string) {
    fmt.Println(str)
}

var exit chan int = make(chan int)

func test(t int) {
    exit<-t
}

func main() {
    fmt.Print(PI)
    i := 0
    for i < 100 {
        i++
    }

    if i > 2 {
    } else {
    }

    d := 0
    switch d {
    case 10:
        break
    case 9:
        break
    default:
        break
    }

    fmt.Println("............")
    hostname, err := os.Hostname()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(hostname)
    }

    mp := make(map[string]string)
    mp["www"] = "dest"
    mp["abc"] = "lopj"
    fmt.Println(mp["www"])

    mm := make(map[int]string)
    mm[1] = "dfg"
    mm[2] = "awqw"
    arr := []int{1, 2, 3, 5, 6, 7, 8, 9}
    for i := 0; i < len(arr); i++ {
        fmt.Println(i, arr[i])
    }

    arr1 := []string{"23456", "123"}
    fmt.Println(arr1[0])

    go func() {
        fmt.Println("hello")
    }()

    kp := make(map[interface{}]interface{})
    kp[1] = "stdc"
    kp[2] = 123

    fmt.Println(kp)

    go dv("nmp")
    time.Sleep(1000)

    go test(1024)

    fmt.Println(<-exit)
    time.Sleep(2000)
}
