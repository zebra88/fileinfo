package fileinfo
//package main

import (
        "fmt"
        "os"
        "io"
//        "errors"
        "path/filepath"
        "strings"
        "crypto/sha1"
       )

func CheckErr(err error) {  
    if nil != err {  
        panic(err)  
    }  
}  


type InfoEntry struct {
    FileName string
    Size    int64
    Hash    string
}

type InfoManager struct {
    infos []InfoEntry
}

func NewInfoManager() *InfoManager {
    return &InfoManager{make([]InfoEntry,0)}
}

func (m *InfoManager) Len() int {
    return len(m.infos)
}
func (m *InfoManager) Add(info *InfoEntry) {
    m.infos = append(m.infos, *info)
}

func (m *InfoManager) Show() {
    if len(m.infos) == 0 {
        return 
    }
    for _, m := range m.infos {
        fmt.Print(m.FileName)
        fmt.Print(" ")
        fmt.Print(m.Size)
        fmt.Print(" ")
        fmt.Println(m.Hash)
       
    }
}

func (m *InfoManager) Write2file(fout *os.File) {
    if len(m.infos) == 0 {
        return 
    }
    for _, m := range m.infos {
        fout.WriteString(m.FileName)
        fout.WriteString(" ")
        str := fmt.Sprintf("%d", m.Size)
        fout.WriteString(str)
        fout.WriteString(" ")
        fout.WriteString(m.Hash)
        fout.WriteString("\n")
    }
}

func (m *InfoManager) CollectInfo(dirPath, suffix string) {
    suffix = strings.ToUpper(suffix)
    filepath.Walk(dirPath,  func(filename string, fi os.FileInfo, err error) error {
    if fi.IsDir() {
        return nil
    }
   if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
        sha1h,_ := hashNum(filename)
        m0 := &InfoEntry{
            fi.Name(),
            fi.Size(), 
            sha1h,
        }
        m.Add(m0)
    }
    return nil
    })
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
//    fmt.Println(hashName)
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
    }
    return nil
    })

    return err
}

