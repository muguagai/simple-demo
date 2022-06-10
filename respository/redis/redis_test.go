/******
** @创建时间 : 2022/6/1 11:11
** @作者 : MUGUAGAI
******/
package redis

import "testing"

func TestInit(t *testing.T) {
	InitClient()
	GetFavouriteVideo(539203490925252608)
}
