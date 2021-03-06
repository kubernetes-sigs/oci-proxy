/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io"
	"sort"
)

const fileHeader = `/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// File generated by ranges2go DO NOT EDIT

package aws

import (
	"net/netip"
)

// regionToRanges contains a preparsed map of AWS regions to netip.Prefix
var regionToRanges = map[string][]netip.Prefix{
`

func generateRangesGo(w io.Writer, rtp regionsToPrefixes) error {
	// generate source file
	if _, err := io.WriteString(w, fileHeader); err != nil {
		return err
	}

	// ensure iteration order is predictable
	regions := make([]string, 0, len(rtp))
	for region := range rtp {
		regions = append(regions, region)
	}
	sort.Strings(regions)
	for _, region := range regions {
		prefixes := rtp[region]
		if _, err := fmt.Fprintf(w, "\t%q: {\n", region); err != nil {
			return err
		}
		for _, prefix := range prefixes {
			addr := prefix.Addr()
			bits := prefix.Bits()
			// Using netip.*From avoids additional runtime allocation.
			//
			// It also means we don't need error checking / parsing cannot fail
			// at runtime, we've already parsed these and re-emitted them
			// as pre-computed IP address / bit mask values.
			if addr.Is4() {
				b := addr.As4()
				if _, err := fmt.Fprintf(w,
					"\t\tnetip.PrefixFrom(netip.AddrFrom4([4]byte{%d, %d, %d, %d}), %d),\n",
					b[0], b[1], b[2], b[3], bits,
				); err != nil {
					return err
				}
			} else {
				b := addr.As16()
				if _, err := fmt.Fprintf(w,
					"\t\tnetip.PrefixFrom(netip.AddrFrom16([16]byte{%d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d}), %d),\n",
					b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], bits,
				); err != nil {
					return err
				}
			}
		}
		if _, err := io.WriteString(w, "\t},\n"); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, "}\n"); err != nil {
		return err
	}

	return nil
}
