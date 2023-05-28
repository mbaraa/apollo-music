import { loadStripe } from "@stripe/stripe-js";
import Desktop from "./Desktop";
import Mobile from "./Mobile";
import { IPlan, startSubscription } from "./shared";

const stripePromise = loadStripe(import.meta.env.VITE_STRIPE_PUBLIC_KEY);

export default function SelectPlan() {
  const selectPlan = async (plan: IPlan, cardToken: string) => {
    await startSubscription(plan, cardToken);
  };
  return (
    <div>
      <div className="md:hidden">
        <Mobile startSubscription={selectPlan} stripePromise={stripePromise} />
      </div>
      <div className="hidden md:block">
        <Desktop />
      </div>
    </div>
  );
}
