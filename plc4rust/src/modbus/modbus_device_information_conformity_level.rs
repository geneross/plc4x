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

use std::io::{Error, ErrorKind, Read, Write};
use crate::{Message, NoOption, plc4x_enum, ReadBuffer, WriteBuffer};

// [enum uint 7 ModbusDeviceInformationConformityLevel
// ['0x01' BASIC_STREAM_ONLY   ]
// ['0x02' REGULAR_STREAM_ONLY ]
// ['0x03' EXTENDED_STREAM_ONLY]
// ]
plc4x_enum!
[enum 7 => u8 : ModbusDeviceInformationConformityLevel
    [0x01 => BASIC_STREAM_ONLY],
    [0x02 => REGULAR_STREAM_ONLY ],
    [0x03 => EXTENDED_STREAM_ONLY]
];