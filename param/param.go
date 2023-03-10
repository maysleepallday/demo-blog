package param

type UserRegister struct {
	Name       string `json:"user_name"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

type UserLogin struct {
	Name     string `json:"user_name"`
	Password string `json:"password"`
}

type PostVote struct {
	PostID int64 `json:"post_id"`
	Value  int   `json:"vote_value"`
}

type PostCreate struct {
	Title       string `json:"post_title" binding:"required"`
	CommunityID int64  `json:"community_id" binding:"required"`
	UserID      int64
}
