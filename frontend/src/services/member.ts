import type {
  IMember,
  IMemberWithMemberPoint,
} from "~/interfaces/member.interface";
import { service as svc } from ".";
import type { Response } from "./service-instance";

type PickedMember = Pick<IMember, "full_name" | "phone_number">;

const prefix = "Members";

const register = async (body: PickedMember) => {
  const { data } = await svc.post<Response<IMember>>(
    `${prefix}/register`,
    body
  );
  return data;
};

const getMember = async (body: PickedMember) => {
  const { data } = await svc.post<Response<IMember>>(`${prefix}/find`, body);
  return data;
};

const getMembers = async <T extends object>(params?: PickedMember | T) => {
  const { data } = await svc.get<Response<IMemberWithMemberPoint[]>>(prefix, {
    params,
  });
  return data;
};

export default { register, getMember, getMembers };
