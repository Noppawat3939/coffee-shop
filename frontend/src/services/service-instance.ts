import axios from "axios";

export type Response<Data = unknown> = { code: number; data: Data };

const serviceInstance = axios.create({
  baseURL: "http://localhost:8081/api/",
  timeout: 1000 * 60,
  headers: {
    "Content-Type": "application/json",
  },
});

export default serviceInstance;
