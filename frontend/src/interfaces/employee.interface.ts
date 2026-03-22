export enum EmployeeRole {
  super_admin = "super_admin",
  admin = "admin",
  staff = "staff",
}

export interface IEmployee {
  id: number;
  username: string;
  name: string;
  active: boolean;
  role: EmployeeRole;
  created_at: string;
  updated_at: string;
}

export interface ICreateEmployee
  extends Pick<IEmployee, "username" | "name" | "role"> {
  password: string;
}
