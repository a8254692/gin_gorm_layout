// Package dto
// Generated by sql2struct-0.0.5 at 2024-06-19 07:59:34
package dto

type SysRoles struct {
	Id int64 `json:"id"`
}

type SysRolesChange struct {
	Role            Role            `json:"role"`
	RoleMenu        RoleMenu        `json:"role_menu"`
	RolePermissions RolePermissions `json:"role_permissions"` // 商户id
}

type Role struct {
	Id         int64  `json:"id"`
	Guid       string `json:"guid"`
	MerchantId int64  `json:"merchant_id"` // 商户id
	Name       string `json:"name"`        // 角色名称
	Code       string `json:"code"`        // 角色标识
	CreateTime int32  `json:"create_time"`
	UpdateTime int32  `json:"update_time"`
}

type RoleMenu struct {
	MenuIds []int64 `json:"menu_ids"`
}

type RolePermissions struct {
	PermissionIds []int64 `json:"permission_ids"`
}
