import { Flex, TextInput } from "@mantine/core";
import { useInputState } from "@mantine/hooks";
import { useCallback, type ChangeEventHandler } from "react";
import type { IMember } from "~/interfaces/member.interface";

type RegisterValues = Pick<IMember, "full_name" | "phone_number">;

type PhoneNumberVerifyProps = {
  onSubmit: (values: Pick<RegisterValues, "phone_number">) => void;
  autoReset?: boolean;
  onChange?: (value: string) => void;
};

export default function PhoneNumberVerify({
  onSubmit,
  autoReset = false,
  onChange,
}: PhoneNumberVerifyProps) {
  const [phone, setPhone] = useInputState("");

  const onPhoneChange = useCallback(
    (value: string) => {
      setPhone(value);
      onChange?.(value);
    },
    [phone, onChange]
  );

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

          if (autoReset) {
            e.currentTarget.reset();
          }
        }}
      >
        <TextInput
          value={phone}
          onChange={({ currentTarget: { value } }) => onPhoneChange(value)}
          label="Phone number"
          name="phone_number"
          type="tel"
          placeholder="Enter your phone number"
        />
      </form>
    </Flex>
  );
}
