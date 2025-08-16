import { Accordion, Typography, Flex, Image, ActionIcon } from "@mantine/core";
import { useMap } from "@mantine/hooks";
import { useNavigate } from "@tanstack/react-router";
import { Coffee } from "lucide-react";
import { Fragment, useEffect } from "react";
import { IncreaseDecreaseInput, MainLayout } from "~/components";
import { priceFormat, sum } from "~/helper";
import { useAxios } from "~/hooks";
import { apis } from "~/services";

export default function MenusPage() {
  const { execute, data } = useAxios(apis.getMenus);

  const navigation = useNavigate();

  const orderMap = useMap();

  useEffect(() => {
    return () => {
      execute();
    };
  }, []);

  const sumOrders = sum([...orderMap.values()] as number[]);

  const goToCheckout = () => {
    const searchParams = {} as Record<string, number>;

    for (const [key, value] of orderMap.entries()) {
      const [, varitaionId] = String(key).split("_");
      const amount = Number(value);

      if (amount > 0) {
        searchParams[`varitaion_id_${varitaionId}`] = amount;
      }
    }

    navigation({ to: "/checkout", search: searchParams });
  };

  return (
    <MainLayout
      title={"Menu"}
      extra={
        <ActionIcon
          variant={sumOrders > 0 ? "filled" : "outline"}
          onClick={goToCheckout}
        >
          <Coffee width={14} />
        </ActionIcon>
      }
    >
      <Accordion mt={12} variant="filled" transitionDuration={300}>
        {data?.data &&
          data.data.map((item) => (
            <Accordion.Item key={item.id} value={item.id.toString()}>
              <Accordion.Control fz="h2" tt="capitalize">
                {item.name}
              </Accordion.Control>
              <Accordion.Panel>
                {item.variations &&
                  item.variations.map((variation, idx) => (
                    <Fragment key={`variation-${idx}`}>
                      <Flex py={4} px={8} columnGap={24} mb={5}>
                        {variation.image && (
                          <Image
                            src={variation.image}
                            alt="menu"
                            h={100}
                            w={100}
                            radius={12}
                            style={{ objectFit: "cover" }}
                            loading="lazy"
                          />
                        )}
                        <Flex direction="column" justify="center" rowGap={2}>
                          <Typography fz="h4" tt="capitalize">
                            {variation.type}
                          </Typography>
                          <Typography c="gray" opacity={0.7} fz="xs">
                            {item.description}
                          </Typography>
                          <Typography fz="h3" fw={500} tt="capitalize">
                            {priceFormat(variation.price)}
                          </Typography>
                        </Flex>
                        <IncreaseDecreaseInput
                          intialValue={0}
                          onChange={(value) => {
                            const name = `${variation.menu_id}_${variation.id}`;

                            orderMap.set(name, value);
                          }}
                        />
                      </Flex>
                    </Fragment>
                  ))}
              </Accordion.Panel>
            </Accordion.Item>
          ))}
      </Accordion>
    </MainLayout>
  );
}
