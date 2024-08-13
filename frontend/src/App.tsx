import './App.css'


const socket = new WebSocket("ws://localhost:8080");


function App() {
    socket.addEventListener("open", () => socket.send("Hello from client!")
)
    return (
        <>
            <label>name: </label>
            <input type={"text"}/>
            <button>Submit</button>
        </>
    )
}

export default App
