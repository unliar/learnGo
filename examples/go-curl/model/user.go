package model

// UserBase 是用户的基础信息
type UserBase struct {
	Id   int    `json:"id" `       			// 用户id
	LoginName string `json:"login_name"`    // 登录用户名
	RealName string `json:"name"`  			// 用户真实姓名
	IDC string `json:"idc"`                 // 用户的身份证号
	Nickname string `json:"nickname"` 		// 用户昵称
	Age  int    `json:"age"` 				// 用户年龄
	Male string `json:"male"` 				// 用户性别
	Avatar string `json:"avatar"`           // 头像信息
 	Location string `json:"location"`   	// 用户的位置信息
	Profession string `json:"profession"`   // 用户的职业信息
	Status int `json:"status"`              // 用户的可用状态
	Token string `json:"token"`             // 用户的当前登录token
	UpdateTime string `json:"update_time"`	// 更新时间
	CreateTime string `json:"create_time"`  // 创建时间
}

// UserAuth 是用户的密码验证信息
type UserAuth struct {
	Id int	`json:"id"` 					// 自增id
	Uid int `json:"uid"` 					// 对应的用户id
	Salt string `json:"salt"`               // 用户加密使用的key
	Password string `json:"password"`		//用户加密后的密码
	IsCurrent int `json:"is_current"` 		// 是否是当前生效的密码
	UpdateTime string `json:"update_time"`	// 用户密码更新时间
	CreateTime string `json:"create_time"`  // 创建密码的时间
}

// UserDataStatistics 是用户的数据统计
type UserDataStatistics struct {
	Id int	`json:"id"` 										// 自增id
	Uid int `json:"uid"` 										// 对应的用户id
	FollowCount int `json:"follow_count"`   					// 关注的人个数
	FansCount int `json:"fans_count"`       					// 粉丝个数
	CommentSendCount int `json:"comment_send_count"` 			// 发出评论个数
	CommentGetCount int `json:"comment_get_count"`   			// 获得评论个数
	LikeSendCount int `json:"like_send_count"`       			// 发出点赞个数
	likeGetCount int `json:"like_get_count"`         			// 获得点赞个数
	CollectionSendCount int `json:"collection_send_count"`      // 发出收藏个数
	CollectionGetCount int `json:"collection_get_count"`        // 获得收藏个数
	UpdateTime string `json:"update_time"`                      // 更新时间
	CreateTime string `json:"create_time"`                      // 创建时间
 }

// UserAccountBind 是第三方账户的绑定信息
 type UserAccountBind struct {
	 Id int	`json:"id"` 							// 自增id
	 Uid int `json:"uid"` 						    // 对应的用户id
	 Email string `json:"email"`                    // 邮件地址
	 EmailStatus int `json:"email_status"`          // 邮箱状态
	 PhoneType int `json:"phone_type"`              // 手机区号
	 Phone string `json:"phone"`                    // 手机号码
	 PhoneStatus int `json:"phone_status""`         // 手机状态
	 WeChatId string `json:"we_chat_id""`           // 微信id
	 WeiBoId string `json:"wei_bo_id"`              // 微博id
	 QQId string `json:"qq_id"`                     // QQ id
	 UpdateTime string `json:"update_time"`         // 更新时间
	 CreateTime string `json:"create_time"`         // 创建时间
 }