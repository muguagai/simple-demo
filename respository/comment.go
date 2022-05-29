package respository

func QueryCommentListByVideoid(videoid int64) []Comment {
	var comments []Comment
	Db.Where("video_id in (?)", videoid).Find(&comments)
	return comments
}
