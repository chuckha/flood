package api

import (
	"strings"
	"testing"
)

var urlTestCases = []struct {
	resource, id, action string
	out                  string
}{
	{
		resource: "droplets",
		id:       "",
		action:   "list",
		out:      base + "/droplets",
	},
	{
		resource: "droplets",
		id:       "",
		action:   "create",
		out:      base + "/droplets/new",
	},
	{
		resource: "droplets",
		id:       "12",
		action:   "reboot",
		out:      base + "/droplets/12/reboot",
	},
	{
		resource: "droplets",
		id:       "13",
		action:   "power_cycle",
		out:      base + "/droplets/13/power_cycle",
	},
	{
		resource: "droplets",
		id:       "44",
		action:   "shutdown",
		out:      base + "/droplets/44/shutdown",
	},
	{
		resource: "droplets",
		id:       "93",
		action:   "power_off",
		out:      base + "/droplets/93/power_off",
	},
	{
		resource: "droplets",
		id:       "21",
		action:   "power_on",
		out:      base + "/droplets/21/power_on",
	},
	{
		resource: "droplets",
		id:       "2",
		action:   "password_reset",
		out:      base + "/droplets/2/password_reset",
	},
	{
		resource: "droplets",
		id:       "1",
		action:   "resize",
		out:      base + "/droplets/1/resize",
	},
	{
		resource: "droplets",
		id:       "4",
		action:   "snapshot",
		out:      base + "/droplets/4/snapshot",
	},
	{
		resource: "droplets",
		id:       "1",
		action:   "restore",
		out:      base + "/droplets/1/restore",
	},
	{
		resource: "droplets",
		id:       "4",
		action:   "rebuild",
		out:      base + "/droplets/4/rebuild",
	},
	{
		resource: "droplets",
		id:       "0",
		action:   "rename",
		out:      base + "/droplets/0/rename",
	},
	{
		resource: "droplets",
		id:       "11",
		action:   "destroy",
		out:      base + "/droplets/11/destroy",
	},
}

// input is always valid
func TestGetUrl(t *testing.T) {
	for _, tc := range urlTestCases {
		u := GetUrl(tc.resource, tc.id, tc.action)
		// ignore the other bits.
		cut := strings.LastIndex(u, "?")
		if u[:cut] != tc.out {
			t.Errorf("Expected %v\nGot: %v\n", tc.out, u)
		}
	}
}
