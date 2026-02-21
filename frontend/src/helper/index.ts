import dayjs, { Dayjs } from "dayjs";

type TFormatDay = string | number | Dayjs | null | undefined;

/**
 * Format a number or string as Thai Baht currency.
 *
 * @param {number | string} amount - The amount to format. Can be a number or a string representing a number.
 * @returns {string} The formatted currency string in THB, e.g. "฿1,234.00".
 *
 * @example
 * priceFormat(1234)       // "฿1,234.00"
 * priceFormat("5678.9")   // "฿5,678.90"
 */
export const priceFormat = (amount: number | string) => {
  return new Intl.NumberFormat("th-TH", {
    style: "currency",
    currency: "THB",
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  }).format(+amount);
};

/**
 * Sum an array of numbers.
 *
 * @param {number[]} nums - The array of numbers to sum.
 * @returns {number} The total sum of all numbers in the array.
 *
 * @example
 * sum([1, 2, 3]) // 6
 * sum([])        // 0
 */
export const sum = (nums: number[]): number => {
  return nums.reduce((total, cur) => total + cur, 0);
};

export const dateFormat = (date?: TFormatDay, format?: string) =>
  dayjs(date).format(format);

export const randomUniqueID = () =>
  (Math.floor(Math.random() * 100) + Date.now()).toString();

export const isExpired = <T>(day: T) => dayjs().isAfter(dayjs(day as Dayjs));

export const isBeforeToDay = (date?: TFormatDay) =>
  dayjs(date).isBefore(dayjs());
