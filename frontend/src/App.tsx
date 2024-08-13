import './App.css'
import React, {useState} from "react";
import {JoinLobbyMessage} from "./proto/user_actions.ts";

function App() {
    // socket.addEventListener("open", () => socket.send("Hello from client!")
    const [socket, setSocket] = useState<WebSocket | null>(null)

    function onSubmit(e: React.SyntheticEvent<HTMLFormElement>) {
        e.preventDefault();
        const formProps = Object.fromEntries(new FormData(e.currentTarget));
        const connectMessage: JoinLobbyMessage = {
            name: formProps.name.toString()
        }
        const protoMsg = JoinLobbyMessage.toBinary(connectMessage);
        const nwSocket = new WebSocket("ws://localhost:8080");

        nwSocket.addEventListener("open", () => {
            setSocket(nwSocket);
            nwSocket.send(protoMsg)
        })
    }

    if (socket === null) {
        return (
            <>
                <form onSubmit={onSubmit}>
                    <label>name: </label>
                    <input type={"text"} id={"name"} name={"name"}/>
                    <button type={"submit"}>Submit</button>
                </form>
            </>
        )
    }
    return (
        <>
            <p>Connected!</p>
        </>
    )
}

export default App
