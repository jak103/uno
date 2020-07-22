import axios from "axios";

export default axios.create({
    baseURL: process.env.NODE_ENV === 'production' ? 'https://uno-fievsqoyiq-uw.a.run.app' : 'http://localhost:8080'
});