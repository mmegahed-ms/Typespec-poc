package models
import (
    "errors"
)
// 
type GadgetCreateOrUpdate_color int

const (
    RED_GADGETCREATEORUPDATE_COLOR GadgetCreateOrUpdate_color = iota
    BLUE_GADGETCREATEORUPDATE_COLOR
)

func (i GadgetCreateOrUpdate_color) String() string {
    return []string{"red", "blue"}[i]
}
func ParseGadgetCreateOrUpdate_color(v string) (any, error) {
    result := RED_GADGETCREATEORUPDATE_COLOR
    switch v {
        case "red":
            result = RED_GADGETCREATEORUPDATE_COLOR
        case "blue":
            result = BLUE_GADGETCREATEORUPDATE_COLOR
        default:
            return 0, errors.New("Unknown GadgetCreateOrUpdate_color value: " + v)
    }
    return &result, nil
}
func SerializeGadgetCreateOrUpdate_color(values []GadgetCreateOrUpdate_color) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
