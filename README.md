# Docker Machine VMware Driver for Windows 7 and above

First, please see [Things you must know before README](https://github.com/AZ-X/docker-machine-driver-vmware/issues/1).

***If this project benefits you, please star it and keep watching***. 

Create Docker machines locally on ~~VMware [Fusion](https://www.vmware.com/products/fusion)
and~~ [Workstation](https://www.vmware.com/products/workstation).

This driver requires VMware Workstation 15 Windows ~~or VMware Fusion 11 (MacOS)~~ to be installed on your host. Earlier versions of Workstation/Fusion might still work
with this driver, but it's not officially supported. The assumable version of underlay kernel inside boot2docker changes to 5.4.

>
> Docker machine has a builtin driver called `vmwarefusion`. The main difference between
> those drivers is that `vmware` also works on VMware Workstation, while `vmwarefusion` only
> works on VMware Fusion.


## Installation

### From a Release

The latest version of the `docker-machine-driver-vmware` binary is available on the
[GithHub Releases](https://github.com/AZ-X/docker-machine-driver-vmware/releases) page.
Download the the binary that corresponds to your OS into a directory residing in your PATH.

### From Source

Make sure you have installed [Go](http://www.golang.org) and configured [GOPATH](http://golang.org/doc/code.html#GOPATH)
properly. ~~For MacOS and Linux, make sure `$GOPATH/bin` is part of your `$PATH` for MacOS and Linux.~~
For Windows, make sure `%GOPATH%\bin` is included in `%PATH%`.

Run the following command:

```shell
go get -u github.com/AZ-X/docker-machine-driver-vmware
```


## Usage

```shell
$ docker-machine create --driver=vmware default
```


## Options

- `--vmware-boot2docker-url`: URL for boot2docker image
- `--vmware-configdrive-url`: URL for cloud-init configdrive
- `--vmware-cpu-count`: Number of CPUs for the machine (-1 to use the number of CPUs available)
- `--vmware-disk-size`: Size of disk for host VM (in MB)
- `--vmware-memory-size`: Size of memory for host VM (in MB)
- `--vmware-no-share`: Disable the mount of your home directory
- `--vmware-ssh-password`: SSH password
- `--vmware-ssh-user`: SSH user
- `--vmware-share-path`: VMware Share Path
- `--vmware-share-name`: VMware Share Name (keep default 'Users' when migration)
- `--vmware-bt2d-data-storage`: vmdk Path for bt2d data storage
- `--ssh-port`: SSH port
- `--dockerd-port`: dockerd port

#### Environment variables and default values

| CLI option                 | Environment variable   | Default                  |
|----------------------------|------------------------|--------------------------|
| `--vmware-boot2docker-url` | VMWARE_BOOT2DOCKER_URL | *Latest boot2docker url* |
| `--vmware-configdrive-url` | VMWARE_CONFIGDRIVE_URL | -                        |
| `--vmware-cpu-count`       | VMWARE_CPU_COUNT       | `1`                      |
| `--vmware-disk-size`       | VMWARE_DISK_SIZE       | `20000`                  |
| `--vmware-memory-size`     | VMWARE_MEMORY_SIZE     | `1024`                   |
| `--vmware-no-share`        | VMWARE_NO_SHARE        | -                        |
| `--vmware-ssh-password`    | VMWARE_SSH_PASSWORD    | `tcuser`                 |
| `--vmware-ssh-user`        | VMWARE_SSH_USER        | `docker`                 |
| `--ssh-port`        | SSH_PORT        | `22`                 |
| `--dockerd-port`        | DOCKERD_PORT        | `2376`                 |
| `--vmware-share-path`        | VMWARE_SHARE_PATH        | `C:\docker\`                 |
| `--vmware-share-name`        | VMWARE_SHARE_NAME        | `Users`                 |
| `--vmware-bt2d-data-storag`        | VMWARE_BT2DDATASTORAGE        | -                 |


## License

See license file [LGPL V3 LICENSE](https://github.com/AZ-X/docker-machine-driver-vmware/blob/master/LICENSE "LICENSE").
