package merkle

import "crypto/sha256"

func BuildMerkleRoot(data [][]byte) []byte {
    if len(data) == 1 {
        return data[0]
    }

    var newLevel [][]byte

    for i := 0; i < len(data); i += 2 {
        if i+1 == len(data) {
            data = append(data, data[i])
        }

        combined := append(data[i], data[i+1]...)
        hash := sha256.Sum256(combined)
        newLevel = append(newLevel, hash[:])
    }

    return BuildMerkleRoot(newLevel)
}