package esent

// CreateDatabaseFlag ...
type CreateDatabaseFlag uintptr

const (
	// Null ...
	Null uintptr = 0
	// DbDefault Empty flags
	DbDefault CreateDatabaseFlag = 0
	// DbOverwriteExisting The old database will be overwritten with a new one if exists
	DbOverwriteExisting CreateDatabaseFlag = 0x00000200
	// DbRecoveryOff Turns off logging
	DbRecoveryOff CreateDatabaseFlag = 0x00000008
	// DbShadowingOff Disable the duplication of some internal database structures (shadowing)
	DbShadowingOff CreateDatabaseFlag = 0x00000080
)
