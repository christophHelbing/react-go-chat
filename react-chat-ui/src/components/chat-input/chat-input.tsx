import "./chat-input.scss"

export function ChatInput({sendMsg}) {

    function send(event) {
        if(event.key == 'Enter') {
            sendMsg(event.target.value);
            event.target.value = "";
        }
    }

    return (
        <div className="ChatInput">
            <input onKeyDown={send}/>
        </div>
    )
}