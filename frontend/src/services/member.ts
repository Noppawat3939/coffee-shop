import type { IMember } from "~/interfaces/member.interface";
import { service as svc } from ".";
import type { Response } from "./service-instance";

const prefix = "members";

const register = async (body: Pick<IMember, "full_name" | "phone_number">) => {
  const { data } = await svc.post<Response<IMember>>(
    `${prefix}/register`,
    body
  );
  return data;
};

const getMember = async (body: Pick<IMember, "phone_number">) => {
  const { data } = await svc.post<Response<IMember>>(prefix, body);
  return data;
};

export default { register, getMember };
