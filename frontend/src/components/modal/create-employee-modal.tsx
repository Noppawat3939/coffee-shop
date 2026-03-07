import {
  Button,
  Group,
  Modal,
  Select,
  Stack,
  Text,
  TextInput,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { z } from "zod";
import { zodValidate } from "~/helper/form";

type FormValues = z.infer<typeof employeeSchema>;
type CreateEmployeeModalProps = {
  open: boolean;
  onClose: () => void;
};

const employeeSchema = z.object({
  username: z
    .string({ error: "username is required" })
    .min(3, "username must be at least 3 characters")
    .max(20, "username too long"),
  name: z
    .string({ error: "name is required" })
    .min(1, "name is required")
    .max(10, "name too long"),
  role: z.enum(["admin", "super_admin", "staff"]),
  password: z.string().min(4, "password must be at least 4 characters"),
});

export default function CreateEmployeeModal({
  onClose,
  open,
}: CreateEmployeeModalProps) {
  const form = useForm<FormValues>({
    mode: "uncontrolled",
    validate: zodValidate(employeeSchema),
    initialValues: { username: "", name: "", role: "admin", password: "" },
  });

  const handleSubmit = (value: typeof form.values) => {
    console.log(value);
  };

  const handleClose = () => {
    form.reset();
    onClose();
  };

  return (
    <Modal
      title={<Text fw="bold">New Employee</Text>}
      centered
      opened={open}
      onClose={handleClose}
    >
      <Stack>
        <form onSubmit={form.onSubmit(handleSubmit)}>
          <TextInput
            label="Username"
            key={form.key("username")}
            {...form.getInputProps("username")}
          />
          <TextInput
            label="Name"
            key={form.key("name")}
            {...form.getInputProps("name")}
          />
          <TextInput
            type="password"
            label="Password"
            key={form.key("password")}
            {...form.getInputProps("password")}
          />
          <Select
            label="Role"
            data={[
              { value: "staff", label: "Staff" },
              { value: "admin", label: "Admin" },
              { value: "super_admin", label: "Super Admin" },
            ]}
            key={form.key("role")}
            {...form.getInputProps("role")}
          />
          <Group justify="center" mt="xl">
            <Button onClick={handleClose} variant="outline" w={150}>
              Cancel
            </Button>
            <Button type="submit" w={150}>
              Save
            </Button>
          </Group>
        </form>
      </Stack>
    </Modal>
  );
}
