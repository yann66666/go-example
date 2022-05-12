// Author: yann
// Date: 2022/5/12
// Desc: errors

package errors

import "fmt"

var (
	ErrSeasTypeCannotNull = fmt.Errorf("SeasType 类型非法 0=内海 2=公海")
)
