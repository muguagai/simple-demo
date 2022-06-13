/******
** @创建时间 : 2022/6/12 12:42
** @作者 : MUGUAGAI
******/
package config

import (
	"github.com/RaymondCode/simple-demo/util"
	"net"
)

var Ip net.IP

func GetAddress() {
	Ip, _ = util.ExternalIP()
}
