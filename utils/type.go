package utils

type Post struct {
	Id            int      `json:"id"`
	UserName      string   `json:"userName"`
	Title         string   `json:"title"`
	Contant       string   `json:"contant"`
	ImgUrl        any      `json:"img"`
	Categories    []string `json:"categories"`
	Reaction      Reaction `json:"reaction"`
	Idscategories []int    `json:"-"`
	Date          string   `json:"date"`
}

type Reaction struct {
	NumLike    int    `json:"numLike"`
	NumDisLike int    `json:"numDisLike"`
	Type       string `json:"type"` // here i can know if user like post or not
}

type Commant struct {
	Id       int      `json:"id"`
	UserId   int      `json:"-"`
	Post_id  int      `json:"-"`
	UserName string   `json:"userName"`
	Contant  string   `json:"contant"`
	Reaction Reaction `json:"reaction"`
	Date     string   `json:"date"`
}

type User struct {
	User_name string `json:"user_name"`
	Email     string `json:"email"`
	Passwd    string `json:"passwd"`
}

type Users struct {
	User_name string `json:"user_name"`
	Uid       string `json:"uid"`
	Online    bool   `json:"online"`
}

type Message struct {
	From     string
	FromName string
	To       string
	Message  string
	Date     string
}
