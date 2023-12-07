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

// Constant values.
const AlarmMessageObjectPushType_VARIABLESPEC uint8 = 0x12

// AlarmMessageObjectPushType is the corresponding interface of AlarmMessageObjectPushType
type AlarmMessageObjectPushType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLengthSpec returns LengthSpec (property field)
	GetLengthSpec() uint8
	// GetSyntaxId returns SyntaxId (property field)
	GetSyntaxId() SyntaxIdType
	// GetNumberOfValues returns NumberOfValues (property field)
	GetNumberOfValues() uint8
	// GetEventId returns EventId (property field)
	GetEventId() uint32
	// GetEventState returns EventState (property field)
	GetEventState() State
	// GetLocalState returns LocalState (property field)
	GetLocalState() State
	// GetAckStateGoing returns AckStateGoing (property field)
	GetAckStateGoing() State
	// GetAckStateComing returns AckStateComing (property field)
	GetAckStateComing() State
	// GetAssociatedValues returns AssociatedValues (property field)
	GetAssociatedValues() []AssociatedValueType
}

// AlarmMessageObjectPushTypeExactly can be used when we want exactly this type and not a type which fulfills AlarmMessageObjectPushType.
// This is useful for switch cases.
type AlarmMessageObjectPushTypeExactly interface {
	AlarmMessageObjectPushType
	isAlarmMessageObjectPushType() bool
}

// _AlarmMessageObjectPushType is the data-structure of this message
type _AlarmMessageObjectPushType struct {
	LengthSpec       uint8
	SyntaxId         SyntaxIdType
	NumberOfValues   uint8
	EventId          uint32
	EventState       State
	LocalState       State
	AckStateGoing    State
	AckStateComing   State
	AssociatedValues []AssociatedValueType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessageObjectPushType) GetLengthSpec() uint8 {
	return m.LengthSpec
}

func (m *_AlarmMessageObjectPushType) GetSyntaxId() SyntaxIdType {
	return m.SyntaxId
}

func (m *_AlarmMessageObjectPushType) GetNumberOfValues() uint8 {
	return m.NumberOfValues
}

func (m *_AlarmMessageObjectPushType) GetEventId() uint32 {
	return m.EventId
}

func (m *_AlarmMessageObjectPushType) GetEventState() State {
	return m.EventState
}

func (m *_AlarmMessageObjectPushType) GetLocalState() State {
	return m.LocalState
}

func (m *_AlarmMessageObjectPushType) GetAckStateGoing() State {
	return m.AckStateGoing
}

func (m *_AlarmMessageObjectPushType) GetAckStateComing() State {
	return m.AckStateComing
}

func (m *_AlarmMessageObjectPushType) GetAssociatedValues() []AssociatedValueType {
	return m.AssociatedValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AlarmMessageObjectPushType) GetVariableSpec() uint8 {
	return AlarmMessageObjectPushType_VARIABLESPEC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlarmMessageObjectPushType factory function for _AlarmMessageObjectPushType
func NewAlarmMessageObjectPushType(lengthSpec uint8, syntaxId SyntaxIdType, numberOfValues uint8, eventId uint32, eventState State, localState State, ackStateGoing State, ackStateComing State, AssociatedValues []AssociatedValueType) *_AlarmMessageObjectPushType {
	return &_AlarmMessageObjectPushType{LengthSpec: lengthSpec, SyntaxId: syntaxId, NumberOfValues: numberOfValues, EventId: eventId, EventState: eventState, LocalState: localState, AckStateGoing: ackStateGoing, AckStateComing: ackStateComing, AssociatedValues: AssociatedValues}
}

// Deprecated: use the interface for direct cast
func CastAlarmMessageObjectPushType(structType any) AlarmMessageObjectPushType {
	if casted, ok := structType.(AlarmMessageObjectPushType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessageObjectPushType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessageObjectPushType) GetTypeName() string {
	return "AlarmMessageObjectPushType"
}

func (m *_AlarmMessageObjectPushType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (variableSpec)
	lengthInBits += 8

	// Simple field (lengthSpec)
	lengthInBits += 8

	// Simple field (syntaxId)
	lengthInBits += 8

	// Simple field (numberOfValues)
	lengthInBits += 8

	// Simple field (eventId)
	lengthInBits += 32

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits(ctx)

	// Simple field (localState)
	lengthInBits += m.LocalState.GetLengthInBits(ctx)

	// Simple field (ackStateGoing)
	lengthInBits += m.AckStateGoing.GetLengthInBits(ctx)

	// Simple field (ackStateComing)
	lengthInBits += m.AckStateComing.GetLengthInBits(ctx)

	// Array field
	if len(m.AssociatedValues) > 0 {
		for _curItem, element := range m.AssociatedValues {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.AssociatedValues), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AlarmMessageObjectPushType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessageObjectPushTypeParse(ctx context.Context, theBytes []byte) (AlarmMessageObjectPushType, error) {
	return AlarmMessageObjectPushTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessageObjectPushTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageObjectPushType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AlarmMessageObjectPushType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageObjectPushType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (variableSpec)
	variableSpec, _variableSpecErr := readBuffer.ReadUint8("variableSpec", 8)
	if _variableSpecErr != nil {
		return nil, errors.Wrap(_variableSpecErr, "Error parsing 'variableSpec' field of AlarmMessageObjectPushType")
	}
	if variableSpec != AlarmMessageObjectPushType_VARIABLESPEC {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AlarmMessageObjectPushType_VARIABLESPEC) + " but got " + fmt.Sprintf("%d", variableSpec))
	}

	// Simple Field (lengthSpec)
	_lengthSpec, _lengthSpecErr := readBuffer.ReadUint8("lengthSpec", 8)
	if _lengthSpecErr != nil {
		return nil, errors.Wrap(_lengthSpecErr, "Error parsing 'lengthSpec' field of AlarmMessageObjectPushType")
	}
	lengthSpec := _lengthSpec

	// Simple Field (syntaxId)
	if pullErr := readBuffer.PullContext("syntaxId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for syntaxId")
	}
	_syntaxId, _syntaxIdErr := SyntaxIdTypeParseWithBuffer(ctx, readBuffer)
	if _syntaxIdErr != nil {
		return nil, errors.Wrap(_syntaxIdErr, "Error parsing 'syntaxId' field of AlarmMessageObjectPushType")
	}
	syntaxId := _syntaxId
	if closeErr := readBuffer.CloseContext("syntaxId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for syntaxId")
	}

	// Simple Field (numberOfValues)
	_numberOfValues, _numberOfValuesErr := readBuffer.ReadUint8("numberOfValues", 8)
	if _numberOfValuesErr != nil {
		return nil, errors.Wrap(_numberOfValuesErr, "Error parsing 'numberOfValues' field of AlarmMessageObjectPushType")
	}
	numberOfValues := _numberOfValues

	// Simple Field (eventId)
	_eventId, _eventIdErr := readBuffer.ReadUint32("eventId", 32)
	if _eventIdErr != nil {
		return nil, errors.Wrap(_eventIdErr, "Error parsing 'eventId' field of AlarmMessageObjectPushType")
	}
	eventId := _eventId

	// Simple Field (eventState)
	if pullErr := readBuffer.PullContext("eventState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventState")
	}
	_eventState, _eventStateErr := StateParseWithBuffer(ctx, readBuffer)
	if _eventStateErr != nil {
		return nil, errors.Wrap(_eventStateErr, "Error parsing 'eventState' field of AlarmMessageObjectPushType")
	}
	eventState := _eventState.(State)
	if closeErr := readBuffer.CloseContext("eventState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventState")
	}

	// Simple Field (localState)
	if pullErr := readBuffer.PullContext("localState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for localState")
	}
	_localState, _localStateErr := StateParseWithBuffer(ctx, readBuffer)
	if _localStateErr != nil {
		return nil, errors.Wrap(_localStateErr, "Error parsing 'localState' field of AlarmMessageObjectPushType")
	}
	localState := _localState.(State)
	if closeErr := readBuffer.CloseContext("localState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for localState")
	}

	// Simple Field (ackStateGoing)
	if pullErr := readBuffer.PullContext("ackStateGoing"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ackStateGoing")
	}
	_ackStateGoing, _ackStateGoingErr := StateParseWithBuffer(ctx, readBuffer)
	if _ackStateGoingErr != nil {
		return nil, errors.Wrap(_ackStateGoingErr, "Error parsing 'ackStateGoing' field of AlarmMessageObjectPushType")
	}
	ackStateGoing := _ackStateGoing.(State)
	if closeErr := readBuffer.CloseContext("ackStateGoing"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ackStateGoing")
	}

	// Simple Field (ackStateComing)
	if pullErr := readBuffer.PullContext("ackStateComing"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ackStateComing")
	}
	_ackStateComing, _ackStateComingErr := StateParseWithBuffer(ctx, readBuffer)
	if _ackStateComingErr != nil {
		return nil, errors.Wrap(_ackStateComingErr, "Error parsing 'ackStateComing' field of AlarmMessageObjectPushType")
	}
	ackStateComing := _ackStateComing.(State)
	if closeErr := readBuffer.CloseContext("ackStateComing"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ackStateComing")
	}

	// Array field (AssociatedValues)
	if pullErr := readBuffer.PullContext("AssociatedValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AssociatedValues")
	}
	// Count array
	AssociatedValues := make([]AssociatedValueType, utils.Max(numberOfValues, 0))
	// This happens when the size is set conditional to 0
	if len(AssociatedValues) == 0 {
		AssociatedValues = nil
	}
	{
		_numItems := uint16(utils.Max(numberOfValues, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := AssociatedValueTypeParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'AssociatedValues' field of AlarmMessageObjectPushType")
			}
			AssociatedValues[_curItem] = _item.(AssociatedValueType)
		}
	}
	if closeErr := readBuffer.CloseContext("AssociatedValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AssociatedValues")
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageObjectPushType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageObjectPushType")
	}

	// Create the instance
	return &_AlarmMessageObjectPushType{
		LengthSpec:       lengthSpec,
		SyntaxId:         syntaxId,
		NumberOfValues:   numberOfValues,
		EventId:          eventId,
		EventState:       eventState,
		LocalState:       localState,
		AckStateGoing:    ackStateGoing,
		AckStateComing:   ackStateComing,
		AssociatedValues: AssociatedValues,
	}, nil
}

func (m *_AlarmMessageObjectPushType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessageObjectPushType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AlarmMessageObjectPushType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageObjectPushType")
	}

	// Const Field (variableSpec)
	_variableSpecErr := writeBuffer.WriteUint8("variableSpec", 8, uint8(0x12))
	if _variableSpecErr != nil {
		return errors.Wrap(_variableSpecErr, "Error serializing 'variableSpec' field")
	}

	// Simple Field (lengthSpec)
	lengthSpec := uint8(m.GetLengthSpec())
	_lengthSpecErr := writeBuffer.WriteUint8("lengthSpec", 8, uint8((lengthSpec)))
	if _lengthSpecErr != nil {
		return errors.Wrap(_lengthSpecErr, "Error serializing 'lengthSpec' field")
	}

	// Simple Field (syntaxId)
	if pushErr := writeBuffer.PushContext("syntaxId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for syntaxId")
	}
	_syntaxIdErr := writeBuffer.WriteSerializable(ctx, m.GetSyntaxId())
	if popErr := writeBuffer.PopContext("syntaxId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for syntaxId")
	}
	if _syntaxIdErr != nil {
		return errors.Wrap(_syntaxIdErr, "Error serializing 'syntaxId' field")
	}

	// Simple Field (numberOfValues)
	numberOfValues := uint8(m.GetNumberOfValues())
	_numberOfValuesErr := writeBuffer.WriteUint8("numberOfValues", 8, uint8((numberOfValues)))
	if _numberOfValuesErr != nil {
		return errors.Wrap(_numberOfValuesErr, "Error serializing 'numberOfValues' field")
	}

	// Simple Field (eventId)
	eventId := uint32(m.GetEventId())
	_eventIdErr := writeBuffer.WriteUint32("eventId", 32, uint32((eventId)))
	if _eventIdErr != nil {
		return errors.Wrap(_eventIdErr, "Error serializing 'eventId' field")
	}

	// Simple Field (eventState)
	if pushErr := writeBuffer.PushContext("eventState"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for eventState")
	}
	_eventStateErr := writeBuffer.WriteSerializable(ctx, m.GetEventState())
	if popErr := writeBuffer.PopContext("eventState"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for eventState")
	}
	if _eventStateErr != nil {
		return errors.Wrap(_eventStateErr, "Error serializing 'eventState' field")
	}

	// Simple Field (localState)
	if pushErr := writeBuffer.PushContext("localState"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for localState")
	}
	_localStateErr := writeBuffer.WriteSerializable(ctx, m.GetLocalState())
	if popErr := writeBuffer.PopContext("localState"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for localState")
	}
	if _localStateErr != nil {
		return errors.Wrap(_localStateErr, "Error serializing 'localState' field")
	}

	// Simple Field (ackStateGoing)
	if pushErr := writeBuffer.PushContext("ackStateGoing"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ackStateGoing")
	}
	_ackStateGoingErr := writeBuffer.WriteSerializable(ctx, m.GetAckStateGoing())
	if popErr := writeBuffer.PopContext("ackStateGoing"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ackStateGoing")
	}
	if _ackStateGoingErr != nil {
		return errors.Wrap(_ackStateGoingErr, "Error serializing 'ackStateGoing' field")
	}

	// Simple Field (ackStateComing)
	if pushErr := writeBuffer.PushContext("ackStateComing"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ackStateComing")
	}
	_ackStateComingErr := writeBuffer.WriteSerializable(ctx, m.GetAckStateComing())
	if popErr := writeBuffer.PopContext("ackStateComing"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ackStateComing")
	}
	if _ackStateComingErr != nil {
		return errors.Wrap(_ackStateComingErr, "Error serializing 'ackStateComing' field")
	}

	// Array Field (AssociatedValues)
	if pushErr := writeBuffer.PushContext("AssociatedValues", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AssociatedValues")
	}
	for _curItem, _element := range m.GetAssociatedValues() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetAssociatedValues()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'AssociatedValues' field")
		}
	}
	if popErr := writeBuffer.PopContext("AssociatedValues", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AssociatedValues")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageObjectPushType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageObjectPushType")
	}
	return nil
}

func (m *_AlarmMessageObjectPushType) isAlarmMessageObjectPushType() bool {
	return true
}

func (m *_AlarmMessageObjectPushType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
