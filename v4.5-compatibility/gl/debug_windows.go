// Code generated by glow (https://github.com/neclepsio/glow). DO NOT EDIT.

package gl

import (
	"syscall"
	"unsafe"
)

type DebugProc func(
	source uint32,
	gltype uint32,
	id uint32,
	severity uint32,
	length int32,
	message string,
	userParam unsafe.Pointer)

func newDebugProcCallback(callback DebugProc) uintptr {
	// I can't find why, but on AMD64 I must use 64-bit types, else I get data corruption.
	// I don't know if this is true for 32-bit. Using uintptrs is an untested guess.
	f := func(source uintptr, gltype uintptr, id uintptr, severity uintptr, length uintptr, message unsafe.Pointer, userParam unsafe.Pointer) uintptr {
		callback(uint32(source), uint32(gltype), uint32(id), uint32(severity), int32(length), GoStr((*uint8)(message)), userParam)
		return 0
	}
	return syscall.NewCallbackCDecl(f)
}
