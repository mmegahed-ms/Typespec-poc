package apisdk.models;

import com.microsoft.kiota.serialization.ValuedEnum;
import java.util.Objects;

public enum WidgetUpdateColor implements ValuedEnum {
    Red("red"),
    Blue("blue");
    public final String value;
    WidgetUpdateColor(final String value) {
        this.value = value;
    }
    @javax.annotation.Nonnull
    public String getValue() { return this.value; }
    @javax.annotation.Nullable
    public static WidgetUpdateColor forValue(@javax.annotation.Nonnull final String searchValue) {
        Objects.requireNonNull(searchValue);
        switch(searchValue) {
            case "red": return Red;
            case "blue": return Blue;
            default: return null;
        }
    }
}
