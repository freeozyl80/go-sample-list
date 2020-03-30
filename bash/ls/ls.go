package main

import (
 "fmt"
 "os/exec"
)

func main() {
  cmd := exec.Command("ls", "-lah")
  out, err := cmd.CombinedOutput()
  must(err)
  fmt.Printf("output of ls:\n%s\n", string(out))
}

func must(err error) {
  if err != nil {
    fmt.Printf("err: %s\n", err)
    panic(err)
  }
}
