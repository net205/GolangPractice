package main

import (
    "fmt"
    "os"
    "sync"
    "path/filepath"
    "./stopwatch"
)

var wg sync.WaitGroup

func main() () {
    //drivers := []string { "C:\\", "D:\\", "E:\\" }
    drivers := getdrives()
    wg.Add(len(drivers))
    
    for _, v := range drivers {
        go fileSearch(v)
    }

    wg.Wait()
}

func fileSearch(driver string) {
    defer wg.Done()
    fileList := []string{}
    s := stopwatch.Start(0)

    filepath.Walk(driver, func(path string, f os.FileInfo, err error) error {
        if path == "E:\\abcdefgh" {
            return filepath.SkipDir
        }

        fileList = append(fileList, path)
        fmt.Println(path)
        return nil
    })

    /*
    for _, file := range fileList {
        fmt.Println(file)
    }
    */

    duration := s.ElapsedTime()
    fmt.Printf("\n搜索 %s\n共 %d 个文件(夹)\n耗时 %s\n", driver, len(fileList), duration)
}

func getdrives() (r []string){
    for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ"{
        _, err := os.Open(string(drive)+":\\")
        if err == nil {
            r = append(r, string(drive)+":\\")
        }
    }
    return
}