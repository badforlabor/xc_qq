package main

import (
	_ "fmt"
)

import (
	"fmt"
	_ "github.com/codyguo/xcgui"
	_ "github.com/codyguo/xcgui/wke"
	xcgui2 "github.com/codyguo/xcgui/xc"
	_ "github.com/lxn/win"
)

const (
	XC_NAME  = "name1"
	XC_NAME2 = "name2"
	XC_NAME3 = "name3"
	XC_NAME4 = "name4"
	XC_NAME5 = "name5"
	XC_NAME6 = "name6"
)

var hWindow xcgui2.HWINDOW
var m_hTree xcgui2.HELE

//var m_pTemplate_group *xcgui2.template_info_i
var m_hVip, m_hQZone, m_hAvatar, m_hAvatarLarge, m_hExpand, m_hExpandNo xcgui2.HIMAGE

func main() {

	xcgui2.XC_EnableDebugFile(true)

	xcgui2.XC_LoadResource("resource.xml", "")

	// xcgui.MsgBox(nil, "警告", "这是一条警告信息!", xcgui.MsgBoxIconWarning)

	fmt.Print("start.")
	hxcgui := xcgui2.XC_LoadLayout("Layout.xml", 0)

	if hxcgui == 0 {
		fmt.Println("error.")
	}

	hWindow = xcgui2.HWINDOW(hxcgui)

	m_hTree = xcgui2.HELE(xcgui2.XC_GetObjectByID(hWindow, 101))
	//XTree_EnableConnectLine(m_hTree,FALSE,FALSE);
	xcgui2.XSView_ShowSBarH(m_hTree, false)
	xcgui2.XSView_ShowSBarH(m_hTree, false)
	xcgui2.XTree_SetIndentation(m_hTree, 0)
	xcgui2.XTree_SetItemHeightDefault(m_hTree, 28, 54)
	xcgui2.XTree_SetItemTemplateXML(m_hTree, "xml-template-test\\Tree_Item_friend.xml")
	xcgui2.XTree_SetItemTemplateXMLSel(m_hTree, "xml-template-test\\Tree_Item_friend_sel.xml")
	//m_pTemplate_group = xcgui2.XC_LoadTemplate(xcgui2.XC_TREE, "xml-template-test\\Tree_Item_group.xml")

	m_hVip = xcgui2.XImage_LoadFile("QQImage\\SuperVIP_LIGHT.png", false)
	m_hQZone = xcgui2.XImage_LoadFile("QQImage\\QQZone.png", false)
	m_hAvatar = xcgui2.XImage_LoadFile("QQImage\\avatar.png", false)
	m_hAvatarLarge = xcgui2.XImage_LoadFile("QQImage\\avatar_large.png", false)

	m_hExpand = xcgui2.XImage_LoadFile("Image\\expand.png", false)
	m_hExpandNo = xcgui2.XImage_LoadFile("Image\\expandno.png", false)
	xcgui2.XImage_EnableAutoDestroy(m_hExpandNo, false)
	xcgui2.XImage_EnableAutoDestroy(m_hExpand, false)

	hAdapter := xcgui2.XAdapterTree_Create()
	xcgui2.XTree_BindAdapter(m_hTree, hAdapter)
	xcgui2.XAdapterTree_AddColumn(hAdapter, XC_NAME)  //昵称 分组名
	xcgui2.XAdapterTree_AddColumn(hAdapter, XC_NAME2) //签名
	xcgui2.XAdapterTree_AddColumn(hAdapter, XC_NAME3) //小头像
	xcgui2.XAdapterTree_AddColumn(hAdapter, XC_NAME4) //大头像
	xcgui2.XAdapterTree_AddColumn(hAdapter, XC_NAME5) //VIP 图标
	xcgui2.XAdapterTree_AddColumn(hAdapter, XC_NAME6) //空间图标

	var nGroupID int
	for iGroup := 0; iGroup < 5; iGroup++ {
		buf := fmt.Sprintf("好友分组-%d", iGroup)
		nGroupID = xcgui2.XAdapterTree_InsertItemText(hAdapter, buf, xcgui2.XC_ID_ROOT, xcgui2.XC_ID_LAST)
		xcgui2.XTree_SetItemHeight(m_hTree, nGroupID, 26, 26)
		for i := 0; i < 3; i++ {
			buf = fmt.Sprintf("我的好友-%d", i)
			nItemID := xcgui2.XAdapterTree_InsertItemText(hAdapter, buf, nGroupID, xcgui2.XC_ID_LAST)
			xcgui2.XAdapterTree_SetItemTextEx(hAdapter, nItemID, XC_NAME2, "我的个性签名!")
			xcgui2.XAdapterTree_SetItemImageEx(hAdapter, nItemID, XC_NAME5, m_hVip)
			xcgui2.XAdapterTree_SetItemImageEx(hAdapter, nItemID, XC_NAME6, m_hQZone)
			xcgui2.XAdapterTree_SetItemImageEx(hAdapter, nItemID, XC_NAME3, m_hAvatar)
			xcgui2.XAdapterTree_SetItemImageEx(hAdapter, nItemID, XC_NAME4, m_hAvatarLarge)
		}
	}

	//xcgui2.XWnd_RegEventC(m_hTree, xcgui2.XE_TREE_SELECT, xcgui2.CallBack(OnTreeSelect))
	//XEle_RegEventCPP(m_hTree,XE_TREE_EXPAND,&CMyWindowTree::OnTreeExpand);
	//XEle_RegEventCPP(m_hTree,XE_TREE_TEMP_CREATE,&CMyWindowTree::OnTemplateCreate);
	//XEle_RegEventCPP(m_hTree,XE_TREE_TEMP_CREATE_END,&CMyWindowTree::OnTreeTemplateCreateEnd);
	//XEle_RegEventCPP(m_hTree,XE_TREE_TEMP_DESTROY,&CMyWindowTree::OnTemplateDestroy);
	//XEle_RegEventCPP(m_hTree,XE_TREE_TEMP_ADJUST_COORDINATE,&CMyWindowTree::OnTemplateAdjustCoordinate);

	xcgui2.XWnd_AdjustLayout(hWindow)
	xcgui2.XWnd_ShowWindow(hWindow, xcgui2.SW_SHOW)

	xcgui2.XRunXCGUI()
	fmt.Println("exit.")
	xcgui2.XExitXCGUI()
}
func OnTreeSelect(nItemID int, pbHandled *bool) int {
	*pbHandled = true
	return 0
}
