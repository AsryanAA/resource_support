import axios from 'axios'

const instance = axios.create({
    baseURL: 'http://localhost:8080/api/v1/',
    headers: {
        Authorization: `Bearer `// + getTokenFromLocalStorage() TODO: реализовать позже
    },
    timeout: 1000
})

export default instance