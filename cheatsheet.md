## Slice
Declaring slice

```go
aSlice := []int{5, 2, 4, 6, 1, 3}
```

Copy slice
``` go
aSlice := []int{5, 2, 4, 6, 1, 3}
copySlice := make([]int, len(aSlice))
copy(copySlice, aSlice) // copySlice = {5, 2, 4, 6, 1, 3}
```

## Loop
Traditional loop
```go
aSlice := []int{5, 2, 4, 6, 1, 3}
for i:=0;i<len(aSlice);i++{
  // do something
}
```

Forever loop
```go
for {
  // do not forget to break
}
```

Modern loop
```go
aSlice := []int{5, 2, 4, 6, 1, 3}
for i,v := range(aSlice) {
  // i -> index
  // v -> value
}
```

Loop with condition
```go
aSlice := []int{5, 2, 4, 6, 1, 3}
for i = 0 && i < len(aSlice) {
  // do something
  // do not forget to change i, otherwise it will loop forever
}
```


## LinkedList
Golang has own linkedlist. It used doubly linkedlist
```go
package main

import (
  "container/list"
  "fmt"
)

func main() {
  l := list.New()

  l.PushBack("A")
  l.PushBack("B")
  l.PushFront("Start")

  // Iterate
  for e := l.Front(); e != nil; e = e.Next() {
      fmt.Println(e.Value)
  }

  l := list.New()
  e := l.PushBack("A")      // A
  l.PushFront("Start")      // Start → A
  l.InsertAfter("B", e)     // Start → A → B
  l.InsertBefore("X", e)    // Start → X → A → B
}
```
