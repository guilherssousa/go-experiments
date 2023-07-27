package main

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "path/filepath"
  "runtime"
)
func openBrowser(url string) {
  var err error
  switch runtime.GOOS {
  case "linux":
    err = exec.Command("xdg-open", url).Start()
  case "windows":
    err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
  case "darwin":
    err = exec.Command("open", url).Start()
  default:
    err = fmt.Errorf("unsupported platform")
  }
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  running_path, err := os.Executable()
  if err != nil {
    log.Println(err)
  }

  dir, _ := filepath.Split(running_path) 

  file_path := fmt.Sprintf("file://%s/cheatsheet.html", dir)

  fmt.Println(`Abrindo a cheatsheet...`)

  openBrowser(file_path)
}
