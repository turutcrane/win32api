// Code generated by "stringer -output win32api_const_string.go -type MessageId"; DO NOT EDIT.

package win32api

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WmActivate-6]
	_ = x[WmActivateapp-28]
	_ = x[WmAfxfirst-864]
	_ = x[WmAfxlast-895]
	_ = x[WmApp-32768]
	_ = x[WmAppcommand-793]
	_ = x[WmAskcbformatname-780]
	_ = x[WmCanceljournal-75]
	_ = x[WmCancelmode-31]
	_ = x[WmCapturechanged-533]
	_ = x[WmChangecbchain-781]
	_ = x[WmChangeuistate-295]
	_ = x[WmChar-258]
	_ = x[WmChartoitem-47]
	_ = x[WmChildactivate-34]
	_ = x[WmClear-771]
	_ = x[WmClipboardupdate-797]
	_ = x[WmClose-16]
	_ = x[WmCommand-273]
	_ = x[WmCommnotify-68]
	_ = x[WmCompacting-65]
	_ = x[WmCompareitem-57]
	_ = x[WmContextmenu-123]
	_ = x[WmCopy-769]
	_ = x[WmCopydata-74]
	_ = x[WmCreate-1]
	_ = x[WmCtlcolorbtn-309]
	_ = x[WmCtlcolordlg-310]
	_ = x[WmCtlcoloredit-307]
	_ = x[WmCtlcolorlistbox-308]
	_ = x[WmCtlcolormsgbox-306]
	_ = x[WmCtlcolorscrollbar-311]
	_ = x[WmCtlcolorstatic-312]
	_ = x[WmCut-768]
	_ = x[WmDeadchar-259]
	_ = x[WmDeleteitem-45]
	_ = x[WmDestroy-2]
	_ = x[WmDestroyclipboard-775]
	_ = x[WmDevicechange-537]
	_ = x[WmDevmodechange-27]
	_ = x[WmDisplaychange-126]
	_ = x[WmDpichanged-736]
	_ = x[WmDrawclipboard-776]
	_ = x[WmDrawitem-43]
	_ = x[WmDropfiles-563]
	_ = x[WmDwmcolorizationcolorchanged-800]
	_ = x[WmDwmcompositionchanged-798]
	_ = x[WmDwmncrenderingchanged-799]
	_ = x[WmDwmsendiconiclivepreviewbitmap-806]
	_ = x[WmDwmsendiconicthumbnail-803]
	_ = x[WmDwmwindowmaximizedchange-801]
	_ = x[WmEnable-10]
	_ = x[WmEndsession-22]
	_ = x[WmEnteridle-289]
	_ = x[WmEntermenuloop-529]
	_ = x[WmEntersizemove-561]
	_ = x[WmErasebkgnd-20]
	_ = x[WmExitmenuloop-530]
	_ = x[WmExitsizemove-562]
	_ = x[WmFontchange-29]
	_ = x[WmGesture-281]
	_ = x[WmGesturenotify-282]
	_ = x[WmGetdlgcode-135]
	_ = x[WmGetfont-49]
	_ = x[WmGethotkey-51]
	_ = x[WmGeticon-127]
	_ = x[WmGetminmaxinfo-36]
	_ = x[WmGetobject-61]
	_ = x[WmGettext-13]
	_ = x[WmGettextlength-14]
	_ = x[WmGettitlebarinfoex-831]
	_ = x[WmHandheldfirst-856]
	_ = x[WmHandheldlast-863]
	_ = x[WmHelp-83]
	_ = x[WmHotkey-786]
	_ = x[WmHscroll-276]
	_ = x[WmHscrollclipboard-782]
	_ = x[WmIconerasebkgnd-39]
	_ = x[WmImeChar-646]
	_ = x[WmImeComposition-271]
	_ = x[WmImeCompositionfull-644]
	_ = x[WmImeControl-643]
	_ = x[WmImeEndcomposition-270]
	_ = x[WmImeKeydown-656]
	_ = x[WmImeKeylast-271]
	_ = x[WmImeKeyup-657]
	_ = x[WmImeNotify-642]
	_ = x[WmImeRequest-648]
	_ = x[WmImeSelect-645]
	_ = x[WmImeSetcontext-641]
	_ = x[WmImeStartcomposition-269]
	_ = x[WmInitdialog-272]
	_ = x[WmInitmenu-278]
	_ = x[WmInitmenupopup-279]
	_ = x[WmInput-255]
	_ = x[WmInputDeviceChange-254]
	_ = x[WmInputlangchange-81]
	_ = x[WmInputlangchangerequest-80]
	_ = x[WmKeydown-256]
	_ = x[WmKeyfirst-256]
	_ = x[WmKeylast-265]
	_ = x[WmKeyup-257]
	_ = x[WmKillfocus-8]
	_ = x[WmLbuttondblclk-515]
	_ = x[WmLbuttondown-513]
	_ = x[WmLbuttonup-514]
	_ = x[WmMbuttondblclk-521]
	_ = x[WmMbuttondown-519]
	_ = x[WmMbuttonup-520]
	_ = x[WmMdiactivate-546]
	_ = x[WmMdicascade-551]
	_ = x[WmMdicreate-544]
	_ = x[WmMdidestroy-545]
	_ = x[WmMdigetactive-553]
	_ = x[WmMdiiconarrange-552]
	_ = x[WmMdimaximize-549]
	_ = x[WmMdinext-548]
	_ = x[WmMdirefreshmenu-564]
	_ = x[WmMdirestore-547]
	_ = x[WmMdisetmenu-560]
	_ = x[WmMditile-550]
	_ = x[WmMeasureitem-44]
	_ = x[WmMenuchar-288]
	_ = x[WmMenucommand-294]
	_ = x[WmMenudrag-291]
	_ = x[WmMenugetobject-292]
	_ = x[WmMenurbuttonup-290]
	_ = x[WmMenuselect-287]
	_ = x[WmMouseactivate-33]
	_ = x[WmMousefirst-512]
	_ = x[WmMousehover-673]
	_ = x[WmMousehwheel-526]
	_ = x[WmMouselast-526]
	_ = x[WmMouseleave-675]
	_ = x[WmMousemove-512]
	_ = x[WmMousewheel-522]
	_ = x[WmMove-3]
	_ = x[WmMoving-534]
	_ = x[WmNcactivate-134]
	_ = x[WmNccalcsize-131]
	_ = x[WmNccreate-129]
	_ = x[WmNcdestroy-130]
	_ = x[WmNchittest-132]
	_ = x[WmNclbuttondblclk-163]
	_ = x[WmNclbuttondown-161]
	_ = x[WmNclbuttonup-162]
	_ = x[WmNcmbuttondblclk-169]
	_ = x[WmNcmbuttondown-167]
	_ = x[WmNcmbuttonup-168]
	_ = x[WmNcmousehover-672]
	_ = x[WmNcmouseleave-674]
	_ = x[WmNcmousemove-160]
	_ = x[WmNcpaint-133]
	_ = x[WmNcrbuttondblclk-166]
	_ = x[WmNcrbuttondown-164]
	_ = x[WmNcrbuttonup-165]
	_ = x[WmNcxbuttondblclk-173]
	_ = x[WmNcxbuttondown-171]
	_ = x[WmNcxbuttonup-172]
	_ = x[WmNextdlgctl-40]
	_ = x[WmNextmenu-531]
	_ = x[WmNotify-78]
	_ = x[WmNotifyformat-85]
	_ = x[WmNull-0]
	_ = x[WmPaint-15]
	_ = x[WmPaintclipboard-777]
	_ = x[WmPainticon-38]
	_ = x[WmPalettechanged-785]
	_ = x[WmPaletteischanging-784]
	_ = x[WmParentnotify-528]
	_ = x[WmPaste-770]
	_ = x[WmPenwinfirst-896]
	_ = x[WmPenwinlast-911]
	_ = x[WmPower-72]
	_ = x[WmPowerbroadcast-536]
	_ = x[WmPrint-791]
	_ = x[WmPrintclient-792]
	_ = x[WmQuerydragicon-55]
	_ = x[WmQueryendsession-17]
	_ = x[WmQuerynewpalette-783]
	_ = x[WmQueryopen-19]
	_ = x[WmQueryuistate-297]
	_ = x[WmQueuesync-35]
	_ = x[WmQuit-18]
	_ = x[WmRbuttondblclk-518]
	_ = x[WmRbuttondown-516]
	_ = x[WmRbuttonup-517]
	_ = x[WmRenderallformats-774]
	_ = x[WmRenderformat-773]
	_ = x[WmSetcursor-32]
	_ = x[WmSetfocus-7]
	_ = x[WmSetfont-48]
	_ = x[WmSethotkey-50]
	_ = x[WmSeticon-128]
	_ = x[WmSetredraw-11]
	_ = x[WmSettext-12]
	_ = x[WmSettingchange-26]
	_ = x[WmShowwindow-24]
	_ = x[WmSize-5]
	_ = x[WmSizeclipboard-779]
	_ = x[WmSizing-532]
	_ = x[WmSpoolerstatus-42]
	_ = x[WmStylechanged-125]
	_ = x[WmStylechanging-124]
	_ = x[WmSyncpaint-136]
	_ = x[WmSyschar-262]
	_ = x[WmSyscolorchange-21]
	_ = x[WmSyscommand-274]
	_ = x[WmSysdeadchar-263]
	_ = x[WmSyskeydown-260]
	_ = x[WmSyskeyup-261]
	_ = x[WmTabletFirst-704]
	_ = x[WmTabletLast-735]
	_ = x[WmTcard-82]
	_ = x[WmThemechanged-794]
	_ = x[WmTimechange-30]
	_ = x[WmTimer-275]
	_ = x[WmTouch-576]
	_ = x[WmUndo-772]
	_ = x[WmUnichar-265]
	_ = x[WmUninitmenupopup-293]
	_ = x[WmUpdateuistate-296]
	_ = x[WmUser-1024]
	_ = x[WmUserchanged-84]
	_ = x[WmVkeytoitem-46]
	_ = x[WmVscroll-277]
	_ = x[WmVscrollclipboard-778]
	_ = x[WmWindowposchanged-71]
	_ = x[WmWindowposchanging-70]
	_ = x[WmWininichange-26]
	_ = x[WmWtssessionChange-689]
	_ = x[WmXbuttondblclk-525]
	_ = x[WmXbuttondown-523]
	_ = x[WmXbuttonup-524]
}

const _MessageId_name = "WmNullWmCreateWmDestroyWmMoveWmSizeWmActivateWmSetfocusWmKillfocusWmEnableWmSetredrawWmSettextWmGettextWmGettextlengthWmPaintWmCloseWmQueryendsessionWmQuitWmQueryopenWmErasebkgndWmSyscolorchangeWmEndsessionWmShowwindowWmSettingchangeWmDevmodechangeWmActivateappWmFontchangeWmTimechangeWmCancelmodeWmSetcursorWmMouseactivateWmChildactivateWmQueuesyncWmGetminmaxinfoWmPainticonWmIconerasebkgndWmNextdlgctlWmSpoolerstatusWmDrawitemWmMeasureitemWmDeleteitemWmVkeytoitemWmChartoitemWmSetfontWmGetfontWmSethotkeyWmGethotkeyWmQuerydragiconWmCompareitemWmGetobjectWmCompactingWmCommnotifyWmWindowposchangingWmWindowposchangedWmPowerWmCopydataWmCanceljournalWmNotifyWmInputlangchangerequestWmInputlangchangeWmTcardWmHelpWmUserchangedWmNotifyformatWmContextmenuWmStylechangingWmStylechangedWmDisplaychangeWmGeticonWmSeticonWmNccreateWmNcdestroyWmNccalcsizeWmNchittestWmNcpaintWmNcactivateWmGetdlgcodeWmSyncpaintWmNcmousemoveWmNclbuttondownWmNclbuttonupWmNclbuttondblclkWmNcrbuttondownWmNcrbuttonupWmNcrbuttondblclkWmNcmbuttondownWmNcmbuttonupWmNcmbuttondblclkWmNcxbuttondownWmNcxbuttonupWmNcxbuttondblclkWmInputDeviceChangeWmInputWmKeydownWmKeyupWmCharWmDeadcharWmSyskeydownWmSyskeyupWmSyscharWmSysdeadcharWmKeylastWmImeStartcompositionWmImeEndcompositionWmImeCompositionWmInitdialogWmCommandWmSyscommandWmTimerWmHscrollWmVscrollWmInitmenuWmInitmenupopupWmGestureWmGesturenotifyWmMenuselectWmMenucharWmEnteridleWmMenurbuttonupWmMenudragWmMenugetobjectWmUninitmenupopupWmMenucommandWmChangeuistateWmUpdateuistateWmQueryuistateWmCtlcolormsgboxWmCtlcoloreditWmCtlcolorlistboxWmCtlcolorbtnWmCtlcolordlgWmCtlcolorscrollbarWmCtlcolorstaticWmMousefirstWmLbuttondownWmLbuttonupWmLbuttondblclkWmRbuttondownWmRbuttonupWmRbuttondblclkWmMbuttondownWmMbuttonupWmMbuttondblclkWmMousewheelWmXbuttondownWmXbuttonupWmXbuttondblclkWmMousehwheelWmParentnotifyWmEntermenuloopWmExitmenuloopWmNextmenuWmSizingWmCapturechangedWmMovingWmPowerbroadcastWmDevicechangeWmMdicreateWmMdidestroyWmMdiactivateWmMdirestoreWmMdinextWmMdimaximizeWmMditileWmMdicascadeWmMdiiconarrangeWmMdigetactiveWmMdisetmenuWmEntersizemoveWmExitsizemoveWmDropfilesWmMdirefreshmenuWmTouchWmImeSetcontextWmImeNotifyWmImeControlWmImeCompositionfullWmImeSelectWmImeCharWmImeRequestWmImeKeydownWmImeKeyupWmNcmousehoverWmMousehoverWmNcmouseleaveWmMouseleaveWmWtssessionChangeWmTabletFirstWmTabletLastWmDpichangedWmCutWmCopyWmPasteWmClearWmUndoWmRenderformatWmRenderallformatsWmDestroyclipboardWmDrawclipboardWmPaintclipboardWmVscrollclipboardWmSizeclipboardWmAskcbformatnameWmChangecbchainWmHscrollclipboardWmQuerynewpaletteWmPaletteischangingWmPalettechangedWmHotkeyWmPrintWmPrintclientWmAppcommandWmThemechangedWmClipboardupdateWmDwmcompositionchangedWmDwmncrenderingchangedWmDwmcolorizationcolorchangedWmDwmwindowmaximizedchangeWmDwmsendiconicthumbnailWmDwmsendiconiclivepreviewbitmapWmGettitlebarinfoexWmHandheldfirstWmHandheldlastWmAfxfirstWmAfxlastWmPenwinfirstWmPenwinlastWmUserWmApp"

var _MessageId_map = map[MessageId]string{
	0:     _MessageId_name[0:6],
	1:     _MessageId_name[6:14],
	2:     _MessageId_name[14:23],
	3:     _MessageId_name[23:29],
	5:     _MessageId_name[29:35],
	6:     _MessageId_name[35:45],
	7:     _MessageId_name[45:55],
	8:     _MessageId_name[55:66],
	10:    _MessageId_name[66:74],
	11:    _MessageId_name[74:85],
	12:    _MessageId_name[85:94],
	13:    _MessageId_name[94:103],
	14:    _MessageId_name[103:118],
	15:    _MessageId_name[118:125],
	16:    _MessageId_name[125:132],
	17:    _MessageId_name[132:149],
	18:    _MessageId_name[149:155],
	19:    _MessageId_name[155:166],
	20:    _MessageId_name[166:178],
	21:    _MessageId_name[178:194],
	22:    _MessageId_name[194:206],
	24:    _MessageId_name[206:218],
	26:    _MessageId_name[218:233],
	27:    _MessageId_name[233:248],
	28:    _MessageId_name[248:261],
	29:    _MessageId_name[261:273],
	30:    _MessageId_name[273:285],
	31:    _MessageId_name[285:297],
	32:    _MessageId_name[297:308],
	33:    _MessageId_name[308:323],
	34:    _MessageId_name[323:338],
	35:    _MessageId_name[338:349],
	36:    _MessageId_name[349:364],
	38:    _MessageId_name[364:375],
	39:    _MessageId_name[375:391],
	40:    _MessageId_name[391:403],
	42:    _MessageId_name[403:418],
	43:    _MessageId_name[418:428],
	44:    _MessageId_name[428:441],
	45:    _MessageId_name[441:453],
	46:    _MessageId_name[453:465],
	47:    _MessageId_name[465:477],
	48:    _MessageId_name[477:486],
	49:    _MessageId_name[486:495],
	50:    _MessageId_name[495:506],
	51:    _MessageId_name[506:517],
	55:    _MessageId_name[517:532],
	57:    _MessageId_name[532:545],
	61:    _MessageId_name[545:556],
	65:    _MessageId_name[556:568],
	68:    _MessageId_name[568:580],
	70:    _MessageId_name[580:599],
	71:    _MessageId_name[599:617],
	72:    _MessageId_name[617:624],
	74:    _MessageId_name[624:634],
	75:    _MessageId_name[634:649],
	78:    _MessageId_name[649:657],
	80:    _MessageId_name[657:681],
	81:    _MessageId_name[681:698],
	82:    _MessageId_name[698:705],
	83:    _MessageId_name[705:711],
	84:    _MessageId_name[711:724],
	85:    _MessageId_name[724:738],
	123:   _MessageId_name[738:751],
	124:   _MessageId_name[751:766],
	125:   _MessageId_name[766:780],
	126:   _MessageId_name[780:795],
	127:   _MessageId_name[795:804],
	128:   _MessageId_name[804:813],
	129:   _MessageId_name[813:823],
	130:   _MessageId_name[823:834],
	131:   _MessageId_name[834:846],
	132:   _MessageId_name[846:857],
	133:   _MessageId_name[857:866],
	134:   _MessageId_name[866:878],
	135:   _MessageId_name[878:890],
	136:   _MessageId_name[890:901],
	160:   _MessageId_name[901:914],
	161:   _MessageId_name[914:929],
	162:   _MessageId_name[929:942],
	163:   _MessageId_name[942:959],
	164:   _MessageId_name[959:974],
	165:   _MessageId_name[974:987],
	166:   _MessageId_name[987:1004],
	167:   _MessageId_name[1004:1019],
	168:   _MessageId_name[1019:1032],
	169:   _MessageId_name[1032:1049],
	171:   _MessageId_name[1049:1064],
	172:   _MessageId_name[1064:1077],
	173:   _MessageId_name[1077:1094],
	254:   _MessageId_name[1094:1113],
	255:   _MessageId_name[1113:1120],
	256:   _MessageId_name[1120:1129],
	257:   _MessageId_name[1129:1136],
	258:   _MessageId_name[1136:1142],
	259:   _MessageId_name[1142:1152],
	260:   _MessageId_name[1152:1164],
	261:   _MessageId_name[1164:1174],
	262:   _MessageId_name[1174:1183],
	263:   _MessageId_name[1183:1196],
	265:   _MessageId_name[1196:1205],
	269:   _MessageId_name[1205:1226],
	270:   _MessageId_name[1226:1245],
	271:   _MessageId_name[1245:1261],
	272:   _MessageId_name[1261:1273],
	273:   _MessageId_name[1273:1282],
	274:   _MessageId_name[1282:1294],
	275:   _MessageId_name[1294:1301],
	276:   _MessageId_name[1301:1310],
	277:   _MessageId_name[1310:1319],
	278:   _MessageId_name[1319:1329],
	279:   _MessageId_name[1329:1344],
	281:   _MessageId_name[1344:1353],
	282:   _MessageId_name[1353:1368],
	287:   _MessageId_name[1368:1380],
	288:   _MessageId_name[1380:1390],
	289:   _MessageId_name[1390:1401],
	290:   _MessageId_name[1401:1416],
	291:   _MessageId_name[1416:1426],
	292:   _MessageId_name[1426:1441],
	293:   _MessageId_name[1441:1458],
	294:   _MessageId_name[1458:1471],
	295:   _MessageId_name[1471:1486],
	296:   _MessageId_name[1486:1501],
	297:   _MessageId_name[1501:1515],
	306:   _MessageId_name[1515:1531],
	307:   _MessageId_name[1531:1545],
	308:   _MessageId_name[1545:1562],
	309:   _MessageId_name[1562:1575],
	310:   _MessageId_name[1575:1588],
	311:   _MessageId_name[1588:1607],
	312:   _MessageId_name[1607:1623],
	512:   _MessageId_name[1623:1635],
	513:   _MessageId_name[1635:1648],
	514:   _MessageId_name[1648:1659],
	515:   _MessageId_name[1659:1674],
	516:   _MessageId_name[1674:1687],
	517:   _MessageId_name[1687:1698],
	518:   _MessageId_name[1698:1713],
	519:   _MessageId_name[1713:1726],
	520:   _MessageId_name[1726:1737],
	521:   _MessageId_name[1737:1752],
	522:   _MessageId_name[1752:1764],
	523:   _MessageId_name[1764:1777],
	524:   _MessageId_name[1777:1788],
	525:   _MessageId_name[1788:1803],
	526:   _MessageId_name[1803:1816],
	528:   _MessageId_name[1816:1830],
	529:   _MessageId_name[1830:1845],
	530:   _MessageId_name[1845:1859],
	531:   _MessageId_name[1859:1869],
	532:   _MessageId_name[1869:1877],
	533:   _MessageId_name[1877:1893],
	534:   _MessageId_name[1893:1901],
	536:   _MessageId_name[1901:1917],
	537:   _MessageId_name[1917:1931],
	544:   _MessageId_name[1931:1942],
	545:   _MessageId_name[1942:1954],
	546:   _MessageId_name[1954:1967],
	547:   _MessageId_name[1967:1979],
	548:   _MessageId_name[1979:1988],
	549:   _MessageId_name[1988:2001],
	550:   _MessageId_name[2001:2010],
	551:   _MessageId_name[2010:2022],
	552:   _MessageId_name[2022:2038],
	553:   _MessageId_name[2038:2052],
	560:   _MessageId_name[2052:2064],
	561:   _MessageId_name[2064:2079],
	562:   _MessageId_name[2079:2093],
	563:   _MessageId_name[2093:2104],
	564:   _MessageId_name[2104:2120],
	576:   _MessageId_name[2120:2127],
	641:   _MessageId_name[2127:2142],
	642:   _MessageId_name[2142:2153],
	643:   _MessageId_name[2153:2165],
	644:   _MessageId_name[2165:2185],
	645:   _MessageId_name[2185:2196],
	646:   _MessageId_name[2196:2205],
	648:   _MessageId_name[2205:2217],
	656:   _MessageId_name[2217:2229],
	657:   _MessageId_name[2229:2239],
	672:   _MessageId_name[2239:2253],
	673:   _MessageId_name[2253:2265],
	674:   _MessageId_name[2265:2279],
	675:   _MessageId_name[2279:2291],
	689:   _MessageId_name[2291:2309],
	704:   _MessageId_name[2309:2322],
	735:   _MessageId_name[2322:2334],
	736:   _MessageId_name[2334:2346],
	768:   _MessageId_name[2346:2351],
	769:   _MessageId_name[2351:2357],
	770:   _MessageId_name[2357:2364],
	771:   _MessageId_name[2364:2371],
	772:   _MessageId_name[2371:2377],
	773:   _MessageId_name[2377:2391],
	774:   _MessageId_name[2391:2409],
	775:   _MessageId_name[2409:2427],
	776:   _MessageId_name[2427:2442],
	777:   _MessageId_name[2442:2458],
	778:   _MessageId_name[2458:2476],
	779:   _MessageId_name[2476:2491],
	780:   _MessageId_name[2491:2508],
	781:   _MessageId_name[2508:2523],
	782:   _MessageId_name[2523:2541],
	783:   _MessageId_name[2541:2558],
	784:   _MessageId_name[2558:2577],
	785:   _MessageId_name[2577:2593],
	786:   _MessageId_name[2593:2601],
	791:   _MessageId_name[2601:2608],
	792:   _MessageId_name[2608:2621],
	793:   _MessageId_name[2621:2633],
	794:   _MessageId_name[2633:2647],
	797:   _MessageId_name[2647:2664],
	798:   _MessageId_name[2664:2687],
	799:   _MessageId_name[2687:2710],
	800:   _MessageId_name[2710:2739],
	801:   _MessageId_name[2739:2765],
	803:   _MessageId_name[2765:2789],
	806:   _MessageId_name[2789:2821],
	831:   _MessageId_name[2821:2840],
	856:   _MessageId_name[2840:2855],
	863:   _MessageId_name[2855:2869],
	864:   _MessageId_name[2869:2879],
	895:   _MessageId_name[2879:2888],
	896:   _MessageId_name[2888:2901],
	911:   _MessageId_name[2901:2913],
	1024:  _MessageId_name[2913:2919],
	32768: _MessageId_name[2919:2924],
}

func (i MessageId) String() string {
	if str, ok := _MessageId_map[i]; ok {
		return str
	}
	return "MessageId(" + strconv.FormatInt(int64(i), 10) + ")"
}
