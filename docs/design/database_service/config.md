# Database Service Configuration

The database service feature introduces the "database_service" section into arc datacenter configuration.
The "database_service" object is a peer with other top level objects such as the "datacenter", "dns" and "notifications"
objects.


## database_service

The "database_service" section is a json object that contains up to two members, the "databases" member and the optional "provider" member.
The "databases" member is a json array of database instances. The "provider" member is a json object that defines the cloud provider data.

```
  "database_service": {
    "databases": [
      ...
    ],

    "provider": {
      ...
    }
  }
```

## databases

Since a datacenter can host multiple database instances, the database element is an array of database elements.

### database

The database member represent a single database instance in the datacenter. It has the following elements

- **database**        (string)  _required_: This is the name of the database instance. This name can be used on the cli to address this database instance.
- **engine**          (string)  _required_: The name of the database engine to use for this instance.
- **version**         (string)  _optional_: The version of the database engine to use for this instance.
- **type**            (string)  _required_: Similar to the pod type, this specifies the db instance class type. ([AWS DB Instance Types](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html))
- **port**            (integer) _optional_: The port the engine will uses for connections.
- **subnet_group**    (string)  _required_: The subnet group that this database instance will use.
- **security_groups** (string)  _required_: The security groups that are applied to the database instance.
- **storage**         (object)  _optional_: The storage allocated to the database instance.
  - **type**            (string)  _required_: The type of storage.
  - **size**            (integer) _required_: The size of the storage in GiB.
  - **iops**            (integer) _optional_: The iops requested.
- **master**          (object)  _optional_
  - **username**:       (string)  _required_: The name of the master user.
  - **password**:       (string)  _required_: The password of the master user.

```json
      {
        "database":        "contacts",
        "engine":          "postgress",
        "version":         "9.6.5",
        "type":            "db.m4.large",
        "port":            5432,
        "subnet_group":    "postgres_subnet",
        "security_groups": [ "postgres_secgroup" ],
        "storage":         { "type": "gp2", "size": 50 },
        "master":          { "username": "myuser", "password": "mypasswd" }
      }
```


## Example database_service configuration

Here is an example of a simple postgres database using the AWS RDS service where the provider is defined outside of the database_service definition. There is a single database instance called "contacts".

```json
  "provider": {
    "vendor": "aws",
    "data": {
      "account": "example-integration",
      "region":  "us-east-1"
    }
  }

  "database_service": {
    "databases": [
      {
        "database":        "contacts",
        "engine":          "postgress",
        "version":         "9.6.5",
        "type":            "db.m4.large",
        "port":            5432,
        "subnet_group":    "postgres_subnet",
        "security_groups": [ "postgres_secgroup" ],
        "storage":         { "type": "gp2", "size": 50 },
        "master":          { "username": "myuser", "password": "mypasswd" }
      }
    ]
  },
```

## provider

See the [provider configuration](../provider/config.md) for more details. The "data" fields required for the database_service
hosted in AWS are "account" and "region".


