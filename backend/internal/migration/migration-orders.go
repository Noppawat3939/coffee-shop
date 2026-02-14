package migration

import (
	"backend/models"
	"fmt"

	"gorm.io/gorm"
)

// Migrate orders column customer not guest to member_id
func migrationOrderCustomerToMember(db *gorm.DB) error {
	fmt.Println("[migrationOrderCustomerToMember]")
	var orders []models.Order

	if err := db.Where("member_id IS NULL").Where("customer <> ?", "guest").Find(&orders).Error; err != nil {
		return err
	}

	if len(orders) == 0 {
		fmt.Println("stop migration: orders not found")
		return nil
	}

	var names []string
	for _, order := range orders {
		names = append(names, order.Customer)
	}

	var members []models.Member

	if err := db.Where("full_name IN ?", names).Find(&members).Error; err != nil {
		return err
	}

	if len(members) == 0 {
		fmt.Println("stop migration: member not found")
		return nil
	}

	memberMap := make(map[string]uint)

	for _, member := range members {
		memberMap[member.FullName] = member.ID
	}

	var totalUpdated int
	for _, order := range orders {
		memberID, exits := memberMap[order.Customer]
		if !exits {
			continue
		}

		res := db.Model(&models.Order{}).Where("id = ?", order.ID).Update("member_id", memberID)
		if res.Error != nil {
			return res.Error
		}

		totalUpdated += int(res.RowsAffected)
	}

	fmt.Printf("migration completed total: %d orders\n", totalUpdated)

	return nil
}

var MigrationOrderCustomerToMember = Migration{
	Name: "orders-customer-to-members",
	Up:   migrationOrderCustomerToMember,
}
