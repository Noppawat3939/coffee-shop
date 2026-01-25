import axios from "axios";
import Cookies from "js-cookie";
import { ACCESS_TOKEN_COOKIE_KEY } from "~/helper/constant";

export type Response<Data = unknown> = { code: number; data: Data };

const authen = Cookies.get(ACCESS_TOKEN_COOKIE_KEY);
const token = `Bearer ${authen}`;

const serviceInstance = axios.create({
  baseURL: "http://localhost:8081/api/",
  timeout: 1000 * 60,
  headers: {
    "Content-Type": "application/json",
    ...(authen && { Authorization: token }),
  },
});

export default serviceInstance;
