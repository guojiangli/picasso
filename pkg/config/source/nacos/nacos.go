package nacos

import (
	"strings"

	"picasso/pkg/utils/kstring"

	"github.com/spf13/viper"

	"picasso/pkg/klog/baselogger"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
)

// Source etcd provider.
type Source struct {
	Client      config_client.IConfigClient
	vip         *viper.Viper
	DataID      string
	Group       string
	Suffix      string
	EnableWatch bool
	Logger      baselogger.Logger
	changed     chan map[string]interface{}
}

// NewSource returns new Source.
func NewSource(client config_client.IConfigClient, opts ...*Option) (*Source, error) {
	sourceOpt := defaultOption().MergeOption(opts...)
	Viper := viper.New()
	Viper.SetConfigType(sourceOpt.Suffix)
	cs := &Source{
		Client:      client,
		DataID:      sourceOpt.DataID,
		Group:       sourceOpt.Group,
		Suffix:      sourceOpt.Suffix,
		EnableWatch: sourceOpt.EnableWatch,
		Logger:      sourceOpt.Logger,
		vip:         Viper,
	}
	if sourceOpt.EnableWatch {
		cs.changed = make(chan map[string]interface{}, 1)
		cs.watch()
	}
	return cs, nil
}

// ReadConfig ...
func (cs *Source) ReadConfig() (content map[string]interface{}, err error) {
	data, err := cs.Client.GetConfig(vo.ConfigParam{
		DataId: cs.DataID,
		Group:  cs.Group,
	})
	if err != nil {
		return nil, err
	}
	reader := strings.NewReader(data)
	cs.vip.ReadConfig(reader)
	return cs.vip.AllSettings(), nil
}

// IsConfigChanged ...
func (cs *Source) ConfigChanged() <-chan map[string]interface{} {
	return cs.changed
}

// Watch etcd and automate update.
func (cs *Source) watch() {
	cs.Client.ListenConfig(vo.ConfigParam{
		DataId: cs.DataID,
		Group:  cs.Group,
		OnChange: func(namespace string, group string, dataId string, data string) {
			cs.Logger.Log("Nacos配置更新", kstring.KVString("namespace", namespace), kstring.KVString("group", group), kstring.KVString("dataId", dataId))
			reader := strings.NewReader(data)
			cs.vip.ReadConfig(reader)
			cs.changed <- cs.vip.AllSettings()
		},
	})
}
