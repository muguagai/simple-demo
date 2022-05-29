package respository

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
