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

type GatewayAddRegionRequest struct {
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type GatewayAddRegionReply struct {
	Id int64 `json:"id"`
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

type GatewayAddIndustryRequest struct {
	IndustryId  string `json:"industry_id"`
	Name        string `json:"name"`
	ParentId    string `json:"parent_id"`
	LevelType   int64  `json:"level_type"`
	Description string `json:"description"`
}

type GatewayAddIndustryReply struct {
	Id int64 `json:"id"`
}

type GatewayGetCategoryRequest struct {
	Id int64 `path:"id"`
}

type GatewayGetCategoryReply struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	ParentId    int64  `json:"parent_id"`
	Description string `json:"description"`
}

type GatewayGetCategoriesRequest struct {
	ParentId int64 `path:"parent_id"`
}

type GatewayGetCategoriesReply struct {
	List []GatewayGetCategoriesChild `json:"list"`
}

type GatewayGetCategoriesChild struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	ParentId    int64  `json:"parent_id"`
	Description string `json:"description"`
}

type GatewayGetCategoryTreeRequest struct {
}

type GatewayGetCategoryTreeReply struct {
	Tree []GatewayCategoryTreeNode `json:"tree"`
}

type GatewayCategoryTreeNode struct {
	Id          int64                     `json:"id"`
	Name        string                    `json:"name"`
	ParentId    int64                     `json:"parent_id"`
	Description string                    `json:"description"`
	Children    []GatewayCategoryTreeNode `json:"children"`
}

type GatewayAddCategoryRequest struct {
	Name        string `json:"name"`
	ParentId    int64  `json:"parent_id"`
	Description string `json:"description"`
}

type GatewayAddCategoryReply struct {
	Id int64 `json:"id"`
}

type GatewayGetClassifyRequest struct {
	Id int64 `path:"id"`
}

type GatewayGetClassifyReply struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	ParentId    int64  `json:"parent_id"`
	Description string `json:"description"`
}

type GatewayGetClassifiesRequest struct {
	ParentId int64 `path:"parent_id"`
}

type GatewayGetClassifiesReply struct {
	List []GatewayGetClassifiesChild `json:"list"`
}

type GatewayGetClassifiesChild struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	ParentId    int64  `json:"parent_id"`
	Description string `json:"description"`
}

type GatewayGetClassifyTreeRequest struct {
}

type GatewayGetClassifyTreeReply struct {
	Tree []GatewayClassifyTreeNode `json:"tree"`
}

type GatewayClassifyTreeNode struct {
	Id          int64                     `json:"id"`
	Name        string                    `json:"name"`
	ParentId    int64                     `json:"parent_id"`
	Description string                    `json:"description"`
	Children    []GatewayClassifyTreeNode `json:"children"`
}

type GatewayAddClassifyRequest struct {
	Name        string `json:"name"`
	ParentId    int64  `json:"parent_id"`
	Description string `json:"description"`
}

type GatewayAddClassifyReply struct {
	Id int64 `json:"id"`
}

type GatewayGetStageRequest struct {
	Id int64 `path:"id"`
}

type GatewayGetStageReply struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GatewayGetStagesRequest struct {
}

type GatewayGetStagesReply struct {
	List []GatewayGetStagesChild `json:"list"`
}

type GatewayGetStagesChild struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GatewayAddStageRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GatewayAddStageReply struct {
	Id int64 `json:"id"`
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
