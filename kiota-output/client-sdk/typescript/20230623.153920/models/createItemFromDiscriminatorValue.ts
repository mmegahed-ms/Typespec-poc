import {deserializeIntoItem} from './deserializeIntoItem';
import {Item} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createItemFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoItem;
}
