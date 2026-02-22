import type { IEmployee } from "~/interfaces/employee.interface";
import { service } from ".";
import type { Response } from "./service-instance";

const prefix = "Employees";

export const getEmployees = async <T extends object>(params: T) => {
  const { data } = await service.get<Response<IEmployee[]>>(prefix, { params });
  return data;
};

export default { getEmployees };
