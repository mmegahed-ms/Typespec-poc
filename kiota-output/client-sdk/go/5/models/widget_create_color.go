package models
import (
    "errors"
)
// 
type WidgetCreate_color int

const (
    RED_WIDGETCREATE_COLOR WidgetCreate_color = iota
    BLUE_WIDGETCREATE_COLOR
)

func (i WidgetCreate_color) String() string {
    return []string{"red", "blue"}[i]
}
func ParseWidgetCreate_color(v string) (any, error) {
    result := RED_WIDGETCREATE_COLOR
    switch v {
        case "red":
            result = RED_WIDGETCREATE_COLOR
        case "blue":
            result = BLUE_WIDGETCREATE_COLOR
        default:
            return 0, errors.New("Unknown WidgetCreate_color value: " + v)
    }
    return &result, nil
}
func SerializeWidgetCreate_color(values []WidgetCreate_color) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
