package migration

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrate member_point_logs and member_points from orders and payment_transaction_logs
func migrationMemberPointLogsByOrders(db *gorm.DB) error {
	fmt.Println("[migrationMemberPointLogsByOrders]")
	// TODO
	return nil
}

var MigrationMemberPointLogsByOrder = Migration{
	Name: "update-member-point-logs-by-orders",
	Up:   migrationMemberPointLogsByOrders,
}
