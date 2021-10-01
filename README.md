# did:ar | multi-chain DIDs backed by arweave

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

```diff
- THIS IS A WIP ABSOLUTELY SHOULD NOT BE USED IN PRODUCTION
```


## Genesis Tx

```
$ go run main.go genesis --arweave-key <YUR_KEY>.json --protocol <ethereum_or_solana> --public-key 0x5f66Ce3....558975d927 --private-key <PRIVATE KEY BE CAREFUL USE A BURNER>
```

## Add Key
```
$ go run main.go add-key --arweave-key <YUR_KEY>.json <ethereum_or_solana> --public-key 0x5f66Ce3....558975d927 --private-key <PRIVATE KEY BE CAREFUL USE A BURNER> --genesis-id <OUTPUT FROM ABOVE> --previous-id <OUTPUT FROM ABOVE>
```