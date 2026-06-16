package p2p

const (
	CmdVersion = "version"
	CmdGetData = "getdata"
	CmdBlock   = "block"
	CmdTx      = "tx"
)

type Message struct {
	Command string
	Payload []byte
}
