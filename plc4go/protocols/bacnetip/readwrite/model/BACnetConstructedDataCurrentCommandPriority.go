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

// BACnetConstructedDataCurrentCommandPriority is the data-structure of this message
type BACnetConstructedDataCurrentCommandPriority struct {
	*BACnetConstructedData
	CurrentCommandPriority *BACnetOptionalUnsigned

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataCurrentCommandPriority is the corresponding interface of BACnetConstructedDataCurrentCommandPriority
type IBACnetConstructedDataCurrentCommandPriority interface {
	IBACnetConstructedData
	// GetCurrentCommandPriority returns CurrentCommandPriority (property field)
	GetCurrentCommandPriority() *BACnetOptionalUnsigned
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

func (m *BACnetConstructedDataCurrentCommandPriority) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataCurrentCommandPriority) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CURRENT_COMMAND_PRIORITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataCurrentCommandPriority) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataCurrentCommandPriority) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataCurrentCommandPriority) GetCurrentCommandPriority() *BACnetOptionalUnsigned {
	return m.CurrentCommandPriority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCurrentCommandPriority factory function for BACnetConstructedDataCurrentCommandPriority
func NewBACnetConstructedDataCurrentCommandPriority(currentCommandPriority *BACnetOptionalUnsigned, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataCurrentCommandPriority {
	_result := &BACnetConstructedDataCurrentCommandPriority{
		CurrentCommandPriority: currentCommandPriority,
		BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataCurrentCommandPriority(structType interface{}) *BACnetConstructedDataCurrentCommandPriority {
	if casted, ok := structType.(BACnetConstructedDataCurrentCommandPriority); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCurrentCommandPriority); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataCurrentCommandPriority(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataCurrentCommandPriority(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataCurrentCommandPriority) GetTypeName() string {
	return "BACnetConstructedDataCurrentCommandPriority"
}

func (m *BACnetConstructedDataCurrentCommandPriority) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataCurrentCommandPriority) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (currentCommandPriority)
	lengthInBits += m.CurrentCommandPriority.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataCurrentCommandPriority) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCurrentCommandPriorityParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataCurrentCommandPriority, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCurrentCommandPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCurrentCommandPriority")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (currentCommandPriority)
	if pullErr := readBuffer.PullContext("currentCommandPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for currentCommandPriority")
	}
	_currentCommandPriority, _currentCommandPriorityErr := BACnetOptionalUnsignedParse(readBuffer)
	if _currentCommandPriorityErr != nil {
		return nil, errors.Wrap(_currentCommandPriorityErr, "Error parsing 'currentCommandPriority' field")
	}
	currentCommandPriority := CastBACnetOptionalUnsigned(_currentCommandPriority)
	if closeErr := readBuffer.CloseContext("currentCommandPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for currentCommandPriority")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCurrentCommandPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCurrentCommandPriority")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataCurrentCommandPriority{
		CurrentCommandPriority: CastBACnetOptionalUnsigned(currentCommandPriority),
		BACnetConstructedData:  &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataCurrentCommandPriority) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCurrentCommandPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCurrentCommandPriority")
		}

		// Simple Field (currentCommandPriority)
		if pushErr := writeBuffer.PushContext("currentCommandPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for currentCommandPriority")
		}
		_currentCommandPriorityErr := m.CurrentCommandPriority.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("currentCommandPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for currentCommandPriority")
		}
		if _currentCommandPriorityErr != nil {
			return errors.Wrap(_currentCommandPriorityErr, "Error serializing 'currentCommandPriority' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCurrentCommandPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCurrentCommandPriority")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataCurrentCommandPriority) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}