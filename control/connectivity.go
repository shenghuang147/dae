/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) since 2023, mzz2017 <mzz@tuta.io>
 */

package control

import (
	"github.com/cilium/ebpf"
	"github.com/sirupsen/logrus"
)

func (c *ControlPlaneCore) OutboundAliveChangeCallback(outbound uint8) func(alive bool, l4proto uint8, ipversion uint8) {
	return func(alive bool, l4proto uint8, ipversion uint8) {
		c.log.WithFields(logrus.Fields{
			"alive":       alive,
			"l4proto":     l4proto,
			"ipversion":   ipversion,
			"outbound_id": outbound,
		}).Tracef("outbound alive state changed")

		value := uint32(0)
		if alive {
			value = 1
		}
		_ = c.bpf.OutboundConnectivityMap.Update(bpfOutboundConnectivityQuery{
			Outbound:  outbound,
			L4proto:   l4proto,
			Ipversion: ipversion,
		}, value, ebpf.UpdateAny)
	}
}