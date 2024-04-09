go mod edit -module github.com/diligencewatchtower-client
-- rename all imported module
find . -type f -name '*.go' \
  -exec sed -i ''  -e 's/witnesschain.com/github.com/diligencewatchtower-client/g' {}; 
