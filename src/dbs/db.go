package dbs

import (
	"src/models"
)

// global "database" of all receipt IDs to JSON receipt objects
var AllReceipts = make(map[string]models.Receipt)