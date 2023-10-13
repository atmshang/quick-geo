package quick_geo

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

type GeoInfo struct {
	ID   string
	Name string
}

// geoParser 返回包含给定经纬度点的标记名称
func geoParser(geoData []byte, lon float64, lat float64) *GeoInfo {

	// 解析为FeatureCollection
	fc := geojson.NewFeatureCollection()
	if err := json.Unmarshal(geoData, fc); err != nil {
		log.Fatal(err)
	}

	// 创建点
	point := orb.Point{lon, lat}

	// 遍历Features
	for _, f := range fc.Features {
		if f.Geometry.Bound().Contains(point) {

			name := f.Properties["name"].(string)
			id := f.ID.(string)
			if id == "TWN" {
				name = "China/Taiwan"
			}

			// 如果包含点,返回名称
			return &GeoInfo{
				Name: name,
				ID:   id,
			}
		}
	}

	// 没找到返回空字符串
	return nil
}

func GetCountryOrRegionName(lon float64, lat float64) string {
	ret := geoParser(worldGeoJson, lon, lat)
	if ret != nil {
		return ret.Name
	} else {
		return ""
	}
}

func example() {
	name := GetCountryOrRegionName(114.03, 22.61)
	fmt.Println(name)

	name = GetCountryOrRegionName(122.96, 12.98)
	fmt.Println(name)
}
