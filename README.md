# Diligence Watchtower Client

Welcome to the official GitHub repository for the Diligence Watchtower Client. Diligence watchtowers are the first line of defense for optimistic rollups. They enable incentive compatible and crypto-economically-secure Proof-of-Diligence (PoD) to make sure watchtowers are working in the happy path for optimistic rollups. 
The watchtower client is the software that entails a crucial responsibility of monitoring the accuracy of L2 state assertions made by the proposer. This guide offers information on setting up a watchtower node for the WitnessChain Network

You can read more about the protocol at [`Watchtower Protocol`](https://docs.witnesschain.com/diligence-watchtowers/introduction)

## Key Features

- **Proof of Diligence**: Proofs-of-Diligence for watching state L2 assertions of OP Stack chains (OP,Base)
- **Proof of Inclusion**: Proofs-of-Inclusion for watching Transactions' inclusions in a Block.

## Getting Started

To get started with the Watchtower Client, please follow the instructions below

### Prerequisites
- **Registered EigenLayer Operator**: [`Register`](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation) as an operator on EL Holesky/Mainnet chain

- Holesky Registration
    - **Registered WitnessChain Operator**: [`Register`](https://docs.witnesschain.com/diligence-watchtowers/for-the-node-operators/watchtower-setup/holesky-setup#step-2-get-whitelisted-on-the-watchtower-network) EL Operator and watchtower on Witness Chain

- Mainnet Registration
    - **Registered WitnessChain Operator**: [`Register`](https://docs.witnesschain.com/diligence-watchtowers/for-the-node-operators/watchtower-setup/mainnet-setup#step-2-get-whitelisted-on-the-watchtower-network) EL Operator and watchtower on Witness Chain

## Building from source

### **Pre-requisite**
- git installed
- go version 1.20
- configuration file set (src/client/config.json)

### Steps
1. Clone the repository

    ```
    git clone https://github.com/witnesschain-com/diligencewatchtower-client.git
    ```

2. Move to the `client` directory

    ```
    cd diligencewatchtower-client/src/client
    ```

3. Building
   -  using run script
      -  build and run the watchtower with the help of the `run` script

    ```
    ./run
    ```

   - using go
      -  build the watchtower binary (optional)

    ```
    go build -o watchtower .
    ```

4. Running the watchtower: If the watchtower was built via go, you can manually start it with the following command

    ```
    ./watchtower
    ```

**Note:** Make sure the `config.json` is filled and present in the directory the watchtower executable is

## Running the containerised version of WitnessChain Watchtower

### **Pre-requisite**
- git installed
- Docker engine and client installed
- configuration file set
  

### Building the container

1. Clone the repository

    ```
    git clone https://github.com/witnesschain-com/diligencewatchtower-client.git
    ```

2. Move to the `client` directory

    ```
    cd diligencewatchtower-client/src/client

    docker build -t witnesschain/watchtower .
    ```

### Running the image
```
docker run --network host witnesschain/watchtower
```

### Running the image with a config file
A volume mount can be used to feed the configuration to the app

`docker run -v PATH_TO_CONFIG:/app/config.json --network host witnesschain/watchtower`
