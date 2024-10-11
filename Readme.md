# Splunk CLI

a quick cli to run command againsst your splunk instance


## Install

```
curl https://raw.githubusercontent.com/cicd-toolkit/spli/master/scripts/install | bash
```

## Documentation

please check [here](./docs/spli.md)

## Test

```
$ docker run -d -p 8000:8000 -p 8088:8088  -p 8089:8089 -e SPLUNK_START_ARGS=--accept-license -e SPLUNK_PASSWORD=Admin-1234 splunk/splunk:9.3 start


$ spli setup
URL: https://localhost:8089
Username: admin
Password:
Logging in with URL: https://localhost:8089, Username: admin
Login data saved to .spli

$ spli version
9.3.1
```





