import type { IMenu } from "~/interfaces/menu.interface";
import { service } from ".";

type Response<Data = unknown> = { code: number; data: Data };

const getMenus = async () => {
  const { data } = await service.get<Response<IMenu[]>>("menu");
  return data;
};

export default { getMenus };
