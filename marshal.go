/*
 *     UUID implementation combining two uint64 types.
 *     Copyright (C) 2024  Dviih
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

//go:build !dviih_uuid_nomarshal

package uuid

import (
	"encoding/binary"
	"errors"
)

var InvalidLength = errors.New("invalid length")

func (uuid *UUID) MarshalBinary() ([]byte, error) {
	return binary.LittleEndian.AppendUint64(binary.LittleEndian.AppendUint64(nil, uuid.a), uuid.b), nil
}

func (uuid *UUID) UnmarshalBinary(data []byte) error {
	if len(data) != 16 {
		return InvalidLength
	}

	uuid.a = binary.LittleEndian.Uint64(data[:8])
	uuid.b = binary.LittleEndian.Uint64(data[8:16])

	return nil
}
