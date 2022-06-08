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

use crate::ModbusTcpADU::ModbusTcpADU;
use crate::ModbusTcpADU::ModbusTcpADUOptions;
use crate::ModbusRtuADU::ModbusRtuADU;
use crate::ModbusRtuADU::ModbusRtuADUOptions;
use crate::ModbusAsciiADU::ModbusAsciiADU;
use crate::ModbusAsciiADU::ModbusAsciiADUOptions;
use crate::DriverType::DriverType;

#[derive(PartialEq, Debug, Clone)]
pub struct ModbusADUOptions {
    pub driverType: DriverType, 
    pub response: bool
}

#[derive(PartialEq, Debug, Clone)]
pub enum ModbusADU {
    ModbusTcpADU(ModbusTcpADU),
    ModbusRtuADU(ModbusRtuADU),
    ModbusAsciiADU(ModbusAsciiADU)
}

impl Message for ModbusADU {
    type M = ModbusADU;
    type P = ModbusADUOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        match self {
            ModbusADU::ModbusTcpADU(msg) => {
                msg.serialize(writer)
            }
            ModbusADU::ModbusRtuADU(msg) => {
                msg.serialize(writer)
            }
            ModbusADU::ModbusAsciiADU(msg) => {
                msg.serialize(writer)
            }
        }
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let driverType = parameter.driverType;
        let response = parameter.response;
        match (driverType) {
            (DriverType::MODBUS_TCP) => {
                Ok(ModbusADU::ModbusTcpADU(ModbusTcpADU::parse::<T>(reader, Some(ModbusTcpADUOptions {
                    driverType, 
                    response
                }))?))
            }
            (DriverType::MODBUS_RTU) => {
                Ok(ModbusADU::ModbusRtuADU(ModbusRtuADU::parse::<T>(reader, Some(ModbusRtuADUOptions {
                    driverType, 
                    response
                }))?))
            }
            (DriverType::MODBUS_ASCII) => {
                Ok(ModbusADU::ModbusAsciiADU(ModbusAsciiADU::parse::<T>(reader, Some(ModbusAsciiADUOptions {
                    driverType, 
                    response
                }))?))
            }
            _ => {
                panic!("Unable to parse!");
            }
        }
    }
}


