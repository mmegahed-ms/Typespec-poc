package models
import (
    "errors"
)
// 
type WidgetUpdate_color int

const (
    RED_WIDGETUPDATE_COLOR WidgetUpdate_color = iota
    BLUE_WIDGETUPDATE_COLOR
)

func (i WidgetUpdate_color) String() string {
    return []string{"red", "blue"}[i]
}
func ParseWidgetUpdate_color(v string) (any, error) {
    result := RED_WIDGETUPDATE_COLOR
    switch v {
        case "red":
            result = RED_WIDGETUPDATE_COLOR
        case "blue":
            result = BLUE_WIDGETUPDATE_COLOR
        default:
            return 0, errors.New("Unknown WidgetUpdate_color value: " + v)
    }
    return &result, nil
}
func SerializeWidgetUpdate_color(values []WidgetUpdate_color) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
