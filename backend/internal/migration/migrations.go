package migration

var AllMigrations = []Migration{
	MigrationOrderCustomerToMember,
	MigrationMemberPointLogsByOrder,
}
