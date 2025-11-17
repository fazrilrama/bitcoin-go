package transaction

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
)

type Wallet struct {
    PrivateKey ecdsa.PrivateKey
    PublicKey  []byte
}

func NewWallet() *Wallet {
    private, public := newKeyPair()
    return &Wallet{private, public}
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
    curve := elliptic.P256()

    private, _ := ecdsa.GenerateKey(curve, rand.Reader)
    pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

    return *private, pubKey
}
