# dcos-diff

DCOS Diff compares the state of a running DCOS cluster with an on-disk representation of that state, or a `manifest`.

## Background

### Problem Statement

DCOS gives users the ability to launch a wide arary of services and configure them at runtime. Each of these services can be defined by a number of configuration files (like `marathon.json` for basic apps, or `options.json` for more complex serivces), resources (like required `artifacts`, or `secrets`), and interactions with other services (like `ingress` rules for EdgeLB, or S3 access `accounts` in Minio).

While it is possible to manage all of these configurations, resources, and interacions completely through the provided Web UI and CLI tools provided by DCOS, operators running real-world workloads will quickly realize the need to keep all of these things organized in source control. Due to the currently available methodologies for deploying services (primarily Marathon App definitions and DCOS Commons SDK based Mesos Frameworks), it can be difficult to create a set of unified, one-size-fits-all scripts and automations for deploying and updating service configurations.

In other words, you need to be a DCOS and/or Apache Mesos expert to successfully manage a large number of services.

### Proposed Solution

One potential solution to this problem is to create a well defined set of file formats and directory layouts that can mirror the running state of a DCOS cluster. We will call this on-disk service definition the `service manifest`, or `manifest`. This `manifest` should be easily persisted in source control so that multiple people can easily update existing service configurations, and launch new services, without the need to know intimate details about how that Mesos Framework or Marathon App were implemented.

Once the `service manifest` has been created, tooling is also needed to calculate the difference between the actual running state of a DCOS cluster, and the local `manifest` representation of those services. This is the primary purpose of this `dcos-diff` tool: To give an actionable `diff` between a local `manifest` and a running set of DCOS services.

## Usage

TBD

## Development

### Testing

```
go test ./...
```

### Building

```
mkdir -p build
go build -o build/dcos-diff ./cmd/dcos-diff/main.go
```

### Running

```
./build/dcos-diff
```