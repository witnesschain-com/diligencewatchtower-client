# op-watchtower

## The First Line of Defence for Optimistic Rollups

----------

# Building from source

**Pre-requisite**
- git installed
- go version 1.20
- configuration file set
  
## Steps
1. Clone the repository

    `git clone https://github.com/witnesschain-com/diligencewatchtower-client.git`

2. Move to the `client` directory

    `cd diligencewatchtower-client/sec/client`

3. a) build and run the watchtower with the help of the `run` script

    `./run`

3. b) build the watchtower binary (optional)

    `go build .`

4. running the watchtower here after
   the watchtower binary should be present in this directory, you can run that by

    `./watchtower`


# Running the containerised version of Witnesschain Watchtower

**Pre-requisite**
- git installed
- Docker engine and client installed
- configuration file set
  

## Building 

1. Clone the repository

    `git clone https://github.com/witnesschain-com/diligencewatchtower-client.git`

2. Move to the `client` directory

    `cd diligencewatchtower-client/sec/client`
    
    `docker build -t witnesschain/watchtower .`

### Running the image
`docker run --network host witnesschain/watchtower`

### Running the image with a config file
A volume mount can be used to feed the configuration to the app

`docker run -v PATH_TO_CONFIG:/app/config.json --network host witnesschain/watchtower`
