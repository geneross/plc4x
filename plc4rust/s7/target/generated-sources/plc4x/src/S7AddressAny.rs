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

use crate::TransportSize::TransportSize;
use crate::MemoryArea::MemoryArea;

#[derive(PartialEq, Debug, Clone)]
pub struct S7AddressAnyOptions {
}

#[derive(PartialEq, Debug, Clone)]
pub struct S7AddressAny {
        // -> DefaultEnumField{fieldName='code'} DefaultTypedNamedField{name='transportSize'} DefaultTypedField{type=DefaultEnumTypeReference{name='TransportSize', params=null}} DefaultField{attributes={}}
    pub numberOfElements: u16,
    pub dbNumber: u16,
    pub area: MemoryArea,
        // -> DefaultReservedField{referenceValue=0x00} DefaultTypedField{type=AbstractSimpleTypeReference{baseType=UINT, sizeInBits=5}} DefaultField{attributes={}}
    pub byteAddress: u16,
    pub bitAddress: u8
}

impl S7AddressAny {
}

impl Message for S7AddressAny {
    type M = S7AddressAny;
    type P = S7AddressAnyOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        self.transportSize.serialize(writer)?;
        writer.write_u16(self.numberOfElements)?;
        writer.write_u16(self.dbNumber)?;
        self.area.serialize(writer)?;
        ---> DefaultReservedField{referenceValue=0x00} DefaultTypedField{type=AbstractSimpleTypeReference{baseType=UINT, sizeInBits=5}} DefaultField{attributes={}}
        writer.write_u16(self.byteAddress)?;
        writer.write_u_n(3, self.self.bitAddress as u64)? as u8;
        Ok(0)
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let transportSize = TransportSize::parse(reader, None)?;
        let numberOfElements = reader.read_u16()?;
        let dbNumber = reader.read_u16()?;
        let area = MemoryArea::parse(reader, None)?;
        // Reserved field
        let _ = reader.read_u_n(5)? as u8;
        let byteAddress = reader.read_u16()?;
        let bitAddress = reader.read_u_n(3)? as u8;
        Ok(Self::M {
            transportSize,
            numberOfElements,
            dbNumber,
            area,
            byteAddress,
            bitAddress
        })
    }
}


