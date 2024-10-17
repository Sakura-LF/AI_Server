package workOrder

import (
	"AI_Server/init/data"
	"AI_Server/internal/models"
)

func CreateWorkOrder(userID uint, roleID uint, reason string, workOrderType int8, aiRoleUpdateData []byte) (*models.AiRoleWorkOrder, error) {
	workOrder := &models.AiRoleWorkOrder{
		UserID:           userID,
		AiRoleID:         roleID,
		Reason:           reason,
		Type:             workOrderType,
		Status:           1,
		AiRoleUpdateData: string(aiRoleUpdateData),
	}
	if err := data.DB.Create(workOrder).Error; err != nil {
		return nil, err
	}
	return workOrder, nil
}

func FindWorkOrder(aiRoleID uint, workOrderType int8) (*models.AiRoleWorkOrder, error) {
	var workOrder models.AiRoleWorkOrder
	if err := data.DB.Take(&workOrder, "ai_role_id = ? and type = ? and status = ?", aiRoleID, workOrderType, 1).Error; err != nil {
		return nil, err
	}
	return &workOrder, nil
}
