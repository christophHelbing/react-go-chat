import "./message.scss"
import {useState} from "react";

export function Message({jsonMessage}) {
    const [message] = useState<{ type: number, body: string }>(JSON.parse(jsonMessage))

    return (
        <div key={message.type} className="Message">{message.body}</div>
    )
}