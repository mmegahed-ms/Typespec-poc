package models
import (
    "errors"
)
// 
type Widget_color int

const (
    RED_WIDGET_COLOR Widget_color = iota
    BLUE_WIDGET_COLOR
)

func (i Widget_color) String() string {
    return []string{"red", "blue"}[i]
}
func ParseWidget_color(v string) (any, error) {
    result := RED_WIDGET_COLOR
    switch v {
        case "red":
            result = RED_WIDGET_COLOR
        case "blue":
            result = BLUE_WIDGET_COLOR
        default:
            return 0, errors.New("Unknown Widget_color value: " + v)
    }
    return &result, nil
}
func SerializeWidget_color(values []Widget_color) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
