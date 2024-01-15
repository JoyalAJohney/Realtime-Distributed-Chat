import React, { createContext, useRef, useEffect } from 'react';

export const WebSocketContext = createContext(null);

export const WebSocketProvider = ({ children }) => {
    const ws = useRef(null);
  
    useEffect(() => {
      const env = process.env.REACT_APP_NGINX_ENV || 'local';
      const host = process.env.REACT_APP_NGINX_HOST || 'localhost';
      const port = process.env.REACT_APP_NGINX_PORT || '3001';

      const serverUrl = env === 'local' ? 
        `ws://${host}:${port}/ws/chat` : `wss://${host}/ws/chat`;
      
      ws.current = new WebSocket(`${serverUrl}`);
  
      return () => {
        console.log('Closing WebSocket connection...');
        if (ws.current) {
          ws.current.close();
        }
      };
    }, []);
  
    return (
      <WebSocketContext.Provider value={ws}>
        {children}
      </WebSocketContext.Provider>
    );
  }; 