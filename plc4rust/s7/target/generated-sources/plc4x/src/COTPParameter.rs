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

use crate::COTPParameterTpduSize::COTPParameterTpduSize;
use crate::COTPParameterTpduSize::COTPParameterTpduSizeOptions;
use crate::COTPParameterCallingTsap::COTPParameterCallingTsap;
use crate::COTPParameterCallingTsap::COTPParameterCallingTsapOptions;
use crate::COTPParameterCalledTsap::COTPParameterCalledTsap;
use crate::COTPParameterCalledTsap::COTPParameterCalledTsapOptions;
use crate::COTPParameterChecksum::COTPParameterChecksum;
use crate::COTPParameterChecksum::COTPParameterChecksumOptions;
use crate::COTPParameterDisconnectAdditionalInformation::COTPParameterDisconnectAdditionalInformation;
use crate::COTPParameterDisconnectAdditionalInformation::COTPParameterDisconnectAdditionalInformationOptions;

#[derive(PartialEq, Debug, Clone)]
pub struct COTPParameterOptions {
    pub rest: u8
}

#[derive(PartialEq, Debug, Clone)]
pub enum COTPParameter {
    COTPParameterTpduSize(COTPParameterTpduSize),
    COTPParameterCallingTsap(COTPParameterCallingTsap),
    COTPParameterCalledTsap(COTPParameterCalledTsap),
    COTPParameterChecksum(COTPParameterChecksum),
    COTPParameterDisconnectAdditionalInformation(COTPParameterDisconnectAdditionalInformation)
}

impl Message for COTPParameter {
    type M = COTPParameter;
    type P = COTPParameterOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        match self {
            COTPParameter::COTPParameterTpduSize(msg) => {
                msg.serialize(writer)
            }
            COTPParameter::COTPParameterCallingTsap(msg) => {
                msg.serialize(writer)
            }
            COTPParameter::COTPParameterCalledTsap(msg) => {
                msg.serialize(writer)
            }
            COTPParameter::COTPParameterChecksum(msg) => {
                msg.serialize(writer)
            }
            COTPParameter::COTPParameterDisconnectAdditionalInformation(msg) => {
                msg.serialize(writer)
            }
        }
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let rest = parameter.rest;
        let parameterType = reader.read_u8()?;
            // -> DefaultImplicitField{serializeExpression=DefaultBinaryTerm{a=DefaultVariableLiteral{name='lengthInBytes', typeReference='null', args=null, index=null, child=null}, b=DefaultNumericLiteral{number=2}, operation='-'}} DefaultTypedNamedField{name='parameterLength'} DefaultTypedField{type=AbstractSimpleTypeReference{baseType=UINT, sizeInBits=8}} DefaultField{attributes={}}
        match (parameterType) {
            (0xC0) => {
                Ok(COTPParameter::COTPParameterTpduSize(COTPParameterTpduSize::parse::<T>(reader, Some(COTPParameterTpduSizeOptions {
                    rest
                }))?))
            }
            (0xC1) => {
                Ok(COTPParameter::COTPParameterCallingTsap(COTPParameterCallingTsap::parse::<T>(reader, Some(COTPParameterCallingTsapOptions {
                    rest
                }))?))
            }
            (0xC2) => {
                Ok(COTPParameter::COTPParameterCalledTsap(COTPParameterCalledTsap::parse::<T>(reader, Some(COTPParameterCalledTsapOptions {
                    rest
                }))?))
            }
            (0xC3) => {
                Ok(COTPParameter::COTPParameterChecksum(COTPParameterChecksum::parse::<T>(reader, Some(COTPParameterChecksumOptions {
                    rest
                }))?))
            }
            (0xE0) => {
                Ok(COTPParameter::COTPParameterDisconnectAdditionalInformation(COTPParameterDisconnectAdditionalInformation::parse::<T>(reader, Some(COTPParameterDisconnectAdditionalInformationOptions {
                    rest
                }))?))
            }
            _ => {
                panic!("Unable to parse!");
            }
        }
    }
}


