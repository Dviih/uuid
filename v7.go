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

import (
	"math/rand/v2"
	"time"
)

func V7() *UUID {
	a := uint64(time.Now().UnixNano() / time.Millisecond.Milliseconds())

	return &UUID{
		a: (a & 0xFFFFFFFFFFFF0FFF) | 0x0000000000007000,
		b: rand.Uint64(),
	}
}
