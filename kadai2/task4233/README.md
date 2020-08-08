# io.Readerとio.Writerの調査
![GitHub Actions](https://github.com/task4233/gopherdojo-studyroom/workflows/Static%20check%20with%20PR%20and%20Add%20comment%20each%20error/badge.svg)
![GitHub Actions](https://github.com/task4233/gopherdojo-studyroom/workflows/make%20godoc%20and%20deploy%20gh-pages%20branch/badge.svg)

## 標準パッケージでどのように使われているかを調査する
### ioパッケージとは
- ioパッケージは、I/Oプリミティブへの基本インタフェースを提供する
- その主な仕事は、osパッケージのような、機能を抽象化した共有のpublicなインタフェースに既存のプリミティブの実装をラップし、他の関連するプリミティブを提供すること
- これらのインタフェースやプリミティブは低レベルな命令を様々な実装でラップするため、クライアントは特に知らされてない限り、平行実行にとって安全であると想定すべきではない
- ref: [package io overview](https://golang.org/pkg/io/#pkg-overview)

### io.Readerとは
 - Readメソッドをラップするインタフェース
 - Readは`p`に含まれる`len(p)`バイトを読み上げる
 - 読んだバイト数(`0 <= n <= len(p)`)と、発生した何らかのエラーを返す
 - ref: [package io Reader](https://golang.org/pkg/io/#Reader)

```golang
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### io.Writerとは
 - Writeメソッドをラップするインタフェース
 - `p`から`len(p)`バイトを基礎となるデータストリームに書き込む
 - 書き込まれたバイト数(`0 <= n <= len(p)`)と、書き込みを停止させた原因となるエラーを返す
 - ref: [package io Writer](https://golang.org/pkg/io/#Writer)
 
```golang
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### osパッケージでの実装例
 - Read
   - ファイルから`len(b)`バイト読み上げ、読んだバイト数と発生した何らかのエラーを返す
   - 終了時に`0, io.EOF`を返す
 - Write
   - ファイルに`len(b)`バイトを書き込み、書かれたバイト数と、発生した何らかのエラーを返す
   - `n != nil`の時にnilでないエラーを返す
 - ref: [package os](https://golang.org/pkg/os/)

## io.Readerとio.Writerがあることによって得られる利点を具体例を挙げて考える
 - I/Oの標準化が出来ることが最も大きなメリットである
   - ファイル、ネットワーク、バッファといったすべての入出力を同じio.Reader/io.Writerで扱うことが可能
 - 例えば、渡されたオブジェクトから読み出しをしたい場合に、ネットワーク系のオブジェクトでもファイルオブジェクトでもバッファでもその中身を考えずに`Read()`を呼び出すだけで良い

# テストを書いてみよう
## テストのしやすさを考えてリファクタリングする & テーブル駆動テストを行う
機能ごとにメソッドを分割し、テストではテーブル駆動のテストケースを用いました。


## テストのカバレッジを取る
元々のカバレッジは71.9%でした。エラー用のテスト追加後は、88.0%に上がりました🎉

### Before
```shell
$ cat tools/getCoverage.sh 
#!/bin/sh

cd kadai2/task4233/eimg
go test -coverprofile=profile ./...
go tool cover -html=profile
/home/al17111/work/go/src/github.com/task4233/gopherdojo-studyroom (kadai2-task4233)
$ sh tools/getCoverage.sh 
ok  	github.com/task4233/gopherdojo-studyroom/kadai1/task4233/eimg	0.050s	coverage: 71.9% of statements
```

### After
```shell
$ sh tools/getCoverage.sh 
ok  	github.com/task4233/gopherdojo-studyroom/kadai2/task4233/eimg	0.059s	coverage: 88.0% of statements

```

## テストヘルパーを作る
テストヘルパー関数に`t.Helper()`を追加しました。

## 感想と課題
前回の課題であったコミットログの問題は解決できたと考えている。
一方で、どう頑張っても通らないエラーハンドリングのせいでカバレッジが高くならなかったのは残念だった。
具体的には、`os.Create()`で作成した後の`os.Open()`のエラーハンドリングがそれに該当する。
また、実装で`os.Exit(1)`になっている部分もテストできないように感じた。
[golangで終了を確認するテストの書き方](https://mattn.kaoriya.net/software/lang/go/20161025113154.htm)という記事があったが、そこまでしてカバレッジを上げる必要があるのか？と疑問に思ったため、これは不採用とした。
