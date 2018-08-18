import { Contributions } from 'cxp/module/protocol/contribution'

/**
 * See the extensions.schema.json JSON Schema for canonical documentation on these types.
 *
 * This file is derived from the extensions.schema.json JSON Schema. It must be updated manually when the JSON
 * Schema or any of its referenced schemas change.
 *
 * TODO: Make this auto-generated. json2ts does not handle the "$ref" well, so it was simpler and faster to just
 * manually duplicate it for now.
 */

export interface CXPExtensionManifest {
    title?: string
    description?: string
    readme?: string
    platform: BundleTarget | DockerTarget | WebSocketTarget | TcpTarget | ExecTarget
    activationEvents: string[]
    args?: {
        [k: string]: any
    }
    contributes?: Contributions & { configuration?: { [key: string]: any } }
}

export interface BundleTarget {
    type: 'bundle'
    contentType?: string
    url: string
}

export interface DockerTarget {
    type: 'docker'
    image: string
}

export interface WebSocketTarget {
    type: 'websocket'
    url: string
}

export interface TcpTarget {
    type: 'tcp'
    address: string
}

export interface ExecTarget {
    type: 'exec'
    command: string
}