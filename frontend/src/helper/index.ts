export const priceFormat = (amount: number | string) => {
  return new Intl.NumberFormat("th-TH", {
    style: "currency",
    currency: "THB",
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  }).format(+amount);
};

export const sum = (nums: number[]) => {
  return nums.reduce((total, cur) => total + cur, 0);
};
