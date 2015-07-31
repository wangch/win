//
//
//

package win

import (
	"syscall"
	"unsafe"
)

const (
	SND_SYNC      = 0x0000 /* play synchronously (default) */
	SND_ASYNC     = 0x0001 /* play asynchronously */
	SND_NODEFAULT = 0x0002 /* silence (!default) if sound not found */
	SND_MEMORY    = 0x0004 /* pszSound points to a memory file */
	SND_LOOP      = 0x0008 /* loop the sound until next sndPlaySound */
	SND_NOSTOP    = 0x0010 /* don't stop any currently playing sound */

	SND_NOWAIT   = 0x00002000 /* don't wait if the driver is busy */
	SND_ALIAS    = 0x00010000 /* name is a registry alias */
	SND_ALIAS_ID = 0x00110000 /* alias is a predefined ID */
	SND_FILENAME = 0x00020000 /* name is file name */
	SND_RESOURCE = 0x00040004 /* name is resource name or atom */
)

var (
	// Library
	libwinmm uintptr

	// Functions
	// mciSendCommand uintptr
	mciSendString uintptr
	playSound     uintptr
)

func init() {
	// Library
	libwinmm = MustLoadLibrary("winmm.dll")

	// Functions
	// mciSendCommand = MustGetProcAddress(libwinmm, "mciSendCommandW")
	mciSendString = MustGetProcAddress(libwinmm, "mciSendStringW")
	playSound = MustGetProcAddress(libwinmm, "PlaySoundW")
}

// Functions
func MciSendString(lpszCommand *uint16) {
	syscall.Syscall6(mciSendString, 4,
		uintptr(unsafe.Pointer(lpszCommand)),
		uintptr(unsafe.Pointer(nil)),
		uintptr(0),
		uintptr(unsafe.Pointer(nil)),
		0,
		0)
}

func PlaySound(pszSound *uint16, fdwSound uint32) bool {
	ret, _, _ := syscall.Syscall(playSound, 3,
		uintptr(unsafe.Pointer(pszSound)),
		uintptr(unsafe.Pointer(nil)),
		uintptr(fdwSound))
	return ret != 0
}
