package fileinfo
//package main

import (
        "fmt"
        "os"
        "io"
        "errors"
        "path/filepath"
        "strings"
        "crypto/sha1"
       )

type result struct {
    fileName string
    size    int 
    hash    string
}

func CheckErr(err error) {  
    if nil != err {  
        panic(err)  
    }  
}  

func hashNum(path string) (string, error) {

    file, err := os.Open(path)
    CheckErr (err)
    defer file.Close()
    sha1h := sha1.New()
    data := make([]byte, 1024)
    for{
        n, err := file.Read(data)

        if n != 0 {
            io.WriteString(sha1h, string(data)) 
        } else {
            break
        }

        if err != nil && err != io.EOF {
            panic(err)
        }
    }

    hashName := fmt.Sprintf("%x", sha1h.Sum(nil))
    fmt.Println(hashName)
    return hashName, nil 
}


func DirIterate(dirPath, suffix string, fout *os.File) (/*files []string,*/ err error) {
    suffix = strings.ToUpper(suffix)
    err = filepath.Walk(dirPath,  func(filename string, fi os.FileInfo, err error) error {
    if fi.IsDir() {
//        return errors.New("is no file")
        return nil
    }
   if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
        fout.WriteString(fi.Name())
        fout.WriteString(" ")
        str := fmt.Sprintf("%d", fi.Size())
        fout.WriteString(str)
        fout.WriteString(" ")
        sha1h,_ := hashNum(filename)
        fout.WriteString(sha1h)
        fout.WriteString("\n")
        fmt.Println(fi.Size())
    }
    return nil
    })

    return err
    //CheckErr(err)
    //return files, err
}
/*
func main() {
        userFile := "out.txt"
        fout,err := os.Create(userFile)
        defer fout.Close()
        if err != nil {
                fmt.Println(userFile,err)
                return
        }
    DirIterate("abc", "txt", fout)
    //fmt.Println(files, err)
}
*/
