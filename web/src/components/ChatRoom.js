import React, { useState, useEffect, useContext } from 'react';
import { WebSocketContext } from '../WebsocketContext';
import { useLocation, useNavigate } from 'react-router-dom';
import data from '@emoji-mart/data';
import Picker from '@emoji-mart/react';


function ChatRoom() {
  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState('');
  const [showPicker, setShowPicker] = useState(false);

  const location = useLocation();
  const navigate = useNavigate();
  const roomName = location.state?.roomName;
  const ws = useContext(WebSocketContext);

  const userID = localStorage.getItem('go-chat-userId');

  useEffect(() => {
    document.addEventListener('mousedown', handleOutsideClick);
    const currentWebSocket = ws.current;
    const roomName = location.state?.roomName;  
    if (!roomName) {
        navigate('/join');
        return;
    }

    const handleMessage = (message) => {
      const data = JSON.parse(message.data);
      setMessages(prevMessages => [...prevMessages, data]);
    };

    const handleClose = () => {
      const errorMessage = `We closed your connection, you've been idle for sometime.`
      navigate('/join', { state: { errorMessage } });
    };

    if (currentWebSocket) {
      currentWebSocket.onmessage = handleMessage;
      currentWebSocket.onclose = handleClose;
    }

    return () => {
      document.removeEventListener('mousedown', handleOutsideClick);
      if (currentWebSocket) {
        currentWebSocket.onmessage = null;
        currentWebSocket.onclose = null;
      }
    };
  }, [location.state, ws, navigate]);

  const handleEmojiPickerToggle = () => {
    setShowPicker(!showPicker);
  };

  const handleOutsideClick = (e) => {
    if (!e.target.closest('.emoji-picker') && !e.target.closest('.toggle-emoji-picker')) {
      setShowPicker(false);
    }
  };

  const addEmoji = (emoji) => {
    setInput(input + emoji.native);
  };

  const sendMessage = () => {
    if (input.trim() !== '') {
      ws.current.send(JSON.stringify({ type: 'chat_message', 'room': roomName, content: input }));
      setInput('');
    }
  };

  const leaveRoom = () => {
    ws.current.send(JSON.stringify({ type: 'leave_room', room: roomName }));
    navigate('/join');
  };
  

  return (
    <div className="bg-container">
      <div className="chat-room">
        <div className="chat-room-header">
          <h3>{roomName}</h3>
          <button className="leave-room-button" onClick={leaveRoom}>Leave Room</button>
        </div>
        <div className="message-box">
          {messages.map((msg, index) => {
            const isJoiningMessage = msg.type === 'join_room';
            const isOwnMessage = msg.sender === userID;

          if (isJoiningMessage) {
            const joiningMessage = isOwnMessage ? 
            `You joined ${roomName} 🎉` : 
            `${msg.senderName} joined from server ${msg.server} 🎉`;
            return (
              <div key={msg.id} className="join-message">
                {joiningMessage}
              </div>
            );
          }

          const messageClass = isOwnMessage ? "message own-message" : "message";
          return (
            <div key={msg.id} className={messageClass}>
              <div className="message-header">
                {isOwnMessage ? <span></span> : <span className="sender-name">{msg.senderName}</span>}
              </div>
              <div className="message-content">{msg.content}</div>
              <div className="message-footer">
              {new Date().toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit', hour12: true })}
              </div>
            </div>
          )})}
      </div>
      <div className="chat-room-footer">
        <button className="toggle-emoji-picker" onClick={handleEmojiPickerToggle}>😃</button>
        <input 
          type="text" 
          value={input} 
          onChange={(e) => setInput(e.target.value)} 
          className="message-input"
          placeholder="What's on your mind?"
        />
        <button className='send-message-button' onClick={sendMessage}>Send</button>
        {showPicker && (
          <div className="emoji-picker">
            <Picker data={data} onEmojiSelect={(emoji) => { 
              addEmoji(emoji);
            }} />
          </div>
        )}
      </div>
      </div>
    </div>
  );
}

export default ChatRoom;
