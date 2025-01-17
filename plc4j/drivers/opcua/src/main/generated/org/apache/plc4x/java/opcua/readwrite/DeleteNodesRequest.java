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
package org.apache.plc4x.java.opcua.readwrite;

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

public class DeleteNodesRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 500;
  }

  // Properties.
  protected final RequestHeader requestHeader;
  protected final List<DeleteNodesItem> nodesToDelete;

  public DeleteNodesRequest(RequestHeader requestHeader, List<DeleteNodesItem> nodesToDelete) {
    super();
    this.requestHeader = requestHeader;
    this.nodesToDelete = nodesToDelete;
  }

  public RequestHeader getRequestHeader() {
    return requestHeader;
  }

  public List<DeleteNodesItem> getNodesToDelete() {
    return nodesToDelete;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("DeleteNodesRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, writeComplex(writeBuffer));

    // Implicit Field (noOfNodesToDelete) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfNodesToDelete =
        (int) ((((getNodesToDelete()) == (null)) ? -(1) : COUNT(getNodesToDelete())));
    writeImplicitField("noOfNodesToDelete", noOfNodesToDelete, writeSignedInt(writeBuffer, 32));

    // Array Field (nodesToDelete)
    writeComplexTypeArrayField("nodesToDelete", nodesToDelete, writeBuffer);

    writeBuffer.popContext("DeleteNodesRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    DeleteNodesRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Implicit Field (noOfNodesToDelete)
    lengthInBits += 32;

    // Array field
    if (nodesToDelete != null) {
      int i = 0;
      for (DeleteNodesItem element : nodesToDelete) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= nodesToDelete.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("DeleteNodesRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    RequestHeader requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () ->
                    (RequestHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (391)),
                readBuffer));

    int noOfNodesToDelete = readImplicitField("noOfNodesToDelete", readSignedInt(readBuffer, 32));

    List<DeleteNodesItem> nodesToDelete =
        readCountArrayField(
            "nodesToDelete",
            readComplex(
                () ->
                    (DeleteNodesItem)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (384)),
                readBuffer),
            noOfNodesToDelete);

    readBuffer.closeContext("DeleteNodesRequest");
    // Create the instance
    return new DeleteNodesRequestBuilderImpl(requestHeader, nodesToDelete);
  }

  public static class DeleteNodesRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final RequestHeader requestHeader;
    private final List<DeleteNodesItem> nodesToDelete;

    public DeleteNodesRequestBuilderImpl(
        RequestHeader requestHeader, List<DeleteNodesItem> nodesToDelete) {
      this.requestHeader = requestHeader;
      this.nodesToDelete = nodesToDelete;
    }

    public DeleteNodesRequest build() {
      DeleteNodesRequest deleteNodesRequest = new DeleteNodesRequest(requestHeader, nodesToDelete);
      return deleteNodesRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DeleteNodesRequest)) {
      return false;
    }
    DeleteNodesRequest that = (DeleteNodesRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getNodesToDelete() == that.getNodesToDelete())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRequestHeader(), getNodesToDelete());
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
