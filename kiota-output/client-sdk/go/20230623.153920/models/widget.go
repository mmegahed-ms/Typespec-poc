package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Widget 
type Widget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The color property
    color *Widget_color
    // The id property
    id *string
    // The weight property
    weight *int32
}
// NewWidget instantiates a new Widget and sets the default values.
func NewWidget()(*Widget) {
    m := &Widget{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWidgetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWidgetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWidget(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Widget) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. The color property
func (m *Widget) GetColor()(*Widget_color) {
    return m.color
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Widget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWidget_color)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val.(*Widget_color))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
// GetId gets the id property value. The id property
func (m *Widget) GetId()(*string) {
    return m.id
}
// GetWeight gets the weight property value. The weight property
func (m *Widget) GetWeight()(*int32) {
    return m.weight
}
// Serialize serializes information the current object
func (m *Widget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetColor() != nil {
        cast := (*m.GetColor()).String()
        err := writer.WriteStringValue("color", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *Widget) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. The color property
func (m *Widget) SetColor(value *Widget_color)() {
    m.color = value
}
// SetId sets the id property value. The id property
func (m *Widget) SetId(value *string)() {
    m.id = value
}
// SetWeight sets the weight property value. The weight property
func (m *Widget) SetWeight(value *int32)() {
    m.weight = value
}
// Widgetable 
type Widgetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*Widget_color)
    GetId()(*string)
    GetWeight()(*int32)
    SetColor(value *Widget_color)()
    SetId(value *string)()
    SetWeight(value *int32)()
}
