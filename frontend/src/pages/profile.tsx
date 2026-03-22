import { Fragment, useCallback, useState } from "react";
import { Stack, TextInput, Button, Flex, Divider } from "@mantine/core";
import { MainLayout } from "~/components";
import { dateFormat } from "~/helper";
import { useNavigate } from "@tanstack/react-router";
import useProfile from "~/hooks/use-profile-hooks";
import { useForm } from "@mantine/form";
import { zodValidate } from "~/helper/form";
import { profileSchema, type TZodSchema } from "~/helper/schemas";

type FormValues = Partial<TZodSchema<typeof profileSchema>>;

export default function Page() {
  const navigate = useNavigate();
  const { data: user } = useProfile();

  const form = useForm<FormValues>({
    mode: "uncontrolled",
    validate: zodValidate(profileSchema),
    initialValues: {
      username: user?.username,
      name: user?.name,
      role: user?.role,
    },
  });

  const [isEditing, setIsEditing] = useState(false);
  const [displayChangePassword, setDisplayChangePassword] = useState(false);

  const handleSubmit = (value: FormValues) => {
    console.log("submit !", value);
    // if password changed should be re-login
  };

  const handleBack = () => {
    form.reset();
    navigate({ to: "/menus", replace: true });
  };

  const onToggleEdit = useCallback(() => {
    isEditing && setDisplayChangePassword(false);

    setIsEditing((prev) => !prev);
  }, [isEditing, displayChangePassword]);

  const onToggleChangePassword = useCallback(() => {
    if (!displayChangePassword) {
      form.setValues({ password: "", confirm_password: "" });
      setIsEditing(true);
    }

    setDisplayChangePassword((prev) => !prev);
  }, [isEditing, displayChangePassword, form]);

  return (
    <MainLayout
      title="Profile"
      extra={
        <Button onClick={onToggleEdit} size="sm" variant="outline">
          {isEditing ? "Editing" : "Edit"}
        </Button>
      }
    >
      <Stack gap="lg">
        <form onSubmit={form.onSubmit(handleSubmit)}>
          <Stack>
            <TextInput
              key={form.key("username")}
              label="Username"
              {...(!isEditing && { variant: "unstyled" })}
              {...(isEditing ? { disabled: true } : { readOnly: true })}
              {...form.getInputProps("username")}
            />

            <TextInput
              key={form.key("name")}
              label="Name"
              {...(!isEditing && { variant: "unstyled" })}
              {...form.getInputProps("name")}
            />

            <TextInput
              key={form.key("role")}
              label="Role"
              {...(isEditing ? { disabled: true } : { readOnly: true })}
              {...(!isEditing && { variant: "unstyled" })}
              {...form.getInputProps("role")}
            />

            <TextInput
              label="Registered at"
              value={dateFormat(user?.created_at, "DD/MM/YYYY")}
              {...(isEditing ? { disabled: true } : { readOnly: true })}
              {...(!isEditing && { variant: "unstyled" })}
            />

            <Divider />
            <Button
              onClick={onToggleChangePassword}
              size="sm"
              variant={displayChangePassword ? "subtle" : "outline"}
            >
              {displayChangePassword
                ? "Close change password"
                : "Change password"}
            </Button>

            {displayChangePassword && (
              <Fragment>
                <TextInput
                  required
                  type="password"
                  key={form.key("password")}
                  label="New password"
                  {...form.getInputProps("password")}
                />
                <TextInput
                  required
                  type="password"
                  key={form.key("confirm_password")}
                  label="Confirm password"
                  {...form.getInputProps("confirm_password")}
                />
              </Fragment>
            )}

            <Flex mt={32} justify="center" gap={20}>
              <Button onClick={handleBack} w={150} variant="outline">
                Back
              </Button>
              <Button disabled={!isEditing} type="submit" w={150}>
                Save
              </Button>
            </Flex>
          </Stack>
        </form>
      </Stack>
    </MainLayout>
  );
}
