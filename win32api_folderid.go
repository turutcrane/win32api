package win32api

import (
	"unsafe"

	"github.com/turutcrane/win32api/internal/win32guid"
)

var (
	FolderidAccountPictures = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidAccountPictures)))
	FolderidAddNewPrograms = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidAddNewPrograms)))
	FolderidAdminTools = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidAdminTools)))
	FolderidAllAppMods = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidAllAppMods)))
	FolderidAppCaptures = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidAppCaptures)))
	FolderidAppDataDesktop = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidAppDataDesktop)))
	FolderidAppDataDocuments = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidAppDataDocuments)))
	FolderidAppDataFavorites = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidAppDataFavorites)))
	FolderidAppDataProgramData = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidAppDataProgramData)))
	FolderidAppsFolder = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidAppsFolder)))
	FolderidApplicationShortcuts = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidApplicationShortcuts)))
	FolderidAppUpdates = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidAppUpdates)))
	FolderidCameraRoll = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidCameraRoll)))
	FolderidCameraRollLibrary = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidCameraRollLibrary)))
	FolderidCDBurning = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidCDBurning)))
	FolderidChangeRemovePrograms = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidChangeRemovePrograms)))
	FolderidCommonAdminTools = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidCommonAdminTools)))
	FolderidCommonOEMLinks = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidCommonOEMLinks)))
	FolderidCommonPrograms = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidCommonPrograms)))
	FolderidCommonStartMenu = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidCommonStartMenu)))
	FolderidCommonStartMenuPlaces = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidCommonStartMenuPlaces)))
	FolderidCommonStartup = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidCommonStartup)))
	FolderidCommonTemplates = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidCommonTemplates)))
	FolderidComputerFolder = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidComputerFolder)))
	FolderidConflictFolder = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidConflictFolder)))
	FolderidConnectionsFolder = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidConnectionsFolder)))
	FolderidContacts = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidContacts)))
	FolderidControlPanelFolder = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidControlPanelFolder)))
	FolderidCookies = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidCookies)))
	FolderidCurrentAppMods = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidCurrentAppMods)))
	FolderidDesktop = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidDesktop)))
	FolderidDevelopmentFiles = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidDevelopmentFiles)))
	FolderidDevice = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidDevice)))
	FolderidDeviceMetadataStore = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidDeviceMetadataStore)))
	FolderidDocuments = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidDocuments)))
	FolderidDocumentsLibrary = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidDocumentsLibrary)))
	FolderidDownloads = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidDownloads)))
	FolderidFavorites = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidFavorites)))
	FolderidFonts = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidFonts)))
	FolderidGames = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidGames)))
	FolderidGameTasks = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidGameTasks)))
	FolderidHistory = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidHistory)))
	FolderidHomeGroup = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidHomeGroup)))
	FolderidHomeGroupCurrentUser = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidHomeGroupCurrentUser)))
	FolderidImplicitAppShortcuts = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidImplicitAppShortcuts)))
	FolderidInternetCache = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidInternetCache)))
	FolderidInternetFolder = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidInternetFolder)))
	FolderidLibraries = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidLibraries)))
	FolderidLinks = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidLinks)))
	FolderidLocalAppData = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidLocalAppData)))
	FolderidLocalAppDataLow = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidLocalAppDataLow)))
	FolderidLocalDocuments = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidLocalDocuments)))
	FolderidLocalDownloads = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidLocalDownloads)))
	FolderidLocalizedResourcesDir = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidLocalizedResourcesDir)))
	FolderidLocalMusic = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidLocalMusic)))
	FolderidLocalPictures = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidLocalPictures)))
	FolderidLocalVideos = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidLocalVideos)))
	FolderidMusic = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidMusic)))
	FolderidMusicLibrary = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidMusicLibrary)))
	FolderidNetHood = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidNetHood)))
	FolderidNetworkFolder = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidNetworkFolder)))
	FolderidObjects3D = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidObjects3D)))
	FolderidOneDrive = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidOneDrive)))
	FolderidOriginalImages = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidOriginalImages)))
	FolderidPhotoAlbums = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPhotoAlbums)))
	FolderidPictures = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPictures)))
	FolderidPicturesLibrary = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPicturesLibrary)))
	FolderidPlaylists = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPlaylists)))
	FolderidPrintHood = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPrintHood)))
	FolderidPrintersFolder = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPrintersFolder)))
	FolderidProfile = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidProfile)))
	FolderidProgramData = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidProgramData)))
	FolderidProgramFiles = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidProgramFiles)))
	FolderidProgramFilesX64 = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidProgramFilesX64)))
	FolderidProgramFilesX86 = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidProgramFilesX86)))
	FolderidProgramFilesCommon = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidProgramFilesCommon)))
	FolderidProgramFilesCommonX64 = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidProgramFilesCommonX64)))
	FolderidProgramFilesCommonX86 = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidProgramFilesCommonX86)))
	FolderidPrograms = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPrograms)))
	FolderidPublic = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPublic)))
	FolderidPublicDesktop = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPublicDesktop)))
	FolderidPublicDocuments = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPublicDocuments)))
	FolderidPublicDownloads = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPublicDownloads)))
	FolderidPublicGameTasks = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPublicGameTasks)))
	FolderidPublicLibraries = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPublicLibraries)))
	FolderidPublicMusic = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPublicMusic)))
	FolderidPublicPictures = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPublicPictures)))
	FolderidPublicRingtones = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPublicRingtones)))
	FolderidPublicUserTiles = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPublicUserTiles)))
	FolderidPublicVideos = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidPublicVideos)))
	FolderidQuickLaunch = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidQuickLaunch)))
	FolderidRecent = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidRecent)))
	FolderidRecordedCalls = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidRecordedCalls)))
	FolderidRecordedTVLibrary = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidRecordedTVLibrary)))
	FolderidRecycleBinFolder = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidRecycleBinFolder)))
	FolderidResourceDir = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidResourceDir)))
	FolderidRetailDemo = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidRetailDemo)))
	FolderidRingtones = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidRingtones)))
	FolderidRoamingAppData = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidRoamingAppData)))
	FolderidRoamingTiles = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidRoamingTiles)))
	FolderidRoamedTileImages = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidRoamedTileImages)))
	FolderidSampleMusic = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSampleMusic)))
	FolderidSamplePictures = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSamplePictures)))
	FolderidSamplePlaylists = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSamplePlaylists)))
	FolderidSampleVideos = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSampleVideos)))
	FolderidSavedGames = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSavedGames)))
	FolderidSavedPictures = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSavedPictures)))
	FolderidSavedPicturesLibrary = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSavedPicturesLibrary)))
	FolderidSavedSearches = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSavedSearches)))
	FolderidScreenshots = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidScreenshots)))
	FolderidSEARCH_MAPI = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSEARCH_MAPI)))
	FolderidSEARCH_CSC = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSEARCH_CSC)))
	FolderidSearchHistory = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSearchHistory)))
	FolderidSearchHome = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSearchHome)))
	FolderidSearchTemplates = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSearchTemplates)))
	FolderidSendTo = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSendTo)))
	FolderidSidebarDefaultParts = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSidebarDefaultParts)))
	FolderidSidebarParts = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSidebarParts)))
	FolderidSkyDrive = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSkyDrive)))
	FolderidSkyDriveCameraRoll = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSkyDriveCameraRoll)))
	FolderidSkyDriveDocuments = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSkyDriveDocuments)))
	FolderidSkyDriveMusic = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSkyDriveMusic)))
	FolderidSkyDrivePictures = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSkyDrivePictures)))
	FolderidStartMenu = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidStartMenu)))
	FolderidStartMenuAllPrograms = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidStartMenuAllPrograms)))
	FolderidStartup = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidStartup)))
	FolderidSyncManagerFolder = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSyncManagerFolder)))
	FolderidSyncResultsFolder = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSyncResultsFolder)))
	FolderidSyncSetupFolder = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSyncSetupFolder)))
	FolderidSystem = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSystem)))
	FolderidSystemX86 = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidSystemX86)))
	FolderidTemplates = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidTemplates)))
	FolderidUserPinned = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidUserPinned)))
	FolderidUserProfiles = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidUserProfiles)))
	FolderidUserProgramFiles = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidUserProgramFiles)))
	FolderidUserProgramFilesCommon = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidUserProgramFilesCommon)))
	FolderidUsersFiles = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidUsersFiles)))
	FolderidUsersLibraries = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidUsersLibraries)))
	FolderidVideos = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidVideos)))
	FolderidVideosLibrary = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidVideosLibrary)))
	FolderidWindows = REFKNOWNFOLDERID(uintptr(unsafe.Pointer(win32guid.FolderidWindows)))
)
