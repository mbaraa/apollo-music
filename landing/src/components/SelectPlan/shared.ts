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
    .then(async (resp) => {
      const respBody = await resp.json();
      if (!resp.ok) {
        window.alert(respBody["errorMsg"]);
        return;
      }
      localStorage.removeItem("checkoutToken");
      const sessionToken = respBody["data"]["token"];
      localStorage.setItem("token", sessionToken);
      window.open(
        `${import.meta.env.VITE_APP_ADDRESS}/sign-in?token=${sessionToken}`,
        "_self"
      );
    })
    .catch((err) => {
      window.alert(err);
    });
}
