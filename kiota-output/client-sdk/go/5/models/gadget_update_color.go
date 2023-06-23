package models
import (
    "errors"
)
// 
type GadgetUpdate_color int

const (
    RED_GADGETUPDATE_COLOR GadgetUpdate_color = iota
    BLUE_GADGETUPDATE_COLOR
)

func (i GadgetUpdate_color) String() string {
    return []string{"red", "blue"}[i]
}
func ParseGadgetUpdate_color(v string) (any, error) {
    result := RED_GADGETUPDATE_COLOR
    switch v {
        case "red":
            result = RED_GADGETUPDATE_COLOR
        case "blue":
            result = BLUE_GADGETUPDATE_COLOR
        default:
            return 0, errors.New("Unknown GadgetUpdate_color value: " + v)
    }
    return &result, nil
}
func SerializeGadgetUpdate_color(values []GadgetUpdate_color) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
