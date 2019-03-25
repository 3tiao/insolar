// Code generated by "stringer -type=NetworkState"; DO NOT EDIT.

package core

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoNetworkState-0]
	_ = x[VoidNetworkState-1]
	_ = x[JetlessNetworkState-2]
	_ = x[AuthorizationNetworkState-3]
	_ = x[CompleteNetworkState-4]
}

const _NetworkState_name = "NoNetworkStateVoidNetworkStateJetlessNetworkStateAuthorizationNetworkStateCompleteNetworkState"

var _NetworkState_index = [...]uint8{0, 14, 30, 49, 74, 94}

func (i NetworkState) String() string {
	if i < 0 || i >= NetworkState(len(_NetworkState_index)-1) {
		return "NetworkState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NetworkState_name[_NetworkState_index[i]:_NetworkState_index[i+1]]
}