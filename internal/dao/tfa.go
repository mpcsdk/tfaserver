// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"tfaserver/internal/dao/internal"
)

// internalTfaDao is internal type for wrapping internal DAO implements.
type internalTfaDao = *internal.TfaDao

// tfaDao is the data access object for table tfa.
// You can define custom methods on it to extend its functionality as you wish.
type tfaDao struct {
	internalTfaDao
}

var (
	// Tfa is globally public accessible object for table tfa operations.
	Tfa = tfaDao{
		internal.NewTfaDao(),
	}
)

// Fill with you ideas below.
