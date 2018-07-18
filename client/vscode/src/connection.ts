/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Sourcegraph. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
'use strict'

import { DataCallback, Message, MessageReader, MessageWriter, RequestMessage } from 'vscode-jsonrpc'
import { AbstractMessageReader } from 'vscode-jsonrpc/lib/messageReader'
import { AbstractMessageWriter } from 'vscode-jsonrpc/lib/messageWriter'
import { MessageTransports } from 'vscode-languageclient'
import * as WebSocket from 'ws'

/**
 * connectWebSocket can be passed to the LSP client to connect to an LSP server via a
 * WebSocket.
 *
 * For robustness against temporary connection loss or server blips, it keeps attempting
 * to connect for connectionRetryTimeout msec (even if that means retrying to connect
 * multiple times). If connectionRetryTimeout is 0, then it will only attempt connecting
 * once.
 */
export function webSocketStreamOpener({
    url,
    headers,
    requestTracer,
    connectionRetryTimeout = 20000,
}: {
    url: string
    headers?: any
    requestTracer?: (trace: MessageTrace) => void
    connectionRetryTimeout?: number
}): Promise<MessageTransports> {
    const deadline = Date.now() + connectionRetryTimeout
    return new Promise((resolve, reject) => {
        const allHeaders: { [name: string]: string } = {
            'Content-Type': 'application/json; charset=utf-8',
            ...headers,
        }

        let socket = new WebSocket(url, [], { headers: allHeaders })
        socket.binaryType = 'arraybuffer'
        let connected = false
        socket.on('open', () => {
            connected = true
            const reader = new WebSocketMessageReader(socket)
            const writer = new WebSocketMessageWriter(socket)
            if (requestTracer) {
                traceJSONRPCRequests(requestTracer, reader, writer)
            }
            resolve({ reader, writer })
        })
        socket.on('close', code => {
            if (connected && code !== 1000 /* Close code: Normal */) {
                console.error('WebSocket closed:', code)
            }
            if (!connected) {
                // Retry to connect?
                if (deadline > Date.now()) {
                    setTimeout(
                        () =>
                            webSocketStreamOpener({
                                url,
                                requestTracer,
                                connectionRetryTimeout: deadline - Date.now(),
                            }).then(resolve, reject),
                        2500
                    )
                    return
                }

                reject(code)
            }
        })
        socket.on('error', (e: any) => {
            console.error('Socket error: ', e)
        })
    })
}

/**
 * WebSocketMessageReader wraps a WebSocket to conform to the MessageReader interface.
 */
class WebSocketMessageReader extends AbstractMessageReader implements MessageReader {
    private callbacks: DataCallback[] = []

    constructor(socket: WebSocket) {
        super()

        socket.on('message', ev => {
            if (this.callbacks.length === 0) {
                this.fireError(new Error('message arrived on WebSocket but there is no listener'))
                return
            }
            try {
                const data = JSON.parse(ev.toString())
                for (const callback of this.callbacks) {
                    callback(data)
                }
            } catch (error) {
                this.fireError(error)
            }
        })
    }

    public listen(callback: DataCallback): void {
        this.callbacks.push(callback)
    }
}

/**
 * WebSocketMessageWriter wraps a WebSocket to conform to the MessageWriter interface.
 */
class WebSocketMessageWriter extends AbstractMessageWriter implements MessageWriter {
    private socket: WebSocket
    private socketClosed: boolean
    private errorCount: number
    private callbacks: DataCallback[] = []

    constructor(socket: WebSocket) {
        super()
        this.socket = socket
        this.socketClosed = false
        this.errorCount = 0

        socket.on('close', () => {
            this.socketClosed = true
            this.fireClose()
        })
        socket.on('error', ev => {
            this.fireError(ev)
        })
    }

    public write(msg: Message): void {
        if (this.socketClosed) {
            this.errorCount++
            this.fireError(new Error('Write on closed WebSocket'), msg, this.errorCount)
            return
        }
        this.errorCount = 0
        for (const callback of this.callbacks) {
            callback(msg)
        }
        this.socket.send(JSON.stringify(msg))
    }

    public listen(callback: DataCallback): void {
        this.callbacks.push(callback)
    }
}

export interface MessageTrace {
    startTime: number
    endTime: number
    request: RequestMessage
    response: ResponseMessage
}

// vscode-jsonrpc2 doesn't export this interface, so we define it here.
// We also add meta which is a Sourcegraph specific extension.
export interface ResponseMessage extends Message {
    id: number | string
    result?: any
    error?: any
    meta?: { [key: string]: string }
}

interface Listener {
    listen(callback: DataCallback): void
}

function traceJSONRPCRequests(tracer: (trace: MessageTrace) => void, reader: Listener, writer: Listener): void {
    const inflight = new Map<string | number, [RequestMessage, number]>()
    writer.listen((data: Message) => {
        const msg = data as RequestMessage
        if (msg.id !== undefined) {
            inflight.set(msg.id, [msg, Date.now()])
        }
    })
    reader.listen((data: Message) => {
        const response = data as ResponseMessage
        if (response.id !== undefined) {
            const msg = inflight.get(response.id)
            if (msg !== undefined) {
                inflight.delete(response.id)
                const [request, startTime] = msg
                tracer({
                    startTime,
                    endTime: Date.now(),
                    request,
                    response,
                })
            }
        }
    })
}