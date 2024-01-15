import React, { useState } from 'react';
import API from '../api/api';
import HomeButton from './HomeButton';
import { useNavigate } from 'react-router-dom';

function Login() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
        const response = await API.post('/api/auth/login', {
            username,
            password
        });
        if (response.data.error) {
          console.log(response.data)
          setErrorMessage(response.data.message || 'Signup failed, Please try again.');
          return;
        }

        localStorage.setItem('go-chat-token', response.data.token);
        localStorage.setItem('go-chat-userId', getUserIDFromToken(response.data.token));

        navigate('/join');
    } catch (error) {
      console.error('Login failed', error);
    }
  };

  const getUserIDFromToken = (token) => {
    try {
      const payloadEncoded = token.split('.')[1];
      const payloadDecoded = atob(payloadEncoded);
      const payload = JSON.parse(payloadDecoded);
      return payload.userID;
    } catch (err) {
      console.error("Failed to parse JWT:", err);
      return null;
    }
  }

  return (
    <div className="bg-container">
      <div className="chat-window">
        <HomeButton />
        <form onSubmit={handleSubmit} className='form'>
          <div className='headings'>Hurry up, your friends have missed you! 🏃‍♀️</div>
          {errorMessage && <span className="error-message">*{errorMessage}</span>}
          <br />
          <br />
          <div className='form-input-box'>
            <input className='form-input' type="text" value={username} onChange={e => setUsername(e.target.value)} placeholder="Name - let's goo" />
            <input className='form-input' type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Password" />
          </div>
          <button type="submit" className='form-button'>Login</button>
        </form>
      </div>
    </div>
  );
}

export default Login;
