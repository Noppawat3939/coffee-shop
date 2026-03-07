import type { ZodSchema } from "zod";

export function zodValidate<T extends ZodSchema>(schema: T) {
  return (values: Parameters<typeof schema.safeParse>[0]) => {
    const result = schema.safeParse(values);

    if (result.success) return {};

    const errors: any = {};

    result.error.issues.forEach((issue) => {
      let current = errors;

      issue.path.forEach((key, index) => {
        if (index === issue.path.length - 1) {
          current[key] = issue.message;
        } else {
          current[key] = current[key] || {};
          current = current[key];
        }
      });
    });

    return errors;
  };
}
