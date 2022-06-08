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
// package org.apache.plc4x.rust.s7.readwrite;

// Code generated by code-generation. DO NOT EDIT.
use std::io::{Error, ErrorKind, Read, Write};
use plc4rust::{Message, NoOption};
use plc4rust::read_buffer::ReadBuffer;
use plc4rust::write_buffer::WriteBuffer;

use crate::DataTransportErrorCode::DataTransportErrorCode;
use crate::DataTransportSize::DataTransportSize;

#[derive(PartialEq, Debug, Clone)]
pub struct S7MessageObjectResponseOptions {
    pub cpuFunctionType: u8
}

#[derive(PartialEq, Debug, Clone)]
pub struct S7MessageObjectResponse {
    pub returnCode: DataTransportErrorCode,
    pub transportSize: DataTransportSize,
        // -> DefaultReservedField{referenceValue=0x00} DefaultTypedField{type=AbstractSimpleTypeReference{baseType=UINT, sizeInBits=8}} DefaultField{attributes={}}
}

impl S7MessageObjectResponse {
}

impl Message for S7MessageObjectResponse {
    type M = S7MessageObjectResponse;
    type P = S7MessageObjectResponseOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        self.returnCode.serialize(writer)?;
        self.transportSize.serialize(writer)?;
        ---> DefaultReservedField{referenceValue=0x00} DefaultTypedField{type=AbstractSimpleTypeReference{baseType=UINT, sizeInBits=8}} DefaultField{attributes={}}
        Ok(0)
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let cpuFunctionType = parameter.cpuFunctionType;
        let returnCode = DataTransportErrorCode::parse(reader, None)?;
        let transportSize = DataTransportSize::parse(reader, None)?;
        // Reserved field
        let _ = reader.read_u8()?;
        Ok(Self::M {
            returnCode,
            transportSize,
        })
    }
}


