import { Button, Card, Flex, Typography } from "@mantine/core";
import { PhoneNumberVerify } from ".";
import { useAxios } from "~/hooks";
import { member } from "~/services";
import { dateFormat } from "~/helper";
import { useDisclosure } from "@mantine/hooks";
import { useMemo } from "react";
import type { IMember } from "~/interfaces/member.interface";

type MemberCheckerProps = {
  withMemberClick?: (data?: IMember) => void;
  noMemberClick?: () => void;
};

export default function MemberChecker({
  withMemberClick,
  noMemberClick,
}: MemberCheckerProps) {
  const [display, { close: onHide, open: onDisplay }] = useDisclosure(false);

  const { execute: getMember, data: memberRes } = useAxios(member.getMember, {
    onFinish: onDisplay,
  });

  const haveMember = useMemo(
    () => typeof memberRes?.data?.phone_number === "string",
    [memberRes?.data]
  );

  return (
    <Flex direction="column" rowGap={10} w="100%">
      <PhoneNumberVerify
        onSubmit={getMember}
        onChange={(value) => {
          if (!value && memberRes?.code && display) {
            onHide();
          }
        }}
      />
      {display && haveMember && (
        <Card>
          <Flex gap={6} direction="column">
            <Typography>{`Name: ${memberRes?.data.full_name}`}</Typography>
            <Typography c="gray">{`Registered: ${dateFormat(memberRes?.data.created_at, "DD/MM/YYYY")}`}</Typography>
          </Flex>
        </Card>
      )}
      {display && !memberRes?.data?.phone_number && (
        <Typography c="gray">{"No member"}</Typography>
      )}
      {haveMember && (
        <Button onClick={() => withMemberClick?.(memberRes?.data)}>
          {"Continue"}
        </Button>
      )}
      <Button variant="outline" onClick={noMemberClick}>
        {"No member"}
      </Button>
    </Flex>
  );
}
