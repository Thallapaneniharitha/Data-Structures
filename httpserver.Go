//1.Http server
1.package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    f, err := os.Open("Akash")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    data := make([]byte, 4096)
    zeroes := 0
    for {
        data = data[:cap(data)]
        n, err := f.Read(data)
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println(err)
            return
        }
        data = data[:n]
        for _, b := range data {
            if b == 0 {
                zeroes++
            }
        }
    }
    fmt.Println("zeroes:", zeroes)
}
package main
 
import (
    "fmt"
    "io/ioutil"
    "net/http"
)
 
func loadFile(fileName string) (string, error) {
    bytes, err := ioutil.ReadFile(fileName)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}
 
func handler(w http.ResponseWriter, r *http.Request) {
    var body, _ = loadFile("page.html")
    fmt.Fprintf(w, body)
}
 
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":9000", nil)
}
package main
 
import (
    "fmt"
    "net/http"
)
 
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Merhaba %s", r.URL.Path[1:])
}
 
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":9000", nil)
 
    fmt.Println("Web Server")
}
