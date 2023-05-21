import { Elements } from "@stripe/react-stripe-js";
import CheckoutForm from "./CheckoutForm";
import { Token, loadStripe } from "@stripe/stripe-js";
import { useState } from "react";
import Plan from "./Plan";

const stripePromise = loadStripe(import.meta.env.VITE_STRIPE_PUBLIC_KEY);

interface IPlan {
  title: string;
  price: string;
  size: string;
  requiresPayment: boolean;
  planEnum: EPlan;
}

enum EPlan {
  Free = "FREE",
  TwintyGB = "20GB",
  HundredGB = "100GB",
}

export default function SelectPlan() {
  const [plans] = useState<IPlan[]>([
    {
      title: "Free",
      price: "$0.00",
      size: "500MB",
      requiresPayment: false,
      planEnum: EPlan.Free,
    },
    {
      title: "20GB",
      price: "$3.00/month",
      size: "20GB",
      requiresPayment: true,
      planEnum: EPlan.TwintyGB,
    },
    {
      title: "100GB",
      price: "$6.00/month",
      size: "100GB",
      requiresPayment: true,
      planEnum: EPlan.HundredGB,
    },
  ]);
  const [plan, setPlan] = useState(plans[0]);
  const [planSelected, setPlanSelected] = useState(false);

  const startSubscription = async (token: Token) => {
    await fetch(`${import.meta.env.VITE_BACKEND_ADDRESS}/subscription/start`, {
      method: "POST",
      mode: "cors",
      body: JSON.stringify({
        cardToken: token.id ?? "",
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
  };

  return (
    <div className="w-[100vw] h-[100vh] bg-dark-primary text-dark-secondary p-[15px]">
      <h1 className="text-[30px] p-[15px]">
        Choose the library size that suites you the most.
      </h1>
      {!planSelected &&
        plans.map((p) => (
          <button
            className="p-[15px]"
            onClick={() => {
              setPlanSelected(true);
              setPlan(p);
            }}
          >
            <Plan title={p.title} size={p.size} price={p.price} />
          </button>
        ))}
      {planSelected && (
        <>
          <Plan title={plan.title} size={plan.size} price={plan.price} />
          <button
            className="py-[20px] text-[18px] underline"
            onClick={() => {
              setPlanSelected(false);
              setPlan({} as IPlan);
            }}
          >
            Choose another plan
          </button>
        </>
      )}
      {planSelected && (
        <>
          {plan.requiresPayment ? (
            <Elements stripe={stripePromise}>
              <CheckoutForm
                stripeTokenHandler={(token) => startSubscription(token)}
              />
            </Elements>
          ) : (
            <button
              className="bg-dark-accent text-dark-neutral w-[330px] h-[48px] rounded-[20px] mt-[25px] text-[20px] cursor-pointer"
              onClick={() => startSubscription({} as Token)}
            >
              Finish Signin up
            </button>
          )}
        </>
      )}
    </div>
  );
}
