/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package packets

import (
	"crypto"
	"io"
	"strconv"

	"github.com/pkg/errors"

	"github.com/insolar/insolar/core"
)

type PacketRoutable interface {
	GetOrigin() core.ShortNodeID
	GetTarget() core.ShortNodeID
	SetRouting(origin, target core.ShortNodeID)
}

type Serializer interface {
	Serialize() ([]byte, error)
	Deserialize(data io.Reader) error
}

type HeaderSkipDeserializer interface {
	DeserializeWithoutHeader(data io.Reader, header *PacketHeader) error
}

type SignedPacket interface {
	Verify(cryptographyService core.CryptographyService, key crypto.PublicKey) error
	Sign(core.CryptographyService) error
}

type ConsensusPacket interface {
	GetType() PacketType

	SignedPacket
	HeaderSkipDeserializer
	Serializer
	PacketRoutable
}

func ExtractPacket(reader io.Reader) (ConsensusPacket, error) {
	header := PacketHeader{}
	err := header.Deserialize(reader)
	if err != nil {
		return nil, errors.New("[ ExtractPacket ] Can't read packet header")
	}

	var packet ConsensusPacket
	switch header.PacketT {
	case Phase1:
		packet = &Phase1Packet{}
	case Phase2:
		packet = &Phase2Packet{}
	case Phase3:
		packet = &Phase3Packet{}
	default:
		return nil, errors.New("[ ExtractPacket ] Unknown extract packet type. " + strconv.Itoa(int(header.PacketT)))
	}

	err = packet.DeserializeWithoutHeader(reader, &header)
	if err != nil {
		return nil, errors.Wrap(err, "[ ExtractPacket ] Can't DeserializeWithoutHeader packet")
	}

	return packet, nil
}
