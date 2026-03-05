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
import z from "zod";
import { zodResolver } from "mantine-form-zod-resolver";

type FormValues = z.infer<typeof employeeSchema>;
type CreateEmployeeModalProps = {
  open: boolean;
  onClose: () => void;
};

const employeeSchema = z.object({
  username: z
    .string()
    .min(3, "username must be at least 3 characters")
    .max(50, "username too long"),
  name: z.string().min(1, "name is required").max(100, "name too long"),
  role: z.enum(["admin", "super_admin", "staff"]),
  password: z.string().min(4, "password must be at least 4 characters"),
});

export default function CreateEmployeeModal({
  onClose,
  open,
}: CreateEmployeeModalProps) {
  const form = useForm<FormValues>({
    mode: "uncontrolled",
    validate: zodResolver(employeeSchema),
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
            data={["staff", "admin", "super_admin"]}
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
