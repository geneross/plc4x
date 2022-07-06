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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CALReplyReply is the corresponding interface of CALReplyReply
type CALReplyReply interface {
	utils.LengthAware
	utils.Serializable
	NormalReply
	// GetCalReply returns CalReply (property field)
	GetCalReply() CALReply
	// GetPayloadLength returns PayloadLength (virtual field)
	GetPayloadLength() uint16
}

// CALReplyReplyExactly can be used when we want exactly this type and not a type which fulfills CALReplyReply.
// This is useful for switch cases.
type CALReplyReplyExactly interface {
	CALReplyReply
	isCALReplyReply() bool
}

// _CALReplyReply is the data-structure of this message
type _CALReplyReply struct {
	*_NormalReply
	CalReply CALReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALReplyReply) InitializeParent(parent NormalReply, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_CALReplyReply) GetParent() NormalReply {
	return m._NormalReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALReplyReply) GetCalReply() CALReply {
	return m.CalReply
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_CALReplyReply) GetPayloadLength() uint16 {
	return uint16(m.ReplyLength)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALReplyReply factory function for _CALReplyReply
func NewCALReplyReply(calReply CALReply, peekedByte byte, cBusOptions CBusOptions, replyLength uint16, requestContext RequestContext) *_CALReplyReply {
	_result := &_CALReplyReply{
		CalReply:     calReply,
		_NormalReply: NewNormalReply(peekedByte, cBusOptions, replyLength, requestContext),
	}
	_result._NormalReply._NormalReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALReplyReply(structType interface{}) CALReplyReply {
	if casted, ok := structType.(CALReplyReply); ok {
		return casted
	}
	if casted, ok := structType.(*CALReplyReply); ok {
		return *casted
	}
	return nil
}

func (m *_CALReplyReply) GetTypeName() string {
	return "CALReplyReply"
}

func (m *_CALReplyReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALReplyReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Manual Field (calReply)
	lengthInBits += uint16(int32(m.GetLengthInBytes()) * int32(int32(2)))

	return lengthInBits
}

func (m *_CALReplyReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALReplyReplyParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions, replyLength uint16, requestContext RequestContext) (CALReplyReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALReplyReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALReplyReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_payloadLength := replyLength
	payloadLength := uint16(_payloadLength)
	_ = payloadLength

	// Manual Field (calReply)
	_calReply, _calReplyErr := ReadCALReply(readBuffer, payloadLength)
	if _calReplyErr != nil {
		return nil, errors.Wrap(_calReplyErr, "Error parsing 'calReply' field of CALReplyReply")
	}
	calReply := _calReply.(CALReply)

	if closeErr := readBuffer.CloseContext("CALReplyReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALReplyReply")
	}

	// Create a partially initialized instance
	_child := &_CALReplyReply{
		CalReply: calReply,
		_NormalReply: &_NormalReply{
			CBusOptions:    cBusOptions,
			ReplyLength:    replyLength,
			RequestContext: requestContext,
		},
	}
	_child._NormalReply._NormalReplyChildRequirements = _child
	return _child, nil
}

func (m *_CALReplyReply) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALReplyReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALReplyReply")
		}
		// Virtual field
		if _payloadLengthErr := writeBuffer.WriteVirtual("payloadLength", m.GetPayloadLength()); _payloadLengthErr != nil {
			return errors.Wrap(_payloadLengthErr, "Error serializing 'payloadLength' field")
		}

		// Manual Field (calReply)
		_calReplyErr := WriteCALReply(writeBuffer, m.GetCalReply())
		if _calReplyErr != nil {
			return errors.Wrap(_calReplyErr, "Error serializing 'calReply' field")
		}

		if popErr := writeBuffer.PopContext("CALReplyReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALReplyReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALReplyReply) isCALReplyReply() bool {
	return true
}

func (m *_CALReplyReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
