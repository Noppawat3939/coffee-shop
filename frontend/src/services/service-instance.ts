import axios from "axios";

const serviceInstance = axios.create({
  baseURL: "http://localhost:8081/api/",
  timeout: 1000 * 60,
  headers: {
    "Content-Type": "application/json",
  },
});

export default serviceInstance;
