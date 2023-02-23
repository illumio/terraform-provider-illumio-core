// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"reflect"
	"testing"
)

func testResourceIllumioVirtualServiceStateDataV1() map[string]any {
	return map[string]any{
		"name": "VS-CRM-DB",
		"href": "/orgs/1/sec_policy/draft/virtual_services/1edaa88f-e905-4682-8623-df55e3d2a966",
		"id":   "/orgs/1/sec_policy/draft/virtual_services/1edaa88f-e905-4682-8623-df55e3d2a966",
		// other fields omitted for brevity, having name/href/id
		// shows that other fields are kept as-is
		"service_addresses": []map[string]any{
			{
				"ip":           "127.0.0.1",
				"network_href": "/orgs/1/networks/a2f3fdfc-c179-4700-8525-133727c3bab4",
			},
		},
	}
}

func testResourceIllumioVirtualServiceStateDataV2() map[string]any {
	return map[string]any{
		"name": "VS-CRM-DB",
		"href": "/orgs/1/sec_policy/draft/virtual_services/1edaa88f-e905-4682-8623-df55e3d2a966",
		"id":   "/orgs/1/sec_policy/draft/virtual_services/1edaa88f-e905-4682-8623-df55e3d2a966",
		"service_addresses": []map[string]any{
			{
				"ip":      "127.0.0.1",
				"network": []map[string]any{{"href": "/orgs/1/networks/a2f3fdfc-c179-4700-8525-133727c3bab4"}},
			},
		},
	}
}

func TestResourceIllumioVirtualServiceStateUpgradeV1(t *testing.T) {
	ctx := context.Background()

	expected := testResourceIllumioVirtualServiceStateDataV2()
	actual, err := resourceIllumioVirtualServiceStateUpgradeV1(ctx, testResourceIllumioVirtualServiceStateDataV1(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}
