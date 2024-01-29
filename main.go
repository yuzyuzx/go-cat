package main

/**
作成するcatコマンドの仕様
引数でファイルパスの一覧を貰い、そのファイルを与えられた順に標準出力にそのまま出力するコマンドを作ってください。
また、-nオプションを指定すると、行番号を各行につけて表示されるようにしてください。
なお、行番号はすべてのファイルで通し番号にしてください。

flag.Args
flag.Bool / flag.BoolVar
os.Open
bufio.Scanner
os.Stdout / os.Stderr
fmt.Fprintln

*/

import (
	"flag"
	"fmt"
  "log"
  "os"
  "bufio"
)

func main() {
  files := []string{}

  flag.Parse()

  // 引数のファイルを取得する
  args := flag.Args()
  for _, v := range args {
    // TODO: ファイルが存在しない場合はエラー出力
    files = append(files, v)

    file, err := os.Open(v)
    if err != nil {
      log.Fatal(err)
    }

    // ファイルの内容を標準出力に表示
    scanner := bufio.NewScanner(file)
    row := 1
    for scanner.Scan() {
      fmt.Println(row, scanner.Text())
      row++
    }

    if err := scanner.Err(); err != nil {
      fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }

    fmt.Println()
  }
}
