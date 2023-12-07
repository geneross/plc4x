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

// VariantBoolean is the corresponding interface of VariantBoolean
type VariantBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []byte
}

// VariantBooleanExactly can be used when we want exactly this type and not a type which fulfills VariantBoolean.
// This is useful for switch cases.
type VariantBooleanExactly interface {
	VariantBoolean
	isVariantBoolean() bool
}

// _VariantBoolean is the data-structure of this message
type _VariantBoolean struct {
	*_Variant
	ArrayLength *int32
	Value       []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantBoolean) GetVariantType() uint8 {
	return uint8(1)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantBoolean) InitializeParent(parent Variant, arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool) {
	m.ArrayLengthSpecified = arrayLengthSpecified
	m.ArrayDimensionsSpecified = arrayDimensionsSpecified
	m.NoOfArrayDimensions = noOfArrayDimensions
	m.ArrayDimensions = arrayDimensions
}

func (m *_VariantBoolean) GetParent() Variant {
	return m._Variant
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantBoolean) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantBoolean) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewVariantBoolean factory function for _VariantBoolean
func NewVariantBoolean(arrayLength *int32, value []byte, arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool) *_VariantBoolean {
	_result := &_VariantBoolean{
		ArrayLength: arrayLength,
		Value:       value,
		_Variant:    NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
	}
	_result._Variant._VariantChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastVariantBoolean(structType any) VariantBoolean {
	if casted, ok := structType.(VariantBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*VariantBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_VariantBoolean) GetTypeName() string {
	return "VariantBoolean"
}

func (m *_VariantBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_VariantBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func VariantBooleanParse(ctx context.Context, theBytes []byte, arrayLengthSpecified bool) (VariantBoolean, error) {
	return VariantBooleanParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), arrayLengthSpecified)
}

func VariantBooleanParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, arrayLengthSpecified bool) (VariantBoolean, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("VariantBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Optional Field (arrayLength) (Can be skipped, if a given expression evaluates to false)
	var arrayLength *int32 = nil
	if arrayLengthSpecified {
		_val, _err := readBuffer.ReadInt32("arrayLength", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'arrayLength' field of VariantBoolean")
		}
		arrayLength = &_val
	}
	// Byte Array field (value)
	numberOfBytesvalue := int(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return uint16(uint16(1)) }, func() any { return uint16((*arrayLength)) }).(uint16))
	value, _readArrayErr := readBuffer.ReadByteArray("value", numberOfBytesvalue)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'value' field of VariantBoolean")
	}

	if closeErr := readBuffer.CloseContext("VariantBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantBoolean")
	}

	// Create a partially initialized instance
	_child := &_VariantBoolean{
		_Variant:    &_Variant{},
		ArrayLength: arrayLength,
		Value:       value,
	}
	_child._Variant._VariantChildRequirements = _child
	return _child, nil
}

func (m *_VariantBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantBoolean")
		}

		// Optional Field (arrayLength) (Can be skipped, if the value is null)
		var arrayLength *int32 = nil
		if m.GetArrayLength() != nil {
			arrayLength = m.GetArrayLength()
			_arrayLengthErr := writeBuffer.WriteInt32("arrayLength", 32, int32(*(arrayLength)))
			if _arrayLengthErr != nil {
				return errors.Wrap(_arrayLengthErr, "Error serializing 'arrayLength' field")
			}
		}

		// Array Field (value)
		// Byte Array field (value)
		if err := writeBuffer.WriteByteArray("value", m.GetValue()); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantBoolean")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantBoolean) isVariantBoolean() bool {
	return true
}

func (m *_VariantBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
