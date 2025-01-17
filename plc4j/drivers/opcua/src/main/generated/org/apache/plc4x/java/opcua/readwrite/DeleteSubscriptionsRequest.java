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

public class DeleteSubscriptionsRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 847;
  }

  // Properties.
  protected final RequestHeader requestHeader;
  protected final List<Long> subscriptionIds;

  public DeleteSubscriptionsRequest(RequestHeader requestHeader, List<Long> subscriptionIds) {
    super();
    this.requestHeader = requestHeader;
    this.subscriptionIds = subscriptionIds;
  }

  public RequestHeader getRequestHeader() {
    return requestHeader;
  }

  public List<Long> getSubscriptionIds() {
    return subscriptionIds;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("DeleteSubscriptionsRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, writeComplex(writeBuffer));

    // Implicit Field (noOfSubscriptionIds) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfSubscriptionIds =
        (int) ((((getSubscriptionIds()) == (null)) ? -(1) : COUNT(getSubscriptionIds())));
    writeImplicitField("noOfSubscriptionIds", noOfSubscriptionIds, writeSignedInt(writeBuffer, 32));

    // Array Field (subscriptionIds)
    writeSimpleTypeArrayField(
        "subscriptionIds", subscriptionIds, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("DeleteSubscriptionsRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    DeleteSubscriptionsRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Implicit Field (noOfSubscriptionIds)
    lengthInBits += 32;

    // Array field
    if (subscriptionIds != null) {
      lengthInBits += 32 * subscriptionIds.size();
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("DeleteSubscriptionsRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    RequestHeader requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () ->
                    (RequestHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (391)),
                readBuffer));

    int noOfSubscriptionIds =
        readImplicitField("noOfSubscriptionIds", readSignedInt(readBuffer, 32));

    List<Long> subscriptionIds =
        readCountArrayField(
            "subscriptionIds", readUnsignedLong(readBuffer, 32), noOfSubscriptionIds);

    readBuffer.closeContext("DeleteSubscriptionsRequest");
    // Create the instance
    return new DeleteSubscriptionsRequestBuilderImpl(requestHeader, subscriptionIds);
  }

  public static class DeleteSubscriptionsRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final RequestHeader requestHeader;
    private final List<Long> subscriptionIds;

    public DeleteSubscriptionsRequestBuilderImpl(
        RequestHeader requestHeader, List<Long> subscriptionIds) {
      this.requestHeader = requestHeader;
      this.subscriptionIds = subscriptionIds;
    }

    public DeleteSubscriptionsRequest build() {
      DeleteSubscriptionsRequest deleteSubscriptionsRequest =
          new DeleteSubscriptionsRequest(requestHeader, subscriptionIds);
      return deleteSubscriptionsRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DeleteSubscriptionsRequest)) {
      return false;
    }
    DeleteSubscriptionsRequest that = (DeleteSubscriptionsRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getSubscriptionIds() == that.getSubscriptionIds())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRequestHeader(), getSubscriptionIds());
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
