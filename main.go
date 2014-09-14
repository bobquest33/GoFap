package main

import (
  "path/filepath"
  "fmt"
  "os"
  "strings"
  "time"
  "sort"
  "net/http"
  "net/url"
  "flag"
)

var filesystemPath = flag.String("path", ".", "Path to your videos")

type Video struct {
    Name string
    LastModified time.Time
}

type Videos []Video

func (slice Videos) Len() int {
    return len(slice)
}

func (slice Videos) Less(i, j int) bool {
    return slice[i].LastModified.After(slice[j].LastModified)
}

func (slice Videos) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

func getVideos() Videos {
    root := *filesystemPath
    filetypes := []string{".mp4", ".wmv", ".flv", ".mov", ".m4v", ".avi"}
    videos := Videos{}
    err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
        path = strings.TrimLeft(path, root)
        for _, suffix := range filetypes {
            if strings.HasSuffix(path, suffix) {
                video := Video{Name: path, LastModified: f.ModTime()}
                videos = append(videos, video)
            }
        }
        return nil
    })
    fmt.Printf("filepath.Walk() returned %v\n", err)
    sort.Sort(videos)
    return videos
}

func generatePlaylist(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "#EXTM3U")
    videos := getVideos()
    for _, video := range videos {
        Url, err := url.Parse("http://" + r.Host + "/" + video.Name)
        if err != nil {
            fmt.Println("url parse fail")
        }
        fmt.Fprintln(w, "#EXTINF:100," + video.Name)
        fmt.Fprintln(w, Url.String())
    }
}

func main() {
    flag.Parse()
    http.Handle("/", http.FileServer(http.Dir(*filesystemPath)))
    http.HandleFunc("/playlist", generatePlaylist)
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        panic(err)
    }
}