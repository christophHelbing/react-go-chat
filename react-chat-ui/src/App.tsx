import './App.css'
import {connect, sendMsg} from "./api";
import {Header} from "./components/header/header.tsx";
import ChatHistory from "./components/chat-history/chat-history.tsx";
import {useEffect, useState} from "react";

function App() {
    const [messages, addMessage] = useState<string[]>([])

    useEffect(() => {
        connect((msg: string) => {
            addMessage([...messages, msg])
        })
    })

    function send() {
        sendMsg("hello")
    }

    return (
        <>
            <div>
                <Header/>
                <ChatHistory messsageHistory={messages}/>
                <button onClick={send}>Send Hello</button>
            </div>
        </>
    )
}

export default App
