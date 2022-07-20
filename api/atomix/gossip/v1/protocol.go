// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package v1

import "time"

type PartitionID uint64

type ReplicaID string

type MemberID string

type PeerID MemberID

type PeerGroupID uint64

type LogicalTime uint64

type Epoch uint64

type PhysicalTime = time.Time
