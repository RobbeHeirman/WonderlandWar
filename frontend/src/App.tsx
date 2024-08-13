import './App.css'
import { JoinLobbyMessage } from './proto/user_actions.ts'

// const socket = new WebSocket("ws://localhost:8080");

const msg: JoinLobbyMessage =  {
    name: "robbe"
}

const b = JoinLobbyMessage.toBinary(msg)
console.log(b);



function App() {
    // socket.addEventListener("open", () => socket.send("Hello from client!")
// )
    return (
        <>
            <label>name: </label>
            <input type={"text"}/>
            <button>Submit</button>
        </>
    )
}

export default App
