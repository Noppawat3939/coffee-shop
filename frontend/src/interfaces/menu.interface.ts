export interface ITimestampt {
  created_at: string;
  updated_at: string;
  deleted_at: null | string;
}

export interface IMenu extends ITimestampt {
  id: number;
  name: string;
  description: string | null;
  is_available: boolean;
  variations?: IVariation[];
}

export interface IVariation extends ITimestampt {
  id: number;
  menu_id: number;
  price: number;
  type: string;
  image: string | null;
}
