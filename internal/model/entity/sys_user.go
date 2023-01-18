// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id        int    `json:"id"        ` //
	Gid       int    `json:"gid"       ` // 组织id
	GroupRoot bool   `json:"groupRoot" ` // 是否组织总帐号 0否 1是
	Account   string `json:"account"   ` // 帐号
	Password  string `json:"password"  ` // 密码
	Nickname  string `json:"nickname"  ` // 昵称
	Avatar    string `json:"avatar"    ` // 头像
	Status    bool   `json:"status"    ` // 状态 0禁用 1正常
	Token     string `json:"token"     ` // 登陆标志
}