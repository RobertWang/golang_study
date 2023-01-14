// subpack/subpack.go
package subpack

import (
	"fmt"
)

var VERSION string = "0.6"

func Sub() string { //注意这里的首字母大写
	fmt.Println("subpack_name sub func")
	return "subpack_name sub func"
}
