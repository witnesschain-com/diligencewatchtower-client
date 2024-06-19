## [1.0.0]

This release focus on stability, security improvements, and bug fixes. 

### Upgrade Steps
* Delete the old container images.
```
docker rm <stale_container_id>
docker rmi -f <stale_image_name>
```
* Download and run diligencewatchtower-cli with latest latest release 
  tag.
```
docker run https://hub.docker.com/r/witnesschain/watchtower:1.0.0
```


For support or to report issues, please visit our [https://github.com/witnesschain-com/diligencewatchtower-client](https://github.com/witnesschain-com/diligencewatchtower-client) and [create a new issue](https://github.com/witnesschain-com/diligencewatchtower-client/issues).
