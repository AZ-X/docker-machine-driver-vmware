module github.com/machine-drivers/docker-machine-driver-vmware

require (
	github.com/docker/machine v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200427165652-729f1e841bcc
	golang.org/x/sys v0.0.0-20200428200454-593003d681fa
)

replace github.com/docker/docker v1.13.1 => github.com/docker/engine v17.12.0-ce-rc1.0.20200309214505-aa6a9891b09c+incompatible

replace github.com/docker/machine => ../../docker/machine

go 1.14
