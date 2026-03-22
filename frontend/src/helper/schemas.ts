import { z } from "zod";
import { EmployeeRole } from "~/interfaces/employee.interface";

export type TZodSchema<T> = z.infer<T>;

export const employeeSchema = z.object({
  username: z
    .string({ error: "username is required" })
    .min(3, "username must be at least 3 characters")
    .max(20, "username too long"),
  name: z
    .string({ error: "name is required" })
    .min(1, "name is required")
    .max(10, "name too long"),
  role: EmployeeRole,
  password: z.string().min(4, "password must be at least 4 characters"),
});
