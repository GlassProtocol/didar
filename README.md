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

We live in a multi-chain world. That's a fact.

At Glass we support two chains (ethereum and solana). Maintaining identity across multiple chains is a complex problem. In a mono-chain environment you can rely on name services e.g. ENS or Solana Name Service, but in a multi-chain environment you either have to opt to use one name service or run a centralized server to link identities. In the case of Glass, using a centralized service is a non-option, so we needed to find a way to link multiple identities in a permission-less fashion. 

The following is an EXTREMELY ROUGH MVP for multi-chain decentralized identities. 

# What is a DID?

DID's are a W3 Standard

> Decentralized identifiers (DIDs) are a new type of identifier that enables verifiable, decentralized digital identity. A DID refers to any subject (e.g., a person, organization, thing, data model, abstract entity, etc.) as determined by the controller of the DID. In contrast to typical, federated identifiers, DIDs have been designed so that they may be decoupled from centralized registries, identity providers, and certificate authorities. Specifically, while other parties might be used to help enable the discovery of information related to a DID, the design enables the controller of a DID to prove control over it without requiring permission from any other party. DIDs are URIs that associate a DID subject with a DID document allowing trustable interactions associated with that subject. 
> 
> Each DID document can express cryptographic material, verification methods, or services, which provide a set of mechanisms enabling a DID controller to prove control of the DID. Services enable trusted interactions associated with the DID subject. A DID might provide the means to return the DID subject itself, if the DID subject is an information resource such as a data model.

Read More: https://www.w3.org/TR/did-core/


# Technical Discussion

Part of the reason for writing this document was to remove reliance on IPFS. Arweave offers true permanence without the need for coordinated pinning services.

For an IPFS alternative that is FAR MORE developed, be sure to check out our friends at [Ceramic](https://ceramic.network/).

---

### Genesis Transactions

Every DID requires a uniquely generated ID. We opt to use a small data payload Arweave transaction to initialize an ID. The genesis transaction IS NOT A DID, but rather a precursor. The genesis document looks like such:

```json
{
  "signing_key":  {
    "key_type":  "ETHEREUM",
    "public_key":  "0xAbf798E220c6E44E4F8d720E8095E8dB230E9718"
  }
}
```

Arweave ID: `xrzz5rl5Nr8cj7nf_XbQKbbUVrOsV0xZB6O4Reuja88` 


**Tags**  
Content-Type: `application/json`  
Operation: `GENESIS`  

The genesis document contains the signing key that will be used to initialize the append only log.


NOTE: anyone would be able to resubmit this document to arweave. The reason we add the signing key is to stop someone else from signing the genesis document and claiming that person's DID. The second arweave tx (and creation of a DID document) ABSOLUTELY MUST sign their DID document with the key used in the genesis document.



### DID Document

To create a DID document, the genesis transaction is referenced in the DID document. The way that we ensure security with updating the document is requiring the previous document's ID and the authentication payload to be cryptographically signed by one of the authorized keys. 

```json
{
  "context":  [
    "https://www.w3.org/ns/did/v1"
  ],
  "id":  "did:ar:xrzz5rl5Nr8cj7nf_XbQKbbUVrOsV0xZB6O4Reuja88",
  "authentication":  [
    {
      "key_type":  "SOLANA",
      "public_key":  "9jiixatNTBsLKAnfiv6BztccKai7UVWoEa1g6hKkWxvP"
    },
    {
      "key_type":  "ETHEREUM",
      "public_key":  "0xAbf798E220c6E44E4F8d720E8095E8dB230E9718"
    }
  ],
  "reference":  {
    "previous_document_id":  "xrzz5rl5Nr8cj7nf_XbQKbbUVrOsV0xZB6O4Reuja88",
    "signing_key":  {
      "key_type":  "ETHEREUM",
      "public_key":  "0xAbf798E220c6E44E4F8d720E8095E8dB230E9718"
    },
    "signature":  "0xebad66afc2572ad39bee572b222503a8d53a205b33858d8c2f7f274e422033400f093d3b8667974d59fabfd10e65d5a9798f939899900318d43ded18fca2e81d01"
  }
}
```

Arweave ID: `jlGTTlQt3W5uDLZI6biidGq_neNV3LlQsfXc-JDhTFw`  


**Tags**  
Content-Type: `application/json`  
Genesis-ID: `xrzz5rl5Nr8cj7nf_XbQKbbUVrOsV0xZB6O4Reuja88`  
Previous-ID: `xrzz5rl5Nr8cj7nf_XbQKbbUVrOsV0xZB6O4Reuja88`  
Operation: `ADD_KEY`  



The log of updates are then cryptographically linked back to the genesis ID. To reiterate, every new data update must not only sign the new data using an approved key but also MUST sign the reference to the previous entry in the DID log. Following this structure allows us to update a DID to add and delete authorized keys.

This method of cryptographically signing the previous log is not novel, but rather a fundamental primative for most blockchains. 


### Potential Attacks

The most pressing attack is a DDOS attack. One way to remedy this is to appove certain arweave addresses to submit transactions on the DID controller's behalf.


As further discussions develop regarding this approach to managing DID's they will be listed here.



# Command Line Tool

## Genesis Tx

``` bash
$ go run main.go add-key --arweave-key YOUR_ARWEAVE_KEY.json
```

## Add Key

```bash
$ go run main.go genesis --arweave-key YOUR_ARWEAVE_KEY.json
```

## Full Example

> did:ar:xrzz5rl5Nr8cj7nf_XbQKbbUVrOsV0xZB6O4Reuja88

```
$ go run main.go genesis --arweave-key YOUR_ARWEAVE_KEY.json
✔ Ethereum
Public Key: 0xAbf798E220c6E44E4F8d720E8095E8dB230E9718
Tx data size: 0.000114MB 
uplaodTx; body: OK, status: 200, txId: xrzz5rl5Nr8cj7nf_XbQKbbUVrOsV0xZB6O4Reuja88 
100.000000% completes, 1/1 

GENESIS ID: xrzz5rl5Nr8cj7nf_XbQKbbUVrOsV0xZB6O4Reuja88



$ go run main.go add-key --arweave-key YOUR_ARWEAVE_KEY.json
✔ Ethereum
Public Key: 0xAbf798E220c6E44E4F8d720E8095E8dB230E9718
Private Key: ****************************************************************
Genesis ID: xrzz5rl5Nr8cj7nf_XbQKbbUVrOsV0xZB6O4Reuja88
Previous ID (Genesis if that was last): xrzz5rl5Nr8cj7nf_XbQKbbUVrOsV0xZB6O4Reuja88
✔ Add Key
✔ Solana
Public Key: 9jiixatNTBsLKAnfiv6BztccKai7UVWoEa1g6hKkWxvP
✔ Finalize
Tx data size: 0.000711MB 
uplaodTx; body: OK, status: 200, txId: jlGTTlQt3W5uDLZI6biidGq_neNV3LlQsfXc-JDhTFw 
100.000000% completes, 1/1 

NEW DOC: jlGTTlQt3W5uDLZI6biidGq_neNV3LlQsfXc-JDhTFw

```

Genesis Data: https://arweave.net/xrzz5rl5Nr8cj7nf_XbQKbbUVrOsV0xZB6O4Reuja88

  
Genesis Explorer: https://viewblock.io/arweave/tx/xrzz5rl5Nr8cj7nf_XbQKbbUVrOsV0xZB6O4Reuja88

---
  

DID Doc: https://arweave.net/jlGTTlQt3W5uDLZI6biidGq_neNV3LlQsfXc-JDhTFw


DID Doc Explorer: https://viewblock.io/arweave/tx/jlGTTlQt3W5uDLZI6biidGq_neNV3LlQsfXc-JDhTFw


# Feedback

Please send any feedback to via DM to [@sam_sends](https://twitter.com/intent/follow?screen_name=sam_sends)
