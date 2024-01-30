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
  // 行番号表示を表示するかのフラグ変数
  var isDisplayNumber bool

  flag.BoolVar(&isDisplayNumber, "n", false, "行番号を表示する")

  flag.Parse()

  // 引数のファイルを取得する
  args := flag.Args()

  // ファイル数分ループする
  for _, file := range args {

    // 指定したファイルを読み込む
    file, err := os.Open(file)

    if err != nil {
      log.Fatal(err)
    }

    // ファイルの内容を行ごとに読み込む
    scanner := bufio.NewScanner(file)
    // 行番号
    row := 1
    for scanner.Scan() {
      // 行番号表示オプションがあるか
      if(isDisplayNumber) {
        // 行番号を表示
        fmt.Fprint(os.Stdout, row, " ")
        row++
      }

      fmt.Fprintln(os.Stdout, scanner.Text())
    }

    // 正常にファイル終端に到達したか
    if err := scanner.Err(); err != nil {
      // 到達していない
      fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }

    fmt.Println()
  }
}
