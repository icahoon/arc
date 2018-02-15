# Provider Configuration

The provider element is a top level element that contains vendor specific information separate from the resource elements.
For arc it is a peer to elements such as the datacenter, database_service, container_service and dns. For amp it is a peer
to elements such as storage.

Some resource elements, dns in particular, may need to specify it's own provider section to override the global provider element.


## provider

The provider element contains two required members, "vendor" and "data". The "vendor" member is a string indicating the cloud vendor to use.
Currently only the "aws" and "mock" vendors are supported. The "data" member is a json object comprised of key value pairs.
These pairs differ depending on the vendor.

A third member, "image", is also available and used by the arc "datacenter" object. The image object is comprised of key value pairs where the
key is a image name used in the config file by the "pod" element and the value is the vendor specific id. For AWS this image id is the ami of
the image. This image map allows the decoupling of image types with the concrete image ids to provide a more consistent configuration.


### aws vendor

For an aws vendor, the following values are supported:

- **account** (string) _required_: The aws account name
- **region**  (string) _required_: The aws region that this database service where this database service will reside
- **number**  (string) _optional_: The aws account number. This is only used by the dns resource element.

```json
    "provider": {
      "vendor": "aws",
      "data": {
        "account": "example-integration",
        "region":  "us-east-1",
        "number":  "123456789012"
      }
    }
```

For an aws vendor used by an arc datacenter config file, the image member is also required:


```json
    "provider": {
      "vendor": "aws",
      "data": {
        "account": "example-integration",
        "region":  "us-east-1"
      }
      "image": {
        "centos7":    "ami-1234abcd",
        "ubuntu16":   "ami-abcd1234"
      }
    }
```

### mock vendor

For a mock vendor, the data section is not required.

```json
  "database": {
    "provider": { "vendor": "mock" }
  }
```
