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

// QueryDataDescription is the corresponding interface of QueryDataDescription
type QueryDataDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRelativePath returns RelativePath (property field)
	GetRelativePath() ExtensionObjectDefinition
	// GetAttributeId returns AttributeId (property field)
	GetAttributeId() uint32
	// GetIndexRange returns IndexRange (property field)
	GetIndexRange() PascalString
}

// QueryDataDescriptionExactly can be used when we want exactly this type and not a type which fulfills QueryDataDescription.
// This is useful for switch cases.
type QueryDataDescriptionExactly interface {
	QueryDataDescription
	isQueryDataDescription() bool
}

// _QueryDataDescription is the data-structure of this message
type _QueryDataDescription struct {
	*_ExtensionObjectDefinition
	RelativePath ExtensionObjectDefinition
	AttributeId  uint32
	IndexRange   PascalString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_QueryDataDescription) GetIdentifier() string {
	return "572"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_QueryDataDescription) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_QueryDataDescription) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QueryDataDescription) GetRelativePath() ExtensionObjectDefinition {
	return m.RelativePath
}

func (m *_QueryDataDescription) GetAttributeId() uint32 {
	return m.AttributeId
}

func (m *_QueryDataDescription) GetIndexRange() PascalString {
	return m.IndexRange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewQueryDataDescription factory function for _QueryDataDescription
func NewQueryDataDescription(relativePath ExtensionObjectDefinition, attributeId uint32, indexRange PascalString) *_QueryDataDescription {
	_result := &_QueryDataDescription{
		RelativePath:               relativePath,
		AttributeId:                attributeId,
		IndexRange:                 indexRange,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastQueryDataDescription(structType any) QueryDataDescription {
	if casted, ok := structType.(QueryDataDescription); ok {
		return casted
	}
	if casted, ok := structType.(*QueryDataDescription); ok {
		return *casted
	}
	return nil
}

func (m *_QueryDataDescription) GetTypeName() string {
	return "QueryDataDescription"
}

func (m *_QueryDataDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (relativePath)
	lengthInBits += m.RelativePath.GetLengthInBits(ctx)

	// Simple field (attributeId)
	lengthInBits += 32

	// Simple field (indexRange)
	lengthInBits += m.IndexRange.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_QueryDataDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func QueryDataDescriptionParse(ctx context.Context, theBytes []byte, identifier string) (QueryDataDescription, error) {
	return QueryDataDescriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func QueryDataDescriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (QueryDataDescription, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("QueryDataDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QueryDataDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (relativePath)
	if pullErr := readBuffer.PullContext("relativePath"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for relativePath")
	}
	_relativePath, _relativePathErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("542"))
	if _relativePathErr != nil {
		return nil, errors.Wrap(_relativePathErr, "Error parsing 'relativePath' field of QueryDataDescription")
	}
	relativePath := _relativePath.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("relativePath"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for relativePath")
	}

	// Simple Field (attributeId)
	_attributeId, _attributeIdErr := readBuffer.ReadUint32("attributeId", 32)
	if _attributeIdErr != nil {
		return nil, errors.Wrap(_attributeIdErr, "Error parsing 'attributeId' field of QueryDataDescription")
	}
	attributeId := _attributeId

	// Simple Field (indexRange)
	if pullErr := readBuffer.PullContext("indexRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for indexRange")
	}
	_indexRange, _indexRangeErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _indexRangeErr != nil {
		return nil, errors.Wrap(_indexRangeErr, "Error parsing 'indexRange' field of QueryDataDescription")
	}
	indexRange := _indexRange.(PascalString)
	if closeErr := readBuffer.CloseContext("indexRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for indexRange")
	}

	if closeErr := readBuffer.CloseContext("QueryDataDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QueryDataDescription")
	}

	// Create a partially initialized instance
	_child := &_QueryDataDescription{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RelativePath:               relativePath,
		AttributeId:                attributeId,
		IndexRange:                 indexRange,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_QueryDataDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QueryDataDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("QueryDataDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for QueryDataDescription")
		}

		// Simple Field (relativePath)
		if pushErr := writeBuffer.PushContext("relativePath"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for relativePath")
		}
		_relativePathErr := writeBuffer.WriteSerializable(ctx, m.GetRelativePath())
		if popErr := writeBuffer.PopContext("relativePath"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for relativePath")
		}
		if _relativePathErr != nil {
			return errors.Wrap(_relativePathErr, "Error serializing 'relativePath' field")
		}

		// Simple Field (attributeId)
		attributeId := uint32(m.GetAttributeId())
		_attributeIdErr := writeBuffer.WriteUint32("attributeId", 32, uint32((attributeId)))
		if _attributeIdErr != nil {
			return errors.Wrap(_attributeIdErr, "Error serializing 'attributeId' field")
		}

		// Simple Field (indexRange)
		if pushErr := writeBuffer.PushContext("indexRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for indexRange")
		}
		_indexRangeErr := writeBuffer.WriteSerializable(ctx, m.GetIndexRange())
		if popErr := writeBuffer.PopContext("indexRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for indexRange")
		}
		if _indexRangeErr != nil {
			return errors.Wrap(_indexRangeErr, "Error serializing 'indexRange' field")
		}

		if popErr := writeBuffer.PopContext("QueryDataDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for QueryDataDescription")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_QueryDataDescription) isQueryDataDescription() bool {
	return true
}

func (m *_QueryDataDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
