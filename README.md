# Feralfile Exhibition Contract

## Pre-requisite

- npm 14+
- solidity 0.8.0
- truffle 5.4.3
- ganache (for local development)

## Setup project

We use `npm` to manage packages. Before start developing, we need to install all required packages.

```
npm install

```

### Configuration

The project reads `.secret.json` for setting up network. Please copy from `.secret.json.example` and update its value accordingly.

## Compile

To compile the smart contract, we use

```
truffle compile
```

It generates contract ABI files in `build/contracts`.

## Deploy

To deploy a contract, we use

```
truffle deploy
```

By default, it deploys to development. To change the network, simply use the parameter `--network`. For example,

```
truffle deploy --network ropsten
```

### Parameters

Here lists parameters for initiating the exhibition contract:

| parameter | type | description |
| - | - | - |
| exhibition_name | string | The name of an exhibition |
| exhibition_symbol | string | The symbol of the ERC721 token |
| exhibition_curator | string | The curator of an exhibition |
| exhibition_edition_size | int | The edition counts of each artwork in an exhibition |

For example,

```
truffle deploy -f 2 --exhibition_name="Social Codes" --exhibition_symbol="SCC" --exhibition_curator="0x8fd310de32848798eB64Bd88f9C5656Eea32415e"
```

## Execute Contract

### Truffle

Truffle contains a built-in console for interacting with the blockchain. Run

```
truffle console
```

you would enter the shell mode.

### Golang

`go-ethereum` includes a command `abigen` which creates `golang` structures and interfaces based on an ABI file. With this file, we can execute smart contract with `golang` easily.

```shell
jq ".abi" FeralfileExhibition.json | abigen  --abi="-" --type="FeralfileExhibition" --pkg=contract --out="feralaile.go"
```
