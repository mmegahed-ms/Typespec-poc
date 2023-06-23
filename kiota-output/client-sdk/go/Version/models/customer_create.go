package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomerCreate 
type CustomerCreate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The address property
    address *string
    // The city property
    city *string
    // The items property
    items []Itemable
    // The name property
    name *string
    // The state property
    state *string
    // The zip property
    zip *string
}
// NewCustomerCreate instantiates a new CustomerCreate and sets the default values.
func NewCustomerCreate()(*CustomerCreate) {
    m := &CustomerCreate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomerCreateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomerCreateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomerCreate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomerCreate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. The address property
func (m *CustomerCreate) GetAddress()(*string) {
    return m.address
}
// GetCity gets the city property value. The city property
func (m *CustomerCreate) GetCity()(*string) {
    return m.city
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomerCreate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["city"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Itemable, len(val))
            for i, v := range val {
                res[i] = v.(Itemable)
            }
            m.SetItems(res)
        }
        return nil
    }
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["zip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZip(val)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The items property
func (m *CustomerCreate) GetItems()([]Itemable) {
    return m.items
}
// GetName gets the name property value. The name property
func (m *CustomerCreate) GetName()(*string) {
    return m.name
}
// GetState gets the state property value. The state property
func (m *CustomerCreate) GetState()(*string) {
    return m.state
}
// GetZip gets the zip property value. The zip property
func (m *CustomerCreate) GetZip()(*string) {
    return m.zip
}
// Serialize serializes information the current object
func (m *CustomerCreate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("zip", m.GetZip())
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
func (m *CustomerCreate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. The address property
func (m *CustomerCreate) SetAddress(value *string)() {
    m.address = value
}
// SetCity sets the city property value. The city property
func (m *CustomerCreate) SetCity(value *string)() {
    m.city = value
}
// SetItems sets the items property value. The items property
func (m *CustomerCreate) SetItems(value []Itemable)() {
    m.items = value
}
// SetName sets the name property value. The name property
func (m *CustomerCreate) SetName(value *string)() {
    m.name = value
}
// SetState sets the state property value. The state property
func (m *CustomerCreate) SetState(value *string)() {
    m.state = value
}
// SetZip sets the zip property value. The zip property
func (m *CustomerCreate) SetZip(value *string)() {
    m.zip = value
}
// CustomerCreateable 
type CustomerCreateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(*string)
    GetCity()(*string)
    GetItems()([]Itemable)
    GetName()(*string)
    GetState()(*string)
    GetZip()(*string)
    SetAddress(value *string)()
    SetCity(value *string)()
    SetItems(value []Itemable)()
    SetName(value *string)()
    SetState(value *string)()
    SetZip(value *string)()
}
