package file

import (
	"path/filepath"

	"picasso/pkg/klog/baselogger"
	"picasso/pkg/utils/kfile"
	"picasso/pkg/utils/kgo"
	"picasso/pkg/utils/kstring"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Source file provider.
type Source struct {
	vip     *viper.Viper
	changed chan map[string]interface{}
	logger  baselogger.Logger
}

// NewSource returns new Source.
func NewSource(opts ...*Option) (*Source, error) {
	sourceOpt := defaultOption().MergeOption(opts...)
	cs := &Source{logger: sourceOpt.Logger}
	// 查文件后缀
	dir, fileName := filepath.Split(sourceOpt.Path)
	name, suffix := kfile.GetSuffixAndFilename(fileName)
	Viper := viper.New()
	Viper.SetConfigName(name) // 设置配置文件名 (不带后缀)
	Viper.AddConfigPath(dir)  // 第一个搜索路径
	Viper.SetConfigType(suffix)
	cs.vip = Viper
	if sourceOpt.EnableWatch {
		cs.changed = make(chan map[string]interface{}, 1)
		kgo.Go(cs.watch)
	}
	return cs, nil
}

// ReadConfig ...
func (cs *Source) ReadConfig() (content map[string]interface{}, err error) {
	err = cs.vip.ReadInConfig() // 读取配置数据
	if err != nil {
		return nil, err
	}
	content = cs.vip.AllSettings()
	return
}

// IsConfigChanged ...
func (cs *Source) ConfigChanged() <-chan map[string]interface{} {
	return cs.changed
}

// Watch file and automate update.
func (cs *Source) watch() {
	cs.vip.WatchConfig() // 监视配置文件，重新读取配置数据
	cs.vip.OnConfigChange(func(e fsnotify.Event) {
		cs.logger.Log("File配置更新", kstring.KVString("name", e.Name), kstring.KVString("op", e.Op.String()))
		cs.changed <- cs.vip.AllSettings()
	})
}
