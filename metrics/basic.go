package metrics

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	peer "github.com/libp2p/go-libp2p-core/peer"
)

type Basic struct {
	logger log.EventLogger
}

func NewBasic(logger log.EventLogger) *Basic {
	return &Basic{
		logger: logger,
	}
}

func (metrics *Basic) RecordWallet(info WalletInfo) {
	if info.Err != nil {
		metrics.logger.Warnf("Could not load any default wallet address, only free retrievals will be attempted: %v", info.Err)
	} else {
		metrics.logger.Infof("Using default wallet address %s", info.Addr)
	}
}

func (metrics *Basic) RecordGetCandidatesResult(info RequestInfo, result GetCandidatesResult) {
	if result.Err != nil {
		metrics.logger.Errorf("Could not get candidates: %v", result.Err)
	} else {
		if result.Count > 0 {
			metrics.logger.Infof(
				"Got %v candidates for %s",
				result.Count,
				info.RequestCid,
			)
		}
	}
}

func (metrics *Basic) RecordQuery(info CandidateInfo) {

}

func (metrics *Basic) RecordQueryResult(info CandidateInfo, result QueryResult) {
	if result.Err != nil {
		metrics.logger.Errorf(
			"Failed to query miner %s for %s: %v",
			info.PeerID,
			formatCidAndRoot(info.RequestCid, info.RootCid, false),
			result.Err,
		)
	}
}

func (metrics *Basic) RecordRetrieval(info CandidateInfo) {
	metrics.logger.Infof(
		"Attempting retrieval from miner %s for %s",
		info.PeerID,
		formatCidAndRoot(info.RequestCid, info.RootCid, false),
	)
}

func (metrics *Basic) RecordRetrievalResult(info CandidateInfo, result RetrievalResult) {
	if result.Err != nil {
		metrics.logger.Errorf(
			"Failed to retrieve from miner %s for %s: %v",
			info.PeerID,
			formatCidAndRoot(info.RequestCid, info.RootCid, false),
			result.Err,
		)
	} else {
		metrics.logger.Infof(
			"Successfully retrieved from miner %s for %s\n"+
				"\tDuration: %s\n"+
				"\tBytes Received: %s\n"+
				"\tTotal Payment: %s",
			info.PeerID,
			formatCidAndRoot(info.RequestCid, info.RootCid, false),
			result.Duration,
			humanize.IBytes(result.BytesReceived),
			result.TotalPayment,
		)
	}
}

func (metrics *Basic) RecordMinerConnection(peer peer.ID) {
	metrics.logger.Infof("Miner %s connected", peer)
}

func (metrics *Basic) RecordMinerDisconnection(peer peer.ID) {
	metrics.logger.Infof("Miner %s disconnected", peer)
}

func (metrics *Basic) RecordClientConnection(peer peer.ID) {
	metrics.logger.Infof("Client %s connected", peer)
}

func (metrics *Basic) RecordClientDisconnection(peer peer.ID) {
	metrics.logger.Infof("Client %s disconnected", peer)
}

func formatCidAndRoot(cid cid.Cid, root cid.Cid, short bool) string {
	if cid.Equals(root) {
		return formatCid(cid, short)
	} else {
		return fmt.Sprintf("%s (root %s)", formatCid(cid, short), formatCid(root, short))
	}
}

func formatCid(cid cid.Cid, short bool) string {
	str := cid.String()
	if short {
		return "..." + str[len(str)-10:]
	} else {
		return str
	}
}
