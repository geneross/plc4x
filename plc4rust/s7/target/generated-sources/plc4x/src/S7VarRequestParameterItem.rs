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

use crate::S7VarRequestParameterItemAddress::S7VarRequestParameterItemAddress;
use crate::S7VarRequestParameterItemAddress::S7VarRequestParameterItemAddressOptions;

#[derive(PartialEq, Debug, Clone)]
pub struct S7VarRequestParameterItemOptions {
}

#[derive(PartialEq, Debug, Clone)]
pub enum S7VarRequestParameterItem {
    S7VarRequestParameterItemAddress(S7VarRequestParameterItemAddress)
}

impl Message for S7VarRequestParameterItem {
    type M = S7VarRequestParameterItem;
    type P = NoOption;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        match self {
            S7VarRequestParameterItem::S7VarRequestParameterItemAddress(msg) => {
                msg.serialize(writer)
            }
        }
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let itemType = reader.read_u8()?;
        match (itemType) {
            (0x12) => {
                Ok(S7VarRequestParameterItem::S7VarRequestParameterItemAddress(S7VarRequestParameterItemAddress::parse::<T>(reader, Some(S7VarRequestParameterItemAddressOptions {
                }))?))
            }
            _ => {
                panic!("Unable to parse!");
            }
        }
    }
}


