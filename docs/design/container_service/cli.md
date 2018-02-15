# Container Service CLI Design

Two commands will be added to the arc cli to provide management of a container service within a datacenter.



## container

To run an arc command against the container service in a datacenter use the form.  Commands at the container
level will not effect the container repositories or individual container instances. This command will only
effect the top level container service. For AWS this means manipulating the
[ECS Cluster](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ECS_clusters.html).

```shell
arc [dc_name] container [command]
```
The commands available to the container service are

- **create**:  Create the container service.
- **destroy**: Destroy the container service.
- **update**:  Update the container service.
- **audit**:   Audits the container service.
- **info**:    Provides run time information for the container service.
- **config**:  Provides configuration information for the container service.
- **help**:    Provides help with the container service command.

For example, to run an audit command against the container service in the fictional "example-integration" datacenter, run the command

```shell
arc example-integration container audit
```


To run an arc command againt a specific container instance use the form. The current design limits arc to 
addressing a single container at a time.

```shell
arc [dc_name] container [container_name] [command]
```

The commands available to a container are

- **create**:  Create the container.
- **destroy**: Destroy the container.
- **update**:  Update the container.
- **audit**:   Audits the container.
- **info**:    Provides run time information for the container.
- **config**:  Provides configuration information for the container.
- **help**:    Provides help with the containercommand.

For example, to create a "my-microservice" container in the fictional "example-integration" datacenter, run the command

```shell
arc example-integration container my-microservice create
```



## repository

To run an arc command against all container repositories in a datacenter use the form

```shell
arc [dc_name] repository [command]
```

To run an arc command against a single container repository in a datacenter use the form

```shell
arc [dc_name] repository [repo_name] [command]
```

The commands available to the container repository are

- **create**:  Create the container repository.
- **destroy**: Destroy the container repository.
- **update**:  Update the container repository.
- **audit**:   Audits the container repository.
- **info**:    Provides run time information for the container repository.
- **config**:  Provides configuration information for the container repository.
- **help**:    Provides help with the container service repository.

