import axios from 'axios';

const axiosInst = axios.create({
    baseURL: 'http://localhost:8080/',
});

export default axiosInst;
