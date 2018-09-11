import { from, Observable, Unsubscribable } from 'rxjs'
import { MarkupKind } from 'sourcegraph'
import {
    ClientCapabilities,
    HoverRequest,
    TextDocumentPositionParams,
    TextDocumentRegistrationOptions,
} from '../../protocol'
import { Hover } from '../../protocol/plainTypes'
import { Client } from '../client'
import { ProvideTextDocumentHoverSignature } from '../providers/hover'
import { FeatureProviderRegistry } from '../providers/registry'
import { ensure, Feature } from './common'

/**
 * Support for hover messages (textDocument/hover requests to the server).
 */
export class TextDocumentHoverFeature extends Feature<TextDocumentRegistrationOptions> {
    constructor(
        client: Client,
        private registry: FeatureProviderRegistry<TextDocumentRegistrationOptions, ProvideTextDocumentHoverSignature>
    ) {
        super(client)
    }

    public readonly messages = HoverRequest.type

    public fillClientCapabilities(capabilities: ClientCapabilities): void {
        const hoverCapability = ensure(ensure(capabilities, 'textDocument')!, 'hover')!
        hoverCapability.dynamicRegistration = true
        hoverCapability.contentFormat = [MarkupKind.Markdown, MarkupKind.PlainText]
    }

    protected registerProvider(options: TextDocumentRegistrationOptions): Unsubscribable {
        return this.registry.registerProvider(
            options,
            (params: TextDocumentPositionParams): Observable<Hover | null> =>
                from(this.client.sendRequest(HoverRequest.type, params))
        )
    }
}