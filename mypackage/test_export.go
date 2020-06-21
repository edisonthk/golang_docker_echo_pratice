package mypackage

import (
  "fmt"
)

// golangでは、独自のExported identifiers が定義されている
// https://golang.org/ref/spec#Exported_identifiers

// 小文字から始まるため、private関数になる
func print() {
    fmt.Printf("Hello, it is private function!\n")
}

// 大文字から始まるため、public関数になる
func Print() {
  fmt.Printf("Hello, it is public function and exported!\n")
}