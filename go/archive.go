package main

import (
    "os"
    "strings"
    "fmt"
    "strconv"
    "regexp"
)

var tagmap=make(map[int64]string)
var tag=make(map[string]string)
var textmap=make(map[string]string)
var timetitle string=""

func main() {
    tag["闲言碎语"] = "nagging"
    tag["资源共享"] = "share" //依此格式创建标签
    tagtext := ""
    for k, v := range tag {
        lujing := "../htm/"
        dir, _ := os.Open(lujing)
        files, _ := dir.Readdir(0)
        var b string
        var t string
        var c []string
        var n int64=1
        for _, a := range files {
            if !a.IsDir() {
                b = a.Name()
                c = strings.Split(b, "-")
                t = strings.Split(c[1], ".")[0]
                fin,err := os.Open("../htm/" + b)
                defer fin.Close()
                if err != nil {
                    panic(err)
                 }
                 buf := make([]byte, 1024)
                 btext := ""
                 for{
                        n, _ := fin.Read(buf)
                        if 0 == n { break }
                        list := string(buf[:n])
                        btext = btext + list
                }
                reg := regexp.MustCompile(`<h2>([^<>/]+)</h2>`)
                reglist := reg.FindStringSubmatch(btext)
                reg2 := regexp.MustCompile(`written in (\d+)\. Tag`)
                reg2list := reg2.FindStringSubmatch(btext)
                timetitle = reg2list[1] + "-" + reglist[1]
                if t == v {
                    tagmap[n] = timetitle+ "-"+c[0]+"-"+t
                    n = n + 1
                }
            }
        }

        fin2,err2 := os.Open("../mo/tag.mo")
        defer fin2.Close()
        if err2 != nil {
                 panic(err2)
         }
         buf2 := make([]byte, 1024)
         motext := ""
         for{
                n, _ := fin2.Read(buf2)
                if 0 == n { break }
                list := string(buf2[:n])
                motext = motext + list
          }

        m := len(tagmap) % 10
        var pn int64
        if m == 0 {
            pn = int64(len(tagmap)/10)
        } else {
            aa := strconv.FormatFloat(float64(len(tagmap)/10),'f',-1,64)
            bb := strings.Split(aa,".")
            pp, err3 := strconv.ParseInt(string(bb[0]),10,64)
            if err3 != nil {panic(err3)}
            pn = pp+1
        }
        var i int64
        var u int64
        for i = 1; i < pn+1; i++ {
            litext := ""
            for u = int64(len(tagmap))-10*(i-1); u>int64(len(tagmap))-10*(i-1)-10 && u>0; u-- {
                list := strings.Split(tagmap[u],"-")
                litext = litext + "<p><a href=\"../htm/" + list[2] + "-" + list[3] + ".htm\">"
                litext = litext + list[1] + "-" + list[0] + "</a></p>\r\n"
            }
            switch i {
            case 1:
                litext = litext + "<div class=\"listnav\">共"+strconv.FormatInt(pn,10)+"页 "
                litext = litext + "第"+strconv.FormatInt(i,10)+"页"
                if pn>i {
                    litext = litext + " <a href=\"./" + v + strconv.FormatInt(i+1,10) + ".htm\">下一页</a>"
                }
                litext = litext + "</div>"
            case pn:
                litext = litext + "<div class=\"listnav\"><a href=\"./" + v + strconv.FormatInt(i-1,10) + ".htm\">上一页</a>"
                litext = litext + " 共"+strconv.FormatInt(pn,10)+"页 "
                litext = litext + "第"+strconv.FormatInt(i,10)+"页"
                litext = litext + "</div>"
            default:
                litext = litext + "<div class=\"listnav\"><a href=\"./" + v + strconv.FormatInt(i-1,10) + ".htm\">上一页</a>"
                litext = litext + " 共"+strconv.FormatInt(pn,10)+"页 "
                litext = litext + "第"+strconv.FormatInt(i,10)+"页"
                litext = litext + " <a href=\"./" + v + strconv.FormatInt(i+1,10) + ".htm\">下一页</a>"
                litext = litext + "</div>"
            }
            lujing2 := "../archive/"
            newname := v + strconv.FormatInt(i,10) + ".htm"
            wFile := lujing2 + newname
            fout,err4 := os.Create(wFile)
            defer fout.Close()
            if err4 != nil {
                    fmt.Println(wFile,err4)
                    return
            }
            wri := strings.Replace(motext,"#TAG#",litext,1)
            wri = strings.Replace(wri,"#TITLE#",k,2)
            fout.WriteString(wri)
        }
        fmt.Printf("%d page %s archive htm have created.\r\n",pn,k)
        tagtext = tagtext + "<p><a href=\"./" + v + "1.htm\">" + k + "</a>(" + strconv.FormatInt(int64(len(tagmap)),10) + ")</p>\r\n"
    }
    tagindexmo := ""
    fin5,err5 := os.Open("../mo/tagindex.mo")
    defer fin5.Close()
    if err5 != nil {
             panic(err5)
     }
     buf5 := make([]byte, 1024)
     for{
            n, _ := fin5.Read(buf5)
            if 0 == n { break }
            list := string(buf5[:n])
            tagindexmo = tagindexmo + list
     }
    wFile := "../archive/index.htm"
    fout,err := os.Create(wFile)
    defer fout.Close()
    if err != nil {
            fmt.Println(wFile,err)
            return
    }
    wri := strings.Replace(tagindexmo,"#TAGINDEX#",tagtext,1)
    fout.WriteString(wri)
    fmt.Printf("/archive/index.htm have created.")
}
