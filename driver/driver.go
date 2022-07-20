// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	gossipv1 "github.com/atomix/gossip-storage/api/atomix/gossip/v1"
	"github.com/atomix/runtime/sdk/pkg/runtime"
)

const (
	name    = "Gossip"
	version = "v1beta1"
)

var Driver = runtime.NewDriver[*gossipv1.GossipConfig](name, version, func(ctx context.Context, config *gossipv1.GossipConfig) (runtime.Client, error) {
	panic("not implemented")
})
