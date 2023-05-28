import { CardElement } from "@stripe/react-stripe-js";

const inputStyle =
  "font-primary bg-[#f9f9f9] w-full h-[72px] rounded-[8px] p-[26px] text-[14px] placeholder:text-[15px]" +
  " " +
  "font-medium text-[#444] placeholder:font-normal border-[2px] border-[#f9f9f9] tracking-[0.02em]" +
  " " +
  "placeholder:text-[#AEAEAE] focus:border-[#f4f4f4]";

function CardSection() {
  return (
    <div>
      <CardElement options={{ hidePostalCode: true }} className={inputStyle} />
    </div>
  );
}
export default CardSection;
