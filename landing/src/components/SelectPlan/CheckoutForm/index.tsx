import { Stripe } from "@stripe/stripe-js";
import { IPlan } from "../shared";
import Free from "./Free";
import Paid from "./Paid";

interface Props {
  plan: IPlan;
  startSubscription(plan: IPlan, cardToken: string): void;
  stripePromise: Promise<Stripe | null>;
}

export default function CheckoutForm({
  plan,
  startSubscription,
  stripePromise,
}: Props) {
  if (plan.requiresPayment) {
    return (
      <Paid
        plan={plan}
        startSubscription={startSubscription}
        stripePromise={stripePromise}
      />
    );
  }
  return <Free plan={plan} startSubscription={startSubscription} />;
}
