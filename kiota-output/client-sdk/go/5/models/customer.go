package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Customer 
type Customer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The address property
    address *string
    // The city property
    city *string
    // The id property
    id *string
    // The items property
    items []Itemable
    // The name property
    name *string
    // The state property
    state *string
    // The zip property
    zip *string
}
// NewCustomer instantiates a new Customer and sets the default values.
func NewCustomer()(*Customer) {
    m := &Customer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomer(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Customer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. The address property
func (m *Customer) GetAddress()(*string) {
    return m.address
}
// GetCity gets the city property value. The city property
func (m *Customer) GetCity()(*string) {
    return m.city
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Customer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetId gets the id property value. The id property
func (m *Customer) GetId()(*string) {
    return m.id
}
// GetItems gets the items property value. The items property
func (m *Customer) GetItems()([]Itemable) {
    return m.items
}
// GetName gets the name property value. The name property
func (m *Customer) GetName()(*string) {
    return m.name
}
// GetState gets the state property value. The state property
func (m *Customer) GetState()(*string) {
    return m.state
}
// GetZip gets the zip property value. The zip property
func (m *Customer) GetZip()(*string) {
    return m.zip
}
// Serialize serializes information the current object
func (m *Customer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *Customer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. The address property
func (m *Customer) SetAddress(value *string)() {
    m.address = value
}
// SetCity sets the city property value. The city property
func (m *Customer) SetCity(value *string)() {
    m.city = value
}
// SetId sets the id property value. The id property
func (m *Customer) SetId(value *string)() {
    m.id = value
}
// SetItems sets the items property value. The items property
func (m *Customer) SetItems(value []Itemable)() {
    m.items = value
}
// SetName sets the name property value. The name property
func (m *Customer) SetName(value *string)() {
    m.name = value
}
// SetState sets the state property value. The state property
func (m *Customer) SetState(value *string)() {
    m.state = value
}
// SetZip sets the zip property value. The zip property
func (m *Customer) SetZip(value *string)() {
    m.zip = value
}
// Customerable 
type Customerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(*string)
    GetCity()(*string)
    GetId()(*string)
    GetItems()([]Itemable)
    GetName()(*string)
    GetState()(*string)
    GetZip()(*string)
    SetAddress(value *string)()
    SetCity(value *string)()
    SetId(value *string)()
    SetItems(value []Itemable)()
    SetName(value *string)()
    SetState(value *string)()
    SetZip(value *string)()
}
