import type { IMenu, IVariation } from "~/interfaces/menu.interface";
import { service as svc } from ".";

type Response<Data = unknown> = { code: number; data: Data };

const prefix = "menu";

const getMenus = async () => {
  const { data } = await svc.get<Response<IMenu[]>>(prefix);
  return data;
};

export const getVariations = async <T extends object>(params: T) => {
  const { data } = await svc.get<Response<IVariation[]>>(
    `${prefix}/variation`,
    { params }
  );
  return data;
};

export default { getMenus, getVariations };
