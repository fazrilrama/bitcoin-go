# 🧱 Bitcoin-Go

Implementasi minimal dari *Bitcoin Whitepaper (Satoshi Nakamoto)* menggunakan bahasa Go.  
Dibangun untuk mempelajari konsep fundamental Bitcoin secara hands-on.

---

## ✅ Fitur yang Sudah Berjalan

### 🔗 Blockchain
- Struktur block: `Timestamp`, `Data`, `PrevHash`, `Hash`, `Nonce`
- Genesis block otomatis saat blockchain dibuat
- Hashing SHA-256

### ⛏️ Proof of Work
- Difficulty target: `256 - 18` (bit shift)
- Mining loop hingga hash memenuhi target
- Nonce disimpan di block

### 🌳 Merkle Tree
- Fungsi `BuildMerkleRoot` menghasilkan merkle root dari slice `[][]byte`
- Rekursif, menduplikasi node ganjil sesuai standar Bitcoin
- ⚠️ Belum diintegrasikan ke dalam block

### 🔐 Wallet (ECDSA)
- Generate private key + public key dengan Elliptic Curve P-256
- Struct `Wallet`, `TXInput`, `TXOutput`, `Transaction` sudah didefinisikan
- ⚠️ Logic sign, verify, dan UTXO belum diimplementasikan

### 🖥️ CLI Node
- `add <data>` — tambah block baru ke chain
- `print` — tampilkan semua block

---

## 🚧 Dalam Pengembangan

| Fitur | Status |
|---|---|
| Integrasi Merkle Root ke block | 🔲 Belum |
| Transaksi (sign & verify) | 🔲 Belum |
| UTXO model | 🔲 Belum |
| P2P networking | 🔲 Belum (file kosong) |

---

## 📦 Instalasi

```bash
git clone https://github.com/fazrilrama/bitcoin-go.git
cd bitcoin-go
go mod tidy
```

## ▶️ Cara Menjalankan

```bash
# Tambah block baru
go run cmd/node/main.go add "Hello Bitcoin"

# Lihat seluruh blockchain
go run cmd/node/main.go print
```

## 📌 Contoh Output

```
Mined block: 00000af23c...

--- Block ---
Timestamp: 1716000000
Data: Hello Bitcoin
PrevHash: 00000abc...
Hash: 00000af23c...
Nonce: 29144
```

---

## 🗂️ Struktur Project

```
bitcoin-go/
├── cmd/node/main.go          # CLI entrypoint
├── internal/
│   ├── block/                # Block struct + Proof of Work
│   ├── blockchain/           # Chain management
│   ├── merkle/               # Merkle tree
│   ├── transaction/          # Wallet, TXInput, TXOutput (WIP)
│   ├── p2p/                  # P2P networking (WIP)
│   └── utils/                # Helper functions
└── go.mod
```
