package config

const (
	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// Service ports
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	// ElasticSearch
	ElasticIndex = "crawler_dating_profile"

	// RPC Endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
)
