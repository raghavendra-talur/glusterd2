package heketi

import (
	"fmt"
	"os/exec"

	heketi "github.com/gluster/glusterd2/plugins/heketi/api"
	"github.com/gluster/glusterd2/transaction"
)

func txnHeketiPrepareDevice(c transaction.TxnCtx) error {
	var deviceinfo heketi.DeviceInfo

	if err := c.Get("nodeid", &deviceinfo.NodeID); err != nil {
		return err
	}
	if err := c.Get("devicename", &deviceinfo.DeviceName); err != nil {
		return err
	}

	fmt.Printf("rtalur we are now in transaction")

	cmd1 := exec.Command("pvcreate", "--metadatasize=128M", "--dataalignment=256K", "/dev/"+deviceinfo.DeviceName)
	if err := cmd1.Run(); err != nil {
		fmt.Printf("rtalur we failed to run pvcreate")
	}
	cmd2 := exec.Command("vgcreate", "rtalurdisk"+deviceinfo.DeviceName, "/dev/"+deviceinfo.DeviceName)
	if err := cmd2.Run(); err != nil {
		fmt.Printf("rtalur we failed to run vgcreate")
	}
	fmt.Printf("rtalur we are now in transaction and executed commands")

	return nil
}

func txnHeketiCreateBrick(c transaction.TxnCtx) error {

	return nil
}
