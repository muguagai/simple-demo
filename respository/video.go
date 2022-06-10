package respository

import (
	"fmt"
	"strings"

	"github.com/RaymondCode/simple-demo/util"
)

func QueryByCreatedTime() (videos []Video, nexttime int64) {
	var Videos []Video
	Db.Table("videos").Order("create_time desc").Limit(30).Find(&Videos)
	//	fmt.Println(Videos[len(Videos)-1].CreateTime.Unix())
	//将Author和Video批量关联
	//不能使用range遍历
	for i := 0; i < len(Videos); i++ {
		var user User
		Db.Where("id = ?", Videos[i].AuthorID).Find(&user)
		Videos[i].Author = user
	}
	return Videos, Videos[len(Videos)-1].CreateTime.Unix()
}

func QueryVideosListByauthorid(user User) []Video {
	var Videos []Video
	Db.Where("author_id in (?)", user.Id).Find(&Videos)
	return Videos
}

func GetVideoListByIDs(ids []int64) (videos []Video) {
	var Videos []Video
	//Db.Where("id in (?)", ids).Find(&Videos)
	strs := util.Int2String(ids)
	fmt.Println(strs)
	//Db.Raw("select * from videos where id in '?' order by ?", ids, "desc").Scan(&Videos)
	strss := strings.Join(strs, ",")
	strss = "'" + strss + "'"
	Db.Raw("select * from videos where id in ? order by find_in_set(id,?)", ids, strss).Scan(&Videos)
	return Videos
}
