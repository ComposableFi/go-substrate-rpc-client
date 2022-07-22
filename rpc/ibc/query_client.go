package ibc

import (
	"github.com/ComposableFi/go-substrate-rpc-client/v4/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (i IBC) QueryClientStateResponse(
	height int64,
	srcClientID string,
) (
	*clienttypes.QueryClientStateResponse,
	error,
) {
	var res *clienttypes.QueryClientStateResponse
	err := i.client.Call(&res, queryClientStateMethod, height, srcClientID)
	if err != nil {
		return &clienttypes.QueryClientStateResponse{}, err
	}
	return res, nil
}

func (i IBC) QueryClientConsensusState(
	clientid string,
	revisionHeight,
	revisionNumber uint64,
	latestConsensusState bool) (
	*clienttypes.QueryConsensusStateResponse,
	error,
) {
	var res *clienttypes.QueryConsensusStateResponse
	err := i.client.Call(&res,
		queryClientConsensusStateMethod,
		clientid,
		revisionHeight,
		revisionNumber,
		latestConsensusState)
	if err != nil {
		return &clienttypes.QueryConsensusStateResponse{}, err
	}
	return res, nil
}
func (i IBC) QueryUpgradedClient(height int64) (*clienttypes.QueryClientStateResponse, error) {
	var res *clienttypes.QueryClientStateResponse
	err := i.client.Call(&res, queryUpgradedClientMethod, height)
	if err != nil {
		return &clienttypes.QueryClientStateResponse{}, err
	}
	return res, nil
}

func (i IBC) QueryUpgradedConsState(
	height int64,
) (
	*clienttypes.QueryConsensusStateResponse,
	error,
) {
	var res *clienttypes.QueryConsensusStateResponse
	err := i.client.Call(&res, queryUpgradedConnectionStateMethod, height)
	if err != nil {
		return &clienttypes.QueryConsensusStateResponse{}, err
	}
	return res, nil
}

func (i IBC) QueryClients() (
	clienttypes.IdentifiedClientStates,
	error,
) {
	var res clienttypes.IdentifiedClientStates
	var x interface{}
	err := i.client.Call(&x, queryClientsMethod)
	if err != nil {
		return clienttypes.IdentifiedClientStates{}, err
	}
	return res, nil
}

func (i IBC) QueryNewlyCreatedClients(
	blockHash types.H256,
) (
	[]clienttypes.IdentifiedClientState,
	error,
) {
	var res []clienttypes.IdentifiedClientState
	err := i.client.Call(&res, queryNewlyCreatedClientsMethod)
	if err != nil {
		return []clienttypes.IdentifiedClientState{}, err
	}
	return res, nil
}
