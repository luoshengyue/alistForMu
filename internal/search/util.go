package search

import (
	"strings"

	"github.com/alist-org/alist/v3/internal/conf"
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/internal/op"
	"github.com/alist-org/alist/v3/internal/setting"
	"github.com/alist-org/alist/v3/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func Progress() (*model.IndexProgress, error) {
	p := setting.GetStr(conf.IndexProgress)
	var progress model.IndexProgress
	err := utils.Json.UnmarshalFromString(p, &progress)
	return &progress, err
}

func WriteProgress(progress *model.IndexProgress) {
	p, err := utils.Json.MarshalToString(progress)
	if err != nil {
		log.Errorf("marshal progress error: %+v", err)
	}
	err = op.SaveSettingItem(&model.SettingItem{
		Key:   conf.IndexProgress,
		Value: p,
		Type:  conf.TypeText,
		Group: model.SINGLE,
		Flag:  model.PRIVATE,
	})
	if err != nil {
		log.Errorf("save progress error: %+v", err)
	}
}

func isIgnorePath(path string) bool {
	for _, ignorePath := range conf.SlicesMap[conf.IgnorePaths] {
		if strings.HasPrefix(path, ignorePath) {
			return true
		}
	}
	return false
}

func init() {
	op.RegisterSettingItemHook(conf.IgnorePaths, func(item *model.SettingItem) error { return nil })
	op.RegisterStorageHook(func(typ string, storage driver.Driver) {})
}
