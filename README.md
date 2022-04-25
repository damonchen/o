# o 
ternary operator for golang, only for the same return type


## example

```go

import (
	"github.com/damonchen/o"
)

v := o.T(p == nil, "not exist", "exits")

```

lazy evaluate

```go

import (
	"github.com/damonchen/o"
)

v := o.L(p == nil, o.W("not exist"), o.W("exist"))

```

complicated example

```go

import (
	"github.com/damonchen/o"
)

v := o.L(p == nil, o.W("unknown"), func() string{
	return o.T(p.gender == 1 ? "man": "women")
})


```
