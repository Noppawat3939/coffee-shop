import { useCallback, useState, useTransition } from "react";
import { Button, Flex, Typography } from "@mantine/core";
import { Minus, Plus } from "lucide-react";

type IncreaseDecreaseInputProps = {
  intialValue: number;
  onChange?: (value: number) => void;
};

export default function IncreaseDecreaseInput({
  intialValue,
  onChange,
}: IncreaseDecreaseInputProps) {
  const [, startTransition] = useTransition();

  const [count, setCount] = useState(intialValue);

  const increase = useCallback(() => {
    const updated = count + 1;

    updateCount(updated);
    onChange?.(updated);
  }, [count]);

  const decrease = useCallback(() => {
    const updated = count < 1 ? 0 : count - 1;

    updateCount(updated);
    onChange?.(updated);
  }, [count]);

  const updateCount = (c: number) => startTransition(() => setCount(c));

  return (
    <Flex align="center">
      <Button
        disabled={count < 1}
        role="decrease"
        onClick={decrease}
        variant="outline"
        size="xs"
        w="auto"
      >
        <Minus width={14} />
      </Button>
      <Flex w={38} h={30} align="center" justify="center">
        <Typography aria-label="count">{count}</Typography>
      </Flex>
      <Button role="increase" onClick={increase} size="xs">
        <Plus width={14} />
      </Button>
    </Flex>
  );
}
