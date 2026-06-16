package transaction

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"math/big"
)

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = "Coinbase"
	}
	tx := &Transaction{
		Vin:  []TXInput{{Txid: []byte{}, OutIndex: -1, Signature: nil, PubKey: []byte(data)}},
		Vout: []TXOutput{{Value: 50, PubKeyHash: []byte(to)}},
	}
	tx.SetID()
	return tx
}

func (tx *Transaction) SetID() {
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(tx)
	hash := sha256.Sum256(buf.Bytes())
	tx.ID = hash[:]
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && tx.Vin[0].OutIndex == -1
}

func (tx *Transaction) Sign(privKey ecdsa.PrivateKey) {
	if tx.IsCoinbase() {
		return
	}
	hash := sha256.Sum256(tx.ID)
	r, s, _ := ecdsa.Sign(rand.Reader, &privKey, hash[:])
	tx.Vin[0].Signature = append(r.Bytes(), s.Bytes()...)
}

func (tx *Transaction) Verify(pubKey []byte) bool {
	if tx.IsCoinbase() {
		return true
	}
	curve := elliptic.P256()
	keyLen := len(pubKey) / 2
	r := new(big.Int).SetBytes(tx.Vin[0].Signature[:len(tx.Vin[0].Signature)/2])
	s := new(big.Int).SetBytes(tx.Vin[0].Signature[len(tx.Vin[0].Signature)/2:])
	x := new(big.Int).SetBytes(pubKey[:keyLen])
	y := new(big.Int).SetBytes(pubKey[keyLen:])
	pub := ecdsa.PublicKey{Curve: curve, X: x, Y: y}
	hash := sha256.Sum256(tx.ID)
	return ecdsa.Verify(&pub, hash[:], r, s)
}
