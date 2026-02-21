import {
  Button,
  Card,
  Container,
  Flex,
  ScrollArea,
  Stack,
  Text,
  Typography,
} from "@mantine/core";
import { useNavigate } from "@tanstack/react-router";
import { Coffee } from "lucide-react";
import { useCallback, type PropsWithChildren } from "react";
import {
  PORTAL_HEADER_WITH_PATHNAME,
  PORTAL_WITH_PATHNAME,
} from "~/helper/constant";
import { useAuth } from "~/hooks";

type TBaseMenu = { label: string; pathname?: string };

type TMenu = TBaseMenu & { children?: TBaseMenu[] };

const MENUS = [
  {
    label: "Transactions",
    children: [
      { label: "Payments", pathname: PORTAL_WITH_PATHNAME.payments },
      { label: "Orders", pathname: PORTAL_WITH_PATHNAME.orders },
    ],
  },
] satisfies TMenu[];

type PortalLayoutProps = Readonly<PropsWithChildren & { pathname: string }>;

export default function PortalLayout({
  children,
  pathname,
}: PortalLayoutProps) {
  const { username } = useAuth();
  const navigation = useNavigate();

  const goTo = useCallback(
    (path: string) => navigation({ to: `/${path}` }),
    []
  );

  return (
    <Container
      bg={"#F8FAFC"}
      itemID="portal-layout"
      pr={16}
      pl={0}
      maw={2400}
      miw={1024}
    >
      <Stack gap={0}>
        <Flex gap={10} justify="space-between">
          <Flex
            style={{ borderTopRightRadius: 54, borderBottomRightRadius: 54 }}
            direction="column"
            flex={0.15}
            itemID="sidebar-menus"
            bg={"#1A1A19"}
            c="gray.1"
          >
            <Flex px={24} align="center" h="5vh">
              <Flex align="center" gap={5}>
                <Coffee color="#228be6" />
                <Flex direction="column">
                  <Text size="sm" fw="bolder">
                    Portal
                  </Text>
                  <Text c="blue" style={{ fontSize: 9 }} fw={"bolder"}>
                    {username}
                  </Text>
                </Flex>
              </Flex>
            </Flex>
            <ScrollArea h="95dvh" p={24}>
              {MENUS.map((menu, i) => {
                if (menu?.children) {
                  return (
                    <div key={`menu-${i}`}>
                      <Text size="sm" fw="bolder" key={`root-${i}`}>
                        {menu.label}
                      </Text>
                      {menu.children.map((sub, i) => (
                        <Button
                          onClick={() =>
                            sub.pathname !== pathname && goTo(sub.pathname)
                          }
                          w="100%"
                          key={`sub-${i}`}
                          size="sm"
                          display="flex"
                          justify="flex-start"
                          type="button"
                          {...(sub.pathname === pathname
                            ? {
                                c: "blue",
                              }
                            : { c: "gray.7" })}
                          variant={"subtle"}
                        >
                          {sub.label}
                        </Button>
                      ))}
                    </div>
                  );
                }
              })}
            </ScrollArea>
          </Flex>
          <Flex direction="column" flex={0.85}>
            <Flex
              bg={"#F8FAFC"}
              align="center"
              h="5vh"
              pos="sticky"
              top={0}
              style={{ zIndex: 10 }}
            >
              <Typography fw="bold">
                {PORTAL_HEADER_WITH_PATHNAME[pathname]}
              </Typography>
            </Flex>
            <Card withBorder>
              <ScrollArea h={"91dvh"}>{children}</ScrollArea>
            </Card>
          </Flex>
        </Flex>
      </Stack>
    </Container>
  );
}
