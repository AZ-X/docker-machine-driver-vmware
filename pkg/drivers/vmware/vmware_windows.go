/*
Copyright 2017 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vmware

import (
	"path/filepath"
	"strings"
	
	"github.com/docker/machine/libmachine/log"
	"github.com/machine-drivers/docker-machine-driver-vmware/pkg/drivers/vmware/config"
	"golang.org/x/sys/windows/registry"
)

//var windowsInstallDir = `C:\Program Files (x86)\VMware\VMware Workstation`
//var dhcpConfFileDir = `C:\ProgramData\VMware\vmnetdhcp.conf`
//var leaseFileDir = `C:\ProgramData\VMware\vmnetdhcp.leases`

var windowsInstallDir, dhcpConfFileDir, leaseFileDir = getDirs()

func getDirs()(val1 string, val2 string, val3 string) {
	// Parse HKEY_.CLASSES_ROOT\vm\shell\open\command's value like:
	// "C:\Program Files (x86)\VMware\VMware Workstation\vmware.exe" "%1"
	// in order to the Workstation install dir.
	
	_windowsInstallDir := ""
	_dhcpConfFileDir := ""
	_leaseFileDir := ""
	key, err := registry.OpenKey(registry.CLASSES_ROOT, `vm\shell\open\command`, registry.QUERY_VALUE)
	if err != nil {
		log.Errorf(">>>>>>ERROR: %s", err)
		//return
	}
	defer key.Close()

	value, _, err := key.GetStringValue("")
	if err != nil {
		log.Errorf(">>>>>>ERROR: %s", err)
		//return
	}else {
		//log.Info(">>>>>>VMWARE REGISTRY INFO: " + value)
	}
	

	if value[0] == '"' {
		values := strings.Split(value[1:], "\"")
		_windowsInstallDir = filepath.Dir(values[0])
	}else {
		log.Error(">>>>>>VMWARE REGISTRY ERROR: FIRST QUOTES NOT FOUND")
	}
	
	
	key2, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\services\VMnetDHCP\Parameters`, registry.QUERY_VALUE)
	if err != nil {
		log.Errorf(">>>>>>ERROR: %s", err)
		//return
	}
	defer key2.Close()

	value2, _, err := key2.GetStringValue("ConfFile")
	if err != nil {
		log.Errorf(">>>>>>ERROR: %s", err)
		//return
	}else {
		//log.Info(">>>>>>VMWARE REGISTRY INFO: " + value2)
		_dhcpConfFileDir = value2;
	}
	
	value3, _, err := key2.GetStringValue("LeaseFile")
	if err != nil {
		log.Errorf(">>>>>>ERROR: %s", err)
		//return
	}else {
		//log.Info(">>>>>>VMWARE REGISTRY INFO: " + value3)
		_leaseFileDir = value3;
	}
	return _windowsInstallDir, _dhcpConfFileDir, _leaseFileDir
}

func DhcpConfigFiles() string {
	return dhcpConfFileDir
}

func DhcpLeaseFiles() string {
	return leaseFileDir
}

func SetUmask() {
}

func setVmwareCmd(cmd string) string {
	cmd = cmd + ".exe"
	return filepath.Join(windowsInstallDir, cmd)
}

func getShareDriveAndName() (string, string, string) {
	return config.DefaultShareName, config.DefaultSharePath, "/hosthome"
}
