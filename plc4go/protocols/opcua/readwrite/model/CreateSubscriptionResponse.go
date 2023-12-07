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

// CreateSubscriptionResponse is the corresponding interface of CreateSubscriptionResponse
type CreateSubscriptionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetRevisedPublishingInterval returns RevisedPublishingInterval (property field)
	GetRevisedPublishingInterval() float64
	// GetRevisedLifetimeCount returns RevisedLifetimeCount (property field)
	GetRevisedLifetimeCount() uint32
	// GetRevisedMaxKeepAliveCount returns RevisedMaxKeepAliveCount (property field)
	GetRevisedMaxKeepAliveCount() uint32
}

// CreateSubscriptionResponseExactly can be used when we want exactly this type and not a type which fulfills CreateSubscriptionResponse.
// This is useful for switch cases.
type CreateSubscriptionResponseExactly interface {
	CreateSubscriptionResponse
	isCreateSubscriptionResponse() bool
}

// _CreateSubscriptionResponse is the data-structure of this message
type _CreateSubscriptionResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader            ExtensionObjectDefinition
	SubscriptionId            uint32
	RevisedPublishingInterval float64
	RevisedLifetimeCount      uint32
	RevisedMaxKeepAliveCount  uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateSubscriptionResponse) GetIdentifier() string {
	return "790"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateSubscriptionResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_CreateSubscriptionResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateSubscriptionResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_CreateSubscriptionResponse) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_CreateSubscriptionResponse) GetRevisedPublishingInterval() float64 {
	return m.RevisedPublishingInterval
}

func (m *_CreateSubscriptionResponse) GetRevisedLifetimeCount() uint32 {
	return m.RevisedLifetimeCount
}

func (m *_CreateSubscriptionResponse) GetRevisedMaxKeepAliveCount() uint32 {
	return m.RevisedMaxKeepAliveCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCreateSubscriptionResponse factory function for _CreateSubscriptionResponse
func NewCreateSubscriptionResponse(responseHeader ExtensionObjectDefinition, subscriptionId uint32, revisedPublishingInterval float64, revisedLifetimeCount uint32, revisedMaxKeepAliveCount uint32) *_CreateSubscriptionResponse {
	_result := &_CreateSubscriptionResponse{
		ResponseHeader:             responseHeader,
		SubscriptionId:             subscriptionId,
		RevisedPublishingInterval:  revisedPublishingInterval,
		RevisedLifetimeCount:       revisedLifetimeCount,
		RevisedMaxKeepAliveCount:   revisedMaxKeepAliveCount,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCreateSubscriptionResponse(structType any) CreateSubscriptionResponse {
	if casted, ok := structType.(CreateSubscriptionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CreateSubscriptionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CreateSubscriptionResponse) GetTypeName() string {
	return "CreateSubscriptionResponse"
}

func (m *_CreateSubscriptionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (revisedPublishingInterval)
	lengthInBits += 64

	// Simple field (revisedLifetimeCount)
	lengthInBits += 32

	// Simple field (revisedMaxKeepAliveCount)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CreateSubscriptionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CreateSubscriptionResponseParse(ctx context.Context, theBytes []byte, identifier string) (CreateSubscriptionResponse, error) {
	return CreateSubscriptionResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func CreateSubscriptionResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (CreateSubscriptionResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CreateSubscriptionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateSubscriptionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of CreateSubscriptionResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (subscriptionId)
	_subscriptionId, _subscriptionIdErr := readBuffer.ReadUint32("subscriptionId", 32)
	if _subscriptionIdErr != nil {
		return nil, errors.Wrap(_subscriptionIdErr, "Error parsing 'subscriptionId' field of CreateSubscriptionResponse")
	}
	subscriptionId := _subscriptionId

	// Simple Field (revisedPublishingInterval)
	_revisedPublishingInterval, _revisedPublishingIntervalErr := readBuffer.ReadFloat64("revisedPublishingInterval", 64)
	if _revisedPublishingIntervalErr != nil {
		return nil, errors.Wrap(_revisedPublishingIntervalErr, "Error parsing 'revisedPublishingInterval' field of CreateSubscriptionResponse")
	}
	revisedPublishingInterval := _revisedPublishingInterval

	// Simple Field (revisedLifetimeCount)
	_revisedLifetimeCount, _revisedLifetimeCountErr := readBuffer.ReadUint32("revisedLifetimeCount", 32)
	if _revisedLifetimeCountErr != nil {
		return nil, errors.Wrap(_revisedLifetimeCountErr, "Error parsing 'revisedLifetimeCount' field of CreateSubscriptionResponse")
	}
	revisedLifetimeCount := _revisedLifetimeCount

	// Simple Field (revisedMaxKeepAliveCount)
	_revisedMaxKeepAliveCount, _revisedMaxKeepAliveCountErr := readBuffer.ReadUint32("revisedMaxKeepAliveCount", 32)
	if _revisedMaxKeepAliveCountErr != nil {
		return nil, errors.Wrap(_revisedMaxKeepAliveCountErr, "Error parsing 'revisedMaxKeepAliveCount' field of CreateSubscriptionResponse")
	}
	revisedMaxKeepAliveCount := _revisedMaxKeepAliveCount

	if closeErr := readBuffer.CloseContext("CreateSubscriptionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateSubscriptionResponse")
	}

	// Create a partially initialized instance
	_child := &_CreateSubscriptionResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
		SubscriptionId:             subscriptionId,
		RevisedPublishingInterval:  revisedPublishingInterval,
		RevisedLifetimeCount:       revisedLifetimeCount,
		RevisedMaxKeepAliveCount:   revisedMaxKeepAliveCount,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_CreateSubscriptionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateSubscriptionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateSubscriptionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateSubscriptionResponse")
		}

		// Simple Field (responseHeader)
		if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseHeader")
		}
		_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
		if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseHeader")
		}
		if _responseHeaderErr != nil {
			return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
		}

		// Simple Field (subscriptionId)
		subscriptionId := uint32(m.GetSubscriptionId())
		_subscriptionIdErr := writeBuffer.WriteUint32("subscriptionId", 32, uint32((subscriptionId)))
		if _subscriptionIdErr != nil {
			return errors.Wrap(_subscriptionIdErr, "Error serializing 'subscriptionId' field")
		}

		// Simple Field (revisedPublishingInterval)
		revisedPublishingInterval := float64(m.GetRevisedPublishingInterval())
		_revisedPublishingIntervalErr := writeBuffer.WriteFloat64("revisedPublishingInterval", 64, (revisedPublishingInterval))
		if _revisedPublishingIntervalErr != nil {
			return errors.Wrap(_revisedPublishingIntervalErr, "Error serializing 'revisedPublishingInterval' field")
		}

		// Simple Field (revisedLifetimeCount)
		revisedLifetimeCount := uint32(m.GetRevisedLifetimeCount())
		_revisedLifetimeCountErr := writeBuffer.WriteUint32("revisedLifetimeCount", 32, uint32((revisedLifetimeCount)))
		if _revisedLifetimeCountErr != nil {
			return errors.Wrap(_revisedLifetimeCountErr, "Error serializing 'revisedLifetimeCount' field")
		}

		// Simple Field (revisedMaxKeepAliveCount)
		revisedMaxKeepAliveCount := uint32(m.GetRevisedMaxKeepAliveCount())
		_revisedMaxKeepAliveCountErr := writeBuffer.WriteUint32("revisedMaxKeepAliveCount", 32, uint32((revisedMaxKeepAliveCount)))
		if _revisedMaxKeepAliveCountErr != nil {
			return errors.Wrap(_revisedMaxKeepAliveCountErr, "Error serializing 'revisedMaxKeepAliveCount' field")
		}

		if popErr := writeBuffer.PopContext("CreateSubscriptionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateSubscriptionResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateSubscriptionResponse) isCreateSubscriptionResponse() bool {
	return true
}

func (m *_CreateSubscriptionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
