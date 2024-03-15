import axios from 'axios'
import { useJWT } from '../store'

const instance = axios.create({
    baseURL: 'http://localhost:8080/',
    headers: {
        Authorization: `Bearer ` + 'token' // TODO: реализовать позже
    },
    timeout: 1000
})

export default instance