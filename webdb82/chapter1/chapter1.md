## memo

### document

```
$ godoc fmt
```

```
$ godoc -http=":3000"
```

### build & install

importするファイルは$GOPATH/src以下のファイル。

```
$ cd main/src
$ go install

```

bin配下に`main`が作成される

### publish package
Change path

```
src/gosample -> src/github.com/<user>/<repository>/gosample/gosample.go
```

Fix main.go

```
package main

import (
  "fmt"
  "github.com/<user>/<repository"
)

func main() {
  fmt.Println(gosample.Message)
}
```

Then, push to github!
