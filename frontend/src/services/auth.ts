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

export default { employeeLogin };
