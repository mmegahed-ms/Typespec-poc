import {deserializeIntoWidget} from './deserializeIntoWidget';
import {Widget} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createWidgetFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoWidget;
}
