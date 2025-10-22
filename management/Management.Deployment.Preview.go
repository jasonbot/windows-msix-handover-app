package winrt

import (
	"syscall"
	"unsafe"

	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-winrtapi/winrt"
)

// structs

type DeploymentPreviewContract struct {
}

// interfaces

// E2FAD668-882C-4F33-B035-0DF7B90D67E6
var IID_IClassicAppManagerStatics = syscall.GUID{0xE2FAD668, 0x882C, 0x4F33,
	[8]byte{0xB0, 0x35, 0x0D, 0xF7, 0xB9, 0x0D, 0x67, 0xE6}}

type IClassicAppManagerStaticsInterface interface {
	win32.IInspectableInterface
	FindInstalledApp(appUninstallKey string) *IInstalledClassicAppInfo
}

type IClassicAppManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	FindInstalledApp uintptr
}

type IClassicAppManagerStatics struct {
	win32.IInspectable
}

func (this *IClassicAppManagerStatics) Vtbl() *IClassicAppManagerStaticsVtbl {
	return (*IClassicAppManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IClassicAppManagerStatics) FindInstalledApp(appUninstallKey string) *IInstalledClassicAppInfo {
	var _result *IInstalledClassicAppInfo
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindInstalledApp, uintptr(unsafe.Pointer(this)), winrt.NewHStr(appUninstallKey).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0A7D3DA3-65D0-4086-80D6-0610D760207D
var IID_IInstalledClassicAppInfo = syscall.GUID{0x0A7D3DA3, 0x65D0, 0x4086,
	[8]byte{0x80, 0xD6, 0x06, 0x10, 0xD7, 0x60, 0x20, 0x7D}}

type IInstalledClassicAppInfoInterface interface {
	win32.IInspectableInterface
	Get_DisplayName() string
	Get_DisplayVersion() string
}

type IInstalledClassicAppInfoVtbl struct {
	win32.IInspectableVtbl
	Get_DisplayName    uintptr
	Get_DisplayVersion uintptr
}

type IInstalledClassicAppInfo struct {
	win32.IInspectable
}

func (this *IInstalledClassicAppInfo) Vtbl() *IInstalledClassicAppInfoVtbl {
	return (*IInstalledClassicAppInfoVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IInstalledClassicAppInfo) Get_DisplayName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return winrt.HStringToStrAndFree(_result)
}

func (this *IInstalledClassicAppInfo) Get_DisplayVersion() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DisplayVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return winrt.HStringToStrAndFree(_result)
}

// classes

type ClassicAppManager struct {
	winrt.RtClass
}

func NewIClassicAppManagerStatics() *IClassicAppManagerStatics {
	var p *IClassicAppManagerStatics
	hs := winrt.NewHStr("Windows.Management.Deployment.Preview.ClassicAppManager")
	hr := win32.RoGetActivationFactory(hs.Ptr, &IID_IClassicAppManagerStatics, unsafe.Pointer(&p))
	win32.ASSERT_SUCCEEDED(hr)
	com.AddToScope(p)
	return p
}

type InstalledClassicAppInfo struct {
	winrt.RtClass
	*IInstalledClassicAppInfo
}
