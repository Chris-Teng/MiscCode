package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type userInfo struct {
	Phone      string
	AtNanjing  string
	Tracked    string
	Glryid     string
	Id         string
	CtrlMethod string
	CtrlStart  string
	CtrlEnd    string
	Zlbh       string
	Fxrylyxq   string
	Fxryzx     string
	Rksj       string
	Bz2        string
}

//要从服务器获取的id信息  convered from https://mholt.github.io/json-to-go/
type getIdInfo struct {
	Total int `json:"total"`
	Data  []struct {
		SearchValue interface{} `json:"searchValue"`
		CreateBy    interface{} `json:"createBy"`
		CreateTime  interface{} `json:"createTime"`
		UpdateBy    interface{} `json:"updateBy"`
		UpdateTime  interface{} `json:"updateTime"`
		Remark      interface{} `json:"remark"`
		Params      struct {
		} `json:"params"`
		ID           string      `json:"id"`
		Glryid       string      `json:"glryid"`
		Zlbh         string      `json:"zlbh"`
		Fxrysjly     interface{} `json:"fxrysjly"`
		Fxrysjlb     interface{} `json:"fxrysjlb"`
		Fxryly       interface{} `json:"fxryly"`
		Fxrylb       string      `json:"fxrylb"`
		Fxrylyxq     string      `json:"fxrylyxq"`
		Fxryxflx     string      `json:"fxryxflx"`
		Fxryzx       string      `json:"fxryzx"`
		Sjlb         string      `json:"sjlb"`
		Sjly         string      `json:"sjly"`
		Xm           string      `json:"xm"`
		Xb           interface{} `json:"xb"`
		Zjlx         interface{} `json:"zjlx"`
		Zjhm         string      `json:"zjhm"`
		Sjhm         string      `json:"sjhm"`
		Yys          interface{} `json:"yys"`
		Rylb         interface{} `json:"rylb"`
		Rysx         interface{} `json:"rysx"`
		Cph          interface{} `json:"cph"`
		Rylnjtxx     interface{} `json:"rylnjtxx"`
		Ryccjcsj     interface{} `json:"ryccjcsj"`
		Rymcjcsj     interface{} `json:"rymcjcsj"`
		Ryjcfs       interface{} `json:"ryjcfs"`
		Rydz         string      `json:"rydz"`
		Sjdwdz       interface{} `json:"sjdwdz"`
		Ryzzbh       interface{} `json:"ryzzbh"`
		Ryzzmc       string      `json:"ryzzmc"`
		Szxq         string      `json:"szxq"`
		Szxqmc       string      `json:"szxqmc"`
		Szjz         string      `json:"szjz"`
		Szjzmc       string      `json:"szjzmc"`
		Szsq         string      `json:"szsq"`
		Szsqmc       string      `json:"szsqmc"`
		Szwg         interface{} `json:"szwg"`
		Szwgmc       interface{} `json:"szwgmc"`
		Rksj         string      `json:"rksj"`
		Fbsj         interface{} `json:"fbsj"`
		Qssj         interface{} `json:"qssj"`
		Yjsj         interface{} `json:"yjsj"`
		Fksj         interface{} `json:"fksj"`
		Szdw         string      `json:"szdw"`
		Szdwmc       string      `json:"szdwmc"`
		Ryzt         string      `json:"ryzt"`
		Delstatus    interface{} `json:"delstatus"`
		Cjr          interface{} `json:"cjr"`
		Cjrxm        string      `json:"cjrxm"`
		Cjdw         interface{} `json:"cjdw"`
		Cjdwmc       string      `json:"cjdwmc"`
		Cjsj         string      `json:"cjsj"`
		Xgr          interface{} `json:"xgr"`
		Xgrxm        interface{} `json:"xgrxm"`
		Xgdw         interface{} `json:"xgdw"`
		Xgdwmc       interface{} `json:"xgdwmc"`
		Xgsj         interface{} `json:"xgsj"`
		Unitcode     interface{} `json:"unitcode"`
		Chinaname    interface{} `json:"chinaname"`
		Yjmbdw       interface{} `json:"yjmbdw"`
		Yjmbdwmc     interface{} `json:"yjmbdwmc"`
		Fqrxm        interface{} `json:"fqrxm"`
		Fqdwmc       interface{} `json:"fqdwmc"`
		Fqsj         interface{} `json:"fqsj"`
		Fqrlxfs      interface{} `json:"fqrlxfs"`
		Fqyy         interface{} `json:"fqyy"`
		Ypjg         string      `json:"ypjg"`
		Zlly         string      `json:"zlly"`
		Zlqsqx       string      `json:"zlqsqx"`
		Gkqkname     string      `json:"gkqkname"`
		Bzdzcomp     interface{} `json:"bzdzcomp"`
		Zlfbsj       string      `json:"zlfbsj"`
		Yszxqmc      interface{} `json:"yszxqmc"`
		Yszjzmc      interface{} `json:"yszjzmc"`
		Yszsqmc      interface{} `json:"yszsqmc"`
		Xpdw         interface{} `json:"xpdw"`
		Sjlbname     string      `json:"sjlbname"`
		Qxname       interface{} `json:"qxname"`
		Jdname       interface{} `json:"jdname"`
		Sqname       interface{} `json:"sqname"`
		Rysxname     string      `json:"rysxname"`
		Zlqh         string      `json:"zlqh"`
		Zlbt         string      `json:"zlbt"`
		Zlnr         interface{} `json:"zlnr"`
		Zlyq         interface{} `json:"zlyq"`
		Jsdwmc       interface{} `json:"jsdwmc"`
		Code         interface{} `json:"code"`
		Num          interface{} `json:"num"`
		Orgname      interface{} `json:"orgname"`
		Zllx         string      `json:"zllx"`
		Fkqx         interface{} `json:"fkqx"`
		Fqdw         interface{} `json:"fqdw"`
		Yjlx         interface{} `json:"yjlx"`
		Znzt         string      `json:"znzt"`
		Lnrq         interface{} `json:"lnrq"`
		Gkqk         interface{} `json:"gkqk"`
		Ksgkrq       interface{} `json:"ksgkrq"`
		Jsgkrq       interface{} `json:"jsgkrq"`
		Hsjccs       interface{} `json:"hsjccs"`
		Lxfsarray    interface{} `json:"lxfsarray"`
		Szsf         interface{} `json:"szsf"`
		Szs          interface{} `json:"szs"`
		Bz           interface{} `json:"bz"`
		Hsjczhrq     interface{} `json:"hsjczhrq"`
		Jkm          interface{} `json:"jkm"`
		Bz2          string      `json:"bz2"`
		Innjaddr     interface{} `json:"innjaddr"`
		Leavenj      interface{} `json:"leavenj"`
		Ryxm         interface{} `json:"ryxm"`
		Sjlyname     string      `json:"sjlyname"`
		Peoplexm     string      `json:"peoplexm"`
		Peoplesfz    string      `json:"peoplesfz"`
		Peoplephone  string      `json:"peoplephone"`
		Peoplecph    interface{} `json:"peoplecph"`
		Szsfsqx      interface{} `json:"szsfsqx"`
		Sfrj         interface{} `json:"sfrj"`
		Rjcsbm       interface{} `json:"rjcsbm"`
		Rjcsmc       interface{} `json:"rjcsmc"`
		Rjrq         interface{} `json:"rjrq"`
		Gj           interface{} `json:"gj"`
		Hz           interface{} `json:"hz"`
		Lxdh         interface{} `json:"lxdh"`
		Peoplelxdh   interface{} `json:"peoplelxdh"`
		Ypjgname     string      `json:"ypjgname"`
		Szsfcode     interface{} `json:"szsfcode"`
		Szscode      interface{} `json:"szscode"`
		Szsfsqxcode  interface{} `json:"szsfsqxcode"`
		Fqr          interface{} `json:"fqr"`
		Fkid         interface{} `json:"fkid"`
		Hsjcbg       interface{} `json:"hsjcbg"`
		Peoplegj     string      `json:"peoplegj"`
		Zlfbrxm      string      `json:"zlfbrxm"`
		Cbrxm        string      `json:"cbrxm"`
		Fxrylyxqbm   interface{} `json:"fxrylyxqbm"`
		Fxrylyxqbmbk interface{} `json:"fxrylyxqbmbk"`
		Innjaddrcode interface{} `json:"innjaddrcode"`
		Ispc         interface{} `json:"ispc"`
		Tqfm         interface{} `json:"tqfm"`
		Jsdw         interface{} `json:"jsdw"`
		Fqxq         interface{} `json:"fqxq"`
		Fqjd         interface{} `json:"fqjd"`
		Fqsq         interface{} `json:"fqsq"`
		Jsxq         interface{} `json:"jsxq"`
		Jsjd         interface{} `json:"jsjd"`
		Jssq         interface{} `json:"jssq"`
		Bhyy         interface{} `json:"bhyy"`
		Bhly         interface{} `json:"bhly"`
		Bhdw         interface{} `json:"bhdw"`
		Bhdwmc       interface{} `json:"bhdwmc"`
		Lkfxdqrq     interface{} `json:"lkfxdqrq"`
		Gksj         interface{} `json:"gksj"`
		Glfkid       interface{} `json:"glfkid"`
		Shsj         interface{} `json:"shsj"`
		Value        interface{} `json:"value"`
		Text         interface{} `json:"text"`
		Isactive     interface{} `json:"isactive"`
	} `json:"data"`
	Success bool `json:"success"`
}

func main() {
	fmt.Print(`

******************************************************************************************************
宁搏疫信息录入工具 2022-3-30 by TYB
3.30 更新： 使用兼容性更优的管控日期格式化方式
3.29 发布第一版

（可能存在问题，请谨慎使用）
本工具基于以下信息编写，其他格式信息未经过测试：
xxx	3月29日	xxx信息1	0329-59	195xxx51059	滁州	是	是	健康监测持 有效报告需要再检测（3月29日-4月11日）
xxx	3月29日	xxx信息1	0329-59	177xxx57275	滁州	是	是	健康监测持 有效报告需要再检测（3月29日-4月11日）
******************************************************************************************************
	`)
	args := os.Args
	if len(args) < 2 {
		fmt.Println("请传入excel文件！")
		return
	}
	f, err := excelize.OpenFile(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	ctrkMethodMap := map[string]string{
		"超出管控期":     "0405",
		"有效报告需要再检测": "0402",
	}

	r := regexp.MustCompile(`^[0-9]+$`) //检查每行首个单元格是否为数字，若不是则跳过该行
	for _, row := range rows {
		if b := r.MatchString(row[0]); !b {
			continue
		}
		fkInfo := userInfo{
			Phone:     row[4],
			AtNanjing: row[6],
		}

		//提取管控措施、管控日期
		var ctrl []string = strings.Split(strings.Replace(strings.Replace(row[8], "（", " ", -1), "）", "", -1), " ") //e.g.  [健康监测 超出管控期 3月12日-3月25日]
		fkInfo.CtrlMethod = ctrkMethodMap[ctrl[1]]
		r, _ := regexp.Compile(`\d+`)
		dates := r.FindAllString(ctrl[2], -1)
		if dates == nil {
			fmt.Println(fkInfo.Phone + " 录入失败，管控日期获取失败！")
			continue
		}

		for index := range dates {
			//填充长度为个位数的日期
			if len(dates[index]) < 2 {
				dates[index] = "0" + dates[index]
			}
		}

		fkInfo.CtrlStart = "2022" + dates[0] + "-" + dates[1]
		fkInfo.CtrlEnd = "2022" + dates[2] + "-" + dates[3]

		fkInfo.Glryid, fkInfo.Id, fkInfo.Tracked, fkInfo.Zlbh, fkInfo.Fxrylyxq, fkInfo.Fxryzx, fkInfo.Rksj, fkInfo.Bz2 = getInfoList(fkInfo.Phone)
		if fkInfo.Glryid == "" && fkInfo.Id == "" {
			fmt.Println(fkInfo.Phone + " 录入失败,信息条数不够")
			continue
		}

		if saveFkxx(fkInfo) {
			fmt.Println(fkInfo.Phone + " 录入成功")
		} else {
			fmt.Println(fkInfo.Phone + " 录入失败")
		}
	}
}

func saveFkxx(fkInfo userInfo) bool {
	url := "http://qyxf.yqfkpt.njga.gov.cn:8082/instruction/saveFkxx"
	bodyData := `{"czly":"PC","dialogRawDataForm":{"searchValue":null,"createBy":null,"createTime":null,"updateBy":null,"updateTime":null,"remark":null,"params":{},"id":"` + fkInfo.Id + `","glryid":"` + fkInfo.Glryid + `","zlbh":"` + fkInfo.Zlbh + `","fxrysjly":null,"fxrysjlb":null,"fxryly":null,"fxrylb":"` + fkInfo.Tracked + `","fxrylyxq":"` + fkInfo.Fxrylyxq + `","fxryxflx":"2","fxryzx":"` + fkInfo.Fxryzx + `","sjlb":"99","sjly":"` + fkInfo.Tracked + `","xm":"null","xb":null,"zjlx":null,"zjhm":"null","sjhm":"` + fkInfo.Phone + `","yys":null,"rylb":null,"rysx":null,"cph":null,"rylnjtxx":null,"ryccjcsj":null,"rymcjcsj":null,"ryjcfs":null,"rydz":"栖霞区","sjdwdz":null,"ryzzbh":null,"ryzzmc":"栖霞区","szxq":"320113","szxqmc":"栖霞区","szjz":null,"szjzmc":null,"szsq":null,"szsqmc":null,"szwg":null,"szwgmc":null,"rksj":"` + fkInfo.Rksj + `","fbsj":null,"qssj":null,"yjsj":null,"fksj":null,"szdw":"320113","szdwmc":"栖霞区","ryzt":"3","delstatus":null,"cjr":null,"cjrxm":"蒋碧媛","cjdw":"320113","cjdwmc":"栖霞区","cjsj":null,"xgr":null,"xgrxm":null,"xgdw":null,"xgdwmc":null,"xgsj":null,"unitcode":null,"chinaname":null,"yjmbdw":null,"yjmbdwmc":null,"fqrxm":null,"fqdwmc":null,"fqsj":null,"fqrlxfs":null,"fqyy":null,"ypjg":"12","zlly":null,"zlqsqx":null,"gkqkname":null,"bzdzcomp":null,"zlfbsj":null,"yszxqmc":null,"yszjzmc":null,"yszsqmc":null,"xpdw":null,"sjlbname":null,"qxname":null,"jdname":null,"sqname":null,"rysxname":null,"zlqh":null,"zlbt":null,"zlnr":null,"zlyq":null,"jsdwmc":null,"code":null,"num":null,"orgname":null,"zllx":null,"fkqx":null,"fqdw":null,"yjlx":null,"znzt":"` + fkInfo.AtNanjing + `","lnrq":"","gkqk":"` + fkInfo.CtrlMethod + `","ksgkrq":"` + fkInfo.CtrlStart + `","jsgkrq":"` + fkInfo.CtrlEnd + `","hsjccs":null,"lxfsarray":"","szsf":"","szs":"","bz":"","hsjczhrq":"","jkm":"绿码","bz2":"` + fkInfo.Bz2 + `","innjaddr":"栖霞","leavenj":"","ryxm":null,"sjlyname":null,"peoplexm":"null","peoplesfz":"null","peoplephone":"` + fkInfo.Phone + `","peoplecph":null,"szsfsqx":"","sfrj":"0","rjcsbm":null,"rjcsmc":null,"rjrq":"","gj":"CHN","hz":null,"lxdh":null,"peoplelxdh":null,"ypjgname":null,"szsfcode":"","szscode":"","szsfsqxcode":"","fqr":null,"fkid":null,"hsjcbg":null,"peoplegj":"CHN","zlfbrxm":null,"cbrxm":null,"fxrylyxqbm":"980200","fxrylyxqbmbk":null,"innjaddrcode":null,"ispc":"PC","tqfm":null,"jsdw":null,"fqxq":null,"fqjd":null,"fqsq":null,"jsxq":null,"jsjd":null,"jssq":null,"bhyy":null,"bhly":null,"bhdw":null,"bhdwmc":null,"lkfxdqrq":"","gksj":null,"glfkid":null,"shsj":null,"value":null,"text":null,"isactive":null,"lsqk":"01","hsjcjg":"阴性","jkmsxsj":"","sfbxq":"是"},"peopleType":"3","orgCode":"320113","userCode":"320107198411213424","phone":"13400066100","userName":"蒋碧媛","phoneMessage":"0","dxMessage":"南京市疫情防控指挥部温馨提示：根据大数据分析，您近期可能与新冠病毒感染者存在直接或间接接触，存在一定的感染风险。按照疫情防控要求，请在收到本信息后配合开展以下工作：1、立即主动联系您所在的社区（单位、宾馆），及时告知本短信情况和本人相关旅居史情况，落实核酸检测等疫情防控措施；2、如出现发热、干咳、鼻塞、流涕、乏力、咽痛、嗅（味）觉减退等症状，请及时就诊，并告知医生本短信情况和本人相关旅居史情况，就医途中请做好个人防护，不乘坐公共交通工具。如未履行疫情防控义务造成传染病传播扩散，将承担相应的法律责任。感谢您的理解和支持！"}
`
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(bodyData))
	if err != nil {
		fmt.Print(fkInfo.Phone + " 录入失败：")
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36 Edg/99.0.1150.52")
	req.Header.Set("x-auth-token", "04FBC29761578F5D35CE4B383FCDF94D92C8AFF2010D7DEBEB803736F168F031004F4875678CB28334EED45C53B89D6B9CE7D955A61C6A7D7BEAD6F336E70217F29B89FCE000BD19832021ABAEBBE91CE95591F6AD3DBB5019CDA01D904B98CBFB916ECD8D1E04DF075CE89F9D35B253BD06756C5B9304A13511DC36075D6D350566289B469F30230F3BABCD723F6F0E60DFD6E448476A58C76AC395FC6295B68C649ECA0D3AEA21942652CA8D3200CA03EBBC23A0770E79199B4D1DF8E13AA117DF18C8AFA745EA9E453C")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(fkInfo.Phone + " 录入失败：")
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	Success, _ := regexp.Match(`"success":true`, body)
	return Success
}

func getInfoList(phoneNumber string) (a string, b string, c string, d string, e string, f string, g string, h string) {
	url := "http://49.77.124.59:8083/YqzlPeople/ryxxPageList"
	bodyData :=
		`{"currentPage":1,"pageSize":10,"gkcs":"","rylb":"","sfcl":"","ypjg":"","sjlb":"","tssj":"","timesetting":"","znzt":"","tqfm":"","orgCode":"320113","zgfxqry":"","orderby":"","gjz":"3_` + phoneNumber + `","zxgz":"","bjzt":"","zlly":"","yCode":"320113","zlbh":"","fxryzx":"","czkbx":"","sfzc":"","gksj":""}
	`

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(bodyData))
	if err != nil {
		fmt.Print(phoneNumber + " 录入失败：")
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36 Edg/99.0.1150.52")
	req.Header.Set("x-auth-token", "04FBC29761578F5D35CE4B383FCDF94D92C8AFF2010D7DEBEB803736F168F031004F4875678CB28334EED45C53B89D6B9CE7D955A61C6A7D7BEAD6F336E70217F29B89FCE000BD19832021ABAEBBE91CE95591F6AD3DBB5019CDA01D904B98CBFB916ECD8D1E04DF075CE89F9D35B253BD06756C5B9304A13511DC36075D6D350566289B469F30230F3BABCD723F6F0E60DFD6E448476A58C76AC395FC6295B68C649ECA0D3AEA21942652CA8D3200CA03EBBC23A0770E79199B4D1DF8E13AA117DF18C8AFA745EA9E453C")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(phoneNumber + " 录入失败：")
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var idInfo getIdInfo
	json.Unmarshal(body, &idInfo)
	if len(idInfo.Data) > 0 {
		return idInfo.Data[0].Glryid, idInfo.Data[0].ID, idInfo.Data[0].Sjly, idInfo.Data[0].Zlbh, idInfo.Data[0].Fxrylyxq, idInfo.Data[0].Fxryzx, idInfo.Data[0].Rksj, idInfo.Data[0].Bz2
	} else {
		return "", "", "", "", "", "", "", ""
	}
}
