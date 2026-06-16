package transaction

// UTXOSet menyimpan semua unspent outputs per address
type UTXOSet map[string][]TXOutput

// FindUTXO mengembalikan semua unspent outputs milik address tertentu
func FindUTXO(address string, txs []*Transaction) []TXOutput {
	var utxos []TXOutput
	spent := map[string]map[int]bool{}

	for _, tx := range txs {
		txID := string(tx.ID)
		for _, in := range tx.Vin {
			if string(in.PubKey) == address || string(in.Signature) == address {
				if spent[string(in.Txid)] == nil {
					spent[string(in.Txid)] = map[int]bool{}
				}
				spent[string(in.Txid)][in.OutIndex] = true
			}
		}
		for i, out := range tx.Vout {
			if string(out.PubKeyHash) == address {
				if !spent[txID][i] {
					utxos = append(utxos, out)
				}
			}
		}
	}
	return utxos
}

// FindSpendableOutputs mencari outputs yang cukup untuk memenuhi amount
func FindSpendableOutputs(address string, amount int, txs []*Transaction) (int, map[string][]int) {
	unspent := map[string][]int{}
	accumulated := 0
	spent := map[string]map[int]bool{}

	for _, tx := range txs {
		txID := string(tx.ID)
		for _, in := range tx.Vin {
			if string(in.PubKey) == address {
				if spent[string(in.Txid)] == nil {
					spent[string(in.Txid)] = map[int]bool{}
				}
				spent[string(in.Txid)][in.OutIndex] = true
			}
		}
		for i, out := range tx.Vout {
			if string(out.PubKeyHash) == address && !spent[txID][i] {
				accumulated += out.Value
				unspent[txID] = append(unspent[txID], i)
				if accumulated >= amount {
					return accumulated, unspent
				}
			}
		}
	}
	return accumulated, unspent
}
