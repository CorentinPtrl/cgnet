package cgnet

import (
	"errors"
)

// Configure a silce of commands on the device. The commands executed automatically in the 'configure terminal' mode.
func (d Device) Configure(cmds []string) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, err := d.exec("conf t")
	if err != nil {
		return err
	}

	err = nil
	for _, cmd := range cmds {
		_, err = d.exec(cmd)
		if err != nil {
			err = errors.New("error on command " + cmd + ". aborting, " + err.Error())
			break
		}
	}
	_, err = d.exec("end")
	if err != nil {
		return err
	}

	return nil
}
