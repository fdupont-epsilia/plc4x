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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpcuaNodeIdServicesVariableData {
  DataTypeDescriptionType_DataTypeVersion((int) 104L),
  DataTypeDescriptionType_DictionaryFragment((int) 105L),
  DataTypeDictionaryType_DataTypeVersion((int) 106L),
  DataTypeDictionaryType_NamespaceUri((int) 107L),
  DataItemType_Definition((int) 2366L),
  DataItemType_ValuePrecision((int) 2367L),
  DataChangeTrigger_EnumStrings((int) 7609L),
  DataSetFolderType_DataSetFolderName_Placeholder_AddPublishedDataItems_InputArguments(
      (int) 14480L),
  DataSetFolderType_DataSetFolderName_Placeholder_AddPublishedDataItems_OutputArguments(
      (int) 14481L),
  DataSetFolderType_DataSetFolderName_Placeholder_AddPublishedEvents_InputArguments((int) 14483L),
  DataSetFolderType_DataSetFolderName_Placeholder_AddPublishedEvents_OutputArguments((int) 14484L),
  DataSetFolderType_DataSetFolderName_Placeholder_RemovePublishedDataSet_InputArguments(
      (int) 14486L),
  DataSetFolderType_PublishedDataSetName_Placeholder_ConfigurationVersion((int) 14489L),
  DataSetFolderType_AddPublishedDataItems_InputArguments((int) 14494L),
  DataSetFolderType_AddPublishedDataItems_OutputArguments((int) 14495L),
  DataSetFolderType_AddPublishedEvents_InputArguments((int) 14497L),
  DataSetFolderType_AddPublishedEvents_OutputArguments((int) 14498L),
  DataSetFolderType_RemovePublishedDataSet_InputArguments((int) 14500L),
  DataTypeDictionaryType_Deprecated((int) 15001L),
  DataSetFolderType_PublishedDataSetName_Placeholder_DataSetMetaData((int) 15221L),
  DataSetWriterType_Status_State((int) 15300L),
  DataSetReaderType_Status_State((int) 15308L),
  DataSetFolderType_PublishedDataSetName_Placeholder_ExtensionFields_AddExtensionField_InputArguments(
      (int) 15475L),
  DataSetFolderType_PublishedDataSetName_Placeholder_ExtensionFields_AddExtensionField_OutputArguments(
      (int) 15476L),
  DataSetFolderType_PublishedDataSetName_Placeholder_ExtensionFields_RemoveExtensionField_InputArguments(
      (int) 15478L),
  DataSetFieldFlags_OptionSetValues((int) 15577L),
  DataSetFieldContentMask_OptionSetValues((int) 15584L),
  DataSetOrderingType_EnumStrings((int) 15641L),
  DataSetReaderType_SecurityMode((int) 15932L),
  DataSetReaderType_SecurityGroupId((int) 15933L),
  DataSetReaderType_SecurityKeyServices((int) 15934L),
  DataSetFolderType_DataSetFolderName_Placeholder_AddPublishedDataItemsTemplate_InputArguments(
      (int) 16843L),
  DataSetFolderType_DataSetFolderName_Placeholder_AddPublishedDataItemsTemplate_OutputArguments(
      (int) 16853L),
  DataSetFolderType_DataSetFolderName_Placeholder_AddPublishedEventsTemplate_InputArguments(
      (int) 16882L),
  DataSetFolderType_DataSetFolderName_Placeholder_AddPublishedEventsTemplate_OutputArguments(
      (int) 16883L),
  DataSetFolderType_DataSetFolderName_Placeholder_AddDataSetFolder_InputArguments((int) 16894L),
  DataSetFolderType_DataSetFolderName_Placeholder_AddDataSetFolder_OutputArguments((int) 16922L),
  DataSetFolderType_DataSetFolderName_Placeholder_RemoveDataSetFolder_InputArguments((int) 16924L),
  DataSetFolderType_PublishedDataSetName_Placeholder_DataSetClassId((int) 16925L),
  DataSetFolderType_AddPublishedDataItemsTemplate_InputArguments((int) 16958L),
  DataSetFolderType_AddPublishedDataItemsTemplate_OutputArguments((int) 16959L),
  DataSetFolderType_AddPublishedEventsTemplate_InputArguments((int) 16961L),
  DataSetFolderType_AddPublishedEventsTemplate_OutputArguments((int) 16971L),
  DataSetFolderType_AddDataSetFolder_InputArguments((int) 16995L),
  DataSetFolderType_AddDataSetFolder_OutputArguments((int) 16996L),
  DataSetFolderType_RemoveDataSetFolder_InputArguments((int) 17007L),
  DataSetReaderType_CreateTargetVariables_InputArguments((int) 17387L),
  DataSetReaderType_CreateTargetVariables_OutputArguments((int) 17388L),
  DataSetReaderType_CreateDataSetMirror_InputArguments((int) 17390L),
  DataSetReaderType_CreateDataSetMirror_OutputArguments((int) 17391L),
  DataSetReaderTypeCreateTargetVariablesMethodType_InputArguments((int) 17393L),
  DataSetReaderTypeCreateTargetVariablesMethodType_OutputArguments((int) 17394L),
  DataSetReaderTypeCreateDataSetMirrorMethodType_InputArguments((int) 17396L),
  DataSetReaderTypeCreateDataSetMirrorMethodType_OutputArguments((int) 17397L),
  DataSetWriterType_DataSetWriterProperties((int) 17493L),
  DataSetReaderType_DataSetReaderProperties((int) 17494L),
  DataSetReaderType_KeyFrameCount((int) 17563L),
  DataSetReaderType_HeaderLayoutUri((int) 17564L),
  DataSetWriterType_Diagnostics_DiagnosticsLevel((int) 19551L),
  DataSetWriterType_Diagnostics_TotalInformation((int) 19552L),
  DataSetWriterType_Diagnostics_TotalInformation_Active((int) 19553L),
  DataSetWriterType_Diagnostics_TotalInformation_Classification((int) 19554L),
  DataSetWriterType_Diagnostics_TotalInformation_DiagnosticsLevel((int) 19555L),
  DataSetWriterType_Diagnostics_TotalInformation_TimeFirstChange((int) 19556L),
  DataSetWriterType_Diagnostics_TotalError((int) 19557L),
  DataSetWriterType_Diagnostics_TotalError_Active((int) 19558L),
  DataSetWriterType_Diagnostics_TotalError_Classification((int) 19559L),
  DataSetWriterType_Diagnostics_TotalError_DiagnosticsLevel((int) 19560L),
  DataSetWriterType_Diagnostics_TotalError_TimeFirstChange((int) 19561L),
  DataSetWriterType_Diagnostics_SubError((int) 19563L),
  DataSetWriterType_Diagnostics_Counters_StateError((int) 19565L),
  DataSetWriterType_Diagnostics_Counters_StateError_Active((int) 19566L),
  DataSetWriterType_Diagnostics_Counters_StateError_Classification((int) 19567L),
  DataSetWriterType_Diagnostics_Counters_StateError_DiagnosticsLevel((int) 19568L),
  DataSetWriterType_Diagnostics_Counters_StateError_TimeFirstChange((int) 19569L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalByMethod((int) 19570L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalByMethod_Active((int) 19571L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalByMethod_Classification((int) 19572L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalByMethod_DiagnosticsLevel((int) 19573L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalByMethod_TimeFirstChange((int) 19574L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalByParent((int) 19575L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalByParent_Active((int) 19576L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalByParent_Classification((int) 19577L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalByParent_DiagnosticsLevel((int) 19578L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalByParent_TimeFirstChange((int) 19579L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalFromError((int) 19580L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalFromError_Active((int) 19581L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalFromError_Classification((int) 19582L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalFromError_DiagnosticsLevel((int) 19583L),
  DataSetWriterType_Diagnostics_Counters_StateOperationalFromError_TimeFirstChange((int) 19584L),
  DataSetWriterType_Diagnostics_Counters_StatePausedByParent((int) 19585L),
  DataSetWriterType_Diagnostics_Counters_StatePausedByParent_Active((int) 19586L),
  DataSetWriterType_Diagnostics_Counters_StatePausedByParent_Classification((int) 19587L),
  DataSetWriterType_Diagnostics_Counters_StatePausedByParent_DiagnosticsLevel((int) 19588L),
  DataSetWriterType_Diagnostics_Counters_StatePausedByParent_TimeFirstChange((int) 19589L),
  DataSetWriterType_Diagnostics_Counters_StateDisabledByMethod((int) 19590L),
  DataSetWriterType_Diagnostics_Counters_StateDisabledByMethod_Active((int) 19591L),
  DataSetWriterType_Diagnostics_Counters_StateDisabledByMethod_Classification((int) 19592L),
  DataSetWriterType_Diagnostics_Counters_StateDisabledByMethod_DiagnosticsLevel((int) 19593L),
  DataSetWriterType_Diagnostics_Counters_StateDisabledByMethod_TimeFirstChange((int) 19594L),
  DataSetWriterType_Diagnostics_Counters_FailedDataSetMessages((int) 19596L),
  DataSetWriterType_Diagnostics_Counters_FailedDataSetMessages_Active((int) 19597L),
  DataSetWriterType_Diagnostics_Counters_FailedDataSetMessages_Classification((int) 19598L),
  DataSetWriterType_Diagnostics_Counters_FailedDataSetMessages_DiagnosticsLevel((int) 19599L),
  DataSetWriterType_Diagnostics_Counters_FailedDataSetMessages_TimeFirstChange((int) 19600L),
  DataSetWriterType_Diagnostics_LiveValues_MessageSequenceNumber((int) 19601L),
  DataSetWriterType_Diagnostics_LiveValues_MessageSequenceNumber_DiagnosticsLevel((int) 19602L),
  DataSetWriterType_Diagnostics_LiveValues_StatusCode((int) 19603L),
  DataSetWriterType_Diagnostics_LiveValues_StatusCode_DiagnosticsLevel((int) 19604L),
  DataSetWriterType_Diagnostics_LiveValues_MajorVersion((int) 19605L),
  DataSetWriterType_Diagnostics_LiveValues_MajorVersion_DiagnosticsLevel((int) 19606L),
  DataSetWriterType_Diagnostics_LiveValues_MinorVersion((int) 19607L),
  DataSetWriterType_Diagnostics_LiveValues_MinorVersion_DiagnosticsLevel((int) 19608L),
  DataSetReaderType_Diagnostics_DiagnosticsLevel((int) 19610L),
  DataSetReaderType_Diagnostics_TotalInformation((int) 19611L),
  DataSetReaderType_Diagnostics_TotalInformation_Active((int) 19612L),
  DataSetReaderType_Diagnostics_TotalInformation_Classification((int) 19613L),
  DataSetReaderType_Diagnostics_TotalInformation_DiagnosticsLevel((int) 19614L),
  DataSetReaderType_Diagnostics_TotalInformation_TimeFirstChange((int) 19615L),
  DataSetReaderType_Diagnostics_TotalError((int) 19616L),
  DataSetReaderType_Diagnostics_TotalError_Active((int) 19617L),
  DataSetReaderType_Diagnostics_TotalError_Classification((int) 19618L),
  DataSetReaderType_Diagnostics_TotalError_DiagnosticsLevel((int) 19619L),
  DataSetReaderType_Diagnostics_TotalError_TimeFirstChange((int) 19620L),
  DataSetReaderType_Diagnostics_SubError((int) 19622L),
  DataSetReaderType_Diagnostics_Counters_StateError((int) 19624L),
  DataSetReaderType_Diagnostics_Counters_StateError_Active((int) 19625L),
  DataSetReaderType_Diagnostics_Counters_StateError_Classification((int) 19626L),
  DataSetReaderType_Diagnostics_Counters_StateError_DiagnosticsLevel((int) 19627L),
  DataSetReaderType_Diagnostics_Counters_StateError_TimeFirstChange((int) 19628L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalByMethod((int) 19629L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalByMethod_Active((int) 19630L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalByMethod_Classification((int) 19631L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalByMethod_DiagnosticsLevel((int) 19632L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalByMethod_TimeFirstChange((int) 19633L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalByParent((int) 19634L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalByParent_Active((int) 19635L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalByParent_Classification((int) 19636L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalByParent_DiagnosticsLevel((int) 19637L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalByParent_TimeFirstChange((int) 19638L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalFromError((int) 19639L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalFromError_Active((int) 19640L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalFromError_Classification((int) 19641L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalFromError_DiagnosticsLevel((int) 19642L),
  DataSetReaderType_Diagnostics_Counters_StateOperationalFromError_TimeFirstChange((int) 19643L),
  DataSetReaderType_Diagnostics_Counters_StatePausedByParent((int) 19644L),
  DataSetReaderType_Diagnostics_Counters_StatePausedByParent_Active((int) 19645L),
  DataSetReaderType_Diagnostics_Counters_StatePausedByParent_Classification((int) 19646L),
  DataSetReaderType_Diagnostics_Counters_StatePausedByParent_DiagnosticsLevel((int) 19647L),
  DataSetReaderType_Diagnostics_Counters_StatePausedByParent_TimeFirstChange((int) 19648L),
  DataSetReaderType_Diagnostics_Counters_StateDisabledByMethod((int) 19649L),
  DataSetReaderType_Diagnostics_Counters_StateDisabledByMethod_Active((int) 19650L),
  DataSetReaderType_Diagnostics_Counters_StateDisabledByMethod_Classification((int) 19651L),
  DataSetReaderType_Diagnostics_Counters_StateDisabledByMethod_DiagnosticsLevel((int) 19652L),
  DataSetReaderType_Diagnostics_Counters_StateDisabledByMethod_TimeFirstChange((int) 19653L),
  DataSetReaderType_Diagnostics_Counters_FailedDataSetMessages((int) 19655L),
  DataSetReaderType_Diagnostics_Counters_FailedDataSetMessages_Active((int) 19656L),
  DataSetReaderType_Diagnostics_Counters_FailedDataSetMessages_Classification((int) 19657L),
  DataSetReaderType_Diagnostics_Counters_FailedDataSetMessages_DiagnosticsLevel((int) 19658L),
  DataSetReaderType_Diagnostics_Counters_FailedDataSetMessages_TimeFirstChange((int) 19659L),
  DataSetReaderType_Diagnostics_Counters_DecryptionErrors((int) 19660L),
  DataSetReaderType_Diagnostics_Counters_DecryptionErrors_Active((int) 19661L),
  DataSetReaderType_Diagnostics_Counters_DecryptionErrors_Classification((int) 19662L),
  DataSetReaderType_Diagnostics_Counters_DecryptionErrors_DiagnosticsLevel((int) 19663L),
  DataSetReaderType_Diagnostics_Counters_DecryptionErrors_TimeFirstChange((int) 19664L),
  DataSetReaderType_Diagnostics_LiveValues_MessageSequenceNumber((int) 19665L),
  DataSetReaderType_Diagnostics_LiveValues_MessageSequenceNumber_DiagnosticsLevel((int) 19666L),
  DataSetReaderType_Diagnostics_LiveValues_StatusCode((int) 19667L),
  DataSetReaderType_Diagnostics_LiveValues_StatusCode_DiagnosticsLevel((int) 19668L),
  DataSetReaderType_Diagnostics_LiveValues_MajorVersion((int) 19669L),
  DataSetReaderType_Diagnostics_LiveValues_MajorVersion_DiagnosticsLevel((int) 19670L),
  DataSetReaderType_Diagnostics_LiveValues_MinorVersion((int) 19671L),
  DataSetReaderType_Diagnostics_LiveValues_MinorVersion_DiagnosticsLevel((int) 19672L),
  DataSetReaderType_Diagnostics_LiveValues_SecurityTokenID((int) 19673L),
  DataSetReaderType_Diagnostics_LiveValues_SecurityTokenID_DiagnosticsLevel((int) 19674L),
  DataSetReaderType_Diagnostics_LiveValues_TimeToNextTokenID((int) 19675L),
  DataSetReaderType_Diagnostics_LiveValues_TimeToNextTokenID_DiagnosticsLevel((int) 19676L),
  DataSetWriterType_DataSetWriterId((int) 21092L),
  DataSetWriterType_DataSetFieldContentMask((int) 21093L),
  DataSetWriterType_KeyFrameCount((int) 21094L),
  DataSetReaderType_PublisherId((int) 21097L),
  DataSetReaderType_WriterGroupId((int) 21098L),
  DataSetReaderType_DataSetWriterId((int) 21099L),
  DataSetReaderType_DataSetMetaData((int) 21100L),
  DataSetReaderType_DataSetFieldContentMask((int) 21101L),
  DataSetReaderType_MessageReceiveTimeout((int) 21102L),
  DataSetFolderType_PublishedDataSetName_Placeholder_CyclicDataSet((int) 25524L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableData> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableData value : OpcuaNodeIdServicesVariableData.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableData(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableData enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
