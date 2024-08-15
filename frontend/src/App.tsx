import './App.css'
import React, {useState} from "react";
import {Envelope, JoinLobbyMessage} from "./proto/user_actions.ts";
import {Any} from "./proto/google/protobuf/any.ts"

type Player = {
    name: string
}

function App() {
    // socket.addEventListener("open", () => socket.send("Hello from client!")
    const [socket, setSocket] = useState<WebSocket | null>(null)
    const [otherPLayers, setotherPLayers] = useState<Player[]>([])

    function onSubmit(e: React.SyntheticEvent<HTMLFormElement>) {
        e.preventDefault();
        const formProps = Object.fromEntries(new FormData(e.currentTarget));
        const connectMessage: JoinLobbyMessage = {
            name: formProps.name.toString()
        }
        const protoMsg = JoinLobbyMessage.toBinary(connectMessage);
        const anyField: Any = {
            typeUrl: "proto_messages.JoinLobbyMessage",
            value: protoMsg
        }
        const envelope: Envelope = {
            data: anyField
        }
        const msgmrshl = Envelope.toBinary(envelope);
        const nwSocket = new WebSocket("ws://localhost:8080");

        nwSocket.addEventListener("open", () => {
            setSocket(nwSocket);
            nwSocket.send(msgmrshl);
        })

        nwSocket.addEventListener("message", event => {
            const envelope = Envelope.fromBinary(event.data)
            if (!envelope.data) {
                return
            }
            switch (envelope.data.typeUrl) {
                case "proto_messages.JoinLobbyMessage":
                    handleJoinsMessage(envelope.data);

            }

        })
    }

    function handleJoinsMessage(data: Any) {
        const msg = JoinLobbyMessage.fromBinary(data.value)
        const player = {
            name: msg.name
        }
        setotherPLayers(lst => [...lst, player])
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
    const vals = otherPLayers.map(value => <li>{value.name}</li>)
    return (
        <>
            <p>Connected!</p>
            <ul>
                {vals}
            </ul>
        </>
    )
}

export default App
