package bootstrap

import (
	"github.com/bh-qt/alist/conf"
	"github.com/bh-qt/alist/utils"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

// read config file
func ReadConf(config string) bool {
	log.Infof("读取配置文件...")
	if !utils.Exists(config) {
		log.Infof("找不到配置文件:%s", config)
		if !Write(config) {
			return false
		}
	}
	confFile, err := ioutil.ReadFile(config)
	if err != nil {
		log.Errorf("读取配置文件时发生错误:%s", err.Error())
		return false
	}
	err = yaml.Unmarshal(confFile, conf.Conf)
	if err != nil {
		log.Errorf("加载配置文件时发生错误:%s", err.Error())
		return false
	}
	log.Debugf("config:%+v", conf.Conf)
	conf.Conf.Info.Roots = utils.GetNames()
	conf.Origins = strings.Split(conf.Conf.Server.SiteUrl, ",")
	conf.AllowProxies = strings.Split(conf.Conf.Server.AllowProxy,",")
	return true
}
func Write(path string) bool {
	log.Infof("创建默认配置文件")
	file, err := utils.CreatNestedFile(path)
	if err != nil {
		log.Errorf("无法创建配置文件, %s", err)
		return false
	}
	defer func() {
		_ = file.Close()
	}()
	str := `
info:
  title: AList #标题
  logo: "" #网站logo 如果填写,则会替换掉默认的
  footer_text: Xhofe's Blog #网页底部文字
  footer_url: https://www.nn.ci #网页底部文字链接
  music_img: https://img.xhofe.top/2020/12/19/0f8b57866bdb5.gif #预览音乐文件时的图片
  check_update: true #前端是否显示更新
  script: #自定义脚本,可以是脚本的链接，也可以直接是脚本内容
  autoplay: true #视频是否自动播放
  preview:
    text: [txt,htm,html,xml,java,properties,sql,js,md,json,conf,ini,vue,php,py,bat,gitignore,yml,go,sh,c,cpp,h,hpp] #要预览的文本文件的后缀，可以自行添加
server:
  address: "0.0.0.0"
  port: "5244"
  search: true
  download: true
  static: dist
  site_url: '*'
  password: password #用于重建目录
  allow_proxy: vtt
ali_drive:
  api_url: https://api.aliyundrive.com/v2
  max_files_count: 100
  drives:
  - refresh_token: xxx #refresh_token
    root_folder: root #根目录的file_id
    name: drive0 #盘名，多个盘不可重复，这里只是示例，不是一定要叫这个名字，可随意修改
    password: pass #该盘密码，空（''）则不设密码，修改需要重建生效
    hide: false #是否在主页隐藏该盘，不可全部隐藏，至少暴露一个
database:
  type: sqlite3
  dBFile: alist.db
`
	_, err = file.WriteString(str)
	if err != nil {
		log.Errorf("无法写入配置文件, %s", err)
		return false
	}
	return true
}
