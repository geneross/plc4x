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

// BACnetConstructedDataGroupMode is the data-structure of this message
type BACnetConstructedDataGroupMode struct {
	*BACnetConstructedData
	GroupMode *BACnetLiftGroupModeTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataGroupMode is the corresponding interface of BACnetConstructedDataGroupMode
type IBACnetConstructedDataGroupMode interface {
	IBACnetConstructedData
	// GetGroupMode returns GroupMode (property field)
	GetGroupMode() *BACnetLiftGroupModeTagged
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

func (m *BACnetConstructedDataGroupMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataGroupMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GROUP_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataGroupMode) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataGroupMode) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataGroupMode) GetGroupMode() *BACnetLiftGroupModeTagged {
	return m.GroupMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataGroupMode factory function for BACnetConstructedDataGroupMode
func NewBACnetConstructedDataGroupMode(groupMode *BACnetLiftGroupModeTagged, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataGroupMode {
	_result := &BACnetConstructedDataGroupMode{
		GroupMode:             groupMode,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataGroupMode(structType interface{}) *BACnetConstructedDataGroupMode {
	if casted, ok := structType.(BACnetConstructedDataGroupMode); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGroupMode); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataGroupMode(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataGroupMode(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataGroupMode) GetTypeName() string {
	return "BACnetConstructedDataGroupMode"
}

func (m *BACnetConstructedDataGroupMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataGroupMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (groupMode)
	lengthInBits += m.GroupMode.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataGroupMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataGroupModeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataGroupMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGroupMode"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (groupMode)
	if pullErr := readBuffer.PullContext("groupMode"); pullErr != nil {
		return nil, pullErr
	}
	_groupMode, _groupModeErr := BACnetLiftGroupModeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _groupModeErr != nil {
		return nil, errors.Wrap(_groupModeErr, "Error parsing 'groupMode' field")
	}
	groupMode := CastBACnetLiftGroupModeTagged(_groupMode)
	if closeErr := readBuffer.CloseContext("groupMode"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGroupMode"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataGroupMode{
		GroupMode:             CastBACnetLiftGroupModeTagged(groupMode),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataGroupMode) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGroupMode"); pushErr != nil {
			return pushErr
		}

		// Simple Field (groupMode)
		if pushErr := writeBuffer.PushContext("groupMode"); pushErr != nil {
			return pushErr
		}
		_groupModeErr := m.GroupMode.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("groupMode"); popErr != nil {
			return popErr
		}
		if _groupModeErr != nil {
			return errors.Wrap(_groupModeErr, "Error serializing 'groupMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGroupMode"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataGroupMode) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}