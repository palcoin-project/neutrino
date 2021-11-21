package neutrino

import (
	"github.com/palcoin-project/neutrino/blockntfns"
	"github.com/palcoin-project/neutrino/pushtx"
	"github.com/palcoin-project/neutrino/query"
	"github.com/palcoin-project/palcd/addrmgr"
	"github.com/palcoin-project/palcd/blockchain"
	"github.com/palcoin-project/palcd/connmgr"
	"github.com/palcoin-project/palcd/peer"
	"github.com/palcoin-project/palcd/txscript"
	"github.com/palcoin-project/palclog"
)

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log palclog.Logger

// The default amount of logging is none.
func init() {
	DisableLog()
}

// DisableLog disables all library log output.  Logging output is disabled
// by default until either UseLogger or SetLogWriter are called.
func DisableLog() {
	log = palclog.Disabled
}

// UseLogger uses a specified Logger to output package logging info.
// This should be used in preference to SetLogWriter if the caller is also
// using palclog.
func UseLogger(logger palclog.Logger) {
	log = logger
	blockchain.UseLogger(logger)
	txscript.UseLogger(logger)
	peer.UseLogger(logger)
	addrmgr.UseLogger(logger)
	blockntfns.UseLogger(logger)
	pushtx.UseLogger(logger)
	connmgr.UseLogger(logger)
	query.UseLogger(logger)
}
