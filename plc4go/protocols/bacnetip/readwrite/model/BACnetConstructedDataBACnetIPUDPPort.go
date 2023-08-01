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

// BACnetConstructedDataBACnetIPUDPPort is the corresponding interface of BACnetConstructedDataBACnetIPUDPPort
type BACnetConstructedDataBACnetIPUDPPort interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpUdpPort returns IpUdpPort (property field)
	GetIpUdpPort() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataBACnetIPUDPPortExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBACnetIPUDPPort.
// This is useful for switch cases.
type BACnetConstructedDataBACnetIPUDPPortExactly interface {
	BACnetConstructedDataBACnetIPUDPPort
	isBACnetConstructedDataBACnetIPUDPPort() bool
}

// _BACnetConstructedDataBACnetIPUDPPort is the data-structure of this message
type _BACnetConstructedDataBACnetIPUDPPort struct {
	*_BACnetConstructedData
	IpUdpPort BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IP_UDP_PORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPUDPPort) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetIpUdpPort() BACnetApplicationTagUnsignedInteger {
	return m.IpUdpPort
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpUdpPort())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBACnetIPUDPPort factory function for _BACnetConstructedDataBACnetIPUDPPort
func NewBACnetConstructedDataBACnetIPUDPPort(ipUdpPort BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPUDPPort {
	_result := &_BACnetConstructedDataBACnetIPUDPPort{
		IpUdpPort:              ipUdpPort,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPUDPPort(structType any) BACnetConstructedDataBACnetIPUDPPort {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPUDPPort); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPUDPPort); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPUDPPort"
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (ipUdpPort)
	lengthInBits += m.IpUdpPort.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBACnetIPUDPPortParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPUDPPort, error) {
	return BACnetConstructedDataBACnetIPUDPPortParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBACnetIPUDPPortParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPUDPPort, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPUDPPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPUDPPort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipUdpPort)
	if pullErr := readBuffer.PullContext("ipUdpPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipUdpPort")
	}
	_ipUdpPort, _ipUdpPortErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _ipUdpPortErr != nil {
		return nil, errors.Wrap(_ipUdpPortErr, "Error parsing 'ipUdpPort' field of BACnetConstructedDataBACnetIPUDPPort")
	}
	ipUdpPort := _ipUdpPort.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("ipUdpPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipUdpPort")
	}

	// Virtual field
	_actualValue := ipUdpPort
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPUDPPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPUDPPort")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBACnetIPUDPPort{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		IpUdpPort: ipUdpPort,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPUDPPort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPUDPPort")
		}

		// Simple Field (ipUdpPort)
		if pushErr := writeBuffer.PushContext("ipUdpPort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipUdpPort")
		}
		_ipUdpPortErr := writeBuffer.WriteSerializable(ctx, m.GetIpUdpPort())
		if popErr := writeBuffer.PopContext("ipUdpPort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipUdpPort")
		}
		if _ipUdpPortErr != nil {
			return errors.Wrap(_ipUdpPortErr, "Error serializing 'ipUdpPort' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPUDPPort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPUDPPort")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) isBACnetConstructedDataBACnetIPUDPPort() bool {
	return true
}

func (m *_BACnetConstructedDataBACnetIPUDPPort) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
