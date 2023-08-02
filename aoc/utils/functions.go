package utils

import (
  "log"
  "os"
)

func Check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func ReadInputs(path string) *os.File {
  input, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm);
  Check(err);
  return input;
}
