package main

import (
    "os"
    "strings"
    "fmt"
    "strconv"
    "regexp"
)

var i int64

func main() {
    index()
}

func index() {
    lujing := "../htm/"
    dir, _ := os.Open(lujing)
    files, _ := dir.Readdir(0)
    name := make(map[int64]string)
    i = 0
    for _, a := range files {
        if !a.IsDir() {
            b := a.Name()
            c := strings.Split(b,"-")
            e := strings.Split(c[1],".")
            d, err := strconv.ParseInt(c[0],10,64)
            if err != nil {
                panic(err)
            }
            name[d] = e[0]
            if i < d {
                i = d
            }

        }
    }
    fmt.Printf("%s\n",name[i])
    ff, err := os.Open("../mo/index.mo")
    defer ff.Close()
    if err != nil {
        panic(err)
    }
    mode := ""
    buf := make([]byte,1024)
    for {
        n, _ := ff.Read(buf)
        if n == 0 { break }
        mode = mode + string(buf[:n])
    }
    ci := i
    index := ""
    for ii := 1; ii < 11; ii++ {
        if ci > 0 {
            readfile := strconv.FormatInt(ci,10) + "-" + name[ci] + ".htm"
            ff2, err := os.Open("../htm/" + readfile)
            if err != nil { panic(err) }
            buf2 := make([]byte,1024)
            rtext := ""
            for {
                n, _ := ff2.Read(buf2)
                if n == 0 { break }
                rtext = rtext + string(buf2[:n])
            }
            reg := regexp.MustCompile("<article>[^<]*<h2>([^<>]+)</h2>")
            reglist := reg.FindStringSubmatch(rtext)
            title := reglist[1]
            reg2 := regexp.MustCompile(`<div class="meta">written[^\r\n]+</a></div>`)
            reglist2 := reg2.FindStringSubmatch(rtext)
            meta := reglist2[0]
            meta = strings.Replace(meta,"meta","time",1)
            index = index + "<h3><a href=\"./htm/" + readfile + "\">"+ title + "</a></h3>\r\n" + meta
            ci = ci - 1
        }

    }
    indextext := strings.Replace(mode,"#INDEX#",index,1)
    ff3,err := os.Create("../index.html")
    defer ff3.Close()
    if err != nil { panic(err) }
    ff3.WriteString(indextext)
}
