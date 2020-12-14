package config

const (
	// Service ports
	ItemSaverPort = "1234"
	WorkerPort0   = "9000"
	// ElasticSearch
	ElasticIndex = "dating_pofile"
	// RPC Endpoint
	ItemSaverRpc = "ItemSaverService.Save"

	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"

	NilParser = "NilParser"

	CrawlServiceRpc = "CrawlService.Process"

	Qps = 20
)
