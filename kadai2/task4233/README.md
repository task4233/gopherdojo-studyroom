# io.Readerとio.Writerの調査
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
## テストのしやすさを考えてリファクタリングする
## テストのカバレッジを取る
## テーブル駆動テストを行う
## テストヘルパーを作る
