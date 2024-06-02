import "./chat-history.scss"
import {Message} from "../message/message.tsx";

export default function ChatHistory({messsageHistory}) {
    const messages = (messsageHistory as string[]).map(msg => (
        <Message jsonMessage={msg}/>
    ));

    return (
        <div className="ChatHistory">
            <h2>Chat History</h2>
            {messages}
        </div>
    )
};