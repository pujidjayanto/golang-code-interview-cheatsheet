## Slice

Declaring slice

```go
aSlice := []int{5, 2, 4, 6, 1, 3}
```

Copy slice

```go
aSlice := []int{5, 2, 4, 6, 1, 3}
copySlice := make([]int, len(aSlice))
copy(copySlice, aSlice) // copySlice = {5, 2, 4, 6, 1, 3}
```

Find middle position

```go
left, right := 0, len(aSlice)-1
mid := left + (right-left)/2
```

Slicing slice

```go
// aSlice[low:high] -> elements from index low to high-1
// Inclusive at start (low), exclusive at end (high)
fmt.Println(aSlice[1:4]) // [2 4 6] → indices 1,2,3

// Omitting low or high
// aSlice[:high] → from 0 to high-1
// aSlice[low:] → from low to end
// aSlice[:] → whole slice
r := []int{0,1,2,3,4}
fmt.Println(r[:3]) // [0 1 2]
fmt.Println(r[2:]) // [2 3 4]
fmt.Println(r[:])  // [0 1 2 3 4]
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

Loop with predefined index looping

```go
i := 0
j := len(aSlice-1)
for ; i < len(a); i++ {
  fmt.Println(aSlice[i])
}
for ; j >=0 ; j-- {
  fmt.Println(aSlice[j])
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

## Swapping

Go can do swapping without temporary variable

```go
a := 10
b := 20
a, b = b, a
fmt.Println(a, b) // 20 10
```

## Converting String to Individual Characters

```go
// Convert string to rune slice so Unicode characters are safe
// []rune can safe more storage than []byte
s := hello
runes := []rune(s)
fmt.Println(string(runes)) // convert back to string
```
