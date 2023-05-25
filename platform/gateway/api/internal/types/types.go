// Code generated by goctl. DO NOT EDIT.
package types

type GatewayGetRegionRequest struct {
	Id int64 `path:"id"`
}

type GatewayGetRegionReply struct {
	Id       int64  `json:"id"`
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type GatewayGetRegionsRequest struct {
	ParentId int64 `path:"parent_id"`
}

type GatewayGetRegionsReply struct {
	List []GatewayGetRegionsChild `json:"list"`
}

type GatewayGetRegionsChild struct {
	Id       int64  `json:"id"`
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type GatewayGetRegionTreeRequest struct {
}

type GatewayGetRegionTreeReply struct {
	Tree []GatewayRegionTreeNode `json:"tree"`
}

type GatewayRegionTreeNode struct {
	Id       int64                   `json:"id"`
	ParentId int64                   `json:"parent_id"`
	Name     string                  `json:"name"`
	Children []GatewayRegionTreeNode `json:"children"`
}

type GatewayGetIndustryRequest struct {
	IndustryId string `path:"industry_id"`
}

type GatewayGetIndustryReply struct {
	Id          int64  `json:"id"`
	IndustryId  string `json:"industry_id"`
	Name        string `json:"name"`
	ParentId    string `json:"parent_id"`
	LevelType   int64  `json:"level_type"`
	Description string `json:"description"`
}

type GatewayGetIndustriesRequest struct {
	ParentId string `path:"parent_id"`
}

type GatewayGetIndustriesReply struct {
	List []GatewayGetIndustriesChild `json:"list"`
}

type GatewayGetIndustriesChild struct {
	Id          int64  `json:"id"`
	IndustryId  string `json:"industry_id"`
	Name        string `json:"name"`
	ParentId    string `json:"parent_id"`
	LevelType   int64  `json:"level_type"`
	Description string `json:"description"`
}

type GatewayGetIndustryTreeRequest struct {
}

type GatewayGetIndustryTreeReply struct {
	Tree []GatewayIndustryTreeNode `json:"tree"`
}

type GatewayIndustryTreeNode struct {
	Id          int64                     `json:"id"`
	IndustryId  string                    `json:"industry_id"`
	Name        string                    `json:"name"`
	ParentId    string                    `json:"parent_id"`
	LevelType   int64                     `json:"level_type"`
	Description string                    `json:"description"`
	Children    []GatewayIndustryTreeNode `json:"children"`
}

type GatewayAddRegionRequest struct {
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type GatewayAddRegionReply struct {
	Id       int64  `json:"id"`
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type GatewayRegisterRequest struct {
	Mobile         string `json:"mobile"`
	Email          string `json:"email"`
	AuthCredential string `json:"auth_credential"`
}

type GatewayRegisterReply struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type GatewayLoginRequest struct {
	Mobile         string `json:"mobile"`
	Email          string `json:"email"`
	AuthCredential string `json:"auth_credential"`
}

type GatewayLoginReply struct {
	Id           int64  `json:"id"`
	Mobile       string `json:"mobile"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type GatewayGetUserRequest struct {
	Id int64 `path:"id"`
}

type GatewayGetUserReply struct {
	Id       int64  `json:"id"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Gender   int64  `json:"gender"`
	Status   int64  `json:"status"`
}

type GatewayUpdateUserRequest struct {
	Id       int64  `path:"id"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Gender   int64  `json:"gender"`
	Status   int64  `json:"status"`
}

type GatewayUpdateUserReply struct {
	Id       int64  `json:"id"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Gender   int64  `json:"gender"`
	Status   int64  `json:"status"`
}
