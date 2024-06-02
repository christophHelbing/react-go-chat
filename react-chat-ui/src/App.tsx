import './App.css'
import {connect, sendMsg} from "./api";
import {Header} from "./components/header/header.tsx";
import ChatHistory from "./components/chat-history/chat-history.tsx";
import {useEffect, useState} from "react";
import {ChatInput} from "./components/chat-input/chat-input.tsx";

function App() {
    const [messages, addMessage] = useState<string[]>([])

    useEffect(() => {
        connect((msg: string) => {
            addMessage([...messages, msg])
        })
    })

    function send(msg: string) {
        sendMsg(msg)
    }

    return (
        <>
            <div className="container">
                <Header/>
                <ChatHistory messsageHistory={messages}/>
                <ChatInput sendMsg={send}/>
            </div>
        </>
    )
}

export default App
