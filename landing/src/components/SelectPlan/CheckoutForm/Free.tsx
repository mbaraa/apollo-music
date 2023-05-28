import { useStripe, useElements, CardElement } from "@stripe/react-stripe-js";
import { TokenResult } from "@stripe/stripe-js";
import { IPlan } from "../shared";
import PlanPreview from "./PlanPreview";

interface Props {
  plan: IPlan;
  startSubscription(plan: IPlan, cardToken?: string): void;
}

export default function Free({ plan, startSubscription }: Props) {
  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    startSubscription(plan, "");
  };

  return (
    <form onSubmit={handleSubmit} className="p-[10px] mt-[25px]">
      <PlanPreview plan={plan} />
      <input
        type="submit"
        className="bg-dark-accent text-dark-neutral w-full h-[48px] rounded-[20px] mt-[25px] text-[20px] cursor-pointer"
        value={"Finish Signin up"}
      />
    </form>
  );
}
