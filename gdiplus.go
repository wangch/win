// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package win

import (
	"syscall"
	"unsafe"
)

type GpStatus int32

const (
	Ok                        GpStatus = 0
	GenericError              GpStatus = 1
	InvalidParameter          GpStatus = 2
	OutOfMemory               GpStatus = 3
	ObjectBusy                GpStatus = 4
	InsufficientBuffer        GpStatus = 5
	NotImplemented            GpStatus = 6
	Win32Error                GpStatus = 7
	WrongState                GpStatus = 8
	Aborted                   GpStatus = 9
	FileNotFound              GpStatus = 10
	ValueOverflow             GpStatus = 11
	AccessDenied              GpStatus = 12
	UnknownImageFormat        GpStatus = 13
	FontFamilyNotFound        GpStatus = 14
	FontStyleNotFound         GpStatus = 15
	NotTrueTypeFont           GpStatus = 16
	UnsupportedGdiplusVersion GpStatus = 17
	GdiplusNotInitialized     GpStatus = 18
	PropertyNotFound          GpStatus = 19
	PropertyNotSupported      GpStatus = 20
	ProfileNotFound           GpStatus = 21
)

func (s GpStatus) String() string {
	switch s {
	case Ok:
		return "Ok"

	case GenericError:
		return "GenericError"

	case InvalidParameter:
		return "InvalidParameter"

	case OutOfMemory:
		return "OutOfMemory"

	case ObjectBusy:
		return "ObjectBusy"

	case InsufficientBuffer:
		return "InsufficientBuffer"

	case NotImplemented:
		return "NotImplemented"

	case Win32Error:
		return "Win32Error"

	case WrongState:
		return "WrongState"

	case Aborted:
		return "Aborted"

	case FileNotFound:
		return "FileNotFound"

	case ValueOverflow:
		return "ValueOverflow"

	case AccessDenied:
		return "AccessDenied"

	case UnknownImageFormat:
		return "UnknownImageFormat"

	case FontFamilyNotFound:
		return "FontFamilyNotFound"

	case FontStyleNotFound:
		return "FontStyleNotFound"

	case NotTrueTypeFont:
		return "NotTrueTypeFont"

	case UnsupportedGdiplusVersion:
		return "UnsupportedGdiplusVersion"

	case GdiplusNotInitialized:
		return "GdiplusNotInitialized"

	case PropertyNotFound:
		return "PropertyNotFound"

	case PropertyNotSupported:
		return "PropertyNotSupported"

	case ProfileNotFound:
		return "ProfileNotFound"
	}

	return "Unknown Status Value"
}

type GdiplusStartupInput struct {
	GdiplusVersion           uint32
	DebugEventCallback       uintptr
	SuppressBackgroundThread BOOL
	SuppressExternalCodecs   BOOL
}

type GdiplusStartupOutput struct {
	NotificationHook   uintptr
	NotificationUnhook uintptr
}

type ARGB uint32
type GpColor struct {
	R, G, B, A uint8
}
type GpRect struct {
	X, Y, W, H int
}
type GpRectF struct {
	X, Y, W, H float32
}
type GpPoint struct {
	X, Y int
}
type GpPointF struct {
	X, Y float32
}
type GpCharacterRange struct {
	Begin, Length int
}

//type (
type GpGraphics uintptr
type GpImage uintptr
type GpBitmap uintptr
type GpPen uintptr
type GpBrush uintptr
type GpFont uintptr
type GpFontFamily uintptr
type GpStringFormat uintptr
type GpRegion uintptr

var (
	// Library
	libgdiplus uintptr

	// Functions
	gdipCreateBitmapFromFile                     uintptr
	gdipCreateBitmapFromHBITMAP                  uintptr
	gdipCreateHBITMAPFromBitmap                  uintptr
	gdipDisposeImage                             uintptr
	gdiplusShutdown                              uintptr
	gdiplusStartup                               uintptr
	gdipCreateFromHDC                            uintptr
	gdipDeleteGraphics                           uintptr
	gdipDrawImage                                uintptr
	gdipDrawImageRect                            uintptr
	gdipDrawImagePointRect                       uintptr
	gdipDrawLine                                 uintptr
	gdipDrawLines                                uintptr
	gdipDrawRect                                 uintptr
	gdipFillRect                                 uintptr
	gdipDrawLineF                                uintptr
	gdipDrawLinesF                               uintptr
	gdipDrawRectF                                uintptr
	gdipFillRectF                                uintptr
	gdipCreatePen                                uintptr
	gdipDeletePen                                uintptr
	gdipSetPenWidth                              uintptr
	gdipCreateSolidBrush                         uintptr
	gdipCreateLineBrush                          uintptr
	gdipCreateHatchBrush                         uintptr
	gdipDeleteBrush                              uintptr
	gdipDrawString                               uintptr
	gdipMeasureCharacterRanges                   uintptr
	gdipCreateStringFormat                       uintptr
	gdipStringFormatGetGenericTypographic        uintptr
	gdipDeleteStringFormat                       uintptr
	gdipStringFormatGetGenericDefault            uintptr
	gdipSetStringFormatFlags                     uintptr
	gdipGetStringFormatFlags                     uintptr
	gdipSetStringFormatAlign                     uintptr
	gdipSetStringFormatMeasurableCharacterRanges uintptr
	gdipCreateFontFamily                         uintptr
	gdipDeleteFontFamily                         uintptr
	gdipCreateFontFromDC                         uintptr
	gdipCreateFont                               uintptr
	gdipDeleteFont                               uintptr
	gdipGetFontStyle                             uintptr
	gdipGetFontSize                              uintptr
	gdipSetClipRect                              uintptr
	gdipResetClip                                uintptr
)

var (
	token uintptr
)

func init() {
	// Library
	libgdiplus = MustLoadLibrary("gdiplus.dll")

	// Functions
	gdipCreateBitmapFromFile = MustGetProcAddress(libgdiplus, "GdipCreateBitmapFromFile")
	gdipCreateBitmapFromHBITMAP = MustGetProcAddress(libgdiplus, "GdipCreateBitmapFromHBITMAP")
	gdipCreateHBITMAPFromBitmap = MustGetProcAddress(libgdiplus, "GdipCreateHBITMAPFromBitmap")
	gdipDisposeImage = MustGetProcAddress(libgdiplus, "GdipDisposeImage")
	gdiplusShutdown = MustGetProcAddress(libgdiplus, "GdiplusShutdown")
	gdiplusStartup = MustGetProcAddress(libgdiplus, "GdiplusStartup")
	gdipCreateFromHDC = MustGetProcAddress(libgdiplus, "GdipCreateFromHDC")
	gdipDeleteGraphics = MustGetProcAddress(libgdiplus, "GdipDeleteGraphics")
	gdipDrawImage = MustGetProcAddress(libgdiplus, "GdipDrawImageI")
	gdipDrawImageRect = MustGetProcAddress(libgdiplus, "GdipDrawImageRectI")
	gdipDrawImagePointRect = MustGetProcAddress(libgdiplus, "GdipDrawImagePointRectI")
	gdipDrawLine = MustGetProcAddress(libgdiplus, "GdipDrawLineI")
	gdipDrawLines = MustGetProcAddress(libgdiplus, "GdipDrawLinesI")
	gdipDrawRect = MustGetProcAddress(libgdiplus, "GdipDrawRectangleI")
	gdipFillRect = MustGetProcAddress(libgdiplus, "GdipFillRectangleI")
	gdipDrawLineF = MustGetProcAddress(libgdiplus, "GdipDrawLine")
	gdipDrawLinesF = MustGetProcAddress(libgdiplus, "GdipDrawLines")
	gdipDrawRectF = MustGetProcAddress(libgdiplus, "GdipDrawRectangle")
	gdipFillRectF = MustGetProcAddress(libgdiplus, "GdipFillRectangle")
	gdipCreatePen = MustGetProcAddress(libgdiplus, "GdipCreatePen1")
	gdipDeletePen = MustGetProcAddress(libgdiplus, "GdipDeletePen")
	gdipSetPenWidth = MustGetProcAddress(libgdiplus, "GdipSetPenWidth")
	gdipCreateSolidBrush = MustGetProcAddress(libgdiplus, "GdipCreateSolidFill")
	gdipCreateLineBrush = MustGetProcAddress(libgdiplus, "GdipCreateLineBrush")
	gdipCreateHatchBrush = MustGetProcAddress(libgdiplus, "GdipCreateHatchBrush")
	gdipDeleteBrush = MustGetProcAddress(libgdiplus, "GdipDeleteBrush")
	gdipDrawString = MustGetProcAddress(libgdiplus, "GdipDrawString")
	gdipMeasureCharacterRanges = MustGetProcAddress(libgdiplus, "GdipMeasureCharacterRanges")
	gdipCreateStringFormat = MustGetProcAddress(libgdiplus, "GdipCreateStringFormat")
	gdipStringFormatGetGenericTypographic = MustGetProcAddress(libgdiplus, "GdipStringFormatGetGenericTypographic")
	gdipDeleteStringFormat = MustGetProcAddress(libgdiplus, "GdipDeleteStringFormat")
	gdipStringFormatGetGenericDefault = MustGetProcAddress(libgdiplus, "GdipStringFormatGetGenericDefault")
	gdipSetStringFormatFlags = MustGetProcAddress(libgdiplus, "GdipSetStringFormatFlags")
	gdipGetStringFormatFlags = MustGetProcAddress(libgdiplus, "GdipGetStringFormatFlags")
	gdipSetStringFormatAlign = MustGetProcAddress(libgdiplus, "GdipSetStringFormatAlign")
	gdipSetStringFormatMeasurableCharacterRanges = MustGetProcAddress(libgdiplus, "GdipSetStringFormatMeasurableCharacterRanges")
	gdipCreateFontFamily = MustGetProcAddress(libgdiplus, "GdipCreateFontFamilyFromName")
	gdipDeleteFontFamily = MustGetProcAddress(libgdiplus, "GdipDeleteFontFamily")
	gdipCreateFontFromDC = MustGetProcAddress(libgdiplus, "GdipCreateFontFromDC")
	gdipCreateFont = MustGetProcAddress(libgdiplus, "GdipCreateFont")
	gdipDeleteFont = MustGetProcAddress(libgdiplus, "GdipDeleteFont")
	gdipGetFontStyle = MustGetProcAddress(libgdiplus, "GdipGetFontStyle")
	gdipGetFontSize = MustGetProcAddress(libgdiplus, "GdipGetFontSize")
	gdipSetClipRect = MustGetProcAddress(libgdiplus, "GdipSetClipRectI")
	gdipResetClip = MustGetProcAddress(libgdiplus, "GdipResetClip")
}

func GdipCreateBitmapFromFile(filename *uint16, bitmap **GpBitmap) GpStatus {
	ret, _, _ := syscall.Syscall(gdipCreateBitmapFromFile, 2,
		uintptr(unsafe.Pointer(filename)),
		uintptr(unsafe.Pointer(bitmap)),
		0)

	return GpStatus(ret)
}

func GdipCreateBitmapFromHBITMAP(hbm HBITMAP, hpal HPALETTE, bitmap **GpBitmap) GpStatus {
	ret, _, _ := syscall.Syscall(gdipCreateBitmapFromHBITMAP, 3,
		uintptr(hbm),
		uintptr(hpal),
		uintptr(unsafe.Pointer(bitmap)))

	return GpStatus(ret)
}

func GdipCreateHBITMAPFromBitmap(bitmap *GpBitmap, hbmReturn *HBITMAP, background ARGB) GpStatus {
	ret, _, _ := syscall.Syscall(gdipCreateHBITMAPFromBitmap, 3,
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(hbmReturn)),
		uintptr(background))

	return GpStatus(ret)
}

func GdipDisposeImage(image *GpImage) GpStatus {
	ret, _, _ := syscall.Syscall(gdipDisposeImage, 1,
		uintptr(unsafe.Pointer(image)),
		0,
		0)

	return GpStatus(ret)
}

func GdiplusShutdown() {
	syscall.Syscall(gdiplusShutdown, 1,
		token,
		0,
		0)
}

func GdiplusStartup(input *GdiplusStartupInput, output *GdiplusStartupOutput) GpStatus {
	ret, _, _ := syscall.Syscall(gdiplusStartup, 3,
		uintptr(unsafe.Pointer(&token)),
		uintptr(unsafe.Pointer(input)),
		uintptr(unsafe.Pointer(output)))

	return GpStatus(ret)
}

/*GdipSaveImageToFile(image *GpImage, filename *uint16, clsidEncoder *CLSID, encoderParams *EncoderParameters) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipSaveImageToFile, 4,
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(filename)),
		uintptr(unsafe.Pointer(clsidEncoder)),
		uintptr(unsafe.Pointer(encoderParams)),
		0,
		0)

	return GpStatus(ret)
}*/
func GdipCreateFromHDC(dc HDC, gg **GpGraphics) GpStatus {
	ret, _, _ := syscall.Syscall(gdipCreateFromHDC, 2,
		uintptr(dc),
		uintptr(unsafe.Pointer(gg)),
		0)
	return GpStatus(ret)
}
func GdipDeleteGraphics(gg *GpGraphics) GpStatus {
	ret, _, _ := syscall.Syscall(gdipDeleteGraphics, 1,
		uintptr(unsafe.Pointer(gg)),
		0,
		0)
	return GpStatus(ret)
}
func GdipDrawImage(gg *GpGraphics, img *GpImage, x, y int) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipDrawImage, 6,
		uintptr(unsafe.Pointer(gg)),
		uintptr(unsafe.Pointer(img)),
		uintptr(x),
		uintptr(y),
		0,
		0)
	return GpStatus(ret)
}
func GdipDrawImageRect(gg *GpGraphics, img *GpImage, x, y, w, h int) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipDrawImageRect, 6,
		uintptr(unsafe.Pointer(gg)),
		uintptr(unsafe.Pointer(img)),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h))
	return GpStatus(ret)
}
func GdipDrawImagePointRect(gg *GpGraphics, img *GpImage, px, py, x, y, w, h int) GpStatus {
	ret, _, _ := syscall.Syscall9(gdipDrawImagePointRect, 9,
		uintptr(unsafe.Pointer(gg)),
		uintptr(unsafe.Pointer(img)),
		uintptr(px),
		uintptr(py),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h),
		2)
	return GpStatus(ret)
}
func GdipDrawLine(gg *GpGraphics, pen *GpPen, x1, y1, x2, y2 int) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipDrawLine, 6,
		uintptr(unsafe.Pointer(gg)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2))
	return GpStatus(ret)
}
func GdipDrawLineF(gg *GpGraphics, pen *GpPen, x1, y1, x2, y2 float32) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipDrawLineF, 6,
		uintptr(unsafe.Pointer(gg)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2))
	return GpStatus(ret)
}
func GdipDrawLines(gg *GpGraphics, pen *GpPen, pts []GpPoint, count int) GpStatus {
	ppt := &pts[0]
	ret, _, _ := syscall.Syscall6(gdipDrawLines, 4,
		uintptr(unsafe.Pointer(gg)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(ppt)),
		uintptr(count),
		0,
		0)
	return GpStatus(ret)
}
func GdipDrawLinesF(gg *GpGraphics, pen *GpPen, pts []GpPointF, count int) GpStatus {
	ppt := &pts[0]
	ret, _, _ := syscall.Syscall6(gdipDrawLinesF, 4,
		uintptr(unsafe.Pointer(gg)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(ppt)),
		uintptr(count),
		0,
		0)
	return GpStatus(ret)
}
func GdipDrawRect(gg *GpGraphics, pen *GpPen, x, y, w, h int) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipDrawRect, 6,
		uintptr(unsafe.Pointer(gg)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h))
	return GpStatus(ret)
}
func GdipDrawRectF(gg *GpGraphics, pen *GpPen, x, y, w, h float32) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipDrawRectF, 6,
		uintptr(unsafe.Pointer(gg)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h))
	return GpStatus(ret)
}
func GdipFillRect(gg *GpGraphics, brush *GpBrush, x, y, w, h int) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipFillRect, 6,
		uintptr(unsafe.Pointer(gg)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h))
	return GpStatus(ret)
}
func GdipFillRectF(gg *GpGraphics, brush *GpBrush, x, y, w, h float32) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipFillRectF, 6,
		uintptr(unsafe.Pointer(gg)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h))
	return GpStatus(ret)
}
func GdipCreatePen(color ARGB, w float32, pen **GpPen) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipCreatePen, 4,
		uintptr(color),
		uintptr(w),
		uintptr(2),
		uintptr(unsafe.Pointer(pen)),
		0,
		0)
	return GpStatus(ret)
}
func GdipDeletePen(pen *GpPen) GpStatus {
	ret, _, _ := syscall.Syscall(gdipDeletePen, 1,
		uintptr(unsafe.Pointer(pen)),
		0,
		0)
	return GpStatus(ret)
}
func GdipSetPenWidth(pen *GpPen, w float32) GpStatus {
	ret, _, _ := syscall.Syscall(gdipSetPenWidth, 2,
		uintptr(unsafe.Pointer(pen)),
		uintptr(w),
		0)
	return GpStatus(ret)
}
func GdipCreateSolidBrush(color ARGB, brush **GpBrush) GpStatus {
	ret, _, _ := syscall.Syscall(gdipCreateSolidBrush, 2,
		uintptr(color),
		uintptr(unsafe.Pointer(brush)),
		0)
	return GpStatus(ret)
}
func GdipCreateLineBrush(pt1, pt2 *GpPointF, color1, color2 ARGB, m int, brush **GpBrush) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipCreateLineBrush, 6,
		uintptr(unsafe.Pointer(pt1)),
		uintptr(unsafe.Pointer(pt2)),
		uintptr(color1),
		uintptr(color2),
		uintptr(m),
		uintptr(unsafe.Pointer(brush)))
	return GpStatus(ret)
}
func GdipCreateHatchBrush(style int, fore, back ARGB, brush **GpBrush) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipCreateHatchBrush, 4,
		uintptr(style),
		uintptr(fore),
		uintptr(back),
		uintptr(unsafe.Pointer(brush)),
		0,
		0)
	return GpStatus(ret)
}
func GdipDeleteBrush(brush *GpBrush) GpStatus {
	ret, _, _ := syscall.Syscall(gdipDeleteBrush, 1,
		uintptr(unsafe.Pointer(brush)),
		0,
		0)
	return GpStatus(ret)
}
func GdipDrawString(gg *GpGraphics, str string, font *GpFont, layoutRect *GpRectF, fmt *GpStringFormat, brush *GpBrush) GpStatus {
	txt := unsafe.Pointer(syscall.StringToUTF16Ptr(str))
	ret, _, _ := syscall.Syscall9(gdipDrawString, 7,
		uintptr(unsafe.Pointer(gg)),
		uintptr(txt),
		uintptr(len(str)),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(fmt)),
		uintptr(unsafe.Pointer(brush)),
		0,
		0)
	return GpStatus(ret)
}
func GdipMeasureCharacterRanges(gg *GpGraphics, str string, font *GpFont, layoutRect *GpRectF, fmt *GpStringFormat, rgns []GpRegion) GpStatus {
	txt := unsafe.Pointer(syscall.StringToUTF16Ptr(str))
	prgn := &rgns[0]
	ret, _, _ := syscall.Syscall9(gdipMeasureCharacterRanges, 8,
		uintptr(unsafe.Pointer(gg)),
		uintptr(txt),
		uintptr(len(str)),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(fmt)),
		uintptr(len(rgns)),
		uintptr(unsafe.Pointer(prgn)),
		0)
	return GpStatus(ret)
}
func GdipCreateStringFormat(flags int, fmt **GpStringFormat) GpStatus {
	ret, _, _ := syscall.Syscall(gdipCreateStringFormat, 3,
		uintptr(flags),
		uintptr(0x00),
		uintptr(unsafe.Pointer(fmt)))
	return GpStatus(ret)
}
func GdipStringFormatGetGenericTypographic(fmt **GpStringFormat) GpStatus {
	ret, _, _ := syscall.Syscall(gdipStringFormatGetGenericTypographic, 1,
		uintptr(unsafe.Pointer(fmt)),
		0,
		0)
	return GpStatus(ret)
}
func GdipDeleteStringFormat(fmt *GpStringFormat) GpStatus {
	ret, _, _ := syscall.Syscall(gdipDeleteStringFormat, 1,
		uintptr(unsafe.Pointer(fmt)),
		0,
		0)
	return GpStatus(ret)
}
func GdipStringFormatGetGenericDefault(fmt **GpStringFormat) GpStatus {
	ret, _, _ := syscall.Syscall(gdipStringFormatGetGenericDefault, 1,
		uintptr(unsafe.Pointer(fmt)),
		0,
		0)
	return GpStatus(ret)
}
func GdipSetStringFormatFlags(fmt *GpStringFormat, flags int) GpStatus {
	ret, _, _ := syscall.Syscall(gdipSetStringFormatFlags, 2,
		uintptr(unsafe.Pointer(fmt)),
		uintptr(flags),
		0)
	return GpStatus(ret)
}
func GdipGetStringFormatFlags(fmt *GpStringFormat, flags *int) GpStatus {
	ret, _, _ := syscall.Syscall(gdipGetStringFormatFlags, 2,
		uintptr(unsafe.Pointer(fmt)),
		uintptr(unsafe.Pointer(flags)),
		0)
	return GpStatus(ret)
}
func GdipSetStringFormatAlign(fmt *GpStringFormat, align int) GpStatus {
	ret, _, _ := syscall.Syscall(gdipSetStringFormatAlign, 2,
		uintptr(unsafe.Pointer(fmt)),
		uintptr(align),
		0)
	return GpStatus(ret)
}
func GdipSetStringFormatMeasurableCharacterRanges(fmt *GpStringFormat, ranges []GpCharacterRange) GpStatus {
	prange := &ranges[0]
	ret, _, _ := syscall.Syscall(gdipSetStringFormatMeasurableCharacterRanges, 3,
		uintptr(unsafe.Pointer(fmt)),
		uintptr(len(ranges)),
		uintptr(unsafe.Pointer(prange)))
	return GpStatus(ret)
}
func GdipCreateFontFamily(name string, fm **GpFontFamily) GpStatus {
	pname := unsafe.Pointer(syscall.StringToUTF16Ptr(name))
	ret, _, _ := syscall.Syscall(gdipCreateFontFamily, 3,
		uintptr(pname),
		uintptr(0),
		uintptr(unsafe.Pointer(fm)))
	return GpStatus(ret)
}
func GdipDeleteFontFamily(fm *GpFontFamily) GpStatus {
	ret, _, _ := syscall.Syscall(gdipDeleteFontFamily, 1,
		uintptr(unsafe.Pointer(fm)),
		0,
		0)
	return GpStatus(ret)
}
func GdipCreateFontFromDC(hdc HDC, font **GpFont) GpStatus {
	ret, _, _ := syscall.Syscall(gdipCreateFontFromDC, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(font)),
		0)
	return GpStatus(ret)
}
func GdipCreateFont(fm *GpFontFamily, size, style int, font **GpFont) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipCreateFont, 5,
		uintptr(unsafe.Pointer(fm)),
		uintptr(float32(size)),
		uintptr(style),
		uintptr(2),
		uintptr(unsafe.Pointer(font)),
		0)
	return GpStatus(ret)
}
func GdipDeleteFont(font *GpFont) GpStatus {
	ret, _, _ := syscall.Syscall(gdipDeleteFont, 1,
		uintptr(unsafe.Pointer(font)),
		0,
		0)
	return GpStatus(ret)
}
func GdipGetFontStyle(font *GpFont, style *int) GpStatus {
	ret, _, _ := syscall.Syscall(gdipGetFontStyle, 2,
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(style)),
		0)
	return GpStatus(ret)
}
func GdipGetFontSize(font *GpFont, size *int) GpStatus {
	ret, _, _ := syscall.Syscall(gdipGetFontSize, 2,
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(size)),
		0)
	return GpStatus(ret)
}
func GdipSetClipRect(gg *GpGraphics, x, y, w, h int, mode int) GpStatus {
	ret, _, _ := syscall.Syscall6(gdipSetClipRect, 6,
		uintptr(unsafe.Pointer(gg)),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h),
		uintptr(mode))
	return GpStatus(ret)
}
func GdipResetClip(gg *GpGraphics, x, y, w, h int) GpStatus {
	ret, _, _ := syscall.Syscall(gdipResetClip, 1,
		uintptr(unsafe.Pointer(gg)),
		0,
		0)
	return GpStatus(ret)
}
