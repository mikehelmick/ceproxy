# ceproxy
Simple cloudevents proxy that reads cloud events and writes them to the configured broker.

# To deploy

Edit the `service.yaml` file to indicate the address to proxy events to.
By default, that is the Kative default broker in the default namespace.

```
ko apply -f service.yaml
```
