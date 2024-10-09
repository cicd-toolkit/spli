

```
docker run -d -p 8000:8000 -p 8088:8088  -p 8089:8089 -e SPLUNK_START_ARGS=--accept-license -e SPLUNK_PASSWORD=Admin-1234 splunk/splunk:9.3 start
curl localhost:8000 -v

curl -k -u   admin:Admin-1234 https://localhost:8089/services/server/info?output_mode=json -G -X GET -v | jq


curl -k -u  admin:Admin-1234 https://localhost:8089/services/apps/local -d name=testapp -v

```


https://docs.splunk.com/Documentation/Splunk/9.3.1/RESTREF/RESTapps#apps.2Flocal
