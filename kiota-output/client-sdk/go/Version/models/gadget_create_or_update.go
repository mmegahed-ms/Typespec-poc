package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GadgetCreateOrUpdate 
type GadgetCreateOrUpdate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The color property
    color *GadgetCreateOrUpdate_color
    // The weight property
    weight *int32
}
// NewGadgetCreateOrUpdate instantiates a new GadgetCreateOrUpdate and sets the default values.
func NewGadgetCreateOrUpdate()(*GadgetCreateOrUpdate) {
    m := &GadgetCreateOrUpdate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGadgetCreateOrUpdateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGadgetCreateOrUpdateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGadgetCreateOrUpdate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GadgetCreateOrUpdate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. The color property
func (m *GadgetCreateOrUpdate) GetColor()(*GadgetCreateOrUpdate_color) {
    return m.color
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GadgetCreateOrUpdate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGadgetCreateOrUpdate_color)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val.(*GadgetCreateOrUpdate_color))
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
func (m *GadgetCreateOrUpdate) GetWeight()(*int32) {
    return m.weight
}
// Serialize serializes information the current object
func (m *GadgetCreateOrUpdate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GadgetCreateOrUpdate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. The color property
func (m *GadgetCreateOrUpdate) SetColor(value *GadgetCreateOrUpdate_color)() {
    m.color = value
}
// SetWeight sets the weight property value. The weight property
func (m *GadgetCreateOrUpdate) SetWeight(value *int32)() {
    m.weight = value
}
// GadgetCreateOrUpdateable 
type GadgetCreateOrUpdateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*GadgetCreateOrUpdate_color)
    GetWeight()(*int32)
    SetColor(value *GadgetCreateOrUpdate_color)()
    SetWeight(value *int32)()
}
