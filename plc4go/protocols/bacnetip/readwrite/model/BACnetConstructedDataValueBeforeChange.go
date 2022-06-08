/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetConstructedDataValueBeforeChange is the data-structure of this message
type BACnetConstructedDataValueBeforeChange struct {
	*BACnetConstructedData
	ValuesBeforeChange *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataValueBeforeChange is the corresponding interface of BACnetConstructedDataValueBeforeChange
type IBACnetConstructedDataValueBeforeChange interface {
	IBACnetConstructedData
	// GetValuesBeforeChange returns ValuesBeforeChange (property field)
	GetValuesBeforeChange() *BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataValueBeforeChange) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataValueBeforeChange) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VALUE_BEFORE_CHANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataValueBeforeChange) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataValueBeforeChange) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataValueBeforeChange) GetValuesBeforeChange() *BACnetApplicationTagUnsignedInteger {
	return m.ValuesBeforeChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataValueBeforeChange factory function for BACnetConstructedDataValueBeforeChange
func NewBACnetConstructedDataValueBeforeChange(valuesBeforeChange *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataValueBeforeChange {
	_result := &BACnetConstructedDataValueBeforeChange{
		ValuesBeforeChange:    valuesBeforeChange,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataValueBeforeChange(structType interface{}) *BACnetConstructedDataValueBeforeChange {
	if casted, ok := structType.(BACnetConstructedDataValueBeforeChange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataValueBeforeChange); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataValueBeforeChange(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataValueBeforeChange(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataValueBeforeChange) GetTypeName() string {
	return "BACnetConstructedDataValueBeforeChange"
}

func (m *BACnetConstructedDataValueBeforeChange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataValueBeforeChange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (valuesBeforeChange)
	lengthInBits += m.ValuesBeforeChange.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataValueBeforeChange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataValueBeforeChangeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataValueBeforeChange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataValueBeforeChange"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (valuesBeforeChange)
	if pullErr := readBuffer.PullContext("valuesBeforeChange"); pullErr != nil {
		return nil, pullErr
	}
	_valuesBeforeChange, _valuesBeforeChangeErr := BACnetApplicationTagParse(readBuffer)
	if _valuesBeforeChangeErr != nil {
		return nil, errors.Wrap(_valuesBeforeChangeErr, "Error parsing 'valuesBeforeChange' field")
	}
	valuesBeforeChange := CastBACnetApplicationTagUnsignedInteger(_valuesBeforeChange)
	if closeErr := readBuffer.CloseContext("valuesBeforeChange"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataValueBeforeChange"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataValueBeforeChange{
		ValuesBeforeChange:    CastBACnetApplicationTagUnsignedInteger(valuesBeforeChange),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataValueBeforeChange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataValueBeforeChange"); pushErr != nil {
			return pushErr
		}

		// Simple Field (valuesBeforeChange)
		if pushErr := writeBuffer.PushContext("valuesBeforeChange"); pushErr != nil {
			return pushErr
		}
		_valuesBeforeChangeErr := m.ValuesBeforeChange.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("valuesBeforeChange"); popErr != nil {
			return popErr
		}
		if _valuesBeforeChangeErr != nil {
			return errors.Wrap(_valuesBeforeChangeErr, "Error serializing 'valuesBeforeChange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataValueBeforeChange"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataValueBeforeChange) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}