package main
import (
    "github.com/zebra88/fileinfo"
    "fmt"
    "flag"
    "os"
)



func main() {

    var fileType = flag.String("type", "txt", "Specified file type")
    var dirPath = flag.String("in", ".", "Specified scanned directory")
    var resultFile = flag.String("out", "out.txt", "Specified scanned directory")
    flag.Parse()
    fout,err := os.Create(*resultFile)
    defer fout.Close()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    err = fileinfo.DirIterate(*dirPath, *fileType, fout)
    if err != nil {
        fmt.Println("111111")
        fmt.Println(err.Error())
        return
    }



}
