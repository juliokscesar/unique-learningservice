import React from 'react';
import './App.css';
import { Toast } from "./components/Toast";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <p>Hello World!</p>
      </header>

      <Toast message="Hello World" author="Julio"/>
    </div>
  );
}

export default App;
