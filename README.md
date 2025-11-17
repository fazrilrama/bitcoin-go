# 🧱 Bitcoin-Go  
Implementasi minimal dari *Bitcoin Whitepaper (Satoshi Nakamoto)* menggunakan bahasa Go.  
Project ini dibangun untuk mempelajari konsep fundamental Bitcoin seperti blockchain, block mining, proof-of-work, merkle tree, transaksi, wallet ECDSA, hingga simple CLI node.

---

## 🚀 Teknologi yang Diimplementasikan

### 🔗 1. Blockchain
- Struktur block:
  - `Timestamp`
  - `Data`
  - `PrevHash`
  - `Hash`
  - `Nonce`
- Genesis block otomatis
- Hashing SHA-256

### 💥 2. Proof of Work (Mining)
- Difficulty target menggunakan shifting bit `256 - difficulty`
- Mining dilakukan hingga hash valid ditemukan

### 🌳 3. Merkle Tree
- Menghasilkan Merkle Root dari transaksi
- Rekursif, sesuai konsep di whitepaper Bitcoin

### 💸 4. Transaksi & UTXO Model
- `TXInput`, `TXOutput`
- Dasar konsep wallet + tanda tangan digital

### 🔐 5. Wallet (ECDSA)
- Elliptic Curve P-256
- Private Key + Public Key generator

### 🗼 6. Simple Node (CLI)
- Perintah:
  - Tambah block: `add`
  - Print chain: `print`

---

## Instalasi & Setup
- Clone repository:
    - git clone https://github.com/fazrilrama/bitcoin-go.git
    - cd bitcoin-go
- Init module (jika belum)
    - go mod init github.com/fazrilrama/bitcoin-go
    - go mod tidy

## ▶️ Cara Menjalankan
- Menambahkan block baru:
    - go run cmd/node/main.go add "Hello Bitcoin"
- Melihat seluruh blockchain:
    - go run cmd/node/main.go print

## 📌 Contoh Output
- Tambah block:
    - Mining block...
    - PoW Success!
    - Block added!
- Print blockchain:
    - ==========================
    - Hash: 00000af23c...
    - Prev: 0000000...
    - Data: Hello Bitcoin
    - Nonce: 29144
    - ==========================