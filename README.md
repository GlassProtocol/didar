<p align="center">
  <img height=100 src="https://arweave.net/sBogY_roIMJWInS0HIEi86eFGzHUnNxUzyKEdOKPWh0" />
</p>

# did:ar | multi-chain DIDs backed by Arweave

```diff
- THIS IS A WIP ABSOLUTELY SHOULD NOT BE USED IN PRODUCTION
```


```

     ___             ___                             
    (   )  .-.      (   )                            
  .-.| |  ( __)   .-.| |          .---.   ___ .-.    
 /   \ |  (''")  /   \ |         / .-, \ (   )   \   
|  .-. |   | |  |  .-. |   .-.  (__) ; |  | ' .-. ;  
| |  | |   | |  | |  | |  (   )   .'`  |  |  / (___) 
| |  | |   | |  | |  | |   `-'   / .'| |  | |        
| |  | |   | |  | |  | |   .-.  | /  | |  | |        
| '  | |   | |  | '  | |  (   ) ; |  ; |  | |        
' `-'  /   | |  ' `-'  /   `-'  ' `-'  |  | |        
 `.__,'   (___)  `.__,'         `.__.'_. (___)       
                                                     
                                                     
```

# Motivation

# Technical Discussion

## Genesis Tx

```
$ go run main.go genesis --arweave-key <YUR_KEY>.json --protocol <ethereum_or_solana> --public-key 0x5f66Ce3....558975d927 --private-key <PRIVATE KEY BE CAREFUL USE A BURNER>
```

## Add Key
```
$ go run main.go add-key --arweave-key <YUR_KEY>.json <ethereum_or_solana> --public-key 0x5f66Ce3....558975d927 --private-key <PRIVATE KEY BE CAREFUL USE A BURNER> --genesis-id <OUTPUT FROM ABOVE> --previous-id <OUTPUT FROM ABOVE>
```

## Full Example

```
$ go run main.go genesis --arweave-key ark.json --protocol ethereum --public-key 0x5f66Ce3fc08Ca73a715B4d00616DBb558975d927 --private-key <SECRET>
Tx data size: 0.000310MB 
uplaodTx; body: OK, status: 200, txId: 3xYtAzdF7JBTUM0cRSeh-FQq7DUr10CoASjDEOJbXSg 
100.000000% completes, 1/1 

GENESIS: 3xYtAzdF7JBTUM0cRSeh-FQq7DUr10CoASjDEOJbXSg



$ go run main.go add-key --arweave-key ark.json --protocol ethereum --public-key 0x5f66Ce3fc08Ca73a715B4d00616DBb558975d927 --private-key <SECRT> --genesis-id 3xYtAzdF7JBTUM0cRSeh-FQq7DUr10CoASjDEOJbXSg --previous-id 3xYtAzdF7JBTUM0cRSeh-FQq7DUr10CoASjDEOJbXSg
✔ Add Key
✔ Ethereum
Public Key: 0x5f66Ce3fc08Ca73a715B4d00616DBb558975d927
✔ Add Key
✔ Solana
Public Key: 9jiixatNTBsLKAnfiv6BztccKai7UVWoEa1g6hKkWxvP
✔ Finalize
Tx data size: 0.000699MB 
uplaodTx; body: OK, status: 200, txId: AlnD2Yq2btI6WMgOc8wINZMHJhYr4qac59eLJcc8Kag 
100.000000% completes, 1/1 

NEW DOC: AlnD2Yq2btI6WMgOc8wINZMHJhYr4qac59eLJcc8Kag

```

Genesis Data: https://arweave.net/3xYtAzdF7JBTUM0cRSeh-FQq7DUr10CoASjDEOJbXSg

  
Genesis Explorer:

---
  

DID Doc: https://arweave.net/AlnD2Yq2btI6WMgOc8wINZMHJhYr4qac59eLJcc8Kag


DID Doc Explorer:

