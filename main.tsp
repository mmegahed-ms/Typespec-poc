import "@typespec/http";
import "./library.tsp";

using TypeSpec.Http;
@service({
    title: "Widget Service",
    version: "1.1.0",
})
@doc("Customer Demo Service API.")
namespace DemoService;

model Widget {
    @visibility("read", "update")
    @path
    id: string;
    weight: int32;
    color: "red" | "blue";
}

// add a new model called gadget that is a copy of widget
model Gadget {
    @visibility("read", "update")
    @path
    id: string;
    weight: int32;
    color: "red" | "blue";
}

//add a new model called customer that is a copy of widget but has differnt properties base on the below sample

@doc("Customer object model.")
model Customer {
    @visibility("read", "update")
    @path
    id: string;
    name: string;
    address: string;
    city: string;
    state: string;
    zip: string;
    items: Item[];
}

@doc("item object model.")
model Item {
    name: string;
    quantity: int32;
    price: float64;
}



@route("/widgets")
@tag("Widgets")
interface Widgets {
    @get list(): Widget[] | Error;
    @get read(@path id: string): Widget | Error;
    @post create(...Widget): Widget | Error;
    @patch update(...Widget): Widget | Error;
    @delete delete(@path id: string): void | Error;
    @route("{id}/analyze") @post analyze(@path id: string): string | Error;
}
// add a new  route called Customers to the service that is a copy of widgets but uses customer model
@doc("Customer Demo Interface.")
@route("/customers")
@tag("Customers")
interface Customers {
    @get list(): Customer[] | Error;
    @get read(@path id: string): Customer | Error;
    @post create(...Customer): Customer | Error;
    @patch update(...Customer): Customer | Error;
    @delete delete(@path id: string): void | Error;
    @route("{id}/analyze") @post analyze(@path id: string): string | Error;
}

// add a new operation
// add a new  route called gadgets to the service that is a copy of widgets but uses gadget model
@route("/gadgets")
@tag("Gadgets")
interface Gadgets extends ResourceInterface<Gadget> {
    @route("{id}/analyze") @post analyze(@path id: string): string | Error;
}

