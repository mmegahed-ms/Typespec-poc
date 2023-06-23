package models
import (
    "errors"
)
// 
type Gadget_color int

const (
    RED_GADGET_COLOR Gadget_color = iota
    BLUE_GADGET_COLOR
)

func (i Gadget_color) String() string {
    return []string{"red", "blue"}[i]
}
func ParseGadget_color(v string) (any, error) {
    result := RED_GADGET_COLOR
    switch v {
        case "red":
            result = RED_GADGET_COLOR
        case "blue":
            result = BLUE_GADGET_COLOR
        default:
            return 0, errors.New("Unknown Gadget_color value: " + v)
    }
    return &result, nil
}
func SerializeGadget_color(values []Gadget_color) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
