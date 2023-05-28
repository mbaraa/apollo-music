import CheckoutForm from "./CheckoutForm";
import { useState } from "react";
import Plan from "./Plan";
import { IPlan, getPlans } from "./shared";
import { Stripe } from "@stripe/stripe-js";

interface Props {
  startSubscription(plan: IPlan, cardToken: string): void;
  stripePromise: Promise<Stripe | null>;
}

export default function Mobile({ startSubscription, stripePromise }: Props) {
  const [plans] = useState<IPlan[]>(getPlans());
  const [plan, setPlan] = useState(plans[0]);
  const [planSelected, setPlanSelected] = useState(false);

  return (
    <>
      <div className="w-[100vw] h-[100vh] bg-dark-primary text-dark-secondary p-[15px]">
        {!planSelected && (
          <>
            <h1 className="text-[30px] p-[15px]">
              Choose the library size that suites you the most.
            </h1>
            {plans.map((p, i) => (
              <button
                key={i}
                className="p-[15px]"
                onClick={() => {
                  setPlanSelected(true);
                  setPlan(p);
                }}
              >
                <Plan plan={p} />
              </button>
            ))}
          </>
        )}
        {planSelected && (
          <>
            <div className="flex justify-between items-center">
              <button
                onClick={() => {
                  setPlanSelected(false);
                }}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  strokeWidth={2.5}
                  stroke="currentColor"
                  className="w-6 h-6 text-dark-accent"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M15.75 19.5L8.25 12l7.5-7.5"
                  />
                </svg>
              </button>

              <h1 className="text-[15px]">Checkout</h1>
              <div />
            </div>
            <CheckoutForm
              plan={plan}
              startSubscription={startSubscription}
              stripePromise={stripePromise}
            />
          </>
        )}
      </div>
    </>
  );
}
