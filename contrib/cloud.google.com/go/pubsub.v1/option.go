// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2022 Datadog, Inc.

package pubsub

import (
	"gopkg.in/DataDog/dd-trace-go.v1/internal/namingschema"
)

type config struct {
	serviceName     string
	publishSpanName string
	receiveSpanName string
	measured        bool
}

func defaultConfig() *config {
	return &config{
		serviceName:     namingschema.ServiceNameOverrideV0("", ""),
		publishSpanName: namingschema.OpName(namingschema.GCPPubSubOutbound),
		receiveSpanName: namingschema.OpName(namingschema.GCPPubSubInbound),
		measured:        false,
	}
}

// Option is used to customize spans started by WrapReceiveHandler or Publish.
type Option = tracing.Option

// Deprecated: ReceiveOption has been deprecated in favor of Option.
type ReceiveOption = Option

// WithServiceName sets the service name tag for traces started by WrapReceiveHandler or Publish.
var WithServiceName = tracing.WithServiceName

// WithMeasured sets the measured tag for traces started by WrapReceiveHandler or Publish.
var WithMeasured = tracing.WithMeasured
