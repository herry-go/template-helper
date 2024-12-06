package internal

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)




type DeployTemplate struct {
	ClusterMode string `yaml:"clusterMode" json:"cluster_mode" dc:"部署版本 single cluster cluster8 cluster11"`
	Spec        Spec   `yaml:"spec" json:"spec" dc:"规格说明" `
	ApiVersion  string `yaml:"apiVersion" json:"api_version" dc:"api版本"`
	Name        string `yaml:"name" json:"name" dc:"模板名称"`
	Kind        string `yaml:"kind" json:"kind" dc:"部署方式 DeploymentInit DeploymentPatch DeploymentUpgrade" `
	Version     string `yaml:"version" json:"version" dc:"版本"`
	Description string `yaml:"description" json:"description" dc:"备注"`
	PluginName  string `yaml:"pluginName" json:"plugin_name" dc:"插件名称 pluginName+version 唯一"`
	Readme      string `yaml:"readme" json:"readme" dc:"readme约束条件"`
}

type Spec struct {
	Host     []Variable `yaml:"hosts" json:"hosts" dc:"部署节点"`
	Resource []Variable `yaml:"resources" json:"resources" dc:"部署资源"`
	Env      Env        `yaml:"env" json:"env" dc:"部署配置"`

	Actions []Actions `yaml:"actions" json:"actions" dc:"部署步骤"`
}

type Env struct {
	Basic    []Variable `yaml:"basic" json:"basic" dc:"基础配置"`
	Advanced []Variable `yaml:"advanced" json:"advanced" dc:"高级配置"`
}

type Variable struct {
	Type        string      `yaml:"type" json:"type" dc:"前端组件类型：input, select, multi-select, switch, file-upload, host-select,resource-select"`
	Label       string      `yaml:"label" json:"label" dc:"label"`
	Description string      `yaml:"description" json:"description" dc:"备注"`
	Required    bool        `yaml:"required" json:"required" dc:"是否必传"`
	Default     interface{} `yaml:"default" json:"default" dc:"默认值"`
	Value       interface{} `yaml:"value" json:"value" dc:"实际值"`
	Name        string      `yaml:"name" json:"name" dc:"name"`
	Options     []Options   `yaml:"options" json:"options" dc:"选项"`
	DataType    string      `yaml:"dataType" json:"data_type" dc:"数据类型 目前支持： TEXT INTEGER BOOLEAN ARRAY"`
}

type Options struct {
	Label string      `yaml:"label" json:"label"`
	Value interface{} `yaml:"value" json:"value"`
}

type Actions struct {
	Name        string    `yaml:"name" json:"name" dc:"name"`
	Label       string    `yaml:"label" json:"label" dc:"label"`
	Description string    `yaml:"description" json:"description" dc:"备注"`
	Type        string    `yaml:"type" json:"type" dc:"类型 目前支持 script_shell ansible_shell"`
	Command     string    `yaml:"command" json:"command" dc:"命令"`
	Args        string    `yaml:"args" json:"args" dc:"命令参数"`
	Required    bool      `yaml:"required" json:"required" dc:"是否必须执行"`
	Rollback    []Actions `yaml:"rollback" json:"rollback" dc:"重置命令"`
	Actions     []Actions `yaml:"actions" json:"actions" dc:"子命令"`
}

// ParsePlugin 解析yaml文件
func ParsePlugin(filePath string) (tmpl *DeployTemplate, err error) {
	tmpl = new(DeployTemplate)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(content, tmpl); err != nil {
		return nil, errors.New("不支持的部署模板！")
	}
	return
}
