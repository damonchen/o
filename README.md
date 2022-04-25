# t
ternary operator for golang, only for the same return type


## example

```go

import (
		"github.com/damonchen/t"
	   )	   

v := t.T(p == nil, "not exist", "exits")

```

lazy evaluate

```go

import (
		"github.com/damonchen/t"
	   )

v := t.T(p == nil, L("not exist"), L("exist"))

```
