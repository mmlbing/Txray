package xray

import (
	"Txray/core"
	"Txray/core/manage"
	"Txray/core/setting"
	"Txray/core/protocols"
	"Txray/log"
)

func Generate(key string) {
	testUrl := setting.TestUrl()
	testTimeout := setting.TestTimeout()
	manager := manage.Manager
	indexList := core.IndexList(key, manager.NodeLen())
	if len(indexList) == 0 {
		log.Warn("没有选取到节点")
	} else if len(indexList) == 1 {
		index := indexList[0]
		node := manager.GetNode(index)
		manager.SetSelectedIndex(index)
		manager.Save()
		cfg := gen(node.Protocol)
		if cfg {
			result, status := TestNode(testUrl, setting.Socks(), testTimeout)
			log.Infof("%6s [ %s ] 延迟: %dms", status, testUrl, result)
			log.Infof("成功")
		} else {
			log.Infof("失败")
		}
	} else {
		log.Infof("Not Supported.")
	}
}

func gen(node protocols.Protocol) bool {
	log.Infof("生成配置文件: [%s]", node.GetName())
	switch node.GetProtocolMode() {
	case protocols.ModeShadowSocks, protocols.ModeTrojan, protocols.ModeVMess, protocols.ModeSocks, protocols.ModeVLESS, protocols.ModeVMessAEAD:
		file := GenConfig(node)
		log.Infof("配置文件生成到: [%s]", file)
		return true
	default:
		log.Infof("暂不支持%v协议", node.GetProtocolMode())
		return false
	}
}