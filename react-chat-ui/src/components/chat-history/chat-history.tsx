import "./chat-history.scss"
export default function ChatHistory({messsageHistory}) {
    const messages = (messsageHistory as string[]).map((msg, index) => (
        <p key={index}>{msg}</p>
    ));

    return (
        <div className="ChatHistory">
            <h2>Chat History</h2>
            {messages}
        </div>
    )
};