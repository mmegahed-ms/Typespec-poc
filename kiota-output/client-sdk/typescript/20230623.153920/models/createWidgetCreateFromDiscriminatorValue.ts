import {deserializeIntoWidgetCreate} from './deserializeIntoWidgetCreate';
import {WidgetCreate} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createWidgetCreateFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoWidgetCreate;
}
