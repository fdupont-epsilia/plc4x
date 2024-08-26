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

package bacnetip

import (
	"bytes"
	"context"
	"fmt"
	"math"

	"github.com/pkg/errors"

	readWriteModel "github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi"
)

// NPDUTypes is a dictionary of message type values and structs
var NPDUTypes map[uint8]func() interface{ Decode(Arg) error }

type NPCI interface {
	PCI

	GetNPDUNetMessage() *uint8

	Encode(pdu Arg) error
	Decode(pdu Arg) error

	setNLM(readWriteModel.NLM)
	getNLM() readWriteModel.NLM
}

type _NPCI struct {
	*_PCI
	*DebugContents

	nlm readWriteModel.NLM
}

var _ NPCI = (*_NPCI)(nil)

func NewNPCI(pduUserData spi.Message, nlm readWriteModel.NLM) NPCI {
	n := &_NPCI{
		nlm: nlm,
	}
	n._PCI = newPCI(pduUserData, nil, nil, nil, false, readWriteModel.NPDUNetworkPriority_NORMAL_MESSAGE)
	switch nlm := pduUserData.(type) {
	case readWriteModel.NLMExactly:
		n.nlm = nlm
	}
	return n
}

func (n *_NPCI) GetNPDUNetMessage() *uint8 {
	if n.nlm == nil {
		return nil
	}
	messageType := n.nlm.GetMessageType()
	return &messageType
}

// Deprecated: check if needed as we do it in update
func (n *_NPCI) setNLM(nlm readWriteModel.NLM) {
	n.nlm = nlm
}

func (n *_NPCI) getNLM() readWriteModel.NLM {
	return n.nlm
}

func (n *_NPCI) Update(npci Arg) error {
	if err := n._PCI.Update(npci); err != nil {
		return errors.Wrap(err, "error updating _PCI")
	}
	switch npci := npci.(type) {
	case NPCI:
		n.nlm = npci.getNLM()
		// TODO: update coordinates...
		return nil
	default:
		return errors.Errorf("invalid NPCI type %T", npci)
	}
}

func (n *_NPCI) Encode(pdu Arg) error {
	if err := pdu.(interface{ Update(Arg) error }).Update(n); err != nil { // TODO: better validate that arg is really PDUData... use switch similar to Update
		return errors.Wrap(err, "error updating pdu")
	}
	// TODO: what should we do here??
	return nil
}

func (n *_NPCI) Decode(pdu Arg) error {
	if err := n._PCI.Update(pdu); err != nil {
		return errors.Wrap(err, "error updating pdu")
	}
	// TODO: what should we do here??
	return nil
}

func (n *_NPCI) deepCopy() *_NPCI {
	return &_NPCI{_PCI: n._PCI.deepCopy(), nlm: n.nlm}
}

type NPDU interface {
	readWriteModel.NPDU
	NPCI
	PDUData

	setNPDU(npdu readWriteModel.NPDU)
	setAPDU(apdu readWriteModel.APDU)
}

type _NPDU struct {
	*_NPCI
	*_PDUData

	// Deprecated: this is the rootMessage so no need to store it here
	npdu readWriteModel.NPDU
	// Deprecated: this is the rootMessage so no need to store it here
	apdu readWriteModel.APDU
}

var _ NPDU = (*_NPDU)(nil)

func NewNPDU(nlm readWriteModel.NLM, apdu readWriteModel.APDU) (NPDU, error) {
	n := &_NPDU{
		apdu: apdu,
	}
	var err error
	n.npdu, err = n.buildNPDU(0, nil, nil, false, readWriteModel.NPDUNetworkPriority_NORMAL_MESSAGE, nlm, apdu)
	if err != nil {
		// TODO: seems to be a legit case, which means we don't build our npdu yet... check that
		//return nil, errors.Wrap(err, "error building NPDU")
	}
	n._NPCI = NewNPCI(n.npdu, nlm).(*_NPCI)
	n._PDUData = NewPDUData(NoArgs).(*_PDUData)
	return n, nil
}

// Deprecated: this is the rootMessage so no need to store it here
func (n *_NPDU) setNPDU(npdu readWriteModel.NPDU) {
	n.npdu = npdu
}

func (n *_NPDU) setAPDU(apdu readWriteModel.APDU) {
	n.apdu = apdu
}

func (n *_NPDU) buildNPDU(hopCount uint8, source *Address, destination *Address, expectingReply bool, networkPriority readWriteModel.NPDUNetworkPriority, nlm readWriteModel.NLM, apdu readWriteModel.APDU) (readWriteModel.NPDU, error) {
	switch {
	case nlm != nil && apdu != nil:
		return nil, errors.New("either specify a NLM or a APDU exclusive")
	case nlm == nil && apdu == nil:
		return nil, errors.New("either specify a NLM or a APDU")
	}
	sourceSpecified := source != nil
	var sourceNetworkAddress *uint16
	var sourceLength *uint8
	var sourceAddress []uint8
	if sourceSpecified {
		sourceSpecified = true
		sourceNetworkAddress = source.AddrNet
		sourceLengthValue := *source.AddrLen
		if sourceLengthValue > math.MaxUint8 {
			return nil, errors.New("source address length overflows")
		}
		sourceLengthValueUint8 := sourceLengthValue
		sourceLength = &sourceLengthValueUint8
		sourceAddress = source.AddrAddress
		if sourceLengthValueUint8 == 0 {
			// If we define the len 0 we must not send the array
			sourceAddress = nil
		}
	}
	destinationSpecified := destination != nil
	var destinationNetworkAddress *uint16
	var destinationLength *uint8
	var destinationAddress []uint8
	var destinationHopCount *uint8
	if destinationSpecified {
		destinationSpecified = true
		destinationNetworkAddress = destination.AddrNet
		destinationLengthValue := *destination.AddrLen
		if destinationLengthValue > math.MaxUint8 {
			return nil, errors.New("source address length overflows")
		}
		destinationLengthValueUint8 := destinationLengthValue
		destinationLength = &destinationLengthValueUint8
		destinationAddress = destination.AddrAddress
		if destinationLengthValueUint8 == 0 {
			// If we define the len 0 we must not send the array
			destinationAddress = nil
		}
		destinationHopCount = &hopCount
	}
	control := readWriteModel.NewNPDUControl(nlm != nil, destinationSpecified, sourceSpecified, expectingReply, networkPriority)
	return readWriteModel.NewNPDU(1, control, destinationNetworkAddress, destinationLength, destinationAddress, sourceNetworkAddress, sourceLength, sourceAddress, destinationHopCount, nlm, apdu, 0), nil
}

func (n *_NPDU) Encode(pdu Arg) error {
	if err := n._NPCI.Encode(pdu); err != nil {
		return errors.Wrap(err, "error encoding _NPCI")
	}
	var err error
	n.npdu, err = n.buildNPDU(0, n.GetPDUSource(), n.GetPDUDestination(), n.GetExpectingReply(), n.GetNetworkPriority(), n.nlm, n.apdu)
	if err != nil {
		return errors.Wrap(err, "error building NPDU")
	}
	serialize, err := n.npdu.Serialize()
	if err != nil {
		return errors.Wrap(err, "error serializing NPDU")
	}
	pdu.(interface{ PutData(n ...byte) }).PutData(serialize...) // TODO: ugly cast...
	return nil
}

func (n *_NPDU) Decode(pdu Arg) error {
	if err := n._NPCI.Decode(pdu); err != nil {
		return errors.Wrap(err, "error decoding _NPCI")
	}
	switch pdu := pdu.(type) {
	case PDUData:
		data := pdu.GetPduData()
		n.PutData(data...)
		var err error
		n.npdu, err = readWriteModel.NPDUParse(context.Background(), data, uint16(len(data)))
		if err != nil {
			return errors.Wrap(err, "error parsing NPDU")
		}
		n.rootMessage = n.npdu
		n.nlm = n.npdu.GetNlm()
		n.apdu = n.npdu.GetApdu()
	}
	return nil
}

func (n *_NPDU) GetRootMessage() spi.Message {
	return n.npdu
}

func (n *_NPDU) GetProtocolVersionNumber() uint8 {
	if n.npdu == nil {
		return 0
	}
	return n.npdu.GetProtocolVersionNumber()
}

func (n *_NPDU) GetControl() readWriteModel.NPDUControl {
	if n.npdu == nil {
		return nil
	}
	return n.npdu.GetControl()
}

func (n *_NPDU) GetDestinationNetworkAddress() *uint16 {
	if n.npdu == nil {
		return nil
	}
	return n.npdu.GetDestinationNetworkAddress()
}

func (n *_NPDU) GetDestinationLength() *uint8 {
	if n.npdu == nil {
		return nil
	}
	return n.npdu.GetDestinationLength()
}

func (n *_NPDU) GetDestinationAddress() []uint8 {
	if n.npdu == nil {
		return nil
	}
	return n.npdu.GetDestinationAddress()
}

func (n *_NPDU) GetSourceNetworkAddress() *uint16 {
	if n.npdu == nil {
		return nil
	}
	return n.npdu.GetSourceNetworkAddress()
}

func (n *_NPDU) GetSourceLength() *uint8 {
	if n.npdu == nil {
		return nil
	}
	return n.npdu.GetSourceLength()
}

func (n *_NPDU) GetSourceAddress() []uint8 {
	if n.npdu == nil {
		return nil
	}
	return n.npdu.GetSourceAddress()
}

func (n *_NPDU) GetHopCount() *uint8 {
	if n.npdu == nil {
		return nil
	}
	return n.npdu.GetHopCount()
}

func (n *_NPDU) GetNlm() readWriteModel.NLM {
	if n.npdu == nil {
		return nil
	}
	return n.npdu.GetNlm()
}

func (n *_NPDU) GetApdu() readWriteModel.APDU {
	if n.npdu == nil {
		return nil
	}
	return n.npdu.GetApdu()
}

func (n *_NPDU) GetDestinationLengthAddon() uint16 {
	if n.npdu == nil {
		return 0
	}
	return n.npdu.GetDestinationLengthAddon()
}

func (n *_NPDU) GetSourceLengthAddon() uint16 {
	if n.npdu == nil {
		return 0
	}
	return n.npdu.GetSourceLengthAddon()
}

func (n *_NPDU) GetPayloadSubtraction() uint16 {
	if n.npdu == nil {
		return 0
	}
	return n.npdu.GetPayloadSubtraction()
}

func (n *_NPDU) deepCopy() *_NPDU {
	return &_NPDU{_NPCI: n._NPCI.deepCopy(), _PDUData: n._PDUData.deepCopy(), npdu: n.npdu, apdu: n.apdu}
}

func (n *_NPDU) DeepCopy() PDU {
	return n.deepCopy()
}

func (n *_NPDU) String() string {
	return fmt.Sprintf("_NPDU{%s}", n._PCI)
}

type WhoIsRouterToNetwork struct {
	*_NPDU

	wirtnNetwork *uint16
}

func NewWhoIsRouterToNetwork(opts ...func(network *WhoIsRouterToNetwork)) (*WhoIsRouterToNetwork, error) {
	w := &WhoIsRouterToNetwork{}
	for _, opt := range opts {
		opt(w)
	}
	npdu, err := NewNPDU(readWriteModel.NewNLMWhoIsRouterToNetwork(w.wirtnNetwork, 0), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	w._NPDU = npdu.(*_NPDU)
	return w, nil
}

func WithWhoIsRouterToNetworkNet(net uint16) func(*WhoIsRouterToNetwork) {
	return func(n *WhoIsRouterToNetwork) {
		n.wirtnNetwork = &net
	}
}

func (n *WhoIsRouterToNetwork) GetWirtnNetwork() *uint16 {
	return n.wirtnNetwork
}

func (n *WhoIsRouterToNetwork) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := npdu.Update(n); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		if n.wirtnNetwork != nil {
			npdu.PutShort(*n.wirtnNetwork)
		}
		npdu.setNPDU(n.npdu)
		npdu.setNLM(n.nlm)
		npdu.setAPDU(n.apdu)
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (n *WhoIsRouterToNetwork) Decode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := n.Update(npdu); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		switch pduUserData := npdu.GetRootMessage().(type) {
		case readWriteModel.NPDUExactly:
			switch nlm := pduUserData.GetNlm().(type) {
			case readWriteModel.NLMWhoIsRouterToNetworkExactly:
				n.setNLM(nlm)
				n.wirtnNetwork = nlm.GetDestinationNetworkAddress()
			}
		}
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (n *WhoIsRouterToNetwork) String() string {
	return fmt.Sprintf("WhoIsRouterToNetwork{%s, wirtnNetwork: %d}", n._NPDU, n.wirtnNetwork)
}

type IAmRouterToNetwork struct {
	*_NPDU

	iartnNetworkList []uint16
}

func NewIAmRouterToNetwork(opts ...func(*IAmRouterToNetwork)) (*IAmRouterToNetwork, error) {
	i := &IAmRouterToNetwork{}
	for _, opt := range opts {
		opt(i)
	}
	npdu, err := NewNPDU(readWriteModel.NewNLMIAmRouterToNetwork(i.iartnNetworkList, 0), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	i._NPDU = npdu.(*_NPDU)
	return i, nil
}

func WithIAmRouterToNetworkNetworkList(iartnNetworkList ...uint16) func(*IAmRouterToNetwork) {
	return func(n *IAmRouterToNetwork) {
		n.iartnNetworkList = iartnNetworkList
	}
}

func (i *IAmRouterToNetwork) GetIartnNetworkList() []uint16 {
	return i.iartnNetworkList
}

func (i *IAmRouterToNetwork) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := npdu.Update(i); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		for _, net := range i.iartnNetworkList {
			npdu.PutShort(net)
		}
		npdu.setNPDU(i.npdu)
		npdu.setNLM(i.nlm)
		npdu.setAPDU(i.apdu)
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (i *IAmRouterToNetwork) Decode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := i.Update(npdu); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		switch pduUserData := npdu.GetRootMessage().(type) {
		case readWriteModel.NPDUExactly:
			switch nlm := pduUserData.GetNlm().(type) {
			case readWriteModel.NLMIAmRouterToNetworkExactly:
				i.setNLM(nlm)
				i.iartnNetworkList = nlm.GetDestinationNetworkAddresses()
			}
		}
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (i *IAmRouterToNetwork) String() string {
	return fmt.Sprintf("IAmRouterToNetwork{%s, iartnNetworkList: %v}", i._NPDU, i.iartnNetworkList)
}

type ICouldBeRouterToNetwork struct {
	*_NPDU

	icbrtnNetwork          uint16
	icbrtnPerformanceIndex uint8
}

func NewICouldBeRouterToNetwork(opts ...func(*ICouldBeRouterToNetwork)) (*ICouldBeRouterToNetwork, error) {
	i := &ICouldBeRouterToNetwork{}
	for _, opt := range opts {
		opt(i)
	}
	npdu, err := NewNPDU(readWriteModel.NewNLMICouldBeRouterToNetwork(i.icbrtnNetwork, i.icbrtnPerformanceIndex, 0), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	i._NPDU = npdu.(*_NPDU)
	return i, nil
}

func WithICouldBeRouterToNetworkNetwork(icbrtnNetwork uint16) func(*ICouldBeRouterToNetwork) {
	return func(n *ICouldBeRouterToNetwork) {
		n.icbrtnNetwork = icbrtnNetwork
	}
}
func WithICouldBeRouterToNetworkPerformanceIndex(icbrtnPerformanceIndex uint8) func(*ICouldBeRouterToNetwork) {
	return func(n *ICouldBeRouterToNetwork) {
		n.icbrtnPerformanceIndex = icbrtnPerformanceIndex
	}
}

func (i *ICouldBeRouterToNetwork) GetIcbrtnNetwork() uint16 {
	return i.icbrtnNetwork
}

func (i *ICouldBeRouterToNetwork) GetIcbrtnPerformanceIndex() uint8 {
	return i.icbrtnPerformanceIndex
}

func (i *ICouldBeRouterToNetwork) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := npdu.Update(i); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		npdu.PutShort(i.icbrtnNetwork)
		npdu.Put(i.icbrtnPerformanceIndex)
		npdu.setNPDU(i.npdu)
		npdu.setNLM(i.nlm)
		npdu.setAPDU(i.apdu)
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (i *ICouldBeRouterToNetwork) Decode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := i.Update(npdu); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		switch pduUserData := npdu.GetRootMessage().(type) {
		case readWriteModel.NPDUExactly:
			switch nlm := pduUserData.GetNlm().(type) {
			case readWriteModel.NLMICouldBeRouterToNetworkExactly:
				i.setNLM(nlm)
				i.icbrtnNetwork = nlm.GetDestinationNetworkAddress()
				i.icbrtnPerformanceIndex = nlm.GetPerformanceIndex()
			}
		}
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (i *ICouldBeRouterToNetwork) String() string {
	return fmt.Sprintf("ICouldBeRouterToNetwork{%s, icbrtnNetwork: %v, icbrtnPerformanceIndex: %v}", i._NPDU, i.icbrtnNetwork, i.icbrtnPerformanceIndex)
}

type RejectMessageToNetwork struct {
	*_NPDU

	rmtnRejectionReason readWriteModel.NLMRejectMessageToNetworkRejectReason
	rmtnDNET            uint16
}

func NewRejectMessageToNetwork(opts ...func(*RejectMessageToNetwork)) (*RejectMessageToNetwork, error) {
	i := &RejectMessageToNetwork{}
	for _, opt := range opts {
		opt(i)
	}
	npdu, err := NewNPDU(readWriteModel.NewNLMRejectMessageToNetwork(i.rmtnRejectionReason, i.rmtnDNET, 0), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	i._NPDU = npdu.(*_NPDU)
	return i, nil
}

func WithRejectMessageToNetworkRejectionReason(reason readWriteModel.NLMRejectMessageToNetworkRejectReason) func(*RejectMessageToNetwork) {
	return func(n *RejectMessageToNetwork) {
		n.rmtnRejectionReason = reason
	}
}

func WithRejectMessageToNetworkDnet(dnet uint16) func(*RejectMessageToNetwork) {
	return func(n *RejectMessageToNetwork) {
		n.rmtnDNET = dnet
	}
}

func (n *RejectMessageToNetwork) GetRmtnRejectionReason() readWriteModel.NLMRejectMessageToNetworkRejectReason {
	return n.rmtnRejectionReason
}

func (n *RejectMessageToNetwork) GetRmtnDNET() uint16 {
	return n.rmtnDNET
}

func (n *RejectMessageToNetwork) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := npdu.Update(n); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		npdu.Put(byte(n.rmtnRejectionReason))
		npdu.PutShort(n.rmtnDNET)
		npdu.setNPDU(n.npdu)
		npdu.setNLM(n.nlm)
		npdu.setAPDU(n.apdu)
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (n *RejectMessageToNetwork) Decode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := n.Update(npdu); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		switch pduUserData := npdu.GetRootMessage().(type) {
		case readWriteModel.NPDUExactly:
			switch nlm := pduUserData.GetNlm().(type) {
			case readWriteModel.NLMRejectMessageToNetworkExactly:
				n.setNLM(nlm)
				n.rmtnRejectionReason = nlm.GetRejectReason()
				n.rmtnDNET = nlm.GetDestinationNetworkAddress()
			}
		}
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (n *RejectMessageToNetwork) String() string {
	return fmt.Sprintf("RejectMessageToNetwork{%s, rmtnRejectionReason: %s, rmtnDNET: %v}", n._NPDU, n.rmtnRejectionReason, n.rmtnDNET)
}

type RouterBusyToNetwork struct {
	*_NPDU

	rbtnNetworkList []uint16
}

func NewRouterBusyToNetwork(opts ...func(*RouterBusyToNetwork)) (*RouterBusyToNetwork, error) {
	i := &RouterBusyToNetwork{}
	for _, opt := range opts {
		opt(i)
	}
	npdu, err := NewNPDU(readWriteModel.NewNLMRouterBusyToNetwork(i.rbtnNetworkList, 0), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	i._NPDU = npdu.(*_NPDU)
	return i, nil
}

func WithRouterBusyToNetworkDnet(networkList []uint16) func(*RouterBusyToNetwork) {
	return func(n *RouterBusyToNetwork) {
		n.rbtnNetworkList = networkList
	}
}

func (r *RouterBusyToNetwork) GetRbtnNetworkList() []uint16 {
	return r.rbtnNetworkList
}

func (r *RouterBusyToNetwork) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := npdu.Update(r); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		for _, net := range r.GetRbtnNetworkList() {
			npdu.PutShort(net)
		}
		npdu.setNPDU(r.npdu)
		npdu.setNLM(r.nlm)
		npdu.setAPDU(r.apdu)
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (r *RouterBusyToNetwork) Decode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := r.Update(npdu); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		switch pduUserData := npdu.GetRootMessage().(type) {
		case readWriteModel.NPDUExactly:
			switch nlm := pduUserData.GetNlm().(type) {
			case readWriteModel.NLMRouterBusyToNetwork:
				r.setNLM(nlm)
				r.rbtnNetworkList = nlm.GetDestinationNetworkAddresses()
			}
		}
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (r *RouterBusyToNetwork) String() string {
	return fmt.Sprintf("RouterBusyToNetwork{%s, rbtnNetworkList: %v}", r._NPDU, r.rbtnNetworkList)
}

type RouterAvailableToNetwork struct {
	*_NPDU

	ratnNetworkList []uint16
}

func NewRouterAvailableToNetwork(opts ...func(*RouterAvailableToNetwork)) (*RouterAvailableToNetwork, error) {
	i := &RouterAvailableToNetwork{}
	for _, opt := range opts {
		opt(i)
	}
	npdu, err := NewNPDU(readWriteModel.NewNLMRouterAvailableToNetwork(i.ratnNetworkList, 0), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	i._NPDU = npdu.(*_NPDU)
	return i, nil
}

func WithRouterAvailableToNetworkDnet(networkList []uint16) func(*RouterAvailableToNetwork) {
	return func(n *RouterAvailableToNetwork) {
		n.ratnNetworkList = networkList
	}
}

func (r *RouterAvailableToNetwork) GetRatnNetworkList() []uint16 {
	return r.ratnNetworkList
}

func (r *RouterAvailableToNetwork) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := npdu.Update(r); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		for _, net := range r.GetRatnNetworkList() {
			npdu.PutShort(net)
		}
		npdu.setNPDU(r.npdu)
		npdu.setNLM(r.nlm)
		npdu.setAPDU(r.apdu)
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (r *RouterAvailableToNetwork) Decode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := r.Update(npdu); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		switch pduUserData := npdu.GetRootMessage().(type) {
		case readWriteModel.NPDUExactly:
			switch nlm := pduUserData.GetNlm().(type) {
			case readWriteModel.NLMRouterAvailableToNetwork:
				r.setNLM(nlm)
				r.ratnNetworkList = nlm.GetDestinationNetworkAddresses()
			}
		}
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (r *RouterAvailableToNetwork) String() string {
	return fmt.Sprintf("RouterAvailableToNetwork{%s, ratnNetworkList: %v}", r._NPDU, r.ratnNetworkList)
}

type RoutingTableEntry struct {
	*DebugContents
	rtDNET     uint16
	rtPortId   uint8
	rtPortInfo []byte
}

func NewRoutingTableEntry(opts ...func(*RoutingTableEntry)) *RoutingTableEntry {
	r := &RoutingTableEntry{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func WithRoutingTableEntryDestinationNetworkAddress(dnet uint16) func(*RoutingTableEntry) {
	return func(r *RoutingTableEntry) {
		r.rtDNET = dnet
	}
}

func WithRoutingTableEntryPortId(id uint8) func(*RoutingTableEntry) {
	return func(r *RoutingTableEntry) {
		r.rtPortId = id
	}
}

func WithRoutingTableEntryPortInfo(portInfo []byte) func(*RoutingTableEntry) {
	return func(r *RoutingTableEntry) {
		r.rtPortInfo = portInfo
	}
}

func (r *RoutingTableEntry) tuple() (destinationNetworkAddress uint16, portId uint8, portInfoLength uint8, portInfo []byte) {
	return r.rtDNET, r.rtPortId, uint8(len(r.rtPortInfo)), r.rtPortInfo
}

func (r *RoutingTableEntry) Equals(other any) bool {
	if r == nil && other == nil {
		return true
	}
	if r == nil {
		return false
	}
	otherEntry, ok := other.(*RoutingTableEntry)
	if !ok {
		return false
	}
	return r.rtDNET == otherEntry.rtDNET &&
		r.rtPortId == otherEntry.rtPortId &&
		bytes.Equal(r.rtPortInfo, otherEntry.rtPortInfo)
}

func (r *RoutingTableEntry) String() string {
	return fmt.Sprintf("RoutingTableEntry{rtDNET: %d, rtPortId: %d, rtPortInfo: %d}", r.rtDNET, r.rtPortId, r.rtPortInfo)
}

type InitializeRoutingTable struct {
	*_NPDU

	irtTable []*RoutingTableEntry
}

func NewInitializeRoutingTable(opts ...func(*InitializeRoutingTable)) (*InitializeRoutingTable, error) {
	i := &InitializeRoutingTable{}
	for _, opt := range opts {
		opt(i)
	}
	npdu, err := NewNPDU(readWriteModel.NewNLMInitializeRoutingTable(i.produceNLMInitializeRoutingTablePortMapping()), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	i._NPDU = npdu.(*_NPDU)
	return i, nil
}

func WithInitializeRoutingTableIrtTable(irtTable ...*RoutingTableEntry) func(*InitializeRoutingTable) {
	return func(r *InitializeRoutingTable) {
		r.irtTable = irtTable
	}
}

func (r *InitializeRoutingTable) GetIrtTable() []*RoutingTableEntry {
	return r.irtTable
}

func (r *InitializeRoutingTable) produceNLMInitializeRoutingTablePortMapping() (numberOfPorts uint8, mappings []readWriteModel.NLMInitializeRoutingTablePortMapping, _ uint16) {
	numberOfPorts = uint8(len(r.irtTable))
	mappings = make([]readWriteModel.NLMInitializeRoutingTablePortMapping, numberOfPorts)
	for i, entry := range r.irtTable {
		mappings[i] = readWriteModel.NewNLMInitializeRoutingTablePortMapping(entry.tuple())
	}
	return
}

func (r *InitializeRoutingTable) produceIRTTable(mappings []readWriteModel.NLMInitializeRoutingTablePortMapping) (irtTable []*RoutingTableEntry) {
	irtTable = make([]*RoutingTableEntry, len(mappings))
	for i, entry := range mappings {
		irtTable[i] = NewRoutingTableEntry(
			WithRoutingTableEntryDestinationNetworkAddress(entry.GetDestinationNetworkAddress()),
			WithRoutingTableEntryPortId(entry.GetPortId()),
			WithRoutingTableEntryPortInfo(entry.GetPortInfo()),
		)
	}
	return
}

func (r *InitializeRoutingTable) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := npdu.Update(r); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		for _, rte := range r.irtTable {
			npdu.PutShort(rte.rtDNET)
			npdu.Put(rte.rtPortId)
			npdu.Put(byte(len(rte.rtPortInfo)))
			npdu.PutData(rte.rtPortInfo...)
		}
		npdu.setNPDU(r.npdu)
		npdu.setNLM(r.nlm)
		npdu.setAPDU(r.apdu)
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (r *InitializeRoutingTable) Decode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := r.Update(npdu); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		switch pduUserData := npdu.GetRootMessage().(type) {
		case readWriteModel.NPDUExactly:
			switch nlm := pduUserData.GetNlm().(type) {
			case readWriteModel.NLMInitializeRoutingTable:
				r.setNLM(nlm)
				r.irtTable = r.produceIRTTable(nlm.GetPortMappings())
			}
		}
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (r *InitializeRoutingTable) String() string {
	return fmt.Sprintf("InitializeRoutingTable{%s, irtTable: %v}", r._NPDU, r.irtTable)
}

type InitializeRoutingTableAck struct {
	*_NPDU

	irtaTable []*RoutingTableEntry
}

func NewInitializeRoutingTableAck(opts ...func(*InitializeRoutingTableAck)) (*InitializeRoutingTableAck, error) {
	i := &InitializeRoutingTableAck{}
	for _, opt := range opts {
		opt(i)
	}
	npdu, err := NewNPDU(readWriteModel.NewNLMInitializeRoutingTableAck(i.produceNLMInitializeRoutingTableAckPortMapping()), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	i._NPDU = npdu.(*_NPDU)
	return i, nil
}

func WithInitializeRoutingTableAckIrtaTable(irtaTable ...*RoutingTableEntry) func(*InitializeRoutingTableAck) {
	return func(r *InitializeRoutingTableAck) {
		r.irtaTable = irtaTable
	}
}

func (r *InitializeRoutingTableAck) GetIrtaTable() []*RoutingTableEntry {
	return r.irtaTable
}

func (r *InitializeRoutingTableAck) produceNLMInitializeRoutingTableAckPortMapping() (numberOfPorts uint8, mappings []readWriteModel.NLMInitializeRoutingTablePortMapping, _ uint16) {
	numberOfPorts = uint8(len(r.irtaTable))
	mappings = make([]readWriteModel.NLMInitializeRoutingTablePortMapping, numberOfPorts)
	for i, entry := range r.irtaTable {
		mappings[i] = readWriteModel.NewNLMInitializeRoutingTablePortMapping(entry.tuple())
	}
	return
}

func (r *InitializeRoutingTableAck) produceIRTTable(mappings []readWriteModel.NLMInitializeRoutingTablePortMapping) (irtTable []*RoutingTableEntry) {
	irtTable = make([]*RoutingTableEntry, len(mappings))
	for i, entry := range mappings {
		irtTable[i] = NewRoutingTableEntry(
			WithRoutingTableEntryDestinationNetworkAddress(entry.GetDestinationNetworkAddress()),
			WithRoutingTableEntryPortId(entry.GetPortId()),
			WithRoutingTableEntryPortInfo(entry.GetPortInfo()),
		)
	}
	return
}

func (r *InitializeRoutingTableAck) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := npdu.Update(r); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		for _, rte := range r.irtaTable {
			npdu.PutShort(rte.rtDNET)
			npdu.Put(rte.rtPortId)
			npdu.Put(byte(len(rte.rtPortInfo)))
			npdu.PutData(rte.rtPortInfo...)
		}
		npdu.setNPDU(r.npdu)
		npdu.setNLM(r.nlm)
		npdu.setAPDU(r.apdu)
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (r *InitializeRoutingTableAck) Decode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := r.Update(npdu); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		switch pduUserData := npdu.GetRootMessage().(type) {
		case readWriteModel.NPDUExactly:
			switch nlm := pduUserData.GetNlm().(type) {
			case readWriteModel.NLMInitializeRoutingTableAck:
				r.setNLM(nlm)
				r.irtaTable = r.produceIRTTable(nlm.GetPortMappings())
			}
		}
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (r *InitializeRoutingTableAck) String() string {
	return fmt.Sprintf("InitializeRoutingTableAck{%s, irtaTable: %v}", r._NPDU, r.irtaTable)
}

type EstablishConnectionToNetwork struct {
	*_NPDU

	ectnDNET            uint16
	ectnTerminationTime uint8
}

func NewEstablishConnectionToNetwork(opts ...func(*EstablishConnectionToNetwork)) (*EstablishConnectionToNetwork, error) {
	i := &EstablishConnectionToNetwork{}
	for _, opt := range opts {
		opt(i)
	}
	npdu, err := NewNPDU(readWriteModel.NewNLMEstablishConnectionToNetwork(i.ectnDNET, i.ectnTerminationTime, 0), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	i._NPDU = npdu.(*_NPDU)
	return i, nil
}

func WithEstablishConnectionToNetworkDNET(dnet uint16) func(*EstablishConnectionToNetwork) {
	return func(n *EstablishConnectionToNetwork) {
		n.ectnDNET = dnet
	}
}

func WithEstablishConnectionToNetworkTerminationTime(terminationTime uint8) func(*EstablishConnectionToNetwork) {
	return func(n *EstablishConnectionToNetwork) {
		n.ectnTerminationTime = terminationTime
	}
}

func (n *EstablishConnectionToNetwork) GetEctnDNET() uint16 {
	return n.ectnDNET
}

func (n *EstablishConnectionToNetwork) GetEctnTerminationTime() uint8 {
	return n.ectnTerminationTime
}

func (n *EstablishConnectionToNetwork) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := npdu.Update(n); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		npdu.PutShort(n.ectnDNET)
		npdu.Put(n.ectnTerminationTime)
		npdu.setNPDU(n.npdu)
		npdu.setNLM(n.nlm)
		npdu.setAPDU(n.apdu)
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (n *EstablishConnectionToNetwork) Decode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := n.Update(npdu); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		switch pduUserData := npdu.GetRootMessage().(type) {
		case readWriteModel.NPDUExactly:
			switch nlm := pduUserData.GetNlm().(type) {
			case readWriteModel.NLMEstablishConnectionToNetworkExactly:
				n.setNLM(nlm)
				n.ectnDNET = nlm.GetDestinationNetworkAddress()
				n.ectnTerminationTime = nlm.GetTerminationTime()
			}
		}
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (n *EstablishConnectionToNetwork) String() string {
	return fmt.Sprintf("EstablishConnectionToNetwork{%s, ectnDNET: %v, ectnTerminationTime: %v}", n._NPDU, n.ectnDNET, n.ectnTerminationTime)
}

type DisconnectConnectionToNetwork struct {
	*_NPDU

	dctnDNET uint16
}

func NewDisconnectConnectionToNetwork(opts ...func(*DisconnectConnectionToNetwork)) (*DisconnectConnectionToNetwork, error) {
	i := &DisconnectConnectionToNetwork{}
	for _, opt := range opts {
		opt(i)
	}
	npdu, err := NewNPDU(readWriteModel.NewNLMDisconnectConnectionToNetwork(i.dctnDNET, 0), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	i._NPDU = npdu.(*_NPDU)
	return i, nil
}

func WithDisconnectConnectionToNetworkDNET(dnet uint16) func(*DisconnectConnectionToNetwork) {
	return func(n *DisconnectConnectionToNetwork) {
		n.dctnDNET = dnet
	}
}

func (n *DisconnectConnectionToNetwork) GetDctnDNET() uint16 {
	return n.dctnDNET
}

func (n *DisconnectConnectionToNetwork) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := npdu.Update(n); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		npdu.PutShort(n.dctnDNET)
		npdu.setNPDU(n.npdu)
		npdu.setNLM(n.nlm)
		npdu.setAPDU(n.apdu)
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (n *DisconnectConnectionToNetwork) Decode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := n.Update(npdu); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		switch pduUserData := npdu.GetRootMessage().(type) {
		case readWriteModel.NPDUExactly:
			switch nlm := pduUserData.GetNlm().(type) {
			case readWriteModel.NLMDisconnectConnectionToNetworkExactly:
				n.setNLM(nlm)
				n.dctnDNET = nlm.GetDestinationNetworkAddress()
			}
		}
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (n *DisconnectConnectionToNetwork) String() string {
	return fmt.Sprintf("DisconnectConnectionToNetwork{%s, dctnDNET: %v}", n._NPDU, n.dctnDNET)
}

type WhatIsNetworkNumber struct {
	*_NPDU
}

func NewWhatIsNetworkNumber(opts ...func(*WhatIsNetworkNumber)) (*WhatIsNetworkNumber, error) {
	i := &WhatIsNetworkNumber{}
	for _, opt := range opts {
		opt(i)
	}
	npdu, err := NewNPDU(readWriteModel.NewNLMWhatIsNetworkNumber(0), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	i._NPDU = npdu.(*_NPDU)
	return i, nil
}

func (n *WhatIsNetworkNumber) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := npdu.Update(n); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		npdu.setNPDU(n.npdu)
		npdu.setNLM(n.nlm)
		npdu.setAPDU(n.apdu)
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (n *WhatIsNetworkNumber) Decode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := n.Update(npdu); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		switch pduUserData := npdu.GetRootMessage().(type) {
		case readWriteModel.NPDUExactly:
			switch nlm := pduUserData.GetNlm().(type) {
			case readWriteModel.NLMWhatIsNetworkNumberExactly:
				n.setNLM(nlm)
			}
		}
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (n *WhatIsNetworkNumber) String() string {
	return fmt.Sprintf("WhatIsNetworkNumber{%s}", n._NPDU)
}

type NetworkNumberIs struct {
	*_NPDU

	nniNet  uint16
	nniFlag bool
}

func NewNetworkNumberIs(opts ...func(*NetworkNumberIs)) (*NetworkNumberIs, error) {
	i := &NetworkNumberIs{}
	for _, opt := range opts {
		opt(i)
	}
	npdu, err := NewNPDU(readWriteModel.NewNLMNetworkNumberIs(i.nniNet, i.nniFlag, 0), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	i._NPDU = npdu.(*_NPDU)
	return i, nil
}

func WithNetworkNumberIsNET(net uint16) func(*NetworkNumberIs) {
	return func(n *NetworkNumberIs) {
		n.nniNet = net
	}
}

func WithNetworkNumberIsTerminationConfigured(configured bool) func(*NetworkNumberIs) {
	return func(n *NetworkNumberIs) {
		n.nniFlag = configured
	}
}

func (n *NetworkNumberIs) GetNniNet() uint16 {
	return n.nniNet
}

func (n *NetworkNumberIs) GetNniFlag() bool {
	return n.nniFlag
}

func (n *NetworkNumberIs) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := npdu.Update(n); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		npdu.PutShort(n.nniNet)
		flag := uint8(0)
		if n.nniFlag {
			flag = 1
		}
		npdu.Put(flag)
		npdu.setNPDU(n.npdu)
		npdu.setNLM(n.nlm)
		npdu.setAPDU(n.apdu)
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (n *NetworkNumberIs) Decode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPDU:
		if err := n.Update(npdu); err != nil {
			return errors.Wrap(err, "error updating _NPCI")
		}
		switch pduUserData := npdu.GetRootMessage().(type) {
		case readWriteModel.NPDUExactly:
			switch nlm := pduUserData.GetNlm().(type) {
			case readWriteModel.NLMNetworkNumberIsExactly:
				n.setNLM(nlm)
				n.nniNet = nlm.GetNetworkNumber()
				n.nniFlag = nlm.GetNetworkNumberConfigured()
			}
		}
		return nil
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
}

func (n *NetworkNumberIs) String() string {
	return fmt.Sprintf("NetworkNumberIs{%s, nniNet: %v, nniFlag: %v}", n._NPDU, n.nniNet, n.nniFlag)
}

func init() {
	NPDUTypes = map[uint8]func() interface{ Decode(Arg) error }{
		0x00: func() interface{ Decode(Arg) error } {
			v, _ := NewWhoIsRouterToNetwork()
			return v
		},
		0x01: func() interface{ Decode(Arg) error } {
			v, _ := NewIAmRouterToNetwork()
			return v
		},
		0x02: func() interface{ Decode(Arg) error } {
			v, _ := NewICouldBeRouterToNetwork()
			return v
		},
		0x03: func() interface{ Decode(Arg) error } {
			v, _ := NewRejectMessageToNetwork()
			return v
		},
		0x04: func() interface{ Decode(Arg) error } {
			v, _ := NewRouterBusyToNetwork()
			return v
		},
		0x05: func() interface{ Decode(Arg) error } {
			v, _ := NewRouterAvailableToNetwork()
			return v
		},
		0x06: func() interface{ Decode(Arg) error } {
			v, _ := NewInitializeRoutingTable()
			return v
		},
		0x07: func() interface{ Decode(Arg) error } {
			v, _ := NewInitializeRoutingTableAck()
			return v
		},
		0x08: func() interface{ Decode(Arg) error } {
			v, _ := NewEstablishConnectionToNetwork()
			return v
		},
		0x09: func() interface{ Decode(Arg) error } {
			v, _ := NewDisconnectConnectionToNetwork()
			return v
		},
		// 0x0A: NewChallengeRequest, // TODO: not present upstream
		// 0x0B: NewSecurityPayload, // TODO: not present upstream
		// 0x0C: NewSecurityResponse, // TODO: not present upstream
		// 0x0D: NewRequestKeyUpdate, // TODO: not present upstream
		// 0x0E: NewUpdateKeyUpdate, // TODO: not present upstream
		// 0x0F: NewUpdateKeyDistributionKey, // TODO: not present upstream
		// 0x10: NewRequestMasterKey, // TODO: not present upstream
		// 0x11: NewSetMasterKey, // TODO: not present upstream
		0x12: func() interface{ Decode(Arg) error } {
			v, _ := NewWhatIsNetworkNumber()
			return v
		},
		0x13: func() interface{ Decode(Arg) error } {
			v, _ := NewNetworkNumberIs()
			return v
		},
	}
}
