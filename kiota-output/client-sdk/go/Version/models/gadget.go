package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Gadget 
type Gadget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The color property
    color *Gadget_color
    // The id property
    id *string
    // The weight property
    weight *int32
}
// NewGadget instantiates a new Gadget and sets the default values.
func NewGadget()(*Gadget) {
    m := &Gadget{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGadgetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGadgetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGadget(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Gadget) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. The color property
func (m *Gadget) GetColor()(*Gadget_color) {
    return m.color
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Gadget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGadget_color)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val.(*Gadget_color))
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
func (m *Gadget) GetId()(*string) {
    return m.id
}
// GetWeight gets the weight property value. The weight property
func (m *Gadget) GetWeight()(*int32) {
    return m.weight
}
// Serialize serializes information the current object
func (m *Gadget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Gadget) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. The color property
func (m *Gadget) SetColor(value *Gadget_color)() {
    m.color = value
}
// SetId sets the id property value. The id property
func (m *Gadget) SetId(value *string)() {
    m.id = value
}
// SetWeight sets the weight property value. The weight property
func (m *Gadget) SetWeight(value *int32)() {
    m.weight = value
}
// Gadgetable 
type Gadgetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*Gadget_color)
    GetId()(*string)
    GetWeight()(*int32)
    SetColor(value *Gadget_color)()
    SetId(value *string)()
    SetWeight(value *int32)()
}
