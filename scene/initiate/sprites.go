package initiate

import (
	"BirdQuest/scene/models"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func prepareSprites(rawSprites rl.Texture2D, jsonObject *jsonSprites, startGid int) *models.Sprites {
	ss := &models.Sprites{
		TileWidth:    jsonObject.TileWidth,
		TileHeight:   jsonObject.TileHeight,
		WidthInTiles: int(rawSprites.Width) / jsonObject.TileWidth,
	}
	ss.Properties = []*models.Property{}
	for _, tile := range jsonObject.Tiles {
		prop := &models.Property{Id: tile.Id + startGid}
		for _, jProperty := range tile.Properties {
			switch jProperty.Name {
			case "AlwaysAfter":
				switch jProperty.Value.(bool) {
				case true:
					prop.AlwaysRenderLast = true
				default:
					prop.AlwaysRenderLast = false
				}
			case "AlwaysFirst":
				switch jProperty.Value.(bool) {
				case true:
					prop.AlwaysRenderFirst = true
				default:
					prop.AlwaysRenderFirst = false
				}
			case "Name":
				prop.Name = jProperty.Value.(string)
			}
		}
		ss.Properties = append(ss.Properties, prop)
	}

	return ss
}
