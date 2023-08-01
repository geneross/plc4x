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

// BACnetConstructedDataNetworkType is the corresponding interface of BACnetConstructedDataNetworkType
type BACnetConstructedDataNetworkType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNetworkType returns NetworkType (property field)
	GetNetworkType() BACnetNetworkTypeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetNetworkTypeTagged
}

// BACnetConstructedDataNetworkTypeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataNetworkType.
// This is useful for switch cases.
type BACnetConstructedDataNetworkTypeExactly interface {
	BACnetConstructedDataNetworkType
	isBACnetConstructedDataNetworkType() bool
}

// _BACnetConstructedDataNetworkType is the data-structure of this message
type _BACnetConstructedDataNetworkType struct {
	*_BACnetConstructedData
	NetworkType BACnetNetworkTypeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNetworkType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNetworkType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NETWORK_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNetworkType) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataNetworkType) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkType) GetNetworkType() BACnetNetworkTypeTagged {
	return m.NetworkType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkType) GetActualValue() BACnetNetworkTypeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetNetworkTypeTagged(m.GetNetworkType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNetworkType factory function for _BACnetConstructedDataNetworkType
func NewBACnetConstructedDataNetworkType(networkType BACnetNetworkTypeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNetworkType {
	_result := &_BACnetConstructedDataNetworkType{
		NetworkType:            networkType,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNetworkType(structType any) BACnetConstructedDataNetworkType {
	if casted, ok := structType.(BACnetConstructedDataNetworkType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNetworkType) GetTypeName() string {
	return "BACnetConstructedDataNetworkType"
}

func (m *_BACnetConstructedDataNetworkType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (networkType)
	lengthInBits += m.NetworkType.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNetworkType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataNetworkTypeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNetworkType, error) {
	return BACnetConstructedDataNetworkTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataNetworkTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNetworkType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNetworkType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkType)
	if pullErr := readBuffer.PullContext("networkType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkType")
	}
	_networkType, _networkTypeErr := BACnetNetworkTypeTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _networkTypeErr != nil {
		return nil, errors.Wrap(_networkTypeErr, "Error parsing 'networkType' field of BACnetConstructedDataNetworkType")
	}
	networkType := _networkType.(BACnetNetworkTypeTagged)
	if closeErr := readBuffer.CloseContext("networkType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkType")
	}

	// Virtual field
	_actualValue := networkType
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNetworkType")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataNetworkType{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NetworkType: networkType,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataNetworkType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNetworkType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNetworkType")
		}

		// Simple Field (networkType)
		if pushErr := writeBuffer.PushContext("networkType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for networkType")
		}
		_networkTypeErr := writeBuffer.WriteSerializable(ctx, m.GetNetworkType())
		if popErr := writeBuffer.PopContext("networkType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for networkType")
		}
		if _networkTypeErr != nil {
			return errors.Wrap(_networkTypeErr, "Error serializing 'networkType' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNetworkType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNetworkType) isBACnetConstructedDataNetworkType() bool {
	return true
}

func (m *_BACnetConstructedDataNetworkType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
