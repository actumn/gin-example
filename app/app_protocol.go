package app

type LoginReq struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpReq struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type BookReq struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
