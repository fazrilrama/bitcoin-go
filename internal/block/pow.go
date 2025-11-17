package block

import (
    "bytes"
    "crypto/sha256"
    "fmt"
    "math/big"
)

const Difficulty = 18

type ProofOfWork struct {
    Block  *Block
    Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
    target := big.NewInt(1)
    target.Lsh(target, uint(256-Difficulty))

    return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) Run() (int, []byte) {
    var hashInt big.Int
    var hash [32]byte
    nonce := 0

    for {
        data := pow.prepareData(nonce)
        hash = sha256.Sum256(data)

        hashInt.SetBytes(hash[:])

        if hashInt.Cmp(pow.Target) == -1 {
            fmt.Printf("Mined block: %x\n", hash)
            break
        } else {
            nonce++
        }
    }

    return nonce, hash[:]
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
    return bytes.Join(
        [][]byte{
            pow.Block.PrevHash,
            pow.Block.Data,
            []byte(fmt.Sprintf("%d", pow.Block.Timestamp)),
            []byte(fmt.Sprintf("%d", nonce)),
        },
        []byte{},
    )
}
