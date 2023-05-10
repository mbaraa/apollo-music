import { useStripe, useElements, CardElement } from "@stripe/react-stripe-js";
import CardSection from "./CardSection";
import { Token, TokenResult } from "@stripe/stripe-js";

interface CheckoutFormParams {
  stripeTokenHandler(token: Token): void;
}

export default function CheckoutForm({
  stripeTokenHandler,
}: CheckoutFormParams) {
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
      stripeTokenHandler(result.token);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <CardSection />
      <button onClick={() => {}}>Checkout & Finish Signup</button>
    </form>
  );
}
