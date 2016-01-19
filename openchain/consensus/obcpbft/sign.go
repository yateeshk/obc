/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package obcpbft

import (
	//"fmt"
	//"reflect"

	"github.com/golang/protobuf/proto"
	//"github.com/openblockchain/obc-peer/openchain/util"
)

type signable interface {
	getSignature() []byte
	setSignature(s []byte)
	getID() uint64
	setID(id uint64)
	serialize() ([]byte, error)
}

func (instance *pbftCore) sign(s signable) error {
	s.setSignature(nil)
	// TTD id := []byte("XXX ID")
	// s.setID(instance.id)
	raw, err := s.serialize()
	if err != nil {
		return err
	}
   signedRaw, err := instance.consumer.getCPI().Sign(raw)
   if err != nil {
      return err  // TTD TODO: what should happen is we have an error in sign() ?
   }
   s.setSignature(signedRaw)
	// s.setSignature(instance.cpi.Sign(raw))
	// s.setSignature(util.ComputeCryptoHash(append(id, raw...)))
	return nil
}

func (instance *pbftCore) verify(s signable) error {
	origSig := s.getSignature()
	s.setSignature(nil)
	raw, err := s.serialize()
	s.setSignature(origSig)
	if err != nil {
		return err
	}
	senderHandle, err := getValidatorHandle(s.getID())
   return instance.consumer.getCPI().Verify(senderHandle, origSig, raw) // TTD

	// XXX check that s.Id() is a valid replica
	// instance.cpi.Verify(s.Id(), origSig, raw)
	//if !reflect.DeepEqual(util.ComputeCryptoHash(append(id, raw...)), origSig) {
	//	return fmt.Errorf("invalid signature")
	//}
	//return nil
}

func (vc *ViewChange) getSignature() []byte {
	return vc.Signature
}

func (vc *ViewChange) setSignature(sig []byte) {
	vc.Signature = sig
}

func (vc *ViewChange) getID() uint64 {
	return vc.ReplicaId
}

func (vc *ViewChange) setID(id uint64) {
	// XXX set id
}

func (vc *ViewChange) serialize() ([]byte, error) {
	return proto.Marshal(vc)
}

func (v *Verify) getSignature() []byte {
	return v.Signature
}

func (v *Verify) setSignature(sig []byte) {
	v.Signature = sig
}

func (v *Verify) getID() uint64 {
	return v.ReplicaId
}

func (v *Verify) setID(id uint64) {
	// XXX set ID
}

func (v *Verify) serialize() ([]byte, error) {
	return proto.Marshal(v)
}

func (msg *VerifySet) getSignature() []byte {
	return msg.Signature
}

func (msg *VerifySet) setSignature(sig []byte) {
	msg.Signature = sig
}

func (msg *VerifySet) getID() uint64 {
	return msg.ReplicaId
}

func (msg *VerifySet) setID(id uint64) {
	// XXX set ID
}

func (msg *VerifySet) serialize() ([]byte, error) {
	return proto.Marshal(msg)
}

func (msg *Flush) getSignature() []byte {
	return msg.Signature
}

func (msg *Flush) setSignature(sig []byte) {
	msg.Signature = sig
}

func (msg *Flush) getID() uint64 {
	return msg.ReplicaId
}

func (msg *Flush) setID(id uint64) {
	// XXX set ID
}

func (msg *Flush) serialize() ([]byte, error) {
	return proto.Marshal(msg)
}
