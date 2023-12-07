/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableUadp is an enum
type OpcuaNodeIdServicesVariableUadp int32

type IOpcuaNodeIdServicesVariableUadp interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableUadp_UadpNetworkMessageContentMask_OptionSetValues          OpcuaNodeIdServicesVariableUadp = 15643
	OpcuaNodeIdServicesVariableUadp_UadpDataSetMessageContentMask_OptionSetValues          OpcuaNodeIdServicesVariableUadp = 15647
	OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetOffset             OpcuaNodeIdServicesVariableUadp = 17477
	OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_GroupVersion                OpcuaNodeIdServicesVariableUadp = 21106
	OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_DataSetOrdering             OpcuaNodeIdServicesVariableUadp = 21107
	OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_NetworkMessageContentMask   OpcuaNodeIdServicesVariableUadp = 21108
	OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_SamplingOffset              OpcuaNodeIdServicesVariableUadp = 21109
	OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_PublishingOffset            OpcuaNodeIdServicesVariableUadp = 21110
	OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_DataSetMessageContentMask OpcuaNodeIdServicesVariableUadp = 21112
	OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_ConfiguredSize            OpcuaNodeIdServicesVariableUadp = 21113
	OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_NetworkMessageNumber      OpcuaNodeIdServicesVariableUadp = 21114
	OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_DataSetOffset             OpcuaNodeIdServicesVariableUadp = 21115
	OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_GroupVersion              OpcuaNodeIdServicesVariableUadp = 21117
	OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_NetworkMessageNumber      OpcuaNodeIdServicesVariableUadp = 21119
	OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetClassId            OpcuaNodeIdServicesVariableUadp = 21120
	OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_NetworkMessageContentMask OpcuaNodeIdServicesVariableUadp = 21121
	OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetMessageContentMask OpcuaNodeIdServicesVariableUadp = 21122
	OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_PublishingInterval        OpcuaNodeIdServicesVariableUadp = 21123
	OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_ProcessingOffset          OpcuaNodeIdServicesVariableUadp = 21124
	OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_ReceiveOffset             OpcuaNodeIdServicesVariableUadp = 21125
)

var OpcuaNodeIdServicesVariableUadpValues []OpcuaNodeIdServicesVariableUadp

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableUadpValues = []OpcuaNodeIdServicesVariableUadp{
		OpcuaNodeIdServicesVariableUadp_UadpNetworkMessageContentMask_OptionSetValues,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetMessageContentMask_OptionSetValues,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetOffset,
		OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_GroupVersion,
		OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_DataSetOrdering,
		OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_NetworkMessageContentMask,
		OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_SamplingOffset,
		OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_PublishingOffset,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_DataSetMessageContentMask,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_ConfiguredSize,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_NetworkMessageNumber,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_DataSetOffset,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_GroupVersion,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_NetworkMessageNumber,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetClassId,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_NetworkMessageContentMask,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetMessageContentMask,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_PublishingInterval,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_ProcessingOffset,
		OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_ReceiveOffset,
	}
}

func OpcuaNodeIdServicesVariableUadpByValue(value int32) (enum OpcuaNodeIdServicesVariableUadp, ok bool) {
	switch value {
	case 15643:
		return OpcuaNodeIdServicesVariableUadp_UadpNetworkMessageContentMask_OptionSetValues, true
	case 15647:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetMessageContentMask_OptionSetValues, true
	case 17477:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetOffset, true
	case 21106:
		return OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_GroupVersion, true
	case 21107:
		return OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_DataSetOrdering, true
	case 21108:
		return OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_NetworkMessageContentMask, true
	case 21109:
		return OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_SamplingOffset, true
	case 21110:
		return OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_PublishingOffset, true
	case 21112:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_DataSetMessageContentMask, true
	case 21113:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_ConfiguredSize, true
	case 21114:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_NetworkMessageNumber, true
	case 21115:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_DataSetOffset, true
	case 21117:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_GroupVersion, true
	case 21119:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_NetworkMessageNumber, true
	case 21120:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetClassId, true
	case 21121:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_NetworkMessageContentMask, true
	case 21122:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetMessageContentMask, true
	case 21123:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_PublishingInterval, true
	case 21124:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_ProcessingOffset, true
	case 21125:
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_ReceiveOffset, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableUadpByName(value string) (enum OpcuaNodeIdServicesVariableUadp, ok bool) {
	switch value {
	case "UadpNetworkMessageContentMask_OptionSetValues":
		return OpcuaNodeIdServicesVariableUadp_UadpNetworkMessageContentMask_OptionSetValues, true
	case "UadpDataSetMessageContentMask_OptionSetValues":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetMessageContentMask_OptionSetValues, true
	case "UadpDataSetReaderMessageType_DataSetOffset":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetOffset, true
	case "UadpWriterGroupMessageType_GroupVersion":
		return OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_GroupVersion, true
	case "UadpWriterGroupMessageType_DataSetOrdering":
		return OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_DataSetOrdering, true
	case "UadpWriterGroupMessageType_NetworkMessageContentMask":
		return OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_NetworkMessageContentMask, true
	case "UadpWriterGroupMessageType_SamplingOffset":
		return OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_SamplingOffset, true
	case "UadpWriterGroupMessageType_PublishingOffset":
		return OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_PublishingOffset, true
	case "UadpDataSetWriterMessageType_DataSetMessageContentMask":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_DataSetMessageContentMask, true
	case "UadpDataSetWriterMessageType_ConfiguredSize":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_ConfiguredSize, true
	case "UadpDataSetWriterMessageType_NetworkMessageNumber":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_NetworkMessageNumber, true
	case "UadpDataSetWriterMessageType_DataSetOffset":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_DataSetOffset, true
	case "UadpDataSetReaderMessageType_GroupVersion":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_GroupVersion, true
	case "UadpDataSetReaderMessageType_NetworkMessageNumber":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_NetworkMessageNumber, true
	case "UadpDataSetReaderMessageType_DataSetClassId":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetClassId, true
	case "UadpDataSetReaderMessageType_NetworkMessageContentMask":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_NetworkMessageContentMask, true
	case "UadpDataSetReaderMessageType_DataSetMessageContentMask":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetMessageContentMask, true
	case "UadpDataSetReaderMessageType_PublishingInterval":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_PublishingInterval, true
	case "UadpDataSetReaderMessageType_ProcessingOffset":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_ProcessingOffset, true
	case "UadpDataSetReaderMessageType_ReceiveOffset":
		return OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_ReceiveOffset, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableUadpKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableUadpValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableUadp(structType any) OpcuaNodeIdServicesVariableUadp {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableUadp {
		if sOpcuaNodeIdServicesVariableUadp, ok := typ.(OpcuaNodeIdServicesVariableUadp); ok {
			return sOpcuaNodeIdServicesVariableUadp
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableUadp) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableUadp) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableUadpParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableUadp, error) {
	return OpcuaNodeIdServicesVariableUadpParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableUadpParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableUadp, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableUadp", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableUadp")
	}
	if enum, ok := OpcuaNodeIdServicesVariableUadpByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableUadp")
		return OpcuaNodeIdServicesVariableUadp(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableUadp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableUadp) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableUadp", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableUadp) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableUadp_UadpNetworkMessageContentMask_OptionSetValues:
		return "UadpNetworkMessageContentMask_OptionSetValues"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetMessageContentMask_OptionSetValues:
		return "UadpDataSetMessageContentMask_OptionSetValues"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetOffset:
		return "UadpDataSetReaderMessageType_DataSetOffset"
	case OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_GroupVersion:
		return "UadpWriterGroupMessageType_GroupVersion"
	case OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_DataSetOrdering:
		return "UadpWriterGroupMessageType_DataSetOrdering"
	case OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_NetworkMessageContentMask:
		return "UadpWriterGroupMessageType_NetworkMessageContentMask"
	case OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_SamplingOffset:
		return "UadpWriterGroupMessageType_SamplingOffset"
	case OpcuaNodeIdServicesVariableUadp_UadpWriterGroupMessageType_PublishingOffset:
		return "UadpWriterGroupMessageType_PublishingOffset"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_DataSetMessageContentMask:
		return "UadpDataSetWriterMessageType_DataSetMessageContentMask"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_ConfiguredSize:
		return "UadpDataSetWriterMessageType_ConfiguredSize"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_NetworkMessageNumber:
		return "UadpDataSetWriterMessageType_NetworkMessageNumber"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetWriterMessageType_DataSetOffset:
		return "UadpDataSetWriterMessageType_DataSetOffset"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_GroupVersion:
		return "UadpDataSetReaderMessageType_GroupVersion"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_NetworkMessageNumber:
		return "UadpDataSetReaderMessageType_NetworkMessageNumber"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetClassId:
		return "UadpDataSetReaderMessageType_DataSetClassId"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_NetworkMessageContentMask:
		return "UadpDataSetReaderMessageType_NetworkMessageContentMask"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_DataSetMessageContentMask:
		return "UadpDataSetReaderMessageType_DataSetMessageContentMask"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_PublishingInterval:
		return "UadpDataSetReaderMessageType_PublishingInterval"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_ProcessingOffset:
		return "UadpDataSetReaderMessageType_ProcessingOffset"
	case OpcuaNodeIdServicesVariableUadp_UadpDataSetReaderMessageType_ReceiveOffset:
		return "UadpDataSetReaderMessageType_ReceiveOffset"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableUadp) String() string {
	return e.PLC4XEnumName()
}
