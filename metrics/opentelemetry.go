package metrics

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// Measures
var (
	BitswapMessagesReceivedCount = stats.Int64("bitswap_messages_received_total", "The total number of bitswap messages received", stats.UnitDimensionless)
	BitswapResponseCount         = stats.Int64("bitswap_response_total", "The total number of bitswap responses", stats.UnitDimensionless)
	BitswapRetrieverRequestCount = stats.Int64("bitswap_retriever_request_total", "The total number of bitswap messages that required a retriever lookup", stats.UnitDimensionless)
	BlockstoreCacheHitCount      = stats.Int64("blockstore_cache_hit_total", "The number of blocks from the local blockstore served to peers", stats.UnitDimensionless)
	BytesTransferredTotal        = stats.Int64("data_transferred_bytes_total", "The number of bytes transferred from storage providers to retrieval clients", stats.UnitBytes)
	RetrievalDealCost            = stats.Int64("retrieval_deal_cost_fil", "The cost in FIL of a retrieval deal with a storage provider", stats.UnitDimensionless)
	RetrievalDealCount           = stats.Int64("retrieval_deal_total", "The number of retrieval deals initiated with storage providers", stats.UnitDimensionless)
	RetrievalDealActiveCount     = stats.Int64("retrieval_deal_active_total", "The total number of active retrieval deals that have not yet succeeded or failed", stats.UnitDimensionless)
	RetrievalDealDuration        = stats.Float64("retrieval_deal_duration_seconds", "The duration in seconds of a retrieval deal with a storage provider", stats.UnitSeconds)
	RetrievalDealFailCount       = stats.Int64("retrieval_deal_fail_total", "The number of failed retrieval deals with storage providers", stats.UnitDimensionless)
	RetrievalDealSize            = stats.Int64("retrieval_deal_size_bytes", "The size in bytes of a retrieval deal with a storage provider", stats.UnitSeconds)
	RetrievalDealSuccessCount    = stats.Int64("retrieval_deal_success_total", "The number of successful retrieval deals with storage providers", stats.UnitDimensionless)

	// TODO: Add counts for known retrieval failure cases
)

// Tags
var (
	BitswapTopic, _ = tag.NewKey("bitswap_topic")
	Error, _        = tag.NewKey("error")
	Method, _       = tag.NewKey("method")
	Status, _       = tag.NewKey("status")
)

// Views
var (
	bitswapMessagesReceivedView = &view.View{
		Measure:     BitswapMessagesReceivedCount,
		Aggregation: view.Count(),
	}
	bitswapResponseView = &view.View{
		Measure:     BitswapResponseCount,
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{BitswapTopic},
	}
	bitswapRetreiverRequestView = &view.View{
		Measure:     BitswapRetrieverRequestCount,
		Aggregation: view.Count(),
	}
	blockstoreCacheHitView = &view.View{
		Measure:     BlockstoreCacheHitCount,
		Aggregation: view.Count(),
	}
	bytesTransferredView = &view.View{
		Measure:     BytesTransferredTotal,
		Aggregation: view.Sum(),
	}
	retrievalDealActiveView = &view.View{
		Measure:     RetrievalDealActiveCount,
		Aggregation: view.Count(),
	}
	retrievalDealCostView = &view.View{
		Measure:     RetrievalDealCost,
		Aggregation: view.Distribution(),
	}
	retrievalDealDurationView = &view.View{
		Measure:     RetrievalDealDuration,
		Aggregation: view.Distribution(0, 10, 20, 30, 40, 50, 60, 120, 240, 480, 540, 600),
	}
	retrievalDealFailView = &view.View{
		Measure:     RetrievalDealFailCount,
		Aggregation: view.Count(),
	}
	retrievalDealSuccessView = &view.View{
		Measure:     RetrievalDealSuccessCount,
		Aggregation: view.Count(),
	}
	retrievalDealSizeView = &view.View{
		Measure:     RetrievalDealSize,
		Aggregation: view.Distribution(),
	}
	retrievalDealTotalView = &view.View{
		Measure:     RetrievalDealCount,
		Aggregation: view.Count(),
	}
)

var DefaultViews = []*view.View{
	bitswapMessagesReceivedView,
	bitswapResponseView,
	bitswapRetreiverRequestView,
	blockstoreCacheHitView,
	bytesTransferredView,
	retrievalDealActiveView,
	retrievalDealCostView,
	retrievalDealDurationView,
	retrievalDealFailView,
	retrievalDealSuccessView,
	retrievalDealSizeView,
	retrievalDealTotalView,
}
