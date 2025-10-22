package winrt

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-winrtapi/winrt"
)

// enums

// enum
// flags
type AddPackageByAppInstallerOptions uint32

const (
	AddPackageByAppInstallerOptions_None                     AddPackageByAppInstallerOptions = 0
	AddPackageByAppInstallerOptions_InstallAllResources      AddPackageByAppInstallerOptions = 32
	AddPackageByAppInstallerOptions_ForceTargetAppShutdown   AddPackageByAppInstallerOptions = 64
	AddPackageByAppInstallerOptions_RequiredContentGroupOnly AddPackageByAppInstallerOptions = 256
	AddPackageByAppInstallerOptions_LimitToExistingPackages  AddPackageByAppInstallerOptions = 512
)

// enum
// flags
type DeploymentOptions uint32

const (
	DeploymentOptions_None                           DeploymentOptions = 0
	DeploymentOptions_ForceApplicationShutdown       DeploymentOptions = 1
	DeploymentOptions_DevelopmentMode                DeploymentOptions = 2
	DeploymentOptions_InstallAllResources            DeploymentOptions = 32
	DeploymentOptions_ForceTargetApplicationShutdown DeploymentOptions = 64
	DeploymentOptions_RequiredContentGroupOnly       DeploymentOptions = 256
	DeploymentOptions_ForceUpdateFromAnyVersion      DeploymentOptions = 262144
	DeploymentOptions_RetainFilesOnFailure           DeploymentOptions = 2097152
	DeploymentOptions_StageInPlace                   DeploymentOptions = 4194304
)

// enum
type DeploymentProgressState int32

const (
	DeploymentProgressState_Queued     DeploymentProgressState = 0
	DeploymentProgressState_Processing DeploymentProgressState = 1
)

// enum
type PackageInstallState int32

const (
	PackageInstallState_NotInstalled PackageInstallState = 0
	PackageInstallState_Staged       PackageInstallState = 1
	PackageInstallState_Installed    PackageInstallState = 2
	PackageInstallState_Paused       PackageInstallState = 6
)

// enum
type PackageState int32

const (
	PackageState_Normal         PackageState = 0
	PackageState_LicenseInvalid PackageState = 1
	PackageState_Modified       PackageState = 2
	PackageState_Tampered       PackageState = 3
)

// enum
// flags
type PackageStatus uint32

const (
	PackageStatus_OK           PackageStatus = 0
	PackageStatus_LicenseIssue PackageStatus = 1
	PackageStatus_Modified     PackageStatus = 2
	PackageStatus_Tampered     PackageStatus = 4
	PackageStatus_Disabled     PackageStatus = 8
)

// enum
type PackageStubPreference int32

const (
	PackageStubPreference_Full PackageStubPreference = 0
	PackageStubPreference_Stub PackageStubPreference = 1
)

// enum
// flags
type PackageTypes uint32

const (
	PackageTypes_None      PackageTypes = 0
	PackageTypes_Main      PackageTypes = 1
	PackageTypes_Framework PackageTypes = 2
	PackageTypes_Resource  PackageTypes = 4
	PackageTypes_Bundle    PackageTypes = 8
	PackageTypes_Xap       PackageTypes = 16
	PackageTypes_Optional  PackageTypes = 32
	PackageTypes_All       PackageTypes = 4294967295
)

// enum
// flags
type RemovalOptions uint32

const (
	RemovalOptions_None                            RemovalOptions = 0
	RemovalOptions_PreserveApplicationData         RemovalOptions = 4096
	RemovalOptions_PreserveRoamableApplicationData RemovalOptions = 128
	RemovalOptions_RemoveForAllUsers               RemovalOptions = 524288
)

// enum
type SharedPackageContainerCreationCollisionOptions int32

const (
	SharedPackageContainerCreationCollisionOptions_FailIfExists      SharedPackageContainerCreationCollisionOptions = 0
	SharedPackageContainerCreationCollisionOptions_MergeWithExisting SharedPackageContainerCreationCollisionOptions = 1
	SharedPackageContainerCreationCollisionOptions_ReplaceExisting   SharedPackageContainerCreationCollisionOptions = 2
)

// enum
type SharedPackageContainerOperationStatus int32

const (
	SharedPackageContainerOperationStatus_Success                               SharedPackageContainerOperationStatus = 0
	SharedPackageContainerOperationStatus_BlockedByPolicy                       SharedPackageContainerOperationStatus = 1
	SharedPackageContainerOperationStatus_AlreadyExists                         SharedPackageContainerOperationStatus = 2
	SharedPackageContainerOperationStatus_PackageFamilyExistsInAnotherContainer SharedPackageContainerOperationStatus = 3
	SharedPackageContainerOperationStatus_NotFound                              SharedPackageContainerOperationStatus = 4
	SharedPackageContainerOperationStatus_UnknownFailure                        SharedPackageContainerOperationStatus = 5
)

// enum
type StubPackageOption int32

const (
	StubPackageOption_Default       StubPackageOption = 0
	StubPackageOption_InstallFull   StubPackageOption = 1
	StubPackageOption_InstallStub   StubPackageOption = 2
	StubPackageOption_UsePreference StubPackageOption = 3
)

// structs

type DeploymentProgress struct {
	State      DeploymentProgressState
	Percentage uint32
}

type SharedPackageContainerContract struct {
}

// interfaces

// 05CEE018-F68F-422B-95A4-66679EC77FC0
var IID_IAddPackageOptions = syscall.GUID{0x05CEE018, 0xF68F, 0x422B,
	[8]byte{0x95, 0xA4, 0x66, 0x67, 0x9E, 0xC7, 0x7F, 0xC0}}

type IAddPackageOptionsInterface interface {
	win32.IInspectableInterface
	Get_DependencyPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass]
	Get_TargetVolume() *IPackageVolume
	Put_TargetVolume(value *IPackageVolume)
	Get_OptionalPackageFamilyNames() *winrt.IVector[string]
	Get_OptionalPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass]
	Get_RelatedPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass]
	Get_ExternalLocationUri() *winrt.IUriRuntimeClass
	Put_ExternalLocationUri(value *winrt.IUriRuntimeClass)
	Get_StubPackageOption() StubPackageOption
	Put_StubPackageOption(value StubPackageOption)
	Get_DeveloperMode() bool
	Put_DeveloperMode(value bool)
	Get_ForceAppShutdown() bool
	Put_ForceAppShutdown(value bool)
	Get_ForceTargetAppShutdown() bool
	Put_ForceTargetAppShutdown(value bool)
	Get_ForceUpdateFromAnyVersion() bool
	Put_ForceUpdateFromAnyVersion(value bool)
	Get_InstallAllResources() bool
	Put_InstallAllResources(value bool)
	Get_RequiredContentGroupOnly() bool
	Put_RequiredContentGroupOnly(value bool)
	Get_RetainFilesOnFailure() bool
	Put_RetainFilesOnFailure(value bool)
	Get_StageInPlace() bool
	Put_StageInPlace(value bool)
	Get_AllowUnsigned() bool
	Put_AllowUnsigned(value bool)
	Get_DeferRegistrationWhenPackagesAreInUse() bool
	Put_DeferRegistrationWhenPackagesAreInUse(value bool)
}

type IAddPackageOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_DependencyPackageUris                 uintptr
	Get_TargetVolume                          uintptr
	Put_TargetVolume                          uintptr
	Get_OptionalPackageFamilyNames            uintptr
	Get_OptionalPackageUris                   uintptr
	Get_RelatedPackageUris                    uintptr
	Get_ExternalLocationUri                   uintptr
	Put_ExternalLocationUri                   uintptr
	Get_StubPackageOption                     uintptr
	Put_StubPackageOption                     uintptr
	Get_DeveloperMode                         uintptr
	Put_DeveloperMode                         uintptr
	Get_ForceAppShutdown                      uintptr
	Put_ForceAppShutdown                      uintptr
	Get_ForceTargetAppShutdown                uintptr
	Put_ForceTargetAppShutdown                uintptr
	Get_ForceUpdateFromAnyVersion             uintptr
	Put_ForceUpdateFromAnyVersion             uintptr
	Get_InstallAllResources                   uintptr
	Put_InstallAllResources                   uintptr
	Get_RequiredContentGroupOnly              uintptr
	Put_RequiredContentGroupOnly              uintptr
	Get_RetainFilesOnFailure                  uintptr
	Put_RetainFilesOnFailure                  uintptr
	Get_StageInPlace                          uintptr
	Put_StageInPlace                          uintptr
	Get_AllowUnsigned                         uintptr
	Put_AllowUnsigned                         uintptr
	Get_DeferRegistrationWhenPackagesAreInUse uintptr
	Put_DeferRegistrationWhenPackagesAreInUse uintptr
}

type IAddPackageOptions struct {
	win32.IInspectable
}

func (this *IAddPackageOptions) Vtbl() *IAddPackageOptionsVtbl {
	return (*IAddPackageOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAddPackageOptions) Get_DependencyPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass] {
	var _result *winrt.IVector[*winrt.IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DependencyPackageUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAddPackageOptions) Get_TargetVolume() *IPackageVolume {
	var _result *IPackageVolume
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TargetVolume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAddPackageOptions) Put_TargetVolume(value *IPackageVolume) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TargetVolume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAddPackageOptions) Get_OptionalPackageFamilyNames() *winrt.IVector[string] {
	var _result *winrt.IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OptionalPackageFamilyNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAddPackageOptions) Get_OptionalPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass] {
	var _result *winrt.IVector[*winrt.IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OptionalPackageUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAddPackageOptions) Get_RelatedPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass] {
	var _result *winrt.IVector[*winrt.IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelatedPackageUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAddPackageOptions) Get_ExternalLocationUri() *winrt.IUriRuntimeClass {
	var _result *winrt.IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExternalLocationUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAddPackageOptions) Put_ExternalLocationUri(value *winrt.IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExternalLocationUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAddPackageOptions) Get_StubPackageOption() StubPackageOption {
	var _result StubPackageOption
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StubPackageOption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAddPackageOptions) Put_StubPackageOption(value StubPackageOption) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StubPackageOption, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAddPackageOptions) Get_DeveloperMode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeveloperMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAddPackageOptions) Put_DeveloperMode(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DeveloperMode, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAddPackageOptions) Get_ForceAppShutdown() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAddPackageOptions) Put_ForceAppShutdown(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForceAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAddPackageOptions) Get_ForceTargetAppShutdown() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceTargetAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAddPackageOptions) Put_ForceTargetAppShutdown(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForceTargetAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAddPackageOptions) Get_ForceUpdateFromAnyVersion() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceUpdateFromAnyVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAddPackageOptions) Put_ForceUpdateFromAnyVersion(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForceUpdateFromAnyVersion, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAddPackageOptions) Get_InstallAllResources() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstallAllResources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAddPackageOptions) Put_InstallAllResources(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InstallAllResources, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAddPackageOptions) Get_RequiredContentGroupOnly() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequiredContentGroupOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAddPackageOptions) Put_RequiredContentGroupOnly(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequiredContentGroupOnly, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAddPackageOptions) Get_RetainFilesOnFailure() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RetainFilesOnFailure, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAddPackageOptions) Put_RetainFilesOnFailure(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RetainFilesOnFailure, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAddPackageOptions) Get_StageInPlace() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StageInPlace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAddPackageOptions) Put_StageInPlace(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StageInPlace, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAddPackageOptions) Get_AllowUnsigned() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowUnsigned, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAddPackageOptions) Put_AllowUnsigned(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowUnsigned, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAddPackageOptions) Get_DeferRegistrationWhenPackagesAreInUse() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeferRegistrationWhenPackagesAreInUse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAddPackageOptions) Put_DeferRegistrationWhenPackagesAreInUse(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DeferRegistrationWhenPackagesAreInUse, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// E7EE21C3-2103-53EE-9B18-68AFEAB0033D
var IID_IAppInstallerManager = syscall.GUID{0xE7EE21C3, 0x2103, 0x53EE,
	[8]byte{0x9B, 0x18, 0x68, 0xAF, 0xEA, 0xB0, 0x03, 0x3D}}

type IAppInstallerManagerInterface interface {
	win32.IInspectableInterface
	SetAutoUpdateSettings(packageFamilyName string, appInstallerInfo *IAutoUpdateSettingsOptions)
	ClearAutoUpdateSettings(packageFamilyName string)
	PauseAutoUpdatesUntil(packageFamilyName string, dateTime winrt.DateTime)
}

type IAppInstallerManagerVtbl struct {
	win32.IInspectableVtbl
	SetAutoUpdateSettings   uintptr
	ClearAutoUpdateSettings uintptr
	PauseAutoUpdatesUntil   uintptr
}

type IAppInstallerManager struct {
	win32.IInspectable
}

func (this *IAppInstallerManager) Vtbl() *IAppInstallerManagerVtbl {
	return (*IAppInstallerManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppInstallerManager) SetAutoUpdateSettings(packageFamilyName string, appInstallerInfo *IAutoUpdateSettingsOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetAutoUpdateSettings, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(appInstallerInfo)))
	_ = _hr
}

func (this *IAppInstallerManager) ClearAutoUpdateSettings(packageFamilyName string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearAutoUpdateSettings, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFamilyName).Ptr)
	_ = _hr
}

func (this *IAppInstallerManager) PauseAutoUpdatesUntil(packageFamilyName string, dateTime winrt.DateTime) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().PauseAutoUpdatesUntil, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFamilyName).Ptr, *(*uintptr)(unsafe.Pointer(&dateTime)))
	_ = _hr
}

// C95A6ED5-FC59-5336-9B2E-2B07C5E61434
var IID_IAppInstallerManagerStatics = syscall.GUID{0xC95A6ED5, 0xFC59, 0x5336,
	[8]byte{0x9B, 0x2E, 0x2B, 0x07, 0xC5, 0xE6, 0x14, 0x34}}

type IAppInstallerManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *IAppInstallerManager
	GetForSystem() *IAppInstallerManager
}

type IAppInstallerManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault   uintptr
	GetForSystem uintptr
}

type IAppInstallerManagerStatics struct {
	win32.IInspectable
}

func (this *IAppInstallerManagerStatics) Vtbl() *IAppInstallerManagerStaticsVtbl {
	return (*IAppInstallerManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAppInstallerManagerStatics) GetDefault() *IAppInstallerManager {
	var _result *IAppInstallerManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAppInstallerManagerStatics) GetForSystem() *IAppInstallerManager {
	var _result *IAppInstallerManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForSystem, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 67491D87-35E1-512A-8968-1AE88D1BE6D3
var IID_IAutoUpdateSettingsOptions = syscall.GUID{0x67491D87, 0x35E1, 0x512A,
	[8]byte{0x89, 0x68, 0x1A, 0xE8, 0x8D, 0x1B, 0xE6, 0xD3}}

type IAutoUpdateSettingsOptionsInterface interface {
	win32.IInspectableInterface
	Get_Version() winrt.PackageVersion
	Put_Version(value winrt.PackageVersion)
	Get_AppInstallerUri() *winrt.IUriRuntimeClass
	Put_AppInstallerUri(value *winrt.IUriRuntimeClass)
	Get_OnLaunch() bool
	Put_OnLaunch(value bool)
	Get_HoursBetweenUpdateChecks() uint32
	Put_HoursBetweenUpdateChecks(value uint32)
	Get_ShowPrompt() bool
	Put_ShowPrompt(value bool)
	Get_UpdateBlocksActivation() bool
	Put_UpdateBlocksActivation(value bool)
	Get_AutomaticBackgroundTask() bool
	Put_AutomaticBackgroundTask(value bool)
	Get_ForceUpdateFromAnyVersion() bool
	Put_ForceUpdateFromAnyVersion(value bool)
	Get_IsAutoRepairEnabled() bool
	Put_IsAutoRepairEnabled(value bool)
	Get_UpdateUris() *winrt.IVector[*winrt.IUriRuntimeClass]
	Get_RepairUris() *winrt.IVector[*winrt.IUriRuntimeClass]
	Get_DependencyPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass]
	Get_OptionalPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass]
}

type IAutoUpdateSettingsOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_Version                   uintptr
	Put_Version                   uintptr
	Get_AppInstallerUri           uintptr
	Put_AppInstallerUri           uintptr
	Get_OnLaunch                  uintptr
	Put_OnLaunch                  uintptr
	Get_HoursBetweenUpdateChecks  uintptr
	Put_HoursBetweenUpdateChecks  uintptr
	Get_ShowPrompt                uintptr
	Put_ShowPrompt                uintptr
	Get_UpdateBlocksActivation    uintptr
	Put_UpdateBlocksActivation    uintptr
	Get_AutomaticBackgroundTask   uintptr
	Put_AutomaticBackgroundTask   uintptr
	Get_ForceUpdateFromAnyVersion uintptr
	Put_ForceUpdateFromAnyVersion uintptr
	Get_IsAutoRepairEnabled       uintptr
	Put_IsAutoRepairEnabled       uintptr
	Get_UpdateUris                uintptr
	Get_RepairUris                uintptr
	Get_DependencyPackageUris     uintptr
	Get_OptionalPackageUris       uintptr
}

type IAutoUpdateSettingsOptions struct {
	win32.IInspectable
}

func (this *IAutoUpdateSettingsOptions) Vtbl() *IAutoUpdateSettingsOptionsVtbl {
	return (*IAutoUpdateSettingsOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAutoUpdateSettingsOptions) Get_Version() winrt.PackageVersion {
	var _result winrt.PackageVersion
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Version, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAutoUpdateSettingsOptions) Put_Version(value winrt.PackageVersion) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Version, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	_ = _hr
}

func (this *IAutoUpdateSettingsOptions) Get_AppInstallerUri() *winrt.IUriRuntimeClass {
	var _result *winrt.IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppInstallerUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAutoUpdateSettingsOptions) Put_AppInstallerUri(value *winrt.IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AppInstallerUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IAutoUpdateSettingsOptions) Get_OnLaunch() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OnLaunch, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAutoUpdateSettingsOptions) Put_OnLaunch(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_OnLaunch, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAutoUpdateSettingsOptions) Get_HoursBetweenUpdateChecks() uint32 {
	var _result uint32
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_HoursBetweenUpdateChecks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAutoUpdateSettingsOptions) Put_HoursBetweenUpdateChecks(value uint32) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_HoursBetweenUpdateChecks, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IAutoUpdateSettingsOptions) Get_ShowPrompt() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ShowPrompt, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAutoUpdateSettingsOptions) Put_ShowPrompt(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ShowPrompt, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAutoUpdateSettingsOptions) Get_UpdateBlocksActivation() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpdateBlocksActivation, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAutoUpdateSettingsOptions) Put_UpdateBlocksActivation(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_UpdateBlocksActivation, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAutoUpdateSettingsOptions) Get_AutomaticBackgroundTask() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AutomaticBackgroundTask, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAutoUpdateSettingsOptions) Put_AutomaticBackgroundTask(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AutomaticBackgroundTask, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAutoUpdateSettingsOptions) Get_ForceUpdateFromAnyVersion() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceUpdateFromAnyVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAutoUpdateSettingsOptions) Put_ForceUpdateFromAnyVersion(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForceUpdateFromAnyVersion, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAutoUpdateSettingsOptions) Get_IsAutoRepairEnabled() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAutoRepairEnabled, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IAutoUpdateSettingsOptions) Put_IsAutoRepairEnabled(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_IsAutoRepairEnabled, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IAutoUpdateSettingsOptions) Get_UpdateUris() *winrt.IVector[*winrt.IUriRuntimeClass] {
	var _result *winrt.IVector[*winrt.IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UpdateUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAutoUpdateSettingsOptions) Get_RepairUris() *winrt.IVector[*winrt.IUriRuntimeClass] {
	var _result *winrt.IVector[*winrt.IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RepairUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAutoUpdateSettingsOptions) Get_DependencyPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass] {
	var _result *winrt.IVector[*winrt.IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DependencyPackageUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IAutoUpdateSettingsOptions) Get_OptionalPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass] {
	var _result *winrt.IVector[*winrt.IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OptionalPackageUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 887B337D-0C05-54D0-BD49-3BB7A2C084CB
var IID_IAutoUpdateSettingsOptionsStatics = syscall.GUID{0x887B337D, 0x0C05, 0x54D0,
	[8]byte{0xBD, 0x49, 0x3B, 0xB7, 0xA2, 0xC0, 0x84, 0xCB}}

type IAutoUpdateSettingsOptionsStaticsInterface interface {
	win32.IInspectableInterface
	CreateFromAppInstallerInfo(appInstallerInfo *winrt.IAppInstallerInfo) *IAutoUpdateSettingsOptions
}

type IAutoUpdateSettingsOptionsStaticsVtbl struct {
	win32.IInspectableVtbl
	CreateFromAppInstallerInfo uintptr
}

type IAutoUpdateSettingsOptionsStatics struct {
	win32.IInspectable
}

func (this *IAutoUpdateSettingsOptionsStatics) Vtbl() *IAutoUpdateSettingsOptionsStaticsVtbl {
	return (*IAutoUpdateSettingsOptionsStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IAutoUpdateSettingsOptionsStatics) CreateFromAppInstallerInfo(appInstallerInfo *winrt.IAppInstallerInfo) *IAutoUpdateSettingsOptions {
	var _result *IAutoUpdateSettingsOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateFromAppInstallerInfo, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(appInstallerInfo)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// C2AB6ECE-F664-5C8E-A4B3-2A33276D3DDE
var IID_ICreateSharedPackageContainerOptions = syscall.GUID{0xC2AB6ECE, 0xF664, 0x5C8E,
	[8]byte{0xA4, 0xB3, 0x2A, 0x33, 0x27, 0x6D, 0x3D, 0xDE}}

type ICreateSharedPackageContainerOptionsInterface interface {
	win32.IInspectableInterface
	Get_Members() *winrt.IVector[*ISharedPackageContainerMember]
	Get_ForceAppShutdown() bool
	Put_ForceAppShutdown(value bool)
	Get_CreateCollisionOption() SharedPackageContainerCreationCollisionOptions
	Put_CreateCollisionOption(value SharedPackageContainerCreationCollisionOptions)
}

type ICreateSharedPackageContainerOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_Members               uintptr
	Get_ForceAppShutdown      uintptr
	Put_ForceAppShutdown      uintptr
	Get_CreateCollisionOption uintptr
	Put_CreateCollisionOption uintptr
}

type ICreateSharedPackageContainerOptions struct {
	win32.IInspectable
}

func (this *ICreateSharedPackageContainerOptions) Vtbl() *ICreateSharedPackageContainerOptionsVtbl {
	return (*ICreateSharedPackageContainerOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateSharedPackageContainerOptions) Get_Members() *winrt.IVector[*ISharedPackageContainerMember] {
	var _result *winrt.IVector[*ISharedPackageContainerMember]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Members, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICreateSharedPackageContainerOptions) Get_ForceAppShutdown() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICreateSharedPackageContainerOptions) Put_ForceAppShutdown(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForceAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *ICreateSharedPackageContainerOptions) Get_CreateCollisionOption() SharedPackageContainerCreationCollisionOptions {
	var _result SharedPackageContainerCreationCollisionOptions
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_CreateCollisionOption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICreateSharedPackageContainerOptions) Put_CreateCollisionOption(value SharedPackageContainerCreationCollisionOptions) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_CreateCollisionOption, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

// CE8810BF-151C-5707-B936-497E564AFC7A
var IID_ICreateSharedPackageContainerResult = syscall.GUID{0xCE8810BF, 0x151C, 0x5707,
	[8]byte{0xB9, 0x36, 0x49, 0x7E, 0x56, 0x4A, 0xFC, 0x7A}}

type ICreateSharedPackageContainerResultInterface interface {
	win32.IInspectableInterface
	Get_Container() *ISharedPackageContainer
	Get_Status() SharedPackageContainerOperationStatus
	Get_ExtendedError() winrt.HResult
}

type ICreateSharedPackageContainerResultVtbl struct {
	win32.IInspectableVtbl
	Get_Container     uintptr
	Get_Status        uintptr
	Get_ExtendedError uintptr
}

type ICreateSharedPackageContainerResult struct {
	win32.IInspectable
}

func (this *ICreateSharedPackageContainerResult) Vtbl() *ICreateSharedPackageContainerResultVtbl {
	return (*ICreateSharedPackageContainerResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ICreateSharedPackageContainerResult) Get_Container() *ISharedPackageContainer {
	var _result *ISharedPackageContainer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Container, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ICreateSharedPackageContainerResult) Get_Status() SharedPackageContainerOperationStatus {
	var _result SharedPackageContainerOperationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *ICreateSharedPackageContainerResult) Get_ExtendedError() winrt.HResult {
	var _result winrt.HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 9D81865F-986E-5138-8B5D-384D8E66ED6C
var IID_IDeleteSharedPackageContainerOptions = syscall.GUID{0x9D81865F, 0x986E, 0x5138,
	[8]byte{0x8B, 0x5D, 0x38, 0x4D, 0x8E, 0x66, 0xED, 0x6C}}

type IDeleteSharedPackageContainerOptionsInterface interface {
	win32.IInspectableInterface
	Get_ForceAppShutdown() bool
	Put_ForceAppShutdown(value bool)
	Get_AllUsers() bool
	Put_AllUsers(value bool)
}

type IDeleteSharedPackageContainerOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_ForceAppShutdown uintptr
	Put_ForceAppShutdown uintptr
	Get_AllUsers         uintptr
	Put_AllUsers         uintptr
}

type IDeleteSharedPackageContainerOptions struct {
	win32.IInspectable
}

func (this *IDeleteSharedPackageContainerOptions) Vtbl() *IDeleteSharedPackageContainerOptionsVtbl {
	return (*IDeleteSharedPackageContainerOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeleteSharedPackageContainerOptions) Get_ForceAppShutdown() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeleteSharedPackageContainerOptions) Put_ForceAppShutdown(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForceAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IDeleteSharedPackageContainerOptions) Get_AllUsers() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllUsers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeleteSharedPackageContainerOptions) Put_AllUsers(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllUsers, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 35398884-5736-517B-85BC-E598C81AB284
var IID_IDeleteSharedPackageContainerResult = syscall.GUID{0x35398884, 0x5736, 0x517B,
	[8]byte{0x85, 0xBC, 0xE5, 0x98, 0xC8, 0x1A, 0xB2, 0x84}}

type IDeleteSharedPackageContainerResultInterface interface {
	win32.IInspectableInterface
	Get_Status() SharedPackageContainerOperationStatus
	Get_ExtendedError() winrt.HResult
}

type IDeleteSharedPackageContainerResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status        uintptr
	Get_ExtendedError uintptr
}

type IDeleteSharedPackageContainerResult struct {
	win32.IInspectable
}

func (this *IDeleteSharedPackageContainerResult) Vtbl() *IDeleteSharedPackageContainerResultVtbl {
	return (*IDeleteSharedPackageContainerResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeleteSharedPackageContainerResult) Get_Status() SharedPackageContainerOperationStatus {
	var _result SharedPackageContainerOperationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeleteSharedPackageContainerResult) Get_ExtendedError() winrt.HResult {
	var _result winrt.HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 2563B9AE-B77D-4C1F-8A7B-20E6AD515EF3
var IID_IDeploymentResult = syscall.GUID{0x2563B9AE, 0xB77D, 0x4C1F,
	[8]byte{0x8A, 0x7B, 0x20, 0xE6, 0xAD, 0x51, 0x5E, 0xF3}}

type IDeploymentResultInterface interface {
	win32.IInspectableInterface
	Get_ErrorText() string
	Get_ActivityId() syscall.GUID
	Get_ExtendedErrorCode() winrt.HResult
}

type IDeploymentResultVtbl struct {
	win32.IInspectableVtbl
	Get_ErrorText         uintptr
	Get_ActivityId        uintptr
	Get_ExtendedErrorCode uintptr
}

type IDeploymentResult struct {
	win32.IInspectable
}

func (this *IDeploymentResult) Vtbl() *IDeploymentResultVtbl {
	return (*IDeploymentResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeploymentResult) Get_ErrorText() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ErrorText, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return winrt.HStringToStrAndFree(_result)
}

func (this *IDeploymentResult) Get_ActivityId() syscall.GUID {
	var _result syscall.GUID
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ActivityId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IDeploymentResult) Get_ExtendedErrorCode() winrt.HResult {
	var _result winrt.HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedErrorCode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// FC0E715C-5A01-4BD7-BCF1-381C8C82E04A
var IID_IDeploymentResult2 = syscall.GUID{0xFC0E715C, 0x5A01, 0x4BD7,
	[8]byte{0xBC, 0xF1, 0x38, 0x1C, 0x8C, 0x82, 0xE0, 0x4A}}

type IDeploymentResult2Interface interface {
	win32.IInspectableInterface
	Get_IsRegistered() bool
}

type IDeploymentResult2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsRegistered uintptr
}

type IDeploymentResult2 struct {
	win32.IInspectable
}

func (this *IDeploymentResult2) Vtbl() *IDeploymentResult2Vtbl {
	return (*IDeploymentResult2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IDeploymentResult2) Get_IsRegistered() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsRegistered, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// B40FC8FE-8384-54CC-817D-AE09D3B6A606
var IID_IFindSharedPackageContainerOptions = syscall.GUID{0xB40FC8FE, 0x8384, 0x54CC,
	[8]byte{0x81, 0x7D, 0xAE, 0x09, 0xD3, 0xB6, 0xA6, 0x06}}

type IFindSharedPackageContainerOptionsInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Put_Name(value string)
	Get_PackageFamilyName() string
	Put_PackageFamilyName(value string)
}

type IFindSharedPackageContainerOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_Name              uintptr
	Put_Name              uintptr
	Get_PackageFamilyName uintptr
	Put_PackageFamilyName uintptr
}

type IFindSharedPackageContainerOptions struct {
	win32.IInspectable
}

func (this *IFindSharedPackageContainerOptions) Vtbl() *IFindSharedPackageContainerOptionsVtbl {
	return (*IFindSharedPackageContainerOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IFindSharedPackageContainerOptions) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return winrt.HStringToStrAndFree(_result)
}

func (this *IFindSharedPackageContainerOptions) Put_Name(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_Name, uintptr(unsafe.Pointer(this)), winrt.NewHStr(value).Ptr)
	_ = _hr
}

func (this *IFindSharedPackageContainerOptions) Get_PackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return winrt.HStringToStrAndFree(_result)
}

func (this *IFindSharedPackageContainerOptions) Put_PackageFamilyName(value string) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_PackageFamilyName, uintptr(unsafe.Pointer(this)), winrt.NewHStr(value).Ptr)
	_ = _hr
}

// DA35AA22-1DE0-5D3E-99FF-D24F3118BF5E
var IID_IPackageAllUserProvisioningOptions = syscall.GUID{0xDA35AA22, 0x1DE0, 0x5D3E,
	[8]byte{0x99, 0xFF, 0xD2, 0x4F, 0x31, 0x18, 0xBF, 0x5E}}

type IPackageAllUserProvisioningOptionsInterface interface {
	win32.IInspectableInterface
	Get_OptionalPackageFamilyNames() *winrt.IVector[string]
	Get_ProjectionOrderPackageFamilyNames() *winrt.IVector[string]
}

type IPackageAllUserProvisioningOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_OptionalPackageFamilyNames        uintptr
	Get_ProjectionOrderPackageFamilyNames uintptr
}

type IPackageAllUserProvisioningOptions struct {
	win32.IInspectable
}

func (this *IPackageAllUserProvisioningOptions) Vtbl() *IPackageAllUserProvisioningOptionsVtbl {
	return (*IPackageAllUserProvisioningOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageAllUserProvisioningOptions) Get_OptionalPackageFamilyNames() *winrt.IVector[string] {
	var _result *winrt.IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OptionalPackageFamilyNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageAllUserProvisioningOptions) Get_ProjectionOrderPackageFamilyNames() *winrt.IVector[string] {
	var _result *winrt.IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ProjectionOrderPackageFamilyNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 9A7D4B65-5E8F-4FC7-A2E5-7F6925CB8B53
var IID_IPackageManager = syscall.GUID{0x9A7D4B65, 0x5E8F, 0x4FC7,
	[8]byte{0xA2, 0xE5, 0x7F, 0x69, 0x25, 0xCB, 0x8B, 0x53}}

type IPackageManagerInterface interface {
	win32.IInspectableInterface
	AddPackageAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	UpdatePackageAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	RemovePackageAsync(packageFullName string) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	StagePackageAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	RegisterPackageAsync(manifestUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	FindPackages() *winrt.IIterable[*winrt.IPackage]
	FindPackagesByUserSecurityId(userSecurityId string) *winrt.IIterable[*winrt.IPackage]
	FindPackagesByNamePublisher(packageName string, packagePublisher string) *winrt.IIterable[*winrt.IPackage]
	FindPackagesByUserSecurityIdNamePublisher(userSecurityId string, packageName string, packagePublisher string) *winrt.IIterable[*winrt.IPackage]
	FindUsers(packageFullName string) *winrt.IIterable[*IPackageUserInformation]
	SetPackageState(packageFullName string, packageState PackageState)
	FindPackageByPackageFullName(packageFullName string) *winrt.IPackage
	CleanupPackageForUserAsync(packageName string, userSecurityId string) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	FindPackagesByPackageFamilyName(packageFamilyName string) *winrt.IIterable[*winrt.IPackage]
	FindPackagesByUserSecurityIdPackageFamilyName(userSecurityId string, packageFamilyName string) *winrt.IIterable[*winrt.IPackage]
	FindPackageByUserSecurityIdPackageFullName(userSecurityId string, packageFullName string) *winrt.IPackage
}

type IPackageManagerVtbl struct {
	win32.IInspectableVtbl
	AddPackageAsync                               uintptr
	UpdatePackageAsync                            uintptr
	RemovePackageAsync                            uintptr
	StagePackageAsync                             uintptr
	RegisterPackageAsync                          uintptr
	FindPackages                                  uintptr
	FindPackagesByUserSecurityId                  uintptr
	FindPackagesByNamePublisher                   uintptr
	FindPackagesByUserSecurityIdNamePublisher     uintptr
	FindUsers                                     uintptr
	SetPackageState                               uintptr
	FindPackageByPackageFullName                  uintptr
	CleanupPackageForUserAsync                    uintptr
	FindPackagesByPackageFamilyName               uintptr
	FindPackagesByUserSecurityIdPackageFamilyName uintptr
	FindPackageByUserSecurityIdPackageFullName    uintptr
}

type IPackageManager struct {
	win32.IInspectable
}

func (this *IPackageManager) Vtbl() *IPackageManagerVtbl {
	return (*IPackageManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageManager) AddPackageAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPackageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) UpdatePackageAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().UpdatePackageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) RemovePackageAsync(packageFullName string) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemovePackageAsync, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFullName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) StagePackageAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StagePackageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) RegisterPackageAsync(manifestUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterPackageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(manifestUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) FindPackages() *winrt.IIterable[*winrt.IPackage] {
	var _result *winrt.IIterable[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) FindPackagesByUserSecurityId(userSecurityId string) *winrt.IIterable[*winrt.IPackage] {
	var _result *winrt.IIterable[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByUserSecurityId, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) FindPackagesByNamePublisher(packageName string, packagePublisher string) *winrt.IIterable[*winrt.IPackage] {
	var _result *winrt.IIterable[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByNamePublisher, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageName).Ptr, winrt.NewHStr(packagePublisher).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) FindPackagesByUserSecurityIdNamePublisher(userSecurityId string, packageName string, packagePublisher string) *winrt.IIterable[*winrt.IPackage] {
	var _result *winrt.IIterable[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByUserSecurityIdNamePublisher, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, winrt.NewHStr(packageName).Ptr, winrt.NewHStr(packagePublisher).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) FindUsers(packageFullName string) *winrt.IIterable[*IPackageUserInformation] {
	var _result *winrt.IIterable[*IPackageUserInformation]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindUsers, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFullName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) SetPackageState(packageFullName string, packageState PackageState) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPackageState, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFullName).Ptr, uintptr(packageState))
	_ = _hr
}

func (this *IPackageManager) FindPackageByPackageFullName(packageFullName string) *winrt.IPackage {
	var _result *winrt.IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackageByPackageFullName, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFullName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) CleanupPackageForUserAsync(packageName string, userSecurityId string) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CleanupPackageForUserAsync, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageName).Ptr, winrt.NewHStr(userSecurityId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) FindPackagesByPackageFamilyName(packageFamilyName string) *winrt.IIterable[*winrt.IPackage] {
	var _result *winrt.IIterable[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByPackageFamilyName, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) FindPackagesByUserSecurityIdPackageFamilyName(userSecurityId string, packageFamilyName string) *winrt.IIterable[*winrt.IPackage] {
	var _result *winrt.IIterable[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByUserSecurityIdPackageFamilyName, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, winrt.NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager) FindPackageByUserSecurityIdPackageFullName(userSecurityId string, packageFullName string) *winrt.IPackage {
	var _result *winrt.IPackage
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackageByUserSecurityIdPackageFullName, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, winrt.NewHStr(packageFullName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// A7D7D07E-2E66-4093-AED5-E093ED87B3BB
var IID_IPackageManager10 = syscall.GUID{0xA7D7D07E, 0x2E66, 0x4093,
	[8]byte{0xAE, 0xD5, 0xE0, 0x93, 0xED, 0x87, 0xB3, 0xBB}}

type IPackageManager10Interface interface {
	win32.IInspectableInterface
	ProvisionPackageForAllUsersWithOptionsAsync(mainPackageFamilyName string, options *IPackageAllUserProvisioningOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
}

type IPackageManager10Vtbl struct {
	win32.IInspectableVtbl
	ProvisionPackageForAllUsersWithOptionsAsync uintptr
}

type IPackageManager10 struct {
	win32.IInspectable
}

func (this *IPackageManager10) Vtbl() *IPackageManager10Vtbl {
	return (*IPackageManager10Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageManager10) ProvisionPackageForAllUsersWithOptionsAsync(mainPackageFamilyName string, options *IPackageAllUserProvisioningOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProvisionPackageForAllUsersWithOptionsAsync, uintptr(unsafe.Pointer(this)), winrt.NewHStr(mainPackageFamilyName).Ptr, uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F7AAD08D-0840-46F2-B5D8-CAD47693A095
var IID_IPackageManager2 = syscall.GUID{0xF7AAD08D, 0x0840, 0x46F2,
	[8]byte{0xB5, 0xD8, 0xCA, 0xD4, 0x76, 0x93, 0xA0, 0x95}}

type IPackageManager2Interface interface {
	win32.IInspectableInterface
	RemovePackageWithOptionsAsync(packageFullName string, removalOptions RemovalOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	StagePackageWithOptionsAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	RegisterPackageByFullNameAsync(mainPackageFullName string, dependencyPackageFullNames *winrt.IIterable[string], deploymentOptions DeploymentOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	FindPackagesWithPackageTypes(packageTypes PackageTypes) *winrt.IIterable[*winrt.IPackage]
	FindPackagesByUserSecurityIdWithPackageTypes(userSecurityId string, packageTypes PackageTypes) *winrt.IIterable[*winrt.IPackage]
	FindPackagesByNamePublisherWithPackageTypes(packageName string, packagePublisher string, packageTypes PackageTypes) *winrt.IIterable[*winrt.IPackage]
	FindPackagesByUserSecurityIdNamePublisherWithPackageTypes(userSecurityId string, packageName string, packagePublisher string, packageTypes PackageTypes) *winrt.IIterable[*winrt.IPackage]
	FindPackagesByPackageFamilyNameWithPackageTypes(packageFamilyName string, packageTypes PackageTypes) *winrt.IIterable[*winrt.IPackage]
	FindPackagesByUserSecurityIdPackageFamilyNameWithPackageTypes(userSecurityId string, packageFamilyName string, packageTypes PackageTypes) *winrt.IIterable[*winrt.IPackage]
	StageUserDataAsync(packageFullName string) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
}

type IPackageManager2Vtbl struct {
	win32.IInspectableVtbl
	RemovePackageWithOptionsAsync                                 uintptr
	StagePackageWithOptionsAsync                                  uintptr
	RegisterPackageByFullNameAsync                                uintptr
	FindPackagesWithPackageTypes                                  uintptr
	FindPackagesByUserSecurityIdWithPackageTypes                  uintptr
	FindPackagesByNamePublisherWithPackageTypes                   uintptr
	FindPackagesByUserSecurityIdNamePublisherWithPackageTypes     uintptr
	FindPackagesByPackageFamilyNameWithPackageTypes               uintptr
	FindPackagesByUserSecurityIdPackageFamilyNameWithPackageTypes uintptr
	StageUserDataAsync                                            uintptr
}

type IPackageManager2 struct {
	win32.IInspectable
}

func (this *IPackageManager2) Vtbl() *IPackageManager2Vtbl {
	return (*IPackageManager2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageManager2) RemovePackageWithOptionsAsync(packageFullName string, removalOptions RemovalOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemovePackageWithOptionsAsync, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFullName).Ptr, uintptr(removalOptions), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager2) StagePackageWithOptionsAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StagePackageWithOptionsAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager2) RegisterPackageByFullNameAsync(mainPackageFullName string, dependencyPackageFullNames *winrt.IIterable[string], deploymentOptions DeploymentOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterPackageByFullNameAsync, uintptr(unsafe.Pointer(this)), winrt.NewHStr(mainPackageFullName).Ptr, uintptr(unsafe.Pointer(dependencyPackageFullNames)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager2) FindPackagesWithPackageTypes(packageTypes PackageTypes) *winrt.IIterable[*winrt.IPackage] {
	var _result *winrt.IIterable[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesWithPackageTypes, uintptr(unsafe.Pointer(this)), uintptr(packageTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager2) FindPackagesByUserSecurityIdWithPackageTypes(userSecurityId string, packageTypes PackageTypes) *winrt.IIterable[*winrt.IPackage] {
	var _result *winrt.IIterable[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByUserSecurityIdWithPackageTypes, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, uintptr(packageTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager2) FindPackagesByNamePublisherWithPackageTypes(packageName string, packagePublisher string, packageTypes PackageTypes) *winrt.IIterable[*winrt.IPackage] {
	var _result *winrt.IIterable[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByNamePublisherWithPackageTypes, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageName).Ptr, winrt.NewHStr(packagePublisher).Ptr, uintptr(packageTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager2) FindPackagesByUserSecurityIdNamePublisherWithPackageTypes(userSecurityId string, packageName string, packagePublisher string, packageTypes PackageTypes) *winrt.IIterable[*winrt.IPackage] {
	var _result *winrt.IIterable[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByUserSecurityIdNamePublisherWithPackageTypes, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, winrt.NewHStr(packageName).Ptr, winrt.NewHStr(packagePublisher).Ptr, uintptr(packageTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager2) FindPackagesByPackageFamilyNameWithPackageTypes(packageFamilyName string, packageTypes PackageTypes) *winrt.IIterable[*winrt.IPackage] {
	var _result *winrt.IIterable[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByPackageFamilyNameWithPackageTypes, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFamilyName).Ptr, uintptr(packageTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager2) FindPackagesByUserSecurityIdPackageFamilyNameWithPackageTypes(userSecurityId string, packageFamilyName string, packageTypes PackageTypes) *winrt.IIterable[*winrt.IPackage] {
	var _result *winrt.IIterable[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByUserSecurityIdPackageFamilyNameWithPackageTypes, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, winrt.NewHStr(packageFamilyName).Ptr, uintptr(packageTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager2) StageUserDataAsync(packageFullName string) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StageUserDataAsync, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFullName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// DAAD9948-36F1-41A7-9188-BC263E0DCB72
var IID_IPackageManager3 = syscall.GUID{0xDAAD9948, 0x36F1, 0x41A7,
	[8]byte{0x91, 0x88, 0xBC, 0x26, 0x3E, 0x0D, 0xCB, 0x72}}

type IPackageManager3Interface interface {
	win32.IInspectableInterface
	AddPackageVolumeAsync(packageStorePath string) *winrt.IAsyncOperation[*IPackageVolume]
	AddPackageToVolumeAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, targetVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	ClearPackageStatus(packageFullName string, status PackageStatus)
	RegisterPackageWithAppDataVolumeAsync(manifestUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, appDataVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	FindPackageVolumeByName(volumeName string) *IPackageVolume
	FindPackageVolumes() *winrt.IIterable[*IPackageVolume]
	GetDefaultPackageVolume() *IPackageVolume
	MovePackageToVolumeAsync(packageFullName string, deploymentOptions DeploymentOptions, targetVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	RemovePackageVolumeAsync(volume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	SetDefaultPackageVolume(volume *IPackageVolume)
	SetPackageStatus(packageFullName string, status PackageStatus)
	SetPackageVolumeOfflineAsync(packageVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	SetPackageVolumeOnlineAsync(packageVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	StagePackageToVolumeAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, targetVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	StageUserDataWithOptionsAsync(packageFullName string, deploymentOptions DeploymentOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
}

type IPackageManager3Vtbl struct {
	win32.IInspectableVtbl
	AddPackageVolumeAsync                 uintptr
	AddPackageToVolumeAsync               uintptr
	ClearPackageStatus                    uintptr
	RegisterPackageWithAppDataVolumeAsync uintptr
	FindPackageVolumeByName               uintptr
	FindPackageVolumes                    uintptr
	GetDefaultPackageVolume               uintptr
	MovePackageToVolumeAsync              uintptr
	RemovePackageVolumeAsync              uintptr
	SetDefaultPackageVolume               uintptr
	SetPackageStatus                      uintptr
	SetPackageVolumeOfflineAsync          uintptr
	SetPackageVolumeOnlineAsync           uintptr
	StagePackageToVolumeAsync             uintptr
	StageUserDataWithOptionsAsync         uintptr
}

type IPackageManager3 struct {
	win32.IInspectable
}

func (this *IPackageManager3) Vtbl() *IPackageManager3Vtbl {
	return (*IPackageManager3Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageManager3) AddPackageVolumeAsync(packageStorePath string) *winrt.IAsyncOperation[*IPackageVolume] {
	var _result *winrt.IAsyncOperation[*IPackageVolume]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPackageVolumeAsync, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageStorePath).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager3) AddPackageToVolumeAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, targetVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPackageToVolumeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(targetVolume)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager3) ClearPackageStatus(packageFullName string, status PackageStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ClearPackageStatus, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFullName).Ptr, uintptr(status))
	_ = _hr
}

func (this *IPackageManager3) RegisterPackageWithAppDataVolumeAsync(manifestUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, appDataVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterPackageWithAppDataVolumeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(manifestUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(appDataVolume)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager3) FindPackageVolumeByName(volumeName string) *IPackageVolume {
	var _result *IPackageVolume
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackageVolumeByName, uintptr(unsafe.Pointer(this)), winrt.NewHStr(volumeName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager3) FindPackageVolumes() *winrt.IIterable[*IPackageVolume] {
	var _result *winrt.IIterable[*IPackageVolume]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackageVolumes, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager3) GetDefaultPackageVolume() *IPackageVolume {
	var _result *IPackageVolume
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefaultPackageVolume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager3) MovePackageToVolumeAsync(packageFullName string, deploymentOptions DeploymentOptions, targetVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().MovePackageToVolumeAsync, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFullName).Ptr, uintptr(deploymentOptions), uintptr(unsafe.Pointer(targetVolume)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager3) RemovePackageVolumeAsync(volume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemovePackageVolumeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(volume)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager3) SetDefaultPackageVolume(volume *IPackageVolume) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetDefaultPackageVolume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(volume)))
	_ = _hr
}

func (this *IPackageManager3) SetPackageStatus(packageFullName string, status PackageStatus) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPackageStatus, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFullName).Ptr, uintptr(status))
	_ = _hr
}

func (this *IPackageManager3) SetPackageVolumeOfflineAsync(packageVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPackageVolumeOfflineAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageVolume)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager3) SetPackageVolumeOnlineAsync(packageVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPackageVolumeOnlineAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageVolume)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager3) StagePackageToVolumeAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, targetVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StagePackageToVolumeAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(targetVolume)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager3) StageUserDataWithOptionsAsync(packageFullName string, deploymentOptions DeploymentOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StageUserDataWithOptionsAsync, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFullName).Ptr, uintptr(deploymentOptions), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 3C719963-BAB6-46BF-8FF7-DA4719230AE6
var IID_IPackageManager4 = syscall.GUID{0x3C719963, 0xBAB6, 0x46BF,
	[8]byte{0x8F, 0xF7, 0xDA, 0x47, 0x19, 0x23, 0x0A, 0xE6}}

type IPackageManager4Interface interface {
	win32.IInspectableInterface
	GetPackageVolumesAsync() *winrt.IAsyncOperation[*winrt.IVectorView[*IPackageVolume]]
}

type IPackageManager4Vtbl struct {
	win32.IInspectableVtbl
	GetPackageVolumesAsync uintptr
}

type IPackageManager4 struct {
	win32.IInspectable
}

func (this *IPackageManager4) Vtbl() *IPackageManager4Vtbl {
	return (*IPackageManager4Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageManager4) GetPackageVolumesAsync() *winrt.IAsyncOperation[*winrt.IVectorView[*IPackageVolume]] {
	var _result *winrt.IAsyncOperation[*winrt.IVectorView[*IPackageVolume]]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPackageVolumesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 711F3117-1AFD-4313-978C-9BB6E1B864A7
var IID_IPackageManager5 = syscall.GUID{0x711F3117, 0x1AFD, 0x4313,
	[8]byte{0x97, 0x8C, 0x9B, 0xB6, 0xE1, 0xB8, 0x64, 0xA7}}

type IPackageManager5Interface interface {
	win32.IInspectableInterface
	AddPackageToVolumeAndOptionalPackagesAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, targetVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string], externalPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	StagePackageToVolumeAndOptionalPackagesAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, targetVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string], externalPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	RegisterPackageByFamilyNameAndOptionalPackagesAsync(mainPackageFamilyName string, dependencyPackageFamilyNames *winrt.IIterable[string], deploymentOptions DeploymentOptions, appDataVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	Get_DebugSettings() *IPackageManagerDebugSettings
}

type IPackageManager5Vtbl struct {
	win32.IInspectableVtbl
	AddPackageToVolumeAndOptionalPackagesAsync          uintptr
	StagePackageToVolumeAndOptionalPackagesAsync        uintptr
	RegisterPackageByFamilyNameAndOptionalPackagesAsync uintptr
	Get_DebugSettings                                   uintptr
}

type IPackageManager5 struct {
	win32.IInspectable
}

func (this *IPackageManager5) Vtbl() *IPackageManager5Vtbl {
	return (*IPackageManager5Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageManager5) AddPackageToVolumeAndOptionalPackagesAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, targetVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string], externalPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPackageToVolumeAndOptionalPackagesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(targetVolume)), uintptr(unsafe.Pointer(optionalPackageFamilyNames)), uintptr(unsafe.Pointer(externalPackageUris)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager5) StagePackageToVolumeAndOptionalPackagesAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, targetVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string], externalPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StagePackageToVolumeAndOptionalPackagesAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(targetVolume)), uintptr(unsafe.Pointer(optionalPackageFamilyNames)), uintptr(unsafe.Pointer(externalPackageUris)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager5) RegisterPackageByFamilyNameAndOptionalPackagesAsync(mainPackageFamilyName string, dependencyPackageFamilyNames *winrt.IIterable[string], deploymentOptions DeploymentOptions, appDataVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterPackageByFamilyNameAndOptionalPackagesAsync, uintptr(unsafe.Pointer(this)), winrt.NewHStr(mainPackageFamilyName).Ptr, uintptr(unsafe.Pointer(dependencyPackageFamilyNames)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(appDataVolume)), uintptr(unsafe.Pointer(optionalPackageFamilyNames)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager5) Get_DebugSettings() *IPackageManagerDebugSettings {
	var _result *IPackageManagerDebugSettings
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DebugSettings, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0847E909-53CD-4E4F-832E-57D180F6E447
var IID_IPackageManager6 = syscall.GUID{0x0847E909, 0x53CD, 0x4E4F,
	[8]byte{0x83, 0x2E, 0x57, 0xD1, 0x80, 0xF6, 0xE4, 0x47}}

type IPackageManager6Interface interface {
	win32.IInspectableInterface
	ProvisionPackageForAllUsersAsync(packageFamilyName string) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	AddPackageByAppInstallerFileAsync(appInstallerFileUri *winrt.IUriRuntimeClass, options AddPackageByAppInstallerOptions, targetVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	RequestAddPackageByAppInstallerFileAsync(appInstallerFileUri *winrt.IUriRuntimeClass, options AddPackageByAppInstallerOptions, targetVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	AddPackageToVolumeAndRelatedSetAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], options DeploymentOptions, targetVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string], packageUrisToInstall *winrt.IIterable[*winrt.IUriRuntimeClass], relatedPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	StagePackageToVolumeAndRelatedSetAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], options DeploymentOptions, targetVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string], packageUrisToInstall *winrt.IIterable[*winrt.IUriRuntimeClass], relatedPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	RequestAddPackageAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, targetVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string], relatedPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
}

type IPackageManager6Vtbl struct {
	win32.IInspectableVtbl
	ProvisionPackageForAllUsersAsync         uintptr
	AddPackageByAppInstallerFileAsync        uintptr
	RequestAddPackageByAppInstallerFileAsync uintptr
	AddPackageToVolumeAndRelatedSetAsync     uintptr
	StagePackageToVolumeAndRelatedSetAsync   uintptr
	RequestAddPackageAsync                   uintptr
}

type IPackageManager6 struct {
	win32.IInspectable
}

func (this *IPackageManager6) Vtbl() *IPackageManager6Vtbl {
	return (*IPackageManager6Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageManager6) ProvisionPackageForAllUsersAsync(packageFamilyName string) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ProvisionPackageForAllUsersAsync, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager6) AddPackageByAppInstallerFileAsync(appInstallerFileUri *winrt.IUriRuntimeClass, options AddPackageByAppInstallerOptions, targetVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPackageByAppInstallerFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(appInstallerFileUri)), uintptr(options), uintptr(unsafe.Pointer(targetVolume)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager6) RequestAddPackageByAppInstallerFileAsync(appInstallerFileUri *winrt.IUriRuntimeClass, options AddPackageByAppInstallerOptions, targetVolume *IPackageVolume) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAddPackageByAppInstallerFileAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(appInstallerFileUri)), uintptr(options), uintptr(unsafe.Pointer(targetVolume)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager6) AddPackageToVolumeAndRelatedSetAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], options DeploymentOptions, targetVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string], packageUrisToInstall *winrt.IIterable[*winrt.IUriRuntimeClass], relatedPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPackageToVolumeAndRelatedSetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(options), uintptr(unsafe.Pointer(targetVolume)), uintptr(unsafe.Pointer(optionalPackageFamilyNames)), uintptr(unsafe.Pointer(packageUrisToInstall)), uintptr(unsafe.Pointer(relatedPackageUris)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager6) StagePackageToVolumeAndRelatedSetAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], options DeploymentOptions, targetVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string], packageUrisToInstall *winrt.IIterable[*winrt.IUriRuntimeClass], relatedPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StagePackageToVolumeAndRelatedSetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(options), uintptr(unsafe.Pointer(targetVolume)), uintptr(unsafe.Pointer(optionalPackageFamilyNames)), uintptr(unsafe.Pointer(packageUrisToInstall)), uintptr(unsafe.Pointer(relatedPackageUris)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager6) RequestAddPackageAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, targetVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string], relatedPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAddPackageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(targetVolume)), uintptr(unsafe.Pointer(optionalPackageFamilyNames)), uintptr(unsafe.Pointer(relatedPackageUris)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F28654F4-2BA7-4B80-88D6-BE15F9A23FBA
var IID_IPackageManager7 = syscall.GUID{0xF28654F4, 0x2BA7, 0x4B80,
	[8]byte{0x88, 0xD6, 0xBE, 0x15, 0xF9, 0xA2, 0x3F, 0xBA}}

type IPackageManager7Interface interface {
	win32.IInspectableInterface
	RequestAddPackageAndRelatedSetAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, targetVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string], relatedPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], packageUrisToInstall *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
}

type IPackageManager7Vtbl struct {
	win32.IInspectableVtbl
	RequestAddPackageAndRelatedSetAsync uintptr
}

type IPackageManager7 struct {
	win32.IInspectable
}

func (this *IPackageManager7) Vtbl() *IPackageManager7Vtbl {
	return (*IPackageManager7Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageManager7) RequestAddPackageAndRelatedSetAsync(packageUri *winrt.IUriRuntimeClass, dependencyPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], deploymentOptions DeploymentOptions, targetVolume *IPackageVolume, optionalPackageFamilyNames *winrt.IIterable[string], relatedPackageUris *winrt.IIterable[*winrt.IUriRuntimeClass], packageUrisToInstall *winrt.IIterable[*winrt.IUriRuntimeClass]) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RequestAddPackageAndRelatedSetAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(dependencyPackageUris)), uintptr(deploymentOptions), uintptr(unsafe.Pointer(targetVolume)), uintptr(unsafe.Pointer(optionalPackageFamilyNames)), uintptr(unsafe.Pointer(relatedPackageUris)), uintptr(unsafe.Pointer(packageUrisToInstall)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// B8575330-1298-4EE2-80EE-7F659C5D2782
var IID_IPackageManager8 = syscall.GUID{0xB8575330, 0x1298, 0x4EE2,
	[8]byte{0x80, 0xEE, 0x7F, 0x65, 0x9C, 0x5D, 0x27, 0x82}}

type IPackageManager8Interface interface {
	win32.IInspectableInterface
	DeprovisionPackageForAllUsersAsync(packageFamilyName string) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
}

type IPackageManager8Vtbl struct {
	win32.IInspectableVtbl
	DeprovisionPackageForAllUsersAsync uintptr
}

type IPackageManager8 struct {
	win32.IInspectable
}

func (this *IPackageManager8) Vtbl() *IPackageManager8Vtbl {
	return (*IPackageManager8Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageManager8) DeprovisionPackageForAllUsersAsync(packageFamilyName string) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeprovisionPackageForAllUsersAsync, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 1AA79035-CC71-4B2E-80A6-C7041D8579A7
var IID_IPackageManager9 = syscall.GUID{0x1AA79035, 0xCC71, 0x4B2E,
	[8]byte{0x80, 0xA6, 0xC7, 0x04, 0x1D, 0x85, 0x79, 0xA7}}

type IPackageManager9Interface interface {
	win32.IInspectableInterface
	FindProvisionedPackages() *winrt.IVector[*winrt.IPackage]
	AddPackageByUriAsync(packageUri *winrt.IUriRuntimeClass, options *IAddPackageOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	StagePackageByUriAsync(packageUri *winrt.IUriRuntimeClass, options *IStagePackageOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	RegisterPackageByUriAsync(manifestUri *winrt.IUriRuntimeClass, options *IRegisterPackageOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	RegisterPackagesByFullNameAsync(packageFullNames *winrt.IIterable[string], options *IRegisterPackageOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	SetPackageStubPreference(packageFamilyName string, useStub PackageStubPreference)
	GetPackageStubPreference(packageFamilyName string) PackageStubPreference
}

type IPackageManager9Vtbl struct {
	win32.IInspectableVtbl
	FindProvisionedPackages         uintptr
	AddPackageByUriAsync            uintptr
	StagePackageByUriAsync          uintptr
	RegisterPackageByUriAsync       uintptr
	RegisterPackagesByFullNameAsync uintptr
	SetPackageStubPreference        uintptr
	GetPackageStubPreference        uintptr
}

type IPackageManager9 struct {
	win32.IInspectable
}

func (this *IPackageManager9) Vtbl() *IPackageManager9Vtbl {
	return (*IPackageManager9Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageManager9) FindProvisionedPackages() *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindProvisionedPackages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager9) AddPackageByUriAsync(packageUri *winrt.IUriRuntimeClass, options *IAddPackageOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().AddPackageByUriAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager9) StagePackageByUriAsync(packageUri *winrt.IUriRuntimeClass, options *IStagePackageOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().StagePackageByUriAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageUri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager9) RegisterPackageByUriAsync(manifestUri *winrt.IUriRuntimeClass, options *IRegisterPackageOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterPackageByUriAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(manifestUri)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager9) RegisterPackagesByFullNameAsync(packageFullNames *winrt.IIterable[string], options *IRegisterPackageOptions) *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress] {
	var _result *winrt.IAsyncOperationWithProgress[*IDeploymentResult, DeploymentProgress]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RegisterPackagesByFullNameAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(packageFullNames)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManager9) SetPackageStubPreference(packageFamilyName string, useStub PackageStubPreference) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetPackageStubPreference, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFamilyName).Ptr, uintptr(useStub))
	_ = _hr
}

func (this *IPackageManager9) GetPackageStubPreference(packageFamilyName string) PackageStubPreference {
	var _result PackageStubPreference
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetPackageStubPreference, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// 1A611683-A988-4FCF-8F0F-CE175898E8EB
var IID_IPackageManagerDebugSettings = syscall.GUID{0x1A611683, 0xA988, 0x4FCF,
	[8]byte{0x8F, 0x0F, 0xCE, 0x17, 0x58, 0x98, 0xE8, 0xEB}}

type IPackageManagerDebugSettingsInterface interface {
	win32.IInspectableInterface
	SetContentGroupStateAsync(package_ *winrt.IPackage, contentGroupName string, state winrt.PackageContentGroupState) *winrt.IAsyncAction
	SetContentGroupStateWithPercentageAsync(package_ *winrt.IPackage, contentGroupName string, state winrt.PackageContentGroupState, completionPercentage float64) *winrt.IAsyncAction
}

type IPackageManagerDebugSettingsVtbl struct {
	win32.IInspectableVtbl
	SetContentGroupStateAsync               uintptr
	SetContentGroupStateWithPercentageAsync uintptr
}

type IPackageManagerDebugSettings struct {
	win32.IInspectable
}

func (this *IPackageManagerDebugSettings) Vtbl() *IPackageManagerDebugSettingsVtbl {
	return (*IPackageManagerDebugSettingsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageManagerDebugSettings) SetContentGroupStateAsync(package_ *winrt.IPackage, contentGroupName string, state winrt.PackageContentGroupState) *winrt.IAsyncAction {
	var _result *winrt.IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetContentGroupStateAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(package_)), winrt.NewHStr(contentGroupName).Ptr, uintptr(state), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageManagerDebugSettings) SetContentGroupStateWithPercentageAsync(package_ *winrt.IPackage, contentGroupName string, state winrt.PackageContentGroupState, completionPercentage float64) *winrt.IAsyncAction {
	var _result *winrt.IAsyncAction
	_hr, _, _ := syscall.SyscallN(this.Vtbl().SetContentGroupStateWithPercentageAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(package_)), winrt.NewHStr(contentGroupName).Ptr, uintptr(state), uintptr(completionPercentage), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// F6383423-FA09-4CBC-9055-15CA275E2E7E
var IID_IPackageUserInformation = syscall.GUID{0xF6383423, 0xFA09, 0x4CBC,
	[8]byte{0x90, 0x55, 0x15, 0xCA, 0x27, 0x5E, 0x2E, 0x7E}}

type IPackageUserInformationInterface interface {
	win32.IInspectableInterface
	Get_UserSecurityId() string
	Get_InstallState() PackageInstallState
}

type IPackageUserInformationVtbl struct {
	win32.IInspectableVtbl
	Get_UserSecurityId uintptr
	Get_InstallState   uintptr
}

type IPackageUserInformation struct {
	win32.IInspectable
}

func (this *IPackageUserInformation) Vtbl() *IPackageUserInformationVtbl {
	return (*IPackageUserInformationVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageUserInformation) Get_UserSecurityId() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_UserSecurityId, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return winrt.HStringToStrAndFree(_result)
}

func (this *IPackageUserInformation) Get_InstallState() PackageInstallState {
	var _result PackageInstallState
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstallState, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// CF2672C3-1A40-4450-9739-2ACE2E898853
var IID_IPackageVolume = syscall.GUID{0xCF2672C3, 0x1A40, 0x4450,
	[8]byte{0x97, 0x39, 0x2A, 0xCE, 0x2E, 0x89, 0x88, 0x53}}

type IPackageVolumeInterface interface {
	win32.IInspectableInterface
	Get_IsOffline() bool
	Get_IsSystemVolume() bool
	Get_MountPoint() string
	Get_Name() string
	Get_PackageStorePath() string
	Get_SupportsHardLinks() bool
	FindPackages() *winrt.IVector[*winrt.IPackage]
	FindPackagesByNamePublisher(packageName string, packagePublisher string) *winrt.IVector[*winrt.IPackage]
	FindPackagesByPackageFamilyName(packageFamilyName string) *winrt.IVector[*winrt.IPackage]
	FindPackagesWithPackageTypes(packageTypes PackageTypes) *winrt.IVector[*winrt.IPackage]
	FindPackagesByNamePublisherWithPackagesTypes(packageTypes PackageTypes, packageName string, packagePublisher string) *winrt.IVector[*winrt.IPackage]
	FindPackagesByPackageFamilyNameWithPackageTypes(packageTypes PackageTypes, packageFamilyName string) *winrt.IVector[*winrt.IPackage]
	FindPackageByPackageFullName(packageFullName string) *winrt.IVector[*winrt.IPackage]
	FindPackagesByUserSecurityId(userSecurityId string) *winrt.IVector[*winrt.IPackage]
	FindPackagesByUserSecurityIdNamePublisher(userSecurityId string, packageName string, packagePublisher string) *winrt.IVector[*winrt.IPackage]
	FindPackagesByUserSecurityIdPackageFamilyName(userSecurityId string, packageFamilyName string) *winrt.IVector[*winrt.IPackage]
	FindPackagesByUserSecurityIdWithPackageTypes(userSecurityId string, packageTypes PackageTypes) *winrt.IVector[*winrt.IPackage]
	FindPackagesByUserSecurityIdNamePublisherWithPackageTypes(userSecurityId string, packageTypes PackageTypes, packageName string, packagePublisher string) *winrt.IVector[*winrt.IPackage]
	FindPackagesByUserSecurityIdPackageFamilyNameWithPackagesTypes(userSecurityId string, packageTypes PackageTypes, packageFamilyName string) *winrt.IVector[*winrt.IPackage]
	FindPackageByUserSecurityIdPackageFullName(userSecurityId string, packageFullName string) *winrt.IVector[*winrt.IPackage]
}

type IPackageVolumeVtbl struct {
	win32.IInspectableVtbl
	Get_IsOffline                                                  uintptr
	Get_IsSystemVolume                                             uintptr
	Get_MountPoint                                                 uintptr
	Get_Name                                                       uintptr
	Get_PackageStorePath                                           uintptr
	Get_SupportsHardLinks                                          uintptr
	FindPackages                                                   uintptr
	FindPackagesByNamePublisher                                    uintptr
	FindPackagesByPackageFamilyName                                uintptr
	FindPackagesWithPackageTypes                                   uintptr
	FindPackagesByNamePublisherWithPackagesTypes                   uintptr
	FindPackagesByPackageFamilyNameWithPackageTypes                uintptr
	FindPackageByPackageFullName                                   uintptr
	FindPackagesByUserSecurityId                                   uintptr
	FindPackagesByUserSecurityIdNamePublisher                      uintptr
	FindPackagesByUserSecurityIdPackageFamilyName                  uintptr
	FindPackagesByUserSecurityIdWithPackageTypes                   uintptr
	FindPackagesByUserSecurityIdNamePublisherWithPackageTypes      uintptr
	FindPackagesByUserSecurityIdPackageFamilyNameWithPackagesTypes uintptr
	FindPackageByUserSecurityIdPackageFullName                     uintptr
}

type IPackageVolume struct {
	win32.IInspectable
}

func (this *IPackageVolume) Vtbl() *IPackageVolumeVtbl {
	return (*IPackageVolumeVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageVolume) Get_IsOffline() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsOffline, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageVolume) Get_IsSystemVolume() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsSystemVolume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageVolume) Get_MountPoint() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_MountPoint, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return winrt.HStringToStrAndFree(_result)
}

func (this *IPackageVolume) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return winrt.HStringToStrAndFree(_result)
}

func (this *IPackageVolume) Get_PackageStorePath() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PackageStorePath, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return winrt.HStringToStrAndFree(_result)
}

func (this *IPackageVolume) Get_SupportsHardLinks() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_SupportsHardLinks, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageVolume) FindPackages() *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackages, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackagesByNamePublisher(packageName string, packagePublisher string) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByNamePublisher, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageName).Ptr, winrt.NewHStr(packagePublisher).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackagesByPackageFamilyName(packageFamilyName string) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByPackageFamilyName, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackagesWithPackageTypes(packageTypes PackageTypes) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesWithPackageTypes, uintptr(unsafe.Pointer(this)), uintptr(packageTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackagesByNamePublisherWithPackagesTypes(packageTypes PackageTypes, packageName string, packagePublisher string) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByNamePublisherWithPackagesTypes, uintptr(unsafe.Pointer(this)), uintptr(packageTypes), winrt.NewHStr(packageName).Ptr, winrt.NewHStr(packagePublisher).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackagesByPackageFamilyNameWithPackageTypes(packageTypes PackageTypes, packageFamilyName string) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByPackageFamilyNameWithPackageTypes, uintptr(unsafe.Pointer(this)), uintptr(packageTypes), winrt.NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackageByPackageFullName(packageFullName string) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackageByPackageFullName, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFullName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackagesByUserSecurityId(userSecurityId string) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByUserSecurityId, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackagesByUserSecurityIdNamePublisher(userSecurityId string, packageName string, packagePublisher string) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByUserSecurityIdNamePublisher, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, winrt.NewHStr(packageName).Ptr, winrt.NewHStr(packagePublisher).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackagesByUserSecurityIdPackageFamilyName(userSecurityId string, packageFamilyName string) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByUserSecurityIdPackageFamilyName, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, winrt.NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackagesByUserSecurityIdWithPackageTypes(userSecurityId string, packageTypes PackageTypes) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByUserSecurityIdWithPackageTypes, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, uintptr(packageTypes), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackagesByUserSecurityIdNamePublisherWithPackageTypes(userSecurityId string, packageTypes PackageTypes, packageName string, packagePublisher string) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByUserSecurityIdNamePublisherWithPackageTypes, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, uintptr(packageTypes), winrt.NewHStr(packageName).Ptr, winrt.NewHStr(packagePublisher).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackagesByUserSecurityIdPackageFamilyNameWithPackagesTypes(userSecurityId string, packageTypes PackageTypes, packageFamilyName string) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackagesByUserSecurityIdPackageFamilyNameWithPackagesTypes, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, uintptr(packageTypes), winrt.NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IPackageVolume) FindPackageByUserSecurityIdPackageFullName(userSecurityId string, packageFullName string) *winrt.IVector[*winrt.IPackage] {
	var _result *winrt.IVector[*winrt.IPackage]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindPackageByUserSecurityIdPackageFullName, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSecurityId).Ptr, winrt.NewHStr(packageFullName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 46ABCF2E-9DD4-47A2-AB8C-C6408349BCD8
var IID_IPackageVolume2 = syscall.GUID{0x46ABCF2E, 0x9DD4, 0x47A2,
	[8]byte{0xAB, 0x8C, 0xC6, 0x40, 0x83, 0x49, 0xBC, 0xD8}}

type IPackageVolume2Interface interface {
	win32.IInspectableInterface
	Get_IsFullTrustPackageSupported() bool
	Get_IsAppxInstallSupported() bool
	GetAvailableSpaceAsync() *winrt.IAsyncOperation[uint64]
}

type IPackageVolume2Vtbl struct {
	win32.IInspectableVtbl
	Get_IsFullTrustPackageSupported uintptr
	Get_IsAppxInstallSupported      uintptr
	GetAvailableSpaceAsync          uintptr
}

type IPackageVolume2 struct {
	win32.IInspectable
}

func (this *IPackageVolume2) Vtbl() *IPackageVolume2Vtbl {
	return (*IPackageVolume2Vtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IPackageVolume2) Get_IsFullTrustPackageSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsFullTrustPackageSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageVolume2) Get_IsAppxInstallSupported() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_IsAppxInstallSupported, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IPackageVolume2) GetAvailableSpaceAsync() *winrt.IAsyncOperation[uint64] {
	var _result *winrt.IAsyncOperation[uint64]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetAvailableSpaceAsync, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 677112A7-50D4-496C-8415-0602B4C6D3BF
var IID_IRegisterPackageOptions = syscall.GUID{0x677112A7, 0x50D4, 0x496C,
	[8]byte{0x84, 0x15, 0x06, 0x02, 0xB4, 0xC6, 0xD3, 0xBF}}

type IRegisterPackageOptionsInterface interface {
	win32.IInspectableInterface
	Get_DependencyPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass]
	Get_AppDataVolume() *IPackageVolume
	Put_AppDataVolume(value *IPackageVolume)
	Get_OptionalPackageFamilyNames() *winrt.IVector[string]
	Get_ExternalLocationUri() *winrt.IUriRuntimeClass
	Put_ExternalLocationUri(value *winrt.IUriRuntimeClass)
	Get_DeveloperMode() bool
	Put_DeveloperMode(value bool)
	Get_ForceAppShutdown() bool
	Put_ForceAppShutdown(value bool)
	Get_ForceTargetAppShutdown() bool
	Put_ForceTargetAppShutdown(value bool)
	Get_ForceUpdateFromAnyVersion() bool
	Put_ForceUpdateFromAnyVersion(value bool)
	Get_InstallAllResources() bool
	Put_InstallAllResources(value bool)
	Get_StageInPlace() bool
	Put_StageInPlace(value bool)
	Get_AllowUnsigned() bool
	Put_AllowUnsigned(value bool)
	Get_DeferRegistrationWhenPackagesAreInUse() bool
	Put_DeferRegistrationWhenPackagesAreInUse(value bool)
}

type IRegisterPackageOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_DependencyPackageUris                 uintptr
	Get_AppDataVolume                         uintptr
	Put_AppDataVolume                         uintptr
	Get_OptionalPackageFamilyNames            uintptr
	Get_ExternalLocationUri                   uintptr
	Put_ExternalLocationUri                   uintptr
	Get_DeveloperMode                         uintptr
	Put_DeveloperMode                         uintptr
	Get_ForceAppShutdown                      uintptr
	Put_ForceAppShutdown                      uintptr
	Get_ForceTargetAppShutdown                uintptr
	Put_ForceTargetAppShutdown                uintptr
	Get_ForceUpdateFromAnyVersion             uintptr
	Put_ForceUpdateFromAnyVersion             uintptr
	Get_InstallAllResources                   uintptr
	Put_InstallAllResources                   uintptr
	Get_StageInPlace                          uintptr
	Put_StageInPlace                          uintptr
	Get_AllowUnsigned                         uintptr
	Put_AllowUnsigned                         uintptr
	Get_DeferRegistrationWhenPackagesAreInUse uintptr
	Put_DeferRegistrationWhenPackagesAreInUse uintptr
}

type IRegisterPackageOptions struct {
	win32.IInspectable
}

func (this *IRegisterPackageOptions) Vtbl() *IRegisterPackageOptionsVtbl {
	return (*IRegisterPackageOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IRegisterPackageOptions) Get_DependencyPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass] {
	var _result *winrt.IVector[*winrt.IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DependencyPackageUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRegisterPackageOptions) Get_AppDataVolume() *IPackageVolume {
	var _result *IPackageVolume
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AppDataVolume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRegisterPackageOptions) Put_AppDataVolume(value *IPackageVolume) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AppDataVolume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IRegisterPackageOptions) Get_OptionalPackageFamilyNames() *winrt.IVector[string] {
	var _result *winrt.IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OptionalPackageFamilyNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRegisterPackageOptions) Get_ExternalLocationUri() *winrt.IUriRuntimeClass {
	var _result *winrt.IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExternalLocationUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IRegisterPackageOptions) Put_ExternalLocationUri(value *winrt.IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExternalLocationUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IRegisterPackageOptions) Get_DeveloperMode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeveloperMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegisterPackageOptions) Put_DeveloperMode(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DeveloperMode, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRegisterPackageOptions) Get_ForceAppShutdown() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegisterPackageOptions) Put_ForceAppShutdown(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForceAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRegisterPackageOptions) Get_ForceTargetAppShutdown() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceTargetAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegisterPackageOptions) Put_ForceTargetAppShutdown(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForceTargetAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRegisterPackageOptions) Get_ForceUpdateFromAnyVersion() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceUpdateFromAnyVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegisterPackageOptions) Put_ForceUpdateFromAnyVersion(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForceUpdateFromAnyVersion, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRegisterPackageOptions) Get_InstallAllResources() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstallAllResources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegisterPackageOptions) Put_InstallAllResources(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InstallAllResources, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRegisterPackageOptions) Get_StageInPlace() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StageInPlace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegisterPackageOptions) Put_StageInPlace(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StageInPlace, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRegisterPackageOptions) Get_AllowUnsigned() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowUnsigned, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegisterPackageOptions) Put_AllowUnsigned(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowUnsigned, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IRegisterPackageOptions) Get_DeferRegistrationWhenPackagesAreInUse() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeferRegistrationWhenPackagesAreInUse, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IRegisterPackageOptions) Put_DeferRegistrationWhenPackagesAreInUse(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DeferRegistrationWhenPackagesAreInUse, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 177F1AA9-151E-5EF7-B1D9-2FBA0B4B0D17
var IID_ISharedPackageContainer = syscall.GUID{0x177F1AA9, 0x151E, 0x5EF7,
	[8]byte{0xB1, 0xD9, 0x2F, 0xBA, 0x0B, 0x4B, 0x0D, 0x17}}

type ISharedPackageContainerInterface interface {
	win32.IInspectableInterface
	Get_Name() string
	Get_Id() string
	GetMembers() *winrt.IVector[*ISharedPackageContainerMember]
	RemovePackageFamily(packageFamilyName string, options *IUpdateSharedPackageContainerOptions) *IUpdateSharedPackageContainerResult
	ResetData() *IUpdateSharedPackageContainerResult
}

type ISharedPackageContainerVtbl struct {
	win32.IInspectableVtbl
	Get_Name            uintptr
	Get_Id              uintptr
	GetMembers          uintptr
	RemovePackageFamily uintptr
	ResetData           uintptr
}

type ISharedPackageContainer struct {
	win32.IInspectable
}

func (this *ISharedPackageContainer) Vtbl() *ISharedPackageContainerVtbl {
	return (*ISharedPackageContainerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISharedPackageContainer) Get_Name() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Name, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return winrt.HStringToStrAndFree(_result)
}

func (this *ISharedPackageContainer) Get_Id() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Id, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return winrt.HStringToStrAndFree(_result)
}

func (this *ISharedPackageContainer) GetMembers() *winrt.IVector[*ISharedPackageContainerMember] {
	var _result *winrt.IVector[*ISharedPackageContainerMember]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetMembers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISharedPackageContainer) RemovePackageFamily(packageFamilyName string, options *IUpdateSharedPackageContainerOptions) *IUpdateSharedPackageContainerResult {
	var _result *IUpdateSharedPackageContainerResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().RemovePackageFamily, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISharedPackageContainer) ResetData() *IUpdateSharedPackageContainerResult {
	var _result *IUpdateSharedPackageContainerResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().ResetData, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// BE353068-1EF7-5AC8-AB3F-0B9F612F0274
var IID_ISharedPackageContainerManager = syscall.GUID{0xBE353068, 0x1EF7, 0x5AC8,
	[8]byte{0xAB, 0x3F, 0x0B, 0x9F, 0x61, 0x2F, 0x02, 0x74}}

type ISharedPackageContainerManagerInterface interface {
	win32.IInspectableInterface
	CreateContainer(name string, options *ICreateSharedPackageContainerOptions) *ICreateSharedPackageContainerResult
	DeleteContainer(id string, options *IDeleteSharedPackageContainerOptions) *IDeleteSharedPackageContainerResult
	GetContainer(id string) *ISharedPackageContainer
	FindContainers() *winrt.IVector[*ISharedPackageContainer]
	FindContainersWithOptions(options *IFindSharedPackageContainerOptions) *winrt.IVector[*ISharedPackageContainer]
}

type ISharedPackageContainerManagerVtbl struct {
	win32.IInspectableVtbl
	CreateContainer           uintptr
	DeleteContainer           uintptr
	GetContainer              uintptr
	FindContainers            uintptr
	FindContainersWithOptions uintptr
}

type ISharedPackageContainerManager struct {
	win32.IInspectable
}

func (this *ISharedPackageContainerManager) Vtbl() *ISharedPackageContainerManagerVtbl {
	return (*ISharedPackageContainerManagerVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISharedPackageContainerManager) CreateContainer(name string, options *ICreateSharedPackageContainerOptions) *ICreateSharedPackageContainerResult {
	var _result *ICreateSharedPackageContainerResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateContainer, uintptr(unsafe.Pointer(this)), winrt.NewHStr(name).Ptr, uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISharedPackageContainerManager) DeleteContainer(id string, options *IDeleteSharedPackageContainerOptions) *IDeleteSharedPackageContainerResult {
	var _result *IDeleteSharedPackageContainerResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().DeleteContainer, uintptr(unsafe.Pointer(this)), winrt.NewHStr(id).Ptr, uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISharedPackageContainerManager) GetContainer(id string) *ISharedPackageContainer {
	var _result *ISharedPackageContainer
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetContainer, uintptr(unsafe.Pointer(this)), winrt.NewHStr(id).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISharedPackageContainerManager) FindContainers() *winrt.IVector[*ISharedPackageContainer] {
	var _result *winrt.IVector[*ISharedPackageContainer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindContainers, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISharedPackageContainerManager) FindContainersWithOptions(options *IFindSharedPackageContainerOptions) *winrt.IVector[*ISharedPackageContainer] {
	var _result *winrt.IVector[*ISharedPackageContainer]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().FindContainersWithOptions, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 2EF56348-838A-5F55-A89E-1198A2C627E6
var IID_ISharedPackageContainerManagerStatics = syscall.GUID{0x2EF56348, 0x838A, 0x5F55,
	[8]byte{0xA8, 0x9E, 0x11, 0x98, 0xA2, 0xC6, 0x27, 0xE6}}

type ISharedPackageContainerManagerStaticsInterface interface {
	win32.IInspectableInterface
	GetDefault() *ISharedPackageContainerManager
	GetForUser(userSid string) *ISharedPackageContainerManager
	GetForProvisioning() *ISharedPackageContainerManager
}

type ISharedPackageContainerManagerStaticsVtbl struct {
	win32.IInspectableVtbl
	GetDefault         uintptr
	GetForUser         uintptr
	GetForProvisioning uintptr
}

type ISharedPackageContainerManagerStatics struct {
	win32.IInspectable
}

func (this *ISharedPackageContainerManagerStatics) Vtbl() *ISharedPackageContainerManagerStaticsVtbl {
	return (*ISharedPackageContainerManagerStaticsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISharedPackageContainerManagerStatics) GetDefault() *ISharedPackageContainerManager {
	var _result *ISharedPackageContainerManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetDefault, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISharedPackageContainerManagerStatics) GetForUser(userSid string) *ISharedPackageContainerManager {
	var _result *ISharedPackageContainerManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForUser, uintptr(unsafe.Pointer(this)), winrt.NewHStr(userSid).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *ISharedPackageContainerManagerStatics) GetForProvisioning() *ISharedPackageContainerManager {
	var _result *ISharedPackageContainerManager
	_hr, _, _ := syscall.SyscallN(this.Vtbl().GetForProvisioning, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// FE0D0438-43C9-5426-B89C-F79BF85DDFF4
var IID_ISharedPackageContainerMember = syscall.GUID{0xFE0D0438, 0x43C9, 0x5426,
	[8]byte{0xB8, 0x9C, 0xF7, 0x9B, 0xF8, 0x5D, 0xDF, 0xF4}}

type ISharedPackageContainerMemberInterface interface {
	win32.IInspectableInterface
	Get_PackageFamilyName() string
}

type ISharedPackageContainerMemberVtbl struct {
	win32.IInspectableVtbl
	Get_PackageFamilyName uintptr
}

type ISharedPackageContainerMember struct {
	win32.IInspectable
}

func (this *ISharedPackageContainerMember) Vtbl() *ISharedPackageContainerMemberVtbl {
	return (*ISharedPackageContainerMemberVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISharedPackageContainerMember) Get_PackageFamilyName() string {
	var _result win32.HSTRING
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_PackageFamilyName, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return winrt.HStringToStrAndFree(_result)
}

// 49B0CEEB-498F-5A62-B738-B3CA0D436704
var IID_ISharedPackageContainerMemberFactory = syscall.GUID{0x49B0CEEB, 0x498F, 0x5A62,
	[8]byte{0xB7, 0x38, 0xB3, 0xCA, 0x0D, 0x43, 0x67, 0x04}}

type ISharedPackageContainerMemberFactoryInterface interface {
	win32.IInspectableInterface
	CreateInstance(packageFamilyName string) *ISharedPackageContainerMember
}

type ISharedPackageContainerMemberFactoryVtbl struct {
	win32.IInspectableVtbl
	CreateInstance uintptr
}

type ISharedPackageContainerMemberFactory struct {
	win32.IInspectable
}

func (this *ISharedPackageContainerMemberFactory) Vtbl() *ISharedPackageContainerMemberFactoryVtbl {
	return (*ISharedPackageContainerMemberFactoryVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *ISharedPackageContainerMemberFactory) CreateInstance(packageFamilyName string) *ISharedPackageContainerMember {
	var _result *ISharedPackageContainerMember
	_hr, _, _ := syscall.SyscallN(this.Vtbl().CreateInstance, uintptr(unsafe.Pointer(this)), winrt.NewHStr(packageFamilyName).Ptr, uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

// 0B110C9C-B95D-4C56-BD36-6D656800D06B
var IID_IStagePackageOptions = syscall.GUID{0x0B110C9C, 0xB95D, 0x4C56,
	[8]byte{0xBD, 0x36, 0x6D, 0x65, 0x68, 0x00, 0xD0, 0x6B}}

type IStagePackageOptionsInterface interface {
	win32.IInspectableInterface
	Get_DependencyPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass]
	Get_TargetVolume() *IPackageVolume
	Put_TargetVolume(value *IPackageVolume)
	Get_OptionalPackageFamilyNames() *winrt.IVector[string]
	Get_OptionalPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass]
	Get_RelatedPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass]
	Get_ExternalLocationUri() *winrt.IUriRuntimeClass
	Put_ExternalLocationUri(value *winrt.IUriRuntimeClass)
	Get_StubPackageOption() StubPackageOption
	Put_StubPackageOption(value StubPackageOption)
	Get_DeveloperMode() bool
	Put_DeveloperMode(value bool)
	Get_ForceUpdateFromAnyVersion() bool
	Put_ForceUpdateFromAnyVersion(value bool)
	Get_InstallAllResources() bool
	Put_InstallAllResources(value bool)
	Get_RequiredContentGroupOnly() bool
	Put_RequiredContentGroupOnly(value bool)
	Get_StageInPlace() bool
	Put_StageInPlace(value bool)
	Get_AllowUnsigned() bool
	Put_AllowUnsigned(value bool)
}

type IStagePackageOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_DependencyPackageUris      uintptr
	Get_TargetVolume               uintptr
	Put_TargetVolume               uintptr
	Get_OptionalPackageFamilyNames uintptr
	Get_OptionalPackageUris        uintptr
	Get_RelatedPackageUris         uintptr
	Get_ExternalLocationUri        uintptr
	Put_ExternalLocationUri        uintptr
	Get_StubPackageOption          uintptr
	Put_StubPackageOption          uintptr
	Get_DeveloperMode              uintptr
	Put_DeveloperMode              uintptr
	Get_ForceUpdateFromAnyVersion  uintptr
	Put_ForceUpdateFromAnyVersion  uintptr
	Get_InstallAllResources        uintptr
	Put_InstallAllResources        uintptr
	Get_RequiredContentGroupOnly   uintptr
	Put_RequiredContentGroupOnly   uintptr
	Get_StageInPlace               uintptr
	Put_StageInPlace               uintptr
	Get_AllowUnsigned              uintptr
	Put_AllowUnsigned              uintptr
}

type IStagePackageOptions struct {
	win32.IInspectable
}

func (this *IStagePackageOptions) Vtbl() *IStagePackageOptionsVtbl {
	return (*IStagePackageOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IStagePackageOptions) Get_DependencyPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass] {
	var _result *winrt.IVector[*winrt.IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DependencyPackageUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStagePackageOptions) Get_TargetVolume() *IPackageVolume {
	var _result *IPackageVolume
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_TargetVolume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStagePackageOptions) Put_TargetVolume(value *IPackageVolume) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_TargetVolume, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IStagePackageOptions) Get_OptionalPackageFamilyNames() *winrt.IVector[string] {
	var _result *winrt.IVector[string]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OptionalPackageFamilyNames, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStagePackageOptions) Get_OptionalPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass] {
	var _result *winrt.IVector[*winrt.IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_OptionalPackageUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStagePackageOptions) Get_RelatedPackageUris() *winrt.IVector[*winrt.IUriRuntimeClass] {
	var _result *winrt.IVector[*winrt.IUriRuntimeClass]
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RelatedPackageUris, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStagePackageOptions) Get_ExternalLocationUri() *winrt.IUriRuntimeClass {
	var _result *winrt.IUriRuntimeClass
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExternalLocationUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	com.AddToScope(_result)
	return _result
}

func (this *IStagePackageOptions) Put_ExternalLocationUri(value *winrt.IUriRuntimeClass) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ExternalLocationUri, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	_ = _hr
}

func (this *IStagePackageOptions) Get_StubPackageOption() StubPackageOption {
	var _result StubPackageOption
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StubPackageOption, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStagePackageOptions) Put_StubPackageOption(value StubPackageOption) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StubPackageOption, uintptr(unsafe.Pointer(this)), uintptr(value))
	_ = _hr
}

func (this *IStagePackageOptions) Get_DeveloperMode() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_DeveloperMode, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStagePackageOptions) Put_DeveloperMode(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_DeveloperMode, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStagePackageOptions) Get_ForceUpdateFromAnyVersion() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceUpdateFromAnyVersion, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStagePackageOptions) Put_ForceUpdateFromAnyVersion(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForceUpdateFromAnyVersion, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStagePackageOptions) Get_InstallAllResources() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_InstallAllResources, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStagePackageOptions) Put_InstallAllResources(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_InstallAllResources, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStagePackageOptions) Get_RequiredContentGroupOnly() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequiredContentGroupOnly, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStagePackageOptions) Put_RequiredContentGroupOnly(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequiredContentGroupOnly, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStagePackageOptions) Get_StageInPlace() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_StageInPlace, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStagePackageOptions) Put_StageInPlace(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_StageInPlace, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IStagePackageOptions) Get_AllowUnsigned() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_AllowUnsigned, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IStagePackageOptions) Put_AllowUnsigned(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_AllowUnsigned, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// 80672E83-7194-59F9-B5B9-DAA5375F130A
var IID_IUpdateSharedPackageContainerOptions = syscall.GUID{0x80672E83, 0x7194, 0x59F9,
	[8]byte{0xB5, 0xB9, 0xDA, 0xA5, 0x37, 0x5F, 0x13, 0x0A}}

type IUpdateSharedPackageContainerOptionsInterface interface {
	win32.IInspectableInterface
	Get_ForceAppShutdown() bool
	Put_ForceAppShutdown(value bool)
	Get_RequirePackagesPresent() bool
	Put_RequirePackagesPresent(value bool)
}

type IUpdateSharedPackageContainerOptionsVtbl struct {
	win32.IInspectableVtbl
	Get_ForceAppShutdown       uintptr
	Put_ForceAppShutdown       uintptr
	Get_RequirePackagesPresent uintptr
	Put_RequirePackagesPresent uintptr
}

type IUpdateSharedPackageContainerOptions struct {
	win32.IInspectable
}

func (this *IUpdateSharedPackageContainerOptions) Vtbl() *IUpdateSharedPackageContainerOptionsVtbl {
	return (*IUpdateSharedPackageContainerOptionsVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUpdateSharedPackageContainerOptions) Get_ForceAppShutdown() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ForceAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUpdateSharedPackageContainerOptions) Put_ForceAppShutdown(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_ForceAppShutdown, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

func (this *IUpdateSharedPackageContainerOptions) Get_RequirePackagesPresent() bool {
	var _result bool
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_RequirePackagesPresent, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUpdateSharedPackageContainerOptions) Put_RequirePackagesPresent(value bool) {
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Put_RequirePackagesPresent, uintptr(unsafe.Pointer(this)), uintptr(*(*byte)(unsafe.Pointer(&value))))
	_ = _hr
}

// AA407DF7-C72D-5458-AEA3-4645B6A8EE99
var IID_IUpdateSharedPackageContainerResult = syscall.GUID{0xAA407DF7, 0xC72D, 0x5458,
	[8]byte{0xAE, 0xA3, 0x46, 0x45, 0xB6, 0xA8, 0xEE, 0x99}}

type IUpdateSharedPackageContainerResultInterface interface {
	win32.IInspectableInterface
	Get_Status() SharedPackageContainerOperationStatus
	Get_ExtendedError() winrt.HResult
}

type IUpdateSharedPackageContainerResultVtbl struct {
	win32.IInspectableVtbl
	Get_Status        uintptr
	Get_ExtendedError uintptr
}

type IUpdateSharedPackageContainerResult struct {
	win32.IInspectable
}

func (this *IUpdateSharedPackageContainerResult) Vtbl() *IUpdateSharedPackageContainerResultVtbl {
	return (*IUpdateSharedPackageContainerResultVtbl)(unsafe.Pointer(this.IUnknown.LpVtbl))
}

func (this *IUpdateSharedPackageContainerResult) Get_Status() SharedPackageContainerOperationStatus {
	var _result SharedPackageContainerOperationStatus
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_Status, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (this *IUpdateSharedPackageContainerResult) Get_ExtendedError() winrt.HResult {
	var _result winrt.HResult
	_hr, _, _ := syscall.SyscallN(this.Vtbl().Get_ExtendedError, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

// classes

type AddPackageOptions struct {
	winrt.RtClass
	*IAddPackageOptions
}

func NewAddPackageOptions() *AddPackageOptions {
	hs := winrt.NewHStr("Windows.Management.Deployment.AddPackageOptions")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &AddPackageOptions{
		RtClass:            winrt.RtClass{PInspect: p},
		IAddPackageOptions: (*IAddPackageOptions)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type DeploymentResult struct {
	winrt.RtClass
	*IDeploymentResult
}

type PackageAllUserProvisioningOptions struct {
	winrt.RtClass
	*IPackageAllUserProvisioningOptions
}

func NewPackageAllUserProvisioningOptions() *PackageAllUserProvisioningOptions {
	hs := winrt.NewHStr("Windows.Management.Deployment.PackageAllUserProvisioningOptions")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PackageAllUserProvisioningOptions{
		RtClass:                            winrt.RtClass{PInspect: p},
		IPackageAllUserProvisioningOptions: (*IPackageAllUserProvisioningOptions)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type PackageManager struct {
	winrt.RtClass
	*IPackageManager
}

func NewPackageManager() *PackageManager {
	hs := winrt.NewHStr("Windows.Management.Deployment.PackageManager")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &PackageManager{
		RtClass:         winrt.RtClass{PInspect: p},
		IPackageManager: (*IPackageManager)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type PackageManagerDebugSettings struct {
	winrt.RtClass
	*IPackageManagerDebugSettings
}

type PackageUserInformation struct {
	winrt.RtClass
	*IPackageUserInformation
}

type PackageVolume struct {
	winrt.RtClass
	*IPackageVolume
}

type RegisterPackageOptions struct {
	winrt.RtClass
	*IRegisterPackageOptions
}

func NewRegisterPackageOptions() *RegisterPackageOptions {
	hs := winrt.NewHStr("Windows.Management.Deployment.RegisterPackageOptions")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &RegisterPackageOptions{
		RtClass:                 winrt.RtClass{PInspect: p},
		IRegisterPackageOptions: (*IRegisterPackageOptions)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}

type StagePackageOptions struct {
	winrt.RtClass
	*IStagePackageOptions
}

func NewStagePackageOptions() *StagePackageOptions {
	hs := winrt.NewHStr("Windows.Management.Deployment.StagePackageOptions")
	var p *win32.IInspectable
	hr := win32.RoActivateInstance(hs.Ptr, &p)
	if win32.FAILED(hr) {
		log.Panic("?")
	}
	result := &StagePackageOptions{
		RtClass:              winrt.RtClass{PInspect: p},
		IStagePackageOptions: (*IStagePackageOptions)(unsafe.Pointer(p))}
	com.AddToScope(result)
	return result
}
