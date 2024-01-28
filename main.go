package main

/**
作成するcatコマンドの仕様
引数でファイルパスの一覧を貰い、そのファイルを与えられた順に標準出力にそのまま出力するコマンドを作ってください。
また、-nオプションを指定すると、行番号を各行につけて表示されるようにしてください。
なお、行番号はすべてのファイルで通し番号にしてください。
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
  }
  // fmt.Println(files)
  fmt.Fprintln(os.Stdout, files)

  // ファイルを開く
  file, err := os.Open(files[0]) // For read access.
  if err != nil {
    log.Fatal(err)
  }

  // ファイルの内容を標準出力に表示
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text()) // Println will add back the final '\n'
  }
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
  }


}
