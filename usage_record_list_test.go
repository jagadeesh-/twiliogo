package twiliogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationUsageRecordList(t *testing.T) {
	CheckTestEnv(t)

	client := NewClient(API_KEY, API_TOKEN)

	usageRecordList, err := GetUsageRecordList(client, "LastMonth")

	if assert.Nil(t, err, "Failed to retrieve usage record list") {
		usageRecords := usageRecordList.GetUsageRecords()
		assert.NotNil(t, usageRecords, "Failed to retrieve usageRecords")
	}
}
