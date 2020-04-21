# Pact Verification Resource

Tracks the [pact verifications](https://docs.pact.io/getting_started/verifying_pacts) published to a [pact-broker](https://docs.pact.io/pact_broker). 


## Source Configuration

* `broker_url`: *Required.* The path of the hosted pact-broker.

* `consumer`: *Required.* The consumer to track provider verifications against.

* `providers`: *Required.* List of providers, along with `consumers`, to track verifications for.

* `tag`: *Optional.* If specified, pulls back only pacts with that [tag](https://docs.pact.io/pact_broker/advanced_topics/using_tags).

* `username`: *Optional.* If specified, along with `password`, used to provide basic auth to the pact-broker.

* `password`: *Optional.* The password to access pact-broker.

### Example

``` yaml
resource_types:
- name: pact-resource
  type: registry-image
  source:
    repository: nenaddev/pact-resource
    tag: latest

resources:
- name: pact
  type: pact-resource
  source:
    broker_url: https://path-to.your.pact-broker.io
    consumers: consumer-name
    providers:
      - provider-1
      - provider-2
    tag: dev
    username: ((you-are.using-a))
    password: ((secret-manager.right))
```
