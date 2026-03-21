import { useRouter } from "@tanstack/react-router";
import { Route } from "~/routes/account/employees";
import { useAxios } from ".";
import { employee } from "~/services";
import type { ICreateEmployee } from "~/interfaces/employee.interface";

// Query employees
export function useEmployees() {
  const { invalidate } = useRouter();

  const { data: initialData } = Route.useLoaderData();

  return { data: initialData, refetch: invalidate };
}

// Create employee
export function useCreateEmployee() {
  const { invalidate } = useRouter();

  const { execute } = useAxios(employee.createEmployee);

  return (data: ICreateEmployee) => {
    execute(data).finally(() => invalidate());
  };
}
