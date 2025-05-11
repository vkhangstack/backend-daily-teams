package utils

import (
	"github.com/godruoyi/go-snowflake"
	"time"
)

// NewSnowflakeService initializes a Snowflake generator
func NewSnowflakeService(nodeID uint16) {
	snowflake.SetMachineID(nodeID) // change to your machineID
	snowflake.SetStartTime(time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC))
}

// GenerateID creates a unique Snowflake ID
func GenerateID() uint64 {
	id := snowflake.ID()
	//idStr := strconv.FormatUint(uint64(id), 10)
	return id
}
