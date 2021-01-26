package testutils

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	// TestDefaultPage is the default number of pages for testing
	TestDefaultPage = 1
	// TestInvalidPage is the invalid number of pages for testing
	TestInvalidPage = 0
	// TestDefaultPerPage is the default number of pages per page for testing
	TestDefaultPerPage = 2
	// TestInvalidPerPage is the invalid number of pages per page for testing
	TestInvalidPerPage = 0
	// TestDefaultLanguageCode is default language for testing
	TestDefaultLanguageCode = "ja"
	// TestInvalidLanguageCode is invalid language for testing
	TestInvalidLanguageCode = "invalid_language_code"
	// TestDefaultCountryID is default country ID for testing
	TestDefaultCountryID = "JPN"
	// TestInvalidCountryID is invalid country ID for testing
	TestInvalidCountryID = "invalid_country_id"
	// TestDefaultIndicatorID is default indicator ID for testing
	TestDefaultIndicatorID = "NY.GDP.MKTP.CD"
	// TestInvalidIndicatorID is invalid indicator ID for testing
	TestInvalidIndicatorID = "INVALID.INDICATOR.ID"
	// TestDefaultTopicID is default topic ID for testing
	TestDefaultTopicID = "1"
	// TestInvalidTopicID is invalid topic ID for testing
	TestInvalidTopicID = "invalid_topic_id"
	// TestDefaultSourceID is default source ID for testing
	TestDefaultSourceID = "2"
	// TestInvalidSourceID is invalid source ID for testing
	TestInvalidSourceID = "invalid_source_id"
	// TestDefaultRegionCode is default region ID for testing
	TestDefaultRegionCode = "AFR"
	// TestInvalidRegionCode is invalid region ID for testing
	TestInvalidRegionCode = "invalid_region_id"
	// TestDefaultIncomeLevelID is default income level ID for testing
	TestDefaultIncomeLevelID = "HIC"
	// TestInvalidIncomeLevelID is invalid income level ID for testing
	TestInvalidIncomeLevelID = "invalid_income_level_id"
	// TestDefaultLendingTypeID is default lending type ID for testing
	TestDefaultLendingTypeID = "IBD"
	// TestInvalidLendingTypeID is invalid lending type ID for testing
	TestInvalidLendingTypeID = "invalid_lending_type_id"
	// TestDefaultDateStart is default date start for testing
	TestDefaultDateStart = "2018"
	// TestDefaultDateEnd is default date end for testing
	TestDefaultDateEnd = "2019"
	// TestInvalidDate is invalid date for testing
	TestInvalidDate = "invalid_date"
)

var (
	// TestDefaultCountryIDs is default country IDs for testing
	TestDefaultCountryIDs = []string{"JPN", "USA"}
	// TestInvalidCountryIDs is invalid country IDs for testing
	TestInvalidCountryIDs = []string{"ABCDEFG", "HIJKLM"}
	// TestDefaultIndicatorIDs is default indicator IDs for testing
	TestDefaultIndicatorIDs = []string{"NY.GDP.MKTP.CD", "SP.POP.TOTL"}
	// TestInvalidIndicatorIDs is invalid indicator IDs for testing
	TestInvalidIndicatorIDs = []string{"INVALID.INDICATOR.ID", "SP.POP.TOTL"}
)

// UpdateFixture removes fixtures when `update` is true.
func UpdateFixture(update *bool) {
	if *update {
		fixtureDir := filepath.Join("testdata", "fixtures")
		os.RemoveAll(fixtureDir)
		if err := os.MkdirAll(fixtureDir, 0755); err != nil {
			panic(fmt.Sprintf("failed to create fixture dirs: %s", err))
		}
	}
}
