export interface ITimestampt {
  created_at: string;
  updated_at: string;
  deleted_at: null | string;
}

export interface IMember extends ITimestampt {
  id: number;
  full_name: string;
  phone_number: string;
  provider: string;
}
