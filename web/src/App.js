import './App.css';
import React from 'react';
import Signup from './components/Signup';
import Login from './components/Login';
import ChatRoom from './components/ChatRoom';
import JoinRoom from './components/JoinRoom';
import LandingPage from './components/LandingPage';
import {  BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { WebSocketProvider } from './WebsocketContext';

function App() {
  return (
    <WebSocketProvider>
    <Router>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login />} />
        <Route path="/chat" element={<ChatRoom />} />
        <Route path="/join" element={<JoinRoom />} />
      </Routes>
    </Router>
    </WebSocketProvider>
  );
}

export default App;
