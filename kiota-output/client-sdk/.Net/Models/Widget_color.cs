using System.Runtime.Serialization;
using System;
namespace ApiSdk.Models {
    public enum Widget_color {
        [EnumMember(Value = "red")]
        Red,
        [EnumMember(Value = "blue")]
        Blue,
    }
}
