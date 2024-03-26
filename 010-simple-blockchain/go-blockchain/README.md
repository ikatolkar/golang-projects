# Concept Explained
Blockchain explained : https://andersbrownworth.com/blockchain/

# Simple Blockchain
- BlockChain contains an array of Blocks
- Each Block contains:
	1. Pos: position in blockchain
	2. Data: BookCheckout, a struct with info on a book transaction
	3. TimeStamp: Block creation timestamp
	4. Hash: SHA256(b.Pos + b.TimeStamp + json.Marshal(b.Data) + b.PrevHash)
	5. PrevHash: Hash of block at pos-1
- Block is validated by doing:
    1. Check if hash of prev block is cur.prevHash
    2. Check if pos is prev pos + 1
    3. Check if hash is valid

# Example
```bash
root@PF2Z4T5C-inl:~$ curl -X POST http://localhost:3000/ -H "Content-type application/json" -d '{"book_id":"423421341324321412", "user":"Mary doe", "checkout_date":"2024-04-26"}'

root@PF2Z4T5C-inl:~$ curl -X GET http://localhost:3000/
[
 {
  "Pos": 1,
  "Data": {
   "book_id": "",
   "user": "",
   "checkout_date": "",
   "is_genesis": false
  },
  "TimeStamp": "2024-03-26 23:45:43.6802082 +0530 IST m=+0.000256701",
  "Hash": "c3ce831ab00f1cad436a00e745ccc2a1176657709a605a51b5e0ba0153db7d25",
  "PrevHash": ""
 },
 {
  "Pos": 2,
  "Data": {
   "book_id": "",
   "user": "",
   "checkout_date": "",
   "is_genesis": false
  },
  "TimeStamp": "2024-03-26 23:47:54.4283197 +0530 IST m=+130.748368201",
  "Hash": "2afca5a0710afafece6a00c0d18f94a9d01f87cb3eb406ab4e1d2c988dc8667e",
  "PrevHash": "c3ce831ab00f1cad436a00e745ccc2a1176657709a605a51b5e0ba0153db7d25"
 }
]
```
