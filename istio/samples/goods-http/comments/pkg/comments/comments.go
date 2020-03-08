package comments

type CommentsInfo struct {
	UserID int64 `json:"user_id"`
	UserName string `json:"user_name"`
	UserPic string `json:"user_pic"`
	Comment string `json:"comment"`
}

