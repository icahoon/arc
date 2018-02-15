# Container Service Configuration

The container service feature introduces the "container_service" section into arc datacenter configuration.
The "container_service" object is a peer with other top level objects such as the "datacenter", "database_service", "dns", "notifications"
and "provider" objects.


## container_service

The "container_service" section is a json object that contains up to two elements, the "containers" element and the optional "provider" element.
The "containers" element is a json array of container instances. The "provider" element is a json object that defines the cloud provider data. 

```
  "container_service": {
    "name": "my-container-service",

    "repositories": [
      {
        "repository": "repostory_name_goes_here",
        ...
      },
      ...
    ],

    "containers": [
      {
        "container": "container_name_goes_here",
        ...
      },
      ...
    ],

    "provider": {
      ...
    }
  }
```


## repositories

Since the container service can user multiple container repositories, the "repositories" member is a json array of "repository" elements.

### repository

The repository element represents a single container repository used to store container images. It has the following elements:

- **repository**        (string)  _required_: This is the name of the container repository. This name can be used on the cli to address this container repository.



## containers

Since the container service can host multiple container instances, the "containers" member is a json array of "container" elements.

### container


## provider

See the [provider configuration](../provider/config.md) for more details. The "data" fields required for the database_service
hosted in AWS are "account" and "region".


