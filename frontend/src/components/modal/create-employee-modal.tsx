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
import { zodValidate } from "~/helper/form";
import { employeeSchema, type TZodSchema } from "~/helper/schemas";

type FormValues = TZodSchema<typeof employeeSchema>;
type CreateEmployeeModalProps = {
  open: boolean;
  onClose: () => void;
};

const { options: roles } = employeeSchema.shape.role;

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
            data={roles.map((role) => ({ label: role, value: role }))}
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
