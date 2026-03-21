export interface IEmployee {
  id: number;
  username: string;
  name: string;
  active: boolean;
  role: string;
  created_at: string;
  updated_at: string;
}

export interface ICreateEmployee
  extends Pick<IEmployee, "username" | "name" | "role"> {
  password: string;
}
