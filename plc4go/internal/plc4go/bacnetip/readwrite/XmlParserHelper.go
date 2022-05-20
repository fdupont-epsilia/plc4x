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

package readwrite

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/bacnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Code generated by code-generation. DO NOT EDIT.

type BacnetipXmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.ParseUint
	_ = strconv.ParseInt
	_ = strings.Join
	_ = utils.Dump
}

func (m BacnetipXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
	switch typeName {
	case "BACnetContextTag":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumberArgument := uint8(parsedUint0)
		dataType := model.BACnetDataTypeByName(parserArguments[1])
		return model.BACnetContextTagParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumberArgument, dataType)
	case "BACnetNotifyTypeTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetNotifyTypeTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "ErrorClassTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.ErrorClassTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetOpeningTag":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumberArgument := uint8(parsedUint0)
		return model.BACnetOpeningTagParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumberArgument)
	case "BACnetStatusFlags":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetStatusFlagsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetPropertyReferenceEnclosed":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetPropertyReferenceEnclosedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetEventTransitionBits":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetEventTransitionBitsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetEventTimestamps":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetEventTimestampsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetTagPayloadReal":
		return model.BACnetTagPayloadRealParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetDeviceObjectReference":
		return model.BACnetDeviceObjectReferenceParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BVLCForeignDeviceTableEntry":
		return model.BVLCForeignDeviceTableEntryParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetReadAccessPropertyReadResult":
		objectType := model.BACnetObjectTypeByName(parserArguments[0])
		propertyIdentifierArgument := model.BACnetPropertyIdentifierByName(parserArguments[1])
		return model.BACnetReadAccessPropertyReadResultParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), objectType, propertyIdentifierArgument)
	case "NLM":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		apduLength := uint16(parsedUint0)
		return model.NLMParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), apduLength)
	case "BACnetActionCommand":
		return model.BACnetActionCommandParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetReliabilityTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetReliabilityTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetTagPayloadDate":
		return model.BACnetTagPayloadDateParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetNotificationParametersExtendedParameters":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetNotificationParametersExtendedParametersParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetReadAccessProperty":
		objectType := model.BACnetObjectTypeByName(parserArguments[0])
		return model.BACnetReadAccessPropertyParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), objectType)
	case "BACnetNotificationParametersChangeOfValueNewValue":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetNotificationParametersChangeOfValueNewValueParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "ErrorCodeTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.ErrorCodeTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetTagPayloadEnumerated":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadEnumeratedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetTagPayloadOctetString":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadOctetStringParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetServiceAckAtomicReadFileStreamOrRecord":
		return model.BACnetServiceAckAtomicReadFileStreamOrRecordParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NPDUControl":
		return model.NPDUControlParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetDeviceObjectPropertyReferenceEnclosed":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetDeviceObjectPropertyReferenceEnclosedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetPropertyStates":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetPropertyStatesParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetReadAccessSpecification":
		return model.BACnetReadAccessSpecificationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetReadAccessResult":
		return model.BACnetReadAccessResultParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetConstructedData":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		objectType := model.BACnetObjectTypeByName(parserArguments[1])
		propertyIdentifierArgument := model.BACnetPropertyIdentifierByName(parserArguments[2])
		return model.BACnetConstructedDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, objectType, propertyIdentifierArgument)
	case "BACnetTimeStampEnclosed":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetTimeStampEnclosedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetEventTypeTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetEventTypeTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetObjectPropertyReference":
		return model.BACnetObjectPropertyReferenceParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetLifeSafetyStateTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetLifeSafetyStateTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetResultFlags":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetResultFlagsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetTagPayloadTime":
		return model.BACnetTagPayloadTimeParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetTagPayloadSignedInteger":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadSignedIntegerParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetEventSummary":
		return model.BACnetEventSummaryParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetUnconfirmedServiceRequest":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		serviceRequestLength := uint16(parsedUint0)
		return model.BACnetUnconfirmedServiceRequestParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), serviceRequestLength)
	case "BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord":
		return model.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BVLC":
		return model.BVLCParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetReadAccessResultListOfResults":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		objectType := model.BACnetObjectTypeByName(parserArguments[1])
		return model.BACnetReadAccessResultListOfResultsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, objectType)
	case "BACnetDateTimeEnclosed":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetDateTimeEnclosedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetRejectReasonTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetRejectReasonTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetTagPayloadObjectIdentifier":
		return model.BACnetTagPayloadObjectIdentifierParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BVLCBroadcastDistributionTableEntry":
		return model.BVLCBroadcastDistributionTableEntryParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetPropertyWriteDefinition":
		objectType := model.BACnetObjectTypeByName(parserArguments[0])
		return model.BACnetPropertyWriteDefinitionParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), objectType)
	case "ListOfCovNotificationsValue":
		objectType := model.BACnetObjectTypeByName(parserArguments[0])
		return model.ListOfCovNotificationsValueParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), objectType)
	case "BACnetBinaryPVTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetBinaryPVTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetDateTime":
		return model.BACnetDateTimeParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ErrorEnclosed":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.ErrorEnclosedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "APDU":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		apduLength := uint16(parsedUint0)
		return model.APDUParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), apduLength)
	case "BACnetSegmentationTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetSegmentationTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetTagPayloadCharacterString":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadCharacterStringParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetConfirmedServiceRequestReadRangeRange":
		return model.BACnetConfirmedServiceRequestReadRangeRangeParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetError":
		errorChoice := model.BACnetConfirmedServiceChoiceByName(parserArguments[0])
		return model.BACnetErrorParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), errorChoice)
	case "BACnetPropertyIdentifierTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetPropertyIdentifierTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetTimeStamp":
		return model.BACnetTimeStampParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetNotificationParameters":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		objectType := model.BACnetObjectTypeByName(parserArguments[1])
		return model.BACnetNotificationParametersParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, objectType)
	case "BACnetTimeStampsEnclosed":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetTimeStampsEnclosedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetClosingTag":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumberArgument := uint8(parsedUint0)
		return model.BACnetClosingTagParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumberArgument)
	case "BACnetConfirmedServiceRequest":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		serviceRequestLength := uint16(parsedUint0)
		return model.BACnetConfirmedServiceRequestParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), serviceRequestLength)
	case "BACnetEventProrities":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetEventProritiesParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "ListOfCovNotificationsList":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.ListOfCovNotificationsListParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetEventSummariesList":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetEventSummariesListParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetAddress":
		return model.BACnetAddressParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetTagPayloadUnsignedInteger":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadUnsignedIntegerParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference":
		return model.BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReferenceParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetApplicationTag":
		return model.BACnetApplicationTagParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetTagPayloadBitString":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadBitStringParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetNetworkTypeTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetNetworkTypeTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetActionTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetActionTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetObjectPropertyReferenceEnclosed":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetObjectPropertyReferenceEnclosedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetDeviceObjectPropertyReference":
		return model.BACnetDeviceObjectPropertyReferenceParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetConstructedDataElement":
		objectType := model.BACnetObjectTypeByName(parserArguments[0])
		propertyIdentifierArgument := model.BACnetPropertyIdentifierByName(parserArguments[1])
		return model.BACnetConstructedDataElementParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), objectType, propertyIdentifierArgument)
	case "BACnetEventStateTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetEventStateTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	case "BACnetPropertyValues":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		objectType := model.BACnetObjectTypeByName(parserArguments[1])
		return model.BACnetPropertyValuesParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, objectType)
	case "BACnetTagHeader":
		return model.BACnetTagHeaderParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetTagPayloadBoolean":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetTagPayloadBooleanParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetTagPayloadDouble":
		return model.BACnetTagPayloadDoubleParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetPropertyValue":
		objectType := model.BACnetObjectTypeByName(parserArguments[0])
		return model.BACnetPropertyValueParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), objectType)
	case "BACnetActionList":
		return model.BACnetActionListParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "VTCloseErrorListOfVTSessionIdentifiers":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.VTCloseErrorListOfVTSessionIdentifiersParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "NLMInitalizeRoutingTablePortMapping":
		return model.NLMInitalizeRoutingTablePortMappingParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SubscribeCOVPropertyMultipleErrorFirstFailedSubscription":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecifications":
		return model.BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetWriteAccessSpecification":
		return model.BACnetWriteAccessSpecificationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetServiceAck":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		serviceRequestLength := uint16(parsedUint0)
		return model.BACnetServiceAckParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), serviceRequestLength)
	case "ListOfCovNotifications":
		return model.ListOfCovNotificationsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetAbortReasonTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
		if err != nil {
			return nil, err
		}
		actualLength := uint32(parsedUint0)
		return model.BACnetAbortReasonTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), actualLength)
	case "BACnetConfirmedServiceRequestCreateObjectObjectSpecifier":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetConfirmedServiceRequestCreateObjectObjectSpecifierParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "Error":
		return model.ErrorParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsListParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "NPDU":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		npduLength := uint16(parsedUint0)
		return model.NPDUParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), npduLength)
	case "BACnetPropertyReference":
		return model.BACnetPropertyReferenceParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetLifeSafetyModeTagged":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		tagClass := model.TagClassByName(parserArguments[1])
		return model.BACnetLifeSafetyModeTaggedParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber, tagClass)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
