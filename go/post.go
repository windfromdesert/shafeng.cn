package main
import (
        "os"
        "fmt"
        "strings"
        "strconv"
        "regexp"
)

var i int64=0
var lujing string=""
var newname string
var lastname string
var mode string=""
var modeFile string=""
var img string
var title string
var wri string=""
var post string
var time string=""
var tag string=""
var tagurl string=""
var tagurlpre string=""
var nav string=""
var body string=""

func main() {
    readtext()
    creattagurl()
    newfilename()
    fmt.Printf("%s\n",newname)
    readmode()
    bodytext()
    addlink()
    writefile()
}

func creattagurl() {        //依此格式创建标签，此格式必须与archive.go保持一致
    if tag == "闲言碎语" {
        tagurl = "nagging"
    }
    if tag == "资源共享" {
        tagurl = "share"
    }
}

func newfilename() {
       lujing = "../htm/"
       dir, _ := os.Open(lujing)
       files, _ := dir.Readdir(0)
       var b string
       var c []string
       for _, a := range files {
           if !a.IsDir() {
               b = a.Name()
               c = strings.Split(b,"-")
               d, err := strconv.ParseInt(c[0],10,64)
               if err != nil {
                   panic(err)
               }
               if i<d {
                   i = d
                   tagurlpre = strings.Split(c[1],".")[0]
               }
           }
       }
       newname = strconv.FormatInt(i+1,10) + "-" + tagurl + ".htm"
       lastname = strconv.FormatInt(i,10) + "-" + tagurlpre + ".htm"
}

func readmode() {
    fin,err := os.Open(modeFile)
    defer fin.Close()
    if err != nil {
             fmt.Println(modeFile,err)
             return
     }
     buf := make([]byte, 1024)
     for{
            n, _ := fin.Read(buf)
            if 0 == n { break }
            list := string(buf[:n])
            mode = mode + list
     }

}

func readtext() {
    textFile := "../post.txt"
    fin,err := os.Open(textFile)
    defer fin.Close()
    if err != nil {
             fmt.Println(textFile,err)
             return
     }
     buf := make([]byte, 1024)
     alltext := ""
     for{
            n, _ := fin.Read(buf)
            if 0 == n { break }
            alltext = alltext + string(buf[:n])
     }
            sp := strings.Split(alltext,"\r\n\r\n")
            sp2 := strings.Split(sp[0],"\r\n")
            modeFile = "../mo/post.mo"
            time = sp2[0]
            tag = sp2[1]
            title = sp2[2]
            temp := sp[0]+"\r\n\r\n"
            post = strings.Replace(alltext,temp,"",1)
            fmt.Printf("tag: %s\n",tag)

}

func writefile() {
    wFile := lujing + newname
    fout,err := os.Create(wFile)
    defer fout.Close()
    if err != nil {
            fmt.Println(wFile,err)
            return
    }
    wri = strings.Replace(mode,"#TITLE#",title,2)
    wri = strings.Replace(wri,"#TIME#",time,1)
    wri = strings.Replace(wri,"#TAG#","<a href=\"../archive/share.htm\">" + tag + "</a>",1)
    wri = strings.Replace(wri,"#POST#",body,1)
    fout.WriteString(wri)
}

func bodytext() {
    body = "<p>"+strings.Replace(post,"\r\n\r\n","</p>\r\n<p>",-1)+"</p>\r\n"
    body = strings.Replace(body," \r\n","<br />\r\n",-1)
    reg := regexp.MustCompile(`\+ +|- +`)
    list := reg.FindAllString(body,-1)
    if len(list) > 0 {
        st2 := ""
        bodylist := strings.Split(body,"</p>")
        reg2 := regexp.MustCompile(`<p>\+ +|<p>- +`)
        st := ""
        st3 := ""
        for k, v := range bodylist {
           switch k {
                case len(bodylist)-1:
                default:
                if len(reg2.FindAllString(v,-1)) > 0 {
                    if strings.Contains(v,"\r\n") {
                        reg3 := regexp.MustCompile(`\r\n\+ +|\r\n- +`)
                        st3 = strings.Replace(v,reg3.FindAllString(v,-1)[0],"</li><li>",-1)
                        st = strings.Replace(st3,reg2.FindAllString(v,-1)[0],"<p><li>",1)
                        st2 = st2 + st + "</li></p>"
                    } else {
                    st = strings.Replace(v,reg2.FindAllString(v,-1)[0],"<p><li>",1)
                    st2 = st2 + st + "</li></p>"
                    }
                } else {
                    st2 = st2 + v + "</p>"
                }
            }
        }
        body = st2
    }
}

func addlink() {
    var n int64
    n = 1
    title := ""
    url := ""
    c := strconv.FormatInt(n,10)
    reg := regexp.MustCompile("\\[" + c + "\\]")
    reglist := reg.FindAllString(body,-1)
    for len(reglist)>0 {
        reg2 := regexp.MustCompile("(\\[)" + "([^\\[\\]]+)" + "(\\])(\\[" + c + "\\])")
        reglist2 := reg2.FindStringSubmatch(body)
        reg3 := regexp.MustCompile("\\[" + c + "\\]:" + "(\\S+)" + "\\s+\\(" +
        "([^\\[]+)" + "\\)")
        reglist3 := reg3.FindStringSubmatch(body)
        title = reglist3[2]
        url = reglist3[1]
        body = strings.Replace(body,reglist2[0],"<a href=\"" + url + "\" title=\"" + title + "\" target=\"_blank\">"+reglist2[2]+"</a>",-1)
        body = strings.Replace(body,reglist3[0],"",1)
        reg4 := regexp.MustCompile("[\\r\\n]+$")
        body = reg4.ReplaceAllString(body,"")
        n++
        c = strconv.FormatInt(n,10)
        reg = regexp.MustCompile("\\[" + c + "\\]")
        reglist = reg.FindAllString(body,-1)
    }
}
