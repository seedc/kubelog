package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetShellder 执行shell命令获取k8s pod名返回结果是uint8
func GetShellder()(intcmd []uint8) {
	strCommand := "kubectl get pod|grep -v 'NAME'|awk '{print $1}'"
	cmd := exec.Command("/bin/sh", "-c", strCommand )
	intcmd, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

// Mapstrder 格式化uint8 转为map[int]string
func Mapstrder() (map[int]string)  {
	var str map[int]string
	str = make(map[int]string)
	var u []uint8
	g := GetShellder()
	a := 0
	for i:=0;i<len(g);i++ {
		if g[i] != 10 {
			u = append(u, g[i])
		}else {
			str[a] = string(u)
			a++
		}
		if g[i] == 10 {
			u = []uint8{}
		}
	}
	return str
}

// ForRangeMapder 接收key值来获取value值
func ForRangeMapder(key int) (v string)  {
	Mapstrder := Mapstrder()
	for k, v := range Mapstrder {
		if k == key {
			return v
		}
	}
	return
}

// Podnameder 单独获取pod项目名
func Podnameder(key int) (name string)  {
	str := ForRangeMapder(key)
	comma := strings.Index(str, "-")
	return str[0:comma]
}

// PodAllNameder 格式化pod名为导航用到的map字典
func PodAllNameder() (map[int]string) {
	var name map[int]string
	list := []string{}
	Mapstrder := Mapstrder()
	name = make(map[int]string)
	m := make(map[string]bool)
	for _, v := range Mapstrder {
		// 提取第一个"-"前面的字符串
		comma := strings.Index(v, "-")
		// 去重切片
		if _, ok := m[v[0:comma]]; !ok {
			list = append(list, v[0:comma])
			m[v[0:comma]] = true
		}
	}
	for k, v := range list {
		name[k] = v
	}
	return  name
}

// LogsTailder 读取pod日志
func LogsTailder(key int, num int) (intcmd []uint8) {
	Pod := ForRangeMapder(key)
	strCommand := fmt.Sprintf("kubectl logs --tail=%d ", num ) + Pod
	cmd := exec.Command("/bin/sh", "-c", strCommand)
	intcmd, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

// LogSeveDownload 指定pod保存日记并下载
func LogSeveDownload(v string)  {
	Pod := v
	strCommand := "kubectl logs " + Pod + " >./tmp/" + Pod + ".log"
	cmd := exec.Command("/bin/sh", "-c", strCommand)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Postname(str string) (map[int]string) {
	var name map[int]string
	name = make(map[int]string)
	pods := Mapstrder()
	a := 0
	for _, v := range pods {
		comma := strings.Index(v, "-")
		if v[0:comma] == str {
			name[a] = v
			a++
		}
	}
	return name
}