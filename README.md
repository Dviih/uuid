# UUID

## UUID implementation for Go, using a `complex128` to represent it.

---

## UUID
### The `UUID` is underlying a `complex` with methods for UUID.

---

# Usage

## UUID
- `MarshalBinary` - Marshal the uuid into a byte array.
- `UnmarshalBinary` - Unmarshal the byte array into uuid.
- `String` - A string representation of the uuid.

## V4
### Generates a new uuid v4 using `math/rand/v2` and returns `UUID`
###### New is the same as calling V4

## V7
### Generates a new uuid v7 using `time` and `math/rand/v2` and returns `UUID`

---

# Example

```go
package main

import (
	"fmt"
	"github.com/Dviih/uuid"
)

func main() {
	fmt.Println(uuid.New())
}
```

---
#### Made for Gophers by @Dviih