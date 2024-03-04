//go:build tools
// +build tools

// This file tracks toolings that aren't directly used by backend
// but that are required to build the project.
// Refer to https://go.dev/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/mailru/easyjson/easyjson"
)
