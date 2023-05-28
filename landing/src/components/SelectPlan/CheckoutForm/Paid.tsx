import {
  useStripe,
  useElements,
  CardElement,
  Elements,
} from "@stripe/react-stripe-js";
import CardSection from "./CardSection";
import { Stripe, TokenResult } from "@stripe/stripe-js";
import { IPlan } from "../shared";
import PlanPreview from "./PlanPreview";

interface Props {
  plan: IPlan;
  startSubscription(plan: IPlan, cardToken: string): void;
}

interface Props2 extends Props {
  stripePromise: Promise<Stripe | null>;
}

function UnderlyingPaidCheckoutForm({ plan, startSubscription }: Props) {
  const stripe = useStripe();
  const elements = useElements();

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    if (!stripe || !elements) {
      return;
    }

    const card = elements.getElement(CardElement);
    let result: TokenResult;
    if (card) {
      result = await stripe.createToken(card);
    } else {
      return;
    }

    if (result.error) {
      window.alert("failed: " + result.error.message);
    } else {
      startSubscription(plan, result.token.id);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="p-[10px] mt-[25px]">
      <PlanPreview plan={plan} />
      <div className="my-[10px]">
        <span>Enter your card details:</span>
        <CardSection />
      </div>
      <input
        type="submit"
        className="bg-dark-accent text-dark-neutral w-full h-[48px] rounded-[20px] mt-[10px] text-[20px] cursor-pointer"
        value={"Checkout & Finish Signin up"}
      />
    </form>
  );
}

export default function Paid({
  plan,
  startSubscription,
  stripePromise,
}: Props2) {
  return (
    <>
      <Elements stripe={stripePromise}>
        <UnderlyingPaidCheckoutForm
          plan={plan}
          startSubscription={startSubscription}
        />
      </Elements>
    </>
  );
}
