package transaction

type TXInput struct {
    Txid      []byte
    OutIndex  int
    Signature []byte
    PubKey    []byte
}
