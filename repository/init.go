package repository

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"robothouse.ui/web3-coding-challenge/config"
	"robothouse.ui/web3-coding-challenge/lib/ethereum/contracts"
	l "robothouse.ui/web3-coding-challenge/lib/log"
	tf "robothouse.ui/web3-coding-challenge/repository/transfers"
)

const blockIncrement = 200_000

var (
	errChan = make(chan error, 1)
	sink    = make(chan *contracts.ERC20Transfer)
)

// InitialiseRepo fetches transfer logs using an ERC20 contract instance and stores them in an in-memory cache
func InitialiseRepo(inst *contracts.ERC20) error {
	var (
		opts bind.FilterOpts
		from []common.Address
		to   []common.Address
	)

	opts.Start = config.ContractStartBlock
	end := opts.Start + blockIncrement
	opts.End = &end

	for {
		logs := make(tf.LogSlice, 0)
		iter, err := inst.FilterTransfer(&opts, from, to)
		if err != nil {
			return err
		}
		for iter.Next() {
			log, err := getAndIndexLog(iter.Event)
			if err != nil {
				l.Error("InitialiseRepo: "+err.Error(), nil)
			}
			logs = append(logs, log)
		}

		if len(logs) == 0 {
			break
		}

		//tf.All = append(tf.All, logs...)

		opts.Start += blockIncrement
		end = opts.Start + blockIncrement
		opts.End = &end
	}

	return nil
}

// InitialiseRepoLiveUpdate creates a subscription to a contract's transfer logs, and indexes logs as they become available
func InitialiseRepoLiveUpdate(inst *contracts.ERC20) {
	ts := transferSubscription{
		opts: new(bind.WatchOpts),
		from: make([]common.Address, 0),
		to:   make([]common.Address, 0),
	}
	go ts.subscribe(inst, sink, errChan)

	for {
		select {
		case err := <-errChan:
			// if there is an error, re-subscribe
			l.Error("InitialiseRepoLiveUpdate: "+err.Error()+" - re-subscribing...", nil)
			go ts.subscribe(inst, sink, errChan)
		case event := <-sink:
			// this is the only write operation to the index maps (post cache population), so a mutex shouldn't be needed
			if _, err := getAndIndexLog(event); err != nil {
				l.Error("InitialiseRepoLiveUpdate: "+err.Error(), nil)
			}
			l.Debug("InitialiseRepoLiveUpdate: new transfer log added to index maps", nil)
			//tf.All = append(tf.All, log)
		}
	}
}

func CloseSubscriptionChannels() {
	close(errChan)
	close(sink)
}

// getAndIndexLog gets a transfer log from an event, indexes it, then returns it
func getAndIndexLog(event *contracts.ERC20Transfer) (*types.Log, error) {
	jsonBytes, err := event.Raw.MarshalJSON()
	if err != nil {
		return nil, err
	}
	log := new(types.Log)
	if err = json.Unmarshal(jsonBytes, log); err != nil {
		return nil, err
	}

	fAddr := getStringAddress(log.Topics[1])
	tAddr := getStringAddress(log.Topics[2])
	bAddr := fAddr + ":" + tAddr

	// index transfer log
	tf.From[fAddr] = append(tf.From[fAddr], log)
	tf.To[tAddr] = append(tf.To[tAddr], log)
	tf.Both[bAddr] = append(tf.Both[bAddr], log)

	return log, nil
}

// getStringAddress takes an address hash and coverts it into a hex string representation
func getStringAddress(a common.Hash) string {
	return "0x" + fmt.Sprintf("%x", a.Big())
}

type transferSubscription struct {
	opts *bind.WatchOpts
	from []common.Address
	to   []common.Address
}

// subscribe calls WatchTransfer on the contract instance - if an error occurs it is passed into the error channel
func (ts *transferSubscription) subscribe(inst *contracts.ERC20, sink chan<- *contracts.ERC20Transfer, errChan chan<- error) {
	sub, err := inst.WatchTransfer(ts.opts, sink, ts.from, ts.to)
	if err != nil {
		errChan <- err
		return
	}

	defer sub.Unsubscribe()

	for {
		select {
		case err = <-sub.Err():
			errChan <- err
			return
		}
	}
}
