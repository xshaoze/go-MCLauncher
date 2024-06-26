package Minecraft

import (
	"io"
	"net/http"
)

type latest struct {
	// 最新发布版
	Release string `json:"release"`
	// 最新快照版
	Snapshot string `json:"snapshot"`
}
type version struct {
	Id          string `json:"id"`   // 版本号
	Type        string `json:"type"` // Type 类别(快照:"snapshot"|发行版:"release"|远古版:"")
	Url         string `json:"url"`  // 该版本的json文件下载地址
	Time        string `json:"time"`
	ReleaseTime string `json:"releaseTime"`
}

type versionManifest struct {
	Latest   latest    `json:"latest"`
	Versions []version `json:"versions"`
}

func GetGameVersionsList() {
	url := "https://piston-meta.mojang.com/mc/game/version_manifest.json"
	get, err := http.Get(url)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(get.Body)

	//body, err := io.ReadAll(get.Body)
	//if err != nil {
	//	return
	//}
	//var versionManifest versionManifest
	//err = json.Unmarshal(body, &versionManifest)
	//if err != nil {
	//	return
	//}
}
