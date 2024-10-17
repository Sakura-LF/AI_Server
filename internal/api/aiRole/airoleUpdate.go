package aiRole

type AiRoleUpdateUserRequest struct {
	RoleID   uint    `json:"roleID" binding:"required"`
	Title    *string `json:"title" mps:"title"`
	Avatar   *string `json:"avatar" mps:"avatar"`
	Category *string `json:"category" mps:"category"` // 角色分类
	Abstract *string `json:"abstract" mps:"abstract"`
	Prompt   *string `json:"prompt" mps:"prompt"` // 提示词
	Reason   string  `json:"reason"`              // 工单理由
}
