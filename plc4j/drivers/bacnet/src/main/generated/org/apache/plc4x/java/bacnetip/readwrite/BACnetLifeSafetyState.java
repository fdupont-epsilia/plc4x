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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum BACnetLifeSafetyState {
  QUIET((int) 0),
  PRE_ALARM((int) 1),
  ALARM((int) 2),
  FAULT((int) 3),
  FAULT_PRE_ALARM((int) 4),
  FAULT_ALARM((int) 5),
  NOT_READY((int) 6),
  ACTIVE((int) 7),
  TAMPER((int) 8),
  TEST_ALARM((int) 9),
  TEST_ACTIVE((int) 10),
  TEST_FAULT((int) 11),
  TEST_FAULT_ALARM((int) 12),
  HOLDUP((int) 13),
  DURESS((int) 14),
  TAMPER_ALARM((int) 15),
  ABNORMAL((int) 16),
  EMERGENCY_POWER((int) 17),
  DELAYED((int) 18),
  BLOCKED((int) 19),
  LOCAL_ALARM((int) 20),
  GENERAL_ALARM((int) 21),
  SUPERVISORY((int) 22),
  TEST_SUPERVISORY((int) 23),
  VENDOR_PROPRIETARY_VALUE((int) 0XFFFF);
  private static final Map<Integer, BACnetLifeSafetyState> map;

  static {
    map = new HashMap<>();
    for (BACnetLifeSafetyState value : BACnetLifeSafetyState.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private int value;

  BACnetLifeSafetyState(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static BACnetLifeSafetyState enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}