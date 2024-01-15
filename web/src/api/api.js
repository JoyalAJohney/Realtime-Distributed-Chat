import axios from 'axios'

const nginxHost = process.env.REACT_APP_NGINX_HOST;
const nginxPort = process.env.REACT_APP_NGINX_PORT;
const isLocalEnv = process.env.REACT_APP_NGINX_ENV === 'local';

let baseURL = '';

if (nginxHost && nginxPort) {
    if (isLocalEnv) {
        baseURL = `http://${nginxHost}:${nginxPort}`;
    } else {
        baseURL = `https://${nginxHost}`;
    }
} else {
    baseURL = 'http://localhost:3001';
}

const API = axios.create({
    baseURL,
    headers: {
        'Content-Type': 'application/json',
    },
});

export default API;