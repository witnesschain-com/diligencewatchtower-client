### Testnet Release: v1.1.0

This release add support to store keys in [web3 secret storage format](https://ethereum.org/en/developers/docs/data-structures-and-encoding/web3-secret-storage/).

**changes need to use keys in web3 secret format**

In you `watchtower.config.json` file add the `encrypted_key` field with 
path to they web3 secret storage key file
```
"encrypted_key": /home/ubuntu/.witnesschain/cli/.w3secretkeys/watchtower1.ecdsa.key.json
```

:whale: DockerHub 
[docker pull witnesschain/watchtower:v1.1.0](https://hub.docker.com/r/witnesschain/watchtower/tags)


For support or to report issues, please visit our [https://github.com/witnesschain-com/diligencewatchtower-client](https://github.com/witnesschain-com/diligencewatchtower-client) and [create a new issue](https://github.com/witnesschain-com/diligencewatchtower-client/issues).
