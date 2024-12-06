package common

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
	"template-helper/internal"
)

var instance *Storage
var once sync.Once

type Storage struct {
	Path string
	Map  map[string]*internal.DeployTemplate
	lock sync.Mutex
}

func GetStorage() *Storage {
	once.Do(func() {
		instance = &Storage{
			Map:  make(map[string]*internal.DeployTemplate),
			lock: sync.Mutex{},
		}
	})
	return instance
}

func (s *Storage) SetPath(path string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Path = path
}

func (s *Storage) SetTmpl(tmpl *internal.DeployTemplate) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Map[s.Path] = tmpl
}

func (s *Storage) GetTmpl() *internal.DeployTemplate {
	if tmpl, ok := s.Map[s.Path]; ok {
		return tmpl
	}
	return nil
}

func (s *Storage) SaveTmpl() error {
	if s.Path == "" {
		return fmt.Errorf("文件不存在!")
	}
	tmpl := GetStorage().GetTmpl()
	bytes, err := yaml.Marshal(tmpl)
	if err != nil {
		return err
	}
	fmt.Print(string(bytes))
	// 将 YAML 数据写入文件
	err = os.WriteFile(s.Path, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) SetTmplHost(host []internal.Variable) {
	s.Map[s.Path].Spec.Host = host
}

func (s *Storage) SetTmplResource(resource []internal.Variable) {
	s.Map[s.Path].Spec.Resource = resource
}

func (s *Storage) SetTmplBasicEnv(env []internal.Variable) {
	s.Map[s.Path].Spec.Env.Basic = env
}


func (s *Storage) SetTmplAdvancedEnv(env []internal.Variable) {
	s.Map[s.Path].Spec.Env.Advanced = env
}


func (s *Storage) SetTmplAction(actios []internal.Actions) {
	s.Map[s.Path].Spec.Actions = actios
}
