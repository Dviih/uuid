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
	"regexp"
	"testing"
)

func TestV4(t *testing.T) {
	t.Parallel()
	v4 := V4()

	matches, err := regexp.MatchString("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}", v4.String())
	if err != nil {
		t.Errorf("error while compiling regex: %v", err)
	}

	if !matches {
		t.Error("uuid v4 does not matches")
	}
}

func TestV7(t *testing.T) {
	t.Parallel()
	v7 := V7()

	matches, err := regexp.MatchString("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}", v7.String())
	if err != nil {
		t.Errorf("error while compiling regex: %v", err)
	}

	if !matches {
		t.Error("uuid v7 does not matches")
	}
}

func TestEqual(t *testing.T) {
	t.Parallel()

	v4 := V4()
	data, _ := v4.MarshalBinary()

	v42 := UUID(0)
	if err := v42.UnmarshalBinary(data); err != nil {
		t.Errorf("failed to unmarshal uuid: %v", err)
	}

	if v4 != v42 {
		t.Errorf("v4 and v42 does not match: %s vs %s", v4, v42)
	}
}

func TestNotEqual(t *testing.T) {
	t.Parallel()

	if V4() == V4() {
		t.Errorf("two random uuids matches")
	}
}

func BenchmarkV4(b *testing.B) {
	_ = V4()
}

func BenchmarkV7(b *testing.B) {
	_ = V7()
}
