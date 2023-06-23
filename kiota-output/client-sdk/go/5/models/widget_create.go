package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WidgetCreate 
type WidgetCreate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The color property
    color *WidgetCreate_color
    // The weight property
    weight *int32
}
// NewWidgetCreate instantiates a new WidgetCreate and sets the default values.
func NewWidgetCreate()(*WidgetCreate) {
    m := &WidgetCreate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWidgetCreateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWidgetCreateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWidgetCreate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WidgetCreate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. The color property
func (m *WidgetCreate) GetColor()(*WidgetCreate_color) {
    return m.color
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WidgetCreate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWidgetCreate_color)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val.(*WidgetCreate_color))
        }
        return nil
    }
    res["weight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeight(val)
        }
        return nil
    }
    return res
}
// GetWeight gets the weight property value. The weight property
func (m *WidgetCreate) GetWeight()(*int32) {
    return m.weight
}
// Serialize serializes information the current object
func (m *WidgetCreate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetColor() != nil {
        cast := (*m.GetColor()).String()
        err := writer.WriteStringValue("color", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("weight", m.GetWeight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WidgetCreate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. The color property
func (m *WidgetCreate) SetColor(value *WidgetCreate_color)() {
    m.color = value
}
// SetWeight sets the weight property value. The weight property
func (m *WidgetCreate) SetWeight(value *int32)() {
    m.weight = value
}
// WidgetCreateable 
type WidgetCreateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*WidgetCreate_color)
    GetWeight()(*int32)
    SetColor(value *WidgetCreate_color)()
    SetWeight(value *int32)()
}
