## memo
### 複数パッケージの取り込み

```
import {
  "fmt"
  "github.com/<user>/gosample"
  "strings"
}

```

組み込みでなければ、GOPATHから解決

### package option

```
package main

import {
  a "fmt"
  _ "github.com/<user>/gosample"
  . "strings"
}

func main() {
  f.Println(ToUpper("hello world"))
}

```

* 任意の文字指定でプログラム内でのパッケージ名を変更できる
* `_` 利用していない場合通常はエラーになるが、無視させる
* `.` パッケージ名が省略できる

#### variables

```
var foo, bar, buz, string = "foo", "bar", "buz"

var (
  a string = "aaa"
  b        = "bbb"
  c        = "ccc"
  d        = "ddd"
  e        = "eee"
)
```

##### 関数内部での宣言

```
func main() {
    //var message string = "hello world"
    message := "hello world"
    fmt.Println(message)
}
```

##### Const

```
func main() {
  const Hello string = "hello"
  Hello = "bye" //cannot assign to Hello

}
```
