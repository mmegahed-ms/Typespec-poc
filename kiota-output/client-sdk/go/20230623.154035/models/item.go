package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Item 
type Item struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name property
    name *string
    // The price property
    price *float64
    // The quantity property
    quantity *int32
}
// NewItem instantiates a new Item and sets the default values.
func NewItem()(*Item) {
    m := &Item{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Item) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Item) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["price"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrice(val)
        }
        return nil
    }
    res["quantity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuantity(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *Item) GetName()(*string) {
    return m.name
}
// GetPrice gets the price property value. The price property
func (m *Item) GetPrice()(*float64) {
    return m.price
}
// GetQuantity gets the quantity property value. The quantity property
func (m *Item) GetQuantity()(*int32) {
    return m.quantity
}
// Serialize serializes information the current object
func (m *Item) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("price", m.GetPrice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("quantity", m.GetQuantity())
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
func (m *Item) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *Item) SetName(value *string)() {
    m.name = value
}
// SetPrice sets the price property value. The price property
func (m *Item) SetPrice(value *float64)() {
    m.price = value
}
// SetQuantity sets the quantity property value. The quantity property
func (m *Item) SetQuantity(value *int32)() {
    m.quantity = value
}
// Itemable 
type Itemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetPrice()(*float64)
    GetQuantity()(*int32)
    SetName(value *string)()
    SetPrice(value *float64)()
    SetQuantity(value *int32)()
}
