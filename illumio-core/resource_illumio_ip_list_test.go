// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"net"
	"reflect"
	"testing"
)

func TestIPListNormalization(t *testing.T) {
	type IPRangeTest struct {
		ipRange  []any
		expected []any
	}
	tt := []IPRangeTest{
		{
			ipRange: []any{
				map[string]any{
					"from_ip":     "10.0.0.14/32",
					"description": "single-ip range",
					"exclusion":   false,
				},
				map[string]any{
					"from_ip":     "10.0.0.14/8",
					"description": "host bits set",
					"exclusion":   false,
				},
				map[string]any{ // exact duplicate network, will be dropped
					"from_ip":     "10.0.0.12/8",
					"description": "host bits set",
					"exclusion":   false,
				},
			},
			expected: []any{
				map[string]any{
					"from_ip":     "10.0.0.14", // single-ip range
					"description": "single-ip range",
					"exclusion":   false,
				},
				map[string]any{
					"from_ip":     "10.0.0.0/8",
					"description": "host bits set",
					"exclusion":   false,
				},
			},
		},
		{
			ipRange: []any{
				map[string]any{
					"from_ip":     "10.0.0.0/8",
					"description": "no change necessary",
					"exclusion":   false,
				},
				map[string]any{
					"from_ip":     "10.0.0.1/8",
					"description": "not exact duplicate",
					"exclusion":   false,
				},
			},
			expected: []any{
				map[string]any{
					"from_ip":     "10.0.0.0/8",
					"description": "no change necessary",
					"exclusion":   false,
				},
				map[string]any{
					"from_ip":     "10.0.0.0/8",
					"description": "not exact duplicate",
					"exclusion":   false,
				},
			},
		},
	}

	for _, test := range tt {
		result, _ := normalizeIPRanges(test.ipRange)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Got %q, expected %q", result, test.expected)
		}
	}
}

func TestCheckSingleAddressRange(t *testing.T) {
	type IPRangeTest struct {
		cidr     string
		expected bool
	}
	tt := []IPRangeTest{
		{
			cidr:     "127.0.0.1/32",
			expected: true,
		},
		{
			cidr:     "127.0.0.1/31",
			expected: false,
		},
		{
			cidr:     "2001:4860:4860::8844/128",
			expected: true,
		},
		{
			cidr:     "::1/128",
			expected: true,
		},
		{
			cidr:     "::1/127",
			expected: false,
		},
	}

	for _, test := range tt {
		ip, ipnet, err := net.ParseCIDR(test.cidr)
		if err != nil {
			t.Fatalf("Failed to parse CIDR range string %q: %v", test.cidr, err)
		}

		result := isSingleAddressRange(ip, ipnet)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Got %v, expected %v", result, test.expected)
		}
	}
}
