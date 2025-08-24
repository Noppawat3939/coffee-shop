import {
  Button,
  Card,
  Flex,
  Image,
  TextInput,
  Typography,
} from "@mantine/core";
import { CircleCheck } from "lucide-react";
import { memo, useState, useTransition } from "react";
import { useAxios, useNotification } from "~/hooks";
import type { IMember } from "~/interfaces/member.interface";
import { member } from "~/services";

type RegisterValues = Pick<IMember, "full_name" | "phone_number">;

export default function MembershipPage() {
  const [md, ctx] = useNotification();

  const [pending, startTransition] = useTransition();

  const [isOpenRegisterForm, setIsOpenRegisterForm] = useState(false);

  const { execute: register, loading } = useAxios(member.register);

  const { execute: getMember, data: memberRes } = useAxios(member.getMember);

  return (
    <Flex
      align="center"
      direction="column"
      py={24}
      maw={380}
      h={"100dvh"}
      mx="auto"
    >
      <Typography fw="bold" fz="xl">
        {"Membership Coffee Shop"}
      </Typography>
      <Image
        my={16}
        alt="banner"
        loading="lazy"
        src="https://images.unsplash.com/photo-1552010266-6458fda4d692?q=80&w=1025&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
        w={280}
      />
      {isOpenRegisterForm ? (
        <Flex direction="column" w="100%">
          <form
            autoCapitalize="off"
            autoComplete="off"
            style={{ width: "100%" }}
            onSubmit={(e) => {
              e.preventDefault();

              let values = {} as RegisterValues;

              startTransition(async () => {
                const formData = new FormData(e.currentTarget);

                values = Object.fromEntries(formData) as RegisterValues;

                if (!values?.full_name || !values?.phone_number.trim()) {
                  md.open({ title: "Please enter all fields" });
                } else {
                  const res = await register(values);
                  if (res?.data) {
                    md.open({
                      title: "Register success",
                    });

                    e.currentTarget.reset();
                  }
                }
              });
            }}
          >
            <TextInput
              name="full_name"
              withAsterisk
              label="Full name"
              placeholder="Enter your full name"
              mb={8}
            />
            <TextInput
              withAsterisk
              name="phone_number"
              label="Phone number"
              placeholder="Enter your phone number"
            />
            <Flex justify="center" columnGap={10} mt={24}>
              <Button loading={pending} w={100} type="submit">
                {"Save"}
              </Button>
              <Button
                w={100}
                type="reset"
                variant="outline"
                onClick={() => setIsOpenRegisterForm((prev) => !prev)}
              >
                {"Cancel"}
              </Button>
            </Flex>
          </form>
        </Flex>
      ) : (
        <Flex direction="column" rowGap={10} w="100%">
          <PhoneNumberVerify onSubmit={getMember} />
          <NoMember
            data={memberRes?.data}
            onClick={() => setIsOpenRegisterForm((prev) => !prev)}
            loading={loading}
          />
        </Flex>
      )}

      {ctx}
    </Flex>
  );
}

const NoMember = memo(function ({
  onClick,
  data,
  loading,
}: Partial<{ onClick: () => void; data?: IMember; loading: boolean }>) {
  return (
    <Flex direction="column" gap={12}>
      {data?.phone_number ? (
        <Card>
          <Flex columnGap={6}>
            <CircleCheck color="teal" />
            <Typography>{`${data.full_name} (${data.phone_number})`}</Typography>
          </Flex>
        </Card>
      ) : (
        <Typography c="gray">{"No membership"}</Typography>
      )}
      <Button
        onClick={onClick}
        loading={loading}
        type="button"
        variant="filled"
      >
        {"Register now"}
      </Button>
    </Flex>
  );
});

const PhoneNumberVerify = memo(function ({
  onSubmit,
}: {
  onSubmit: (value: Pick<RegisterValues, "phone_number">) => void;
}) {
  return (
    <Flex w="100%">
      <form
        style={{ width: "100%" }}
        autoComplete="off"
        autoCapitalize="off"
        onSubmit={(e) => {
          e.preventDefault();

          const values = Object.fromEntries(
            new FormData(e.currentTarget)
          ) as Pick<RegisterValues, "phone_number">;

          if (!values.phone_number.trim()) return;

          onSubmit(values);
          e.currentTarget.reset();
        }}
      >
        <TextInput
          label="Phone number"
          name="phone_number"
          type="tel"
          placeholder="Enter your phone number"
        />
      </form>
    </Flex>
  );
});
