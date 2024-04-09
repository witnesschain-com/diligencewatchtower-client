/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package opchain

import (
	wtCommon "github.com/diligencewatchtower-client/common"
)

// returns `true` if output root is valid
func ValidateOutputRoot(node_output string, L2OO_output string) bool {
	wtCommon.Info("Validating proposed output root")
	if node_output == L2OO_output {
		wtCommon.Success("Proposed output root is valid!")
		return true
	}
	wtCommon.Error("Alert! the proposed output root is Invalid!")
	return false
}
