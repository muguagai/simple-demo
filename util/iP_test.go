/******
** @创建时间 : 2022/6/12 12:36
** @作者 : MUGUAGAI
******/
package util

import (
	"fmt"
	"testing"
)

func TestGetAddress(t *testing.T) {
	ip, _ := ExternalIP()
	s := ip.String()
	fmt.Println(s)
}
