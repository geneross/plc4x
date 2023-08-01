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

// BACnetConstructedDataDeviceType is the corresponding interface of BACnetConstructedDataDeviceType
type BACnetConstructedDataDeviceType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDeviceType returns DeviceType (property field)
	GetDeviceType() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataDeviceTypeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDeviceType.
// This is useful for switch cases.
type BACnetConstructedDataDeviceTypeExactly interface {
	BACnetConstructedDataDeviceType
	isBACnetConstructedDataDeviceType() bool
}

// _BACnetConstructedDataDeviceType is the data-structure of this message
type _BACnetConstructedDataDeviceType struct {
	*_BACnetConstructedData
	DeviceType BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDeviceType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDeviceType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEVICE_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDeviceType) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDeviceType) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDeviceType) GetDeviceType() BACnetApplicationTagCharacterString {
	return m.DeviceType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDeviceType) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetDeviceType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDeviceType factory function for _BACnetConstructedDataDeviceType
func NewBACnetConstructedDataDeviceType(deviceType BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDeviceType {
	_result := &_BACnetConstructedDataDeviceType{
		DeviceType:             deviceType,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDeviceType(structType any) BACnetConstructedDataDeviceType {
	if casted, ok := structType.(BACnetConstructedDataDeviceType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDeviceType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDeviceType) GetTypeName() string {
	return "BACnetConstructedDataDeviceType"
}

func (m *_BACnetConstructedDataDeviceType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (deviceType)
	lengthInBits += m.DeviceType.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDeviceType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataDeviceTypeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDeviceType, error) {
	return BACnetConstructedDataDeviceTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataDeviceTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDeviceType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDeviceType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDeviceType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (deviceType)
	if pullErr := readBuffer.PullContext("deviceType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deviceType")
	}
	_deviceType, _deviceTypeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _deviceTypeErr != nil {
		return nil, errors.Wrap(_deviceTypeErr, "Error parsing 'deviceType' field of BACnetConstructedDataDeviceType")
	}
	deviceType := _deviceType.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("deviceType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deviceType")
	}

	// Virtual field
	_actualValue := deviceType
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDeviceType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDeviceType")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDeviceType{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DeviceType: deviceType,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDeviceType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDeviceType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDeviceType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDeviceType")
		}

		// Simple Field (deviceType)
		if pushErr := writeBuffer.PushContext("deviceType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceType")
		}
		_deviceTypeErr := writeBuffer.WriteSerializable(ctx, m.GetDeviceType())
		if popErr := writeBuffer.PopContext("deviceType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceType")
		}
		if _deviceTypeErr != nil {
			return errors.Wrap(_deviceTypeErr, "Error serializing 'deviceType' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDeviceType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDeviceType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDeviceType) isBACnetConstructedDataDeviceType() bool {
	return true
}

func (m *_BACnetConstructedDataDeviceType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
