# dcos-diff

`dcos-diff` compares the state of a running DCOS cluster with an on-disk representation of that state, or a `manifest`.

## Background

### Problem Statement

DCOS gives users the ability to launch a wide array of services and configure them at runtime. Each of these services can be defined by a number of configuration files (like `marathon.json` for basic apps, or `options.json` for more complex services), resources (like required `artifacts`, or `secrets`), and interactions with other services (like `ingress` rules for EdgeLB, or S3 access `accounts` in Minio).

While it is possible to manage all of these configurations, resources, and interactions completely through the provided Web UI and CLI tools provided by DCOS, operators running real-world workloads will quickly realize the need to keep all of these things organized in source control. Due to the currently available methodologies for deploying services (primarily Marathon App definitions and DCOS Commons SDK based Mesos Frameworks), it can be difficult to create a set of unified, one-size-fits-all scripts and automations for deploying and updating service configurations.

In other words, you need to be a DCOS and/or Apache Mesos expert to successfully manage a large number of services.

### Proposed Solution

One potential solution to this problem is to create a well defined set of file formats and directory layouts that can mirror the running state of a DCOS cluster. We will call this on-disk service definition the `service manifest`, or `manifest`. This `manifest` should be easily persisted in source control so that multiple people can update existing service configurations and launch new services, without the need to know intimate details about how that Mesos Framework or Marathon App were implemented.

Once the `service manifest` has been created, tooling is also needed to calculate the difference between the actual running state of a DCOS cluster, and the local `manifest` representation of those services. This is the primary purpose of this `dcos-diff` tool: To give an actionable `diff` between a local `manifest` and a running set of DCOS services.

## Service Manifest Format

### Folder Layout

The structure of the folders in a `manifest` correspond directly to the service paths / folders / groups that those services will reside in within a running DCOS cluster. Take the following example:

```
data/
├── miniod/
│   └── options.json
│   └── secret_paths.env
│   └── secrets.env
│   └── sa.env
│   └── private.pem
│   └── public.pem
```

The above definition results in the following DCOS resources and deployments:

* Marathon app running the `miniod` package scheduler at the path `/data/miniod/`
* Service account named `miniod-principal`
* Secret containing the private key for that service account at the path `data/miniod/miniod-secret`
* Secrets located at paths `data/miniod/access_key` and `data/miniod/secret_key` as defined in the `secret_paths.env` and `secrets.env`

### Service Definition

An individual DCOS service can be defined by a folder with the following contents:

* `options.json` OR `marathon.json` (Exactly one of these is **required**): The Framework options or Marathon App definition file.
* `service.env`: An **optional** env file containing keys and values needed to run install and update commands against a cluster. These include:
  * `PACKAGE_NAME`: Name of the package from the catalog.
  * `PACKAGE_VERSION`: Version string of the package from the catalog.
* `sa.env`: An **optional** env file. When empty, default values for the `principal` and `secret` names will be used (the name of the folder + `-principal` or `-secret`). These values can be overridden with these keys in `sa.env`
  * `SA_PRINCIPAL`
  * `SA_SECRET`
* `secret_paths.env`:
* `secrets.env`:

### MKE-Based Kubernetes Cluster Definition

A Kubernetes cluster can be defined very much like a service, since they are basically another service definition, but the `options.json` is replaced with a `k8s.json` to ensure that the correct commands are used to deploy / update the cluster. A folder with these files defines a Kubernetes cluster:

* `k8s.json`: Same as `options.json` used for `dcos kubernetes ...` commands
* `service.env`: Same as service definition, default `PACKAGE_NAME` being `kubernetes-cluster`

### Cluster Definition

Each DCOS cluster can be defined by the following files in the root of a `manifest`:

* `package_repos.env`: An **optional** env file containing keys and values corresponding to dcos package repositories. The default contents would be:
  * `Universe='https://universe.mesosphere.com/repo'`
* `urls.env`: An **optional** env file containing keys and values corresponding to masters and public agents. examples might be:
  * `EXTERNAL_URL='https://{{.Service}}.dcos.example.com`

### EdgeLB Pool Definition

An EdgeLB pool of loadbalancers can be defined by a folder with the following files:

* `pool.json`: The **required** rendered pool configuration file that can be used with `dcos edgelb update pool.json`.
* `pool.mustache.json`: An **optional** template file used to render the final `pool.json`.
* `pool.yml`: A more human readable representation of an EdgeLB pool configuration. Useful when exposing many services based on virtual host or path based routing rules.

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