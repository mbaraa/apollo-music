import { Elements } from "@stripe/react-stripe-js";
import CheckoutForm from "./CheckoutForm";
import { Token, loadStripe } from "@stripe/stripe-js";
import { useState } from "react";

const stripePromise = loadStripe(import.meta.env.VITE_STRIPE_PUBLIC_KEY);

interface CheckoutParams {
  handler(token: Token): void;
}

enum Plan {
  Free = "FREE",
  TwintyGB = "20GB",
  HundredGB = "100GB",
}

function getPlan(planText: string): Plan {
  switch (planText) {
    case "20GB":
      return Plan.TwintyGB;
    case "100GB":
      return Plan.HundredGB;
    case "FREE":
    default:
      return Plan.Free;
  }
}

export default function Checkout({ handler }: CheckoutParams) {
  handler({} as Token);
  const [plan, setPlan] = useState(Plan.Free);
  const startSubscription = async (token: Token) => {
    await fetch(`${import.meta.env.VITE_BACKEND_ADDRESS}/subscription/start`, {
      method: "POST",
      mode: "cors",
      body: JSON.stringify({
        cardToken: token.id,
        plan: plan,
      }),
      headers: {
        Authorization: localStorage.getItem("checkoutToken") ?? "",
      },
    })
      .then((resp) => resp.json())
      .then((resp) => {
        localStorage.removeItem("checkoutToken");
        console.log(resp);
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <>
      <select
        onChange={(e) => {
          setPlan(getPlan(e.target.value));
        }}
      >
        {Object.values(Plan).map((v, i) => (
          <option key={i}>{v}</option>
        ))}
      </select>
      <Elements stripe={stripePromise}>
        <CheckoutForm
          stripeTokenHandler={(token) => startSubscription(token)}
        />
      </Elements>
    </>
  );
}
