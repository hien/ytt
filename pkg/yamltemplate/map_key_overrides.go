package yamltemplate

import (
	"fmt"

	"github.com/k14s/ytt/pkg/structmeta"
	"github.com/k14s/ytt/pkg/template"
	"github.com/k14s/ytt/pkg/yamlmeta"
)

const (
	AnnotationMapKeyOverride structmeta.AnnotationName = "yaml/map-key-override"
)

type MapItemOverride struct{}

func (d MapItemOverride) Apply(
	typedMap *yamlmeta.Map, newItem *yamlmeta.MapItem) error {

	itemIndex := map[interface{}]int{}

	for idx, item := range typedMap.Items {
		itemIndex[item.Key] = idx
	}

	if prevIdx, ok := itemIndex[newItem.Key]; ok {
		if template.NewAnnotations(newItem).Has(AnnotationMapKeyOverride) {
			typedMap.Items = append(typedMap.Items[:prevIdx], typedMap.Items[prevIdx+1:]...)
			return nil
		}

		return fmt.Errorf("disallowed to override key '%s' value", newItem.Key)
	}

	return nil
}
