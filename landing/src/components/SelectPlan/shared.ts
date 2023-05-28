export interface IPlan {
  title: string;
  price: string;
  size: string;
  requiresPayment: boolean;
  planEnum: EPlan;
}

export enum EPlan {
  Free = "FREE",
  TwintyGB = "20GB",
  HundredGB = "100GB",
}

const plans: IPlan[] = [
  {
    title: "Free",
    price: "$0.00",
    size: "500MB",
    requiresPayment: false,
    planEnum: EPlan.Free,
  },
  {
    title: "Starter",
    price: "$3.00",
    size: "20GB",
    requiresPayment: true,
    planEnum: EPlan.TwintyGB,
  },
  {
    title: "Premium",
    price: "$6.00",
    size: "100GB",
    requiresPayment: true,
    planEnum: EPlan.HundredGB,
  },
];

export function getPlans(): IPlan[] {
  return plans;
}

export async function startSubscription(plan: IPlan, cardToken?: string) {
  await fetch(`${import.meta.env.VITE_BACKEND_ADDRESS}/subscription/start`, {
    method: "POST",
    mode: "cors",
    body: JSON.stringify({
      cardToken: cardToken ?? "",
      plan: plan.planEnum,
    }),
    headers: {
      Authorization: localStorage.getItem("checkoutToken") ?? "",
    },
  })
    .then(() => {
      localStorage.removeItem("checkoutToken");
      window.open("https://apollo-music.app/sign-in", "_self");
    })
    .catch((err) => {
      console.error(err);
    });
}
