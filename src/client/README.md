# op-watchtower

## The First Line of Defence for Optimistic Rollups

----------

# Running the containerised version of Witnesschain Watchtower

**Pre-requisite**
- Docker engine and client installed
- Configuration file set as per requirement (found at: `eigenlayer-avs-watchtower/src/client/config.json`)
  
## ~~Building from Scratch~~ [Developer only]
~~Clone the repository that contains the watchtower client code~~

`git clone git@github.com:kaleidoscope-blockchain/eigenlayer-avs-watchtower.git`

`cd eigenlayer-avs-watchtower/src/client/`

## Building [Developer only]

`docker build -t witnesschain/watchtower .`

### Running the image
`docker run --network host witnesschain/watchtower`

### Running the image with a config file
A volume mount can be used to feed the configuration to the app

`docker run -v PATH_TO_CONFIG:/app/config.json --network host witnesschain/watchtower`
