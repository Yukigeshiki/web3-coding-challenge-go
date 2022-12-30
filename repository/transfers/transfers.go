package transfers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"strconv"
)

// A very similar caching setup could be done using an external in-memory database like Redis, but I thought for
// simplicity I'd do it here in the application.

const (
	bHex    int = 16
	bitSize int = 64
)

type LogSlice []*types.Log

type LogMap map[string]LogSlice

type FilterOpts struct {
	From, To     string
	Above, Below int64
}

var (
	From LogMap
	To   LogMap
	Both LogMap
	//All  LogSlice
)

// GetLogs fetches logs from the in-memory cache depending on the filter values
func GetLogs(opts *FilterOpts) (LogSlice, error) {

	logs := make(LogSlice, 0)

	f := opts.From
	t := opts.To
	a := opts.Above
	b := opts.Below

	switch {
	case f != "" && t != "":
		logs = Both[f+":"+t]
	case f != "":
		logs = From[f]
	case t != "":
		logs = To[t]
		//default:
		//	logs = All
	}

	if a > 0 || b > 0 {
		logsFiltered := make(LogSlice, 0)
		for _, log := range logs {
			val, err := strconv.ParseInt(common.Bytes2Hex(log.Data), bHex, bitSize)
			if err != nil {
				return nil, err
			}
			if (a > 0 && val < a) || (b > 0 && val > b) {
				continue
			}
			logsFiltered = append(logsFiltered, log)
		}
		return logsFiltered, nil
	}

	return logs, nil
}

func init() {
	From = make(LogMap)
	To = make(LogMap)
	Both = make(LogMap)
	//All = make(LogSlice, 0)
}
