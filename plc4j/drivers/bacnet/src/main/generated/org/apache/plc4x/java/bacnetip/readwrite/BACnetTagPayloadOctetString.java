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
package org.apache.plc4x.java.bacnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class BACnetTagPayloadOctetString implements Message {

  // Properties.
  protected final byte[] octets;

  // Arguments.
  protected final Long actualLength;

  public BACnetTagPayloadOctetString(byte[] octets, Long actualLength) {
    super();
    this.octets = octets;
    this.actualLength = actualLength;
  }

  public byte[] getOctets() {
    return octets;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetTagPayloadOctetString");

    // Array Field (octets)
    writeByteArrayField("octets", octets, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("BACnetTagPayloadOctetString");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetTagPayloadOctetString _value = this;

    // Array field
    if (octets != null) {
      lengthInBits += 8 * octets.length;
    }

    return lengthInBits;
  }

  public static BACnetTagPayloadOctetString staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Long actualLength;
    if (args[0] instanceof Long) {
      actualLength = (Long) args[0];
    } else if (args[0] instanceof String) {
      actualLength = Long.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Long or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, actualLength);
  }

  public static BACnetTagPayloadOctetString staticParse(ReadBuffer readBuffer, Long actualLength)
      throws ParseException {
    readBuffer.pullContext("BACnetTagPayloadOctetString");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    byte[] octets = readBuffer.readByteArray("octets", Math.toIntExact(actualLength));

    readBuffer.closeContext("BACnetTagPayloadOctetString");
    // Create the instance
    BACnetTagPayloadOctetString _bACnetTagPayloadOctetString;
    _bACnetTagPayloadOctetString = new BACnetTagPayloadOctetString(octets, actualLength);
    return _bACnetTagPayloadOctetString;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetTagPayloadOctetString)) {
      return false;
    }
    BACnetTagPayloadOctetString that = (BACnetTagPayloadOctetString) o;
    return (getOctets() == that.getOctets()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getOctets());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}