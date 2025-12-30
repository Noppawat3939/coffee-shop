import type {
  IEmployeeLoggedIn,
  ILoginEmployee,
} from "~/interfaces/auth.interface";
import { service as svc } from ".";
import type { Response } from "./service-instance";

const prefix = "Authen";

const employeeLogin = async (body: ILoginEmployee) => {
  const { data } = await svc.post<Response<IEmployeeLoggedIn>>(
    `${prefix}/employee/login`,
    body
  );
  return data;
};

const verifyToken = async () => {
  const { data } = await svc.post<Response>(`${prefix}/employee/verification`);
  return data;
};

export default { employeeLogin, verifyToken };
