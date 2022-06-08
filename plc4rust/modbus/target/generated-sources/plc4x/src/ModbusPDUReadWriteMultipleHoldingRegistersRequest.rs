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
// package org.apache.plc4x.rust.modbus.readwrite;

// Code generated by code-generation. DO NOT EDIT.
use std::io::{Error, ErrorKind, Read, Write};
use heapless::Vec;
use plc4rust::{Message, NoOption};
use plc4rust::read_buffer::ReadBuffer;
use plc4rust::write_buffer::WriteBuffer;


#[derive(PartialEq, Debug, Clone)]
pub struct ModbusPDUReadWriteMultipleHoldingRegistersRequestOptions {
    pub response: bool
}

#[derive(PartialEq, Debug, Clone)]
pub struct ModbusPDUReadWriteMultipleHoldingRegistersRequest {
    pub readStartingAddress: u16,
    pub readQuantity: u16,
    pub writeStartingAddress: u16,
    pub writeQuantity: u16,
    pub value: Vec<u8, 1024>
}

impl ModbusPDUReadWriteMultipleHoldingRegistersRequest {
    pub fn byteCount(&self) -> u8 {
        self.value.len() as u8
    }
}

impl Message for ModbusPDUReadWriteMultipleHoldingRegistersRequest {
    type M = ModbusPDUReadWriteMultipleHoldingRegistersRequest;
    type P = ModbusPDUReadWriteMultipleHoldingRegistersRequestOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        writer.write_u16(self.readStartingAddress)?;
        writer.write_u16(self.readQuantity)?;
        writer.write_u16(self.writeStartingAddress)?;
        writer.write_u16(self.writeQuantity)?;
        writer.write_u8(self.byteCount())?;
        writer.write_bytes(self.value.as_slice())?;
        Ok(0)
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let response = parameter.response;
        let readStartingAddress = reader.read_u16()?;
        let readQuantity = reader.read_u16()?;
        let writeStartingAddress = reader.read_u16()?;
        let writeQuantity = reader.read_u16()?;
        let byteCount = reader.read_u8()?;
        let value: Vec<_, 1024>  = Vec::new();
        let value_read = 0 as usize;
        // for _ in 0..(DefaultVariableLiteral{name='byteCount', typeReference='null', args=null, index=null, child=null}) {
            // do something
        // }
        Ok(Self::M {
            readStartingAddress,
            readQuantity,
            writeStartingAddress,
            writeQuantity,
            value
        })
    }
}


