/******
** @创建时间 : 2022/6/1 11:11
** @作者 : MUGUAGAI
******/
package redis

import (
	"github.com/RaymondCode/simple-demo/respository"
	"testing"
)

func TestInit(t *testing.T) {
	respository.Init()
	InitClient()
	GetFavouriteVideo(539203490925252608)
	FavouriteToMysql()
}
