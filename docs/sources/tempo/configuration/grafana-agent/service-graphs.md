---
title: Enable service graphs
menuTitle: Enable service graphs
description: Service graphs help to understand the structure of a distributed system, and the connections and dependencies between its components.
weight:
aliases:
- /docs/tempo/grafana-agent/service-graphs
---

# Enable service graphs

A service graph is a visual representation of the interrelationships between various services.
Service graphs help to understand the structure of a distributed system,
and the connections and dependencies between its components.

The same service graph metrics can also be generated by Tempo.
This is more efficient and recommended for larger installations.
For a deep look into service graphs, visit [this section]({{< relref "../../metrics-generator/service_graphs" >}}).

Service graphs are also used in the application performance management dashboard.
For more information, refer to the [service graph view documentation]({{< relref "../../metrics-generator/service-graph-view" >}}).

## Before you begin

Service graphs are generated in the Grafana Agent and pushed to a Prometheus-compatible backend.
Once generated, they can be represented in Grafana as a graph.
You will need these components to fully use service graphs.

### Enable service graphs in Grafana Agent

To start using service graphs, enable the feature in the Grafana Agent config.

```yaml
traces:
  configs:
    - name: default
      ...
      service_graphs:
        enabled: true
```

To see all the available config options, refer to the [configuration reference](/docs/agent/latest/configuration/traces-config).

Metrics are registered in the Agent's default registerer.
Therefore, they are exposed at `/metrics` in the Agent's server port (default 12345).
One option is to use the Agent self-scrape capabilities to export the metrics to a Prometheus-compatible backend.

```yaml
metrics:
  configs:
    - name: default
      scrape_configs:
        - job_name: local_scrape
          static_configs:
            - targets: ['127.0.0.1:12345']
      remote_write:
        - url: <remote_write>
```

### Grafana

The same service graph metrics can also be generated by Tempo.
This is more efficient and recommended for larger installations.

For additional information about viewing service graph metrics in Grafana and calculating cardinality, refer to the [server side documentation]({{< relref "../../metrics-generator/service_graphs/enable-service-graphs" >}}).
