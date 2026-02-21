import type {
  IEmployeeLoggedIn,
  ILoginEmployee,
} from "~/interfaces/auth.interface";
import { service as svc } from ".";
import type { Response } from "./service-instance";

export type TVerifyUserResponse = Response<{
  id: number;
  exp: number;
  username: string;
}>;

const prefix = "Authen";

const employeeLogin = async (body: ILoginEmployee) => {
  const { data } = await svc.post<Response<IEmployeeLoggedIn>>(
    `${prefix}/employee/login`,
    body
  );
  return data;
};

const verifyToken = async () => {
  const { data } = await svc.post<TVerifyUserResponse>(
    `${prefix}/employee/verification`
  );
  return data;
};

const employeeLogout = async () => {
  const { data } = await svc.post<Response>(`${prefix}/employee/logout`);
  return data;
};

export default { employeeLogin, verifyToken, employeeLogout };
