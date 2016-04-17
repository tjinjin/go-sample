## 型システム
### type
Goには型があるので、引数などを通常のintにするだけではなく独自の型を定義させることでコンパイル時に間違いに気づくことができる

ex.

```
func ProcessTask(id, priority int) {}

id := 3//int
priority := 5 //int
ProcessTask(id, priority) //正しい
ProcessTask(priority, id) //間違いだが、コンパイルが通る
```

独自の型を使う

ex.

```
type ID int
type Priority int

func ProcessTask(id ID, priority Priority) {
}

var id ID = 3
var priority Priority = 5
ProcessTask(priority, id) // エラー

```

### 構造体

* 各フィールドの名前でスコープが決まり、大文字ならパブリック、小文字ならパッケージ内に閉じたスコープになる
* ポインタ型を扱うこともできる
* new()で初期化できる。その際はゼロ値で初期化される
* 型にはメソッドも宣言可能
