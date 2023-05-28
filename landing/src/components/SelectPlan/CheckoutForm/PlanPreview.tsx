import { IPlan } from "../shared";

interface Props {
  plan: IPlan;
}

function getNextPaymentDate(): string {
  const date = new Date();
  date.setHours(24 * 30);
  const fullDate = date.toUTCString();

  return fullDate
    .substring(fullDate.indexOf(",") + 1, fullDate.indexOf(":") - 2)
    .trim()
    .split(" ")
    .join("/");
}

export default function PlanPreview({ plan }: Props) {
  return (
    <>
      <h1 className="text-[20px] font-[IBM Plex Sans]">Your selected plan</h1>
      <div>
        <div className="bg-dark-neutral rounded-t-[20px] mt-[15px] p-[25px] flex justify-between">
          <span>{plan.title}</span>
          <span>{plan.size}</span>
        </div>
        <div className="bg-dark-accent2 rounded-b-[20px] p-[25px] border-r-[2px] border-l-[2px] border-b-[2px] border-dark-accent text-[15px]">
          {!plan.requiresPayment && (
            <>
              <h1>What you get:</h1>
              <ul className="list-disc ml-[15px]">
                <li>Free of charge</li>
                <li>Upgrade or cancel at anytime</li>
                <li>
                  Enjoy all of the features with the selected storage size
                </li>
              </ul>
            </>
          )}
          {plan.requiresPayment && (
            <>
              <h1 className="mb-[10px] font-[800]">
                Starting today with {plan.price}
              </h1>
              <h1>What you get:</h1>
              <ul className="list-disc ml-[15px]">
                <li>
                  Next payment due date <br /> <b>{getNextPaymentDate()}</b>
                </li>
                <li>Upgrade or cancel at anytime</li>
                <li>
                  Enjoy all of the features with the selected storage size
                </li>
              </ul>
            </>
          )}
        </div>
      </div>
    </>
  );
}
