import { Elements } from "@stripe/react-stripe-js";
import CheckoutForm from "./CheckoutForm";
import { Token, loadStripe } from "@stripe/stripe-js";

const stripePromise = loadStripe(import.meta.env.VITE_STRIPE_PUBLIC_KEY);

interface CheckoutParams {
  handler(token: Token): void;
}

export default function Checkout({ handler }: CheckoutParams) {
  return (
    <Elements stripe={stripePromise}>
      <CheckoutForm stripeTokenHandler={(token) => handler(token)} />
    </Elements>
  );
}
