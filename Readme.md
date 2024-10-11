# Splunk CLI

a quick cli to run command againsst your splunk instance


## Install

```
curl -sSL https://raw.githubusercontent.com/cicd-toolkit/spli/master/scripts/install | bash
```

## Documentation

please check [here](./docs/spli.md)

## Test

```
$ docker run -d -p 8000:8000 -p 8088:8088  -p 8089:8089 -e SPLUNK_START_ARGS=--accept-license -e SPLUNK_PASSWORD=Admin-1234 splunk/splunk:9.3 start


$ spli configure
Host [default localhost] : splunk.example.com
Username [default admin] :
Password [default admin] :
admin port [default 8089] :
web port [default 8000] :
protocol [default http] :
Done

$ spli app list
[
  "alert_logevent",
  "alert_webhook",
  "appsbrowser",
  "introspection_generator_addon",
  "journald_input",
  "launcher",
  "learned",
  "legacy",
  "python_upgrade_readiness_app",
  "sample_app",
  "search",
  "splunk-dashboard-studio",
  "splunk-rolling-upgrade",
  "splunk-visual-exporter",
  "splunk_archiver",
  "splunk_assist",
  "splunk_enterprise_on_docker",
  "splunk_gdi",
  "splunk_httpinput",
  "splunk_ingest_actions",
  "splunk_instrumentation",
  "splunk_internal_metrics",
  "splunk_metrics_workspace",
  "splunk_monitoring_console",
  "splunk_rapid_diag",
  "splunk_secure_gateway",
  "SplunkDeploymentServerConfig",
  "SplunkForwarder",
  "SplunkLightForwarder"
]
```





