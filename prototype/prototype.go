package prototype

import "encoding/json"

// 克隆接口
type ICloneable interface {
	// 克隆得到自己
	Clone() ICloneable
}

// 用户信息
type UserInfo struct {
	ID       int
	Name     string
	RoleList []string
}

func newUserInfo() *UserInfo {
	return &UserInfo{}
}

func (u *UserInfo) Clone() ICloneable {
	it := &UserInfo{
		ID:       u.ID,
		Name:     u.Name,
		RoleList: make([]string, 0, len(u.RoleList)),
	}
	for _, role := range u.RoleList {
		it.RoleList = append(it.RoleList, role)
	}
	return it
}

var defaultUserFactory = newUserFactory()

// 用户信息工厂
type IUserFactory interface {
	Create() *UserInfo
}

type userFactory struct {
	u *UserInfo
}

func newUserFactory() *userFactory {
	config := mockUserConfig()
	u := &UserInfo{}
	if err := json.Unmarshal(config, u); err != nil {
		panic(err)
	}

	return &userFactory{u: u}
}

func mockUserConfig() []byte {
	config, _ := json.Marshal(&UserInfo{
		ID:       1,
		Name:     "新一",
		RoleList: []string{"guest"},
	})
	return config
}

func (f *userFactory) Create() *UserInfo {
	return f.u.Clone().(*UserInfo)
}


