import type {
  IEmployeeLoggedIn,
  ILoginEmployee,
} from "~/interfaces/auth.interface";
import { service as svc } from ".";
import type { Response } from "./service-instance";
import type { EmployeeRole } from "~/interfaces/employee.interface";

export type TVerifyUserResponse = Response<{
  id: number;
  exp: number;
  username: string;
  role: EmployeeRole;
}>;

const prefix = "Auth";

const versions = {
  v1: "v1",
  v2: "v2",
};

const employeeLogin = async (body: ILoginEmployee) => {
  const { data } = await svc.post<Response<IEmployeeLoggedIn>>(
    `${prefix}/${versions.v2}/employee/login`,
    body
  );
  return data;
};

const verifyToken = async () => {
  const { data } = await svc.post<TVerifyUserResponse>(
    `${prefix}/${versions.v1}/employee/verification`
  );
  return data;
};

const employeeLogout = async () => {
  const { data } = await svc.post<Response>(
    `${prefix}/${versions.v1}/employee/logout`
  );
  return data;
};

const revokeToken = async () => {
  const { data } = await svc.post<Response<IEmployeeLoggedIn>>(
    `${prefix}/${versions.v2}/employee/refresh`
  );
  return data;
};

export default {
  employeeLogin,
  employeeLogout,
  revokeToken,
  verifyToken,
};
