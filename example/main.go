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

    mm := fileinfo.NewInfoManager()
    if mm == nil {
        fmt.Println("NewInfoManager failed.")
    }
    if mm.Len() != 0 {
        fmt.Println("NewInfoManager failed,not empty.")
    }
    mm.CollectInfo(*dirPath, *fileType)

    mm.Write2file(fout)

}
