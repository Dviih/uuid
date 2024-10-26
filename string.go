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

package uuid

const set = "0123456789abcdef"

func format(num uint64) string {
	var data []byte

	for ; num > 0; num /= 16 {
		data = append(data, set[num%16])
	}

	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	return string(data)
}

func (uuid *UUID) String() string {
	s := format(uuid.a) + format(uuid.b)
	if len(s) != 32 {
		for i := 0; i < 32-len(s); i++ {
			s += "0"
		}
	}

	return s[:8] + "-" + s[8:12] + "-" + s[12:16] + "-" + s[16:20] + "-" + s[20:]
}
