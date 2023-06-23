import {deserializeIntoWidgetUpdate} from './deserializeIntoWidgetUpdate';
import {WidgetUpdate} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createWidgetUpdateFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoWidgetUpdate;
}
