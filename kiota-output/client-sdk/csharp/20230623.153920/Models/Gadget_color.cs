using System.Runtime.Serialization;
using System;
namespace ApiSdk.Models {
    public enum Gadget_color {
        [EnumMember(Value = "red")]
        Red,
        [EnumMember(Value = "blue")]
        Blue,
    }
}
