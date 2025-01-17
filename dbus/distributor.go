package dbus

import (
	"github.com/godbus/dbus/v5"
	"unifiedpush.org/go/dbus_connector/definitions"
)

func NewDistributor(object dbus.BusObject) *Distributor {
	return &Distributor{
		object: object,
	}
}

type Distributor struct {
	object dbus.BusObject
}

func (d *Distributor) Register(name, token string) (definitions.RegisterStatus, string) {
	var status, reason string
	err := d.object.Call(definitions.DistributorInterface+".Register", dbus.Flags(0), name, token).Store(&status, &reason)
	if err != nil {
		return definitions.RegisterStatusFailedRequest, ""
	}

	registerStatus, ok := definitions.RegisterStatusMap[status]
	if !ok {
		return definitions.RegisterStatusFailedRequest, ""
	}

	return registerStatus, reason
}

func (d *Distributor) Unregister(token string) (err error) {
	result := d.object.Call(definitions.DistributorInterface+".Unregister", dbus.FlagNoReplyExpected, token)
	if result == nil {
		return nil // this should be some sort of error
	}
	return result.Err
}
