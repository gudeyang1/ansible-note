#!/bin/bash

kit_release_version=`cat main.go  |grep "version =" |awk -F "\"" '{print $2}'`

rm -rf suitectl suitectl.exe suitectl*.tar.gz
#build
go get github.com/spf13/cobra
go get github.com/spf13/viper
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o suitectl.exe main.go
go build -o suitectl main.go

[ -f /tmp/suite-kit-release/${kit_release_version} ] && rm -rf /tmp/suite-kit-release/${kit_release_version}
#release
mkdir -p /tmp/suite-kit-release/suitectl-v${kit_release_version}
cp -r ./ansible /tmp/suite-kit-release/suitectl-v${kit_release_version}
cp -r ./templates /tmp/suite-kit-release/suitectl-v${kit_release_version}
cp suitectl* /tmp/suite-kit-release/suitectl-v${kit_release_version}
cp .suitectl.yaml.default /tmp/suite-kit-release/suitectl-v${kit_release_version}

#delete useless file
rm -rf /tmp/suite-kit-release/suitectl-v${kit_release_version}/ansible/host
rm -rf /tmp/suite-kit-release/suitectl-v${kit_release_version}/ansible/site.retry
rm -rf /tmp/suite-kit-release/suitectl-v${kit_release_version}/ansible/.vagrant

#tar 
cd /tmp/suite-kit-release
tar -zcf suitectl-v${kit_release_version}.tar.gz suitectl-v${kit_release_version}
cd -
mv /tmp/suite-kit-release/suitectl-v${kit_release_version}.tar.gz ./
