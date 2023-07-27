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

// AdditionalParametersType is the corresponding interface of AdditionalParametersType
type AdditionalParametersType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfParameters returns NoOfParameters (property field)
	GetNoOfParameters() int32
	// GetParameters returns Parameters (property field)
	GetParameters() []ExtensionObjectDefinition
}

// AdditionalParametersTypeExactly can be used when we want exactly this type and not a type which fulfills AdditionalParametersType.
// This is useful for switch cases.
type AdditionalParametersTypeExactly interface {
	AdditionalParametersType
	isAdditionalParametersType() bool
}

// _AdditionalParametersType is the data-structure of this message
type _AdditionalParametersType struct {
	*_ExtensionObjectDefinition
	NoOfParameters int32
	Parameters     []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdditionalParametersType) GetIdentifier() string {
	return "16315"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdditionalParametersType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_AdditionalParametersType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdditionalParametersType) GetNoOfParameters() int32 {
	return m.NoOfParameters
}

func (m *_AdditionalParametersType) GetParameters() []ExtensionObjectDefinition {
	return m.Parameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdditionalParametersType factory function for _AdditionalParametersType
func NewAdditionalParametersType(noOfParameters int32, parameters []ExtensionObjectDefinition) *_AdditionalParametersType {
	_result := &_AdditionalParametersType{
		NoOfParameters:             noOfParameters,
		Parameters:                 parameters,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdditionalParametersType(structType any) AdditionalParametersType {
	if casted, ok := structType.(AdditionalParametersType); ok {
		return casted
	}
	if casted, ok := structType.(*AdditionalParametersType); ok {
		return *casted
	}
	return nil
}

func (m *_AdditionalParametersType) GetTypeName() string {
	return "AdditionalParametersType"
}

func (m *_AdditionalParametersType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (noOfParameters)
	lengthInBits += 32

	// Array field
	if len(m.Parameters) > 0 {
		for _curItem, element := range m.Parameters {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Parameters), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AdditionalParametersType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdditionalParametersTypeParse(ctx context.Context, theBytes []byte, identifier string) (AdditionalParametersType, error) {
	return AdditionalParametersTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func AdditionalParametersTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (AdditionalParametersType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AdditionalParametersType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdditionalParametersType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (noOfParameters)
	_noOfParameters, _noOfParametersErr := readBuffer.ReadInt32("noOfParameters", 32)
	if _noOfParametersErr != nil {
		return nil, errors.Wrap(_noOfParametersErr, "Error parsing 'noOfParameters' field of AdditionalParametersType")
	}
	noOfParameters := _noOfParameters

	// Array field (parameters)
	if pullErr := readBuffer.PullContext("parameters", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for parameters")
	}
	// Count array
	parameters := make([]ExtensionObjectDefinition, noOfParameters)
	// This happens when the size is set conditional to 0
	if len(parameters) == 0 {
		parameters = nil
	}
	{
		_numItems := uint16(noOfParameters)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "14535")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'parameters' field of AdditionalParametersType")
			}
			parameters[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("parameters", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for parameters")
	}

	if closeErr := readBuffer.CloseContext("AdditionalParametersType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdditionalParametersType")
	}

	// Create a partially initialized instance
	_child := &_AdditionalParametersType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NoOfParameters:             noOfParameters,
		Parameters:                 parameters,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_AdditionalParametersType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdditionalParametersType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdditionalParametersType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdditionalParametersType")
		}

		// Simple Field (noOfParameters)
		noOfParameters := int32(m.GetNoOfParameters())
		_noOfParametersErr := writeBuffer.WriteInt32("noOfParameters", 32, (noOfParameters))
		if _noOfParametersErr != nil {
			return errors.Wrap(_noOfParametersErr, "Error serializing 'noOfParameters' field")
		}

		// Array Field (parameters)
		if pushErr := writeBuffer.PushContext("parameters", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for parameters")
		}
		for _curItem, _element := range m.GetParameters() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetParameters()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'parameters' field")
			}
		}
		if popErr := writeBuffer.PopContext("parameters", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for parameters")
		}

		if popErr := writeBuffer.PopContext("AdditionalParametersType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdditionalParametersType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdditionalParametersType) isAdditionalParametersType() bool {
	return true
}

func (m *_AdditionalParametersType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}