## Slice
Declaring slice

```go
aSlice := []int{5, 2, 4, 6, 1, 3}
```

Copy slice
``` go
aSlice := []int{5, 2, 4, 6, 1, 3}
copySlice := make([]int, len(aSlice))
copy(copySlice, aSlice)
```
