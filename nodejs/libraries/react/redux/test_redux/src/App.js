import React from "react";
import { useSelector, useDispatch } from "react-redux";
import { increment, decrement } from "./Action";

function App() {
    const counter = useSelector((state) => state.counter);
    const isLogged = useSelector((state) => state.logged);
    const dispatch = useDispatch();

    return (
        <div className="App">
            <h1>Counter {counter}</h1>
            <button onClick={() => dispatch(increment())}>+</button>
            <button onClick={() => dispatch(decrement())}>-</button>

            {isLogged ? <h3>Valuable Information</h3> : ""}
        </div>
    );
}

export default App;
