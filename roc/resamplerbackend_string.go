// Code generated by "stringer -type ResamplerBackend -trimprefix ResamplerBackend"; DO NOT EDIT.

package roc

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ResamplerBackendDefault-0]
	_ = x[ResamplerBackendBuiltin-1]
	_ = x[ResamplerBackendSpeex-2]
}

const _ResamplerBackend_name = "DefaultBuiltinSpeex"

var _ResamplerBackend_index = [...]uint8{0, 7, 14, 19}

func (i ResamplerBackend) String() string {
	if i < 0 || i >= ResamplerBackend(len(_ResamplerBackend_index)-1) {
		return "ResamplerBackend(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ResamplerBackend_name[_ResamplerBackend_index[i]:_ResamplerBackend_index[i+1]]
}
