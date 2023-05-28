import { IPlan } from "./shared";

interface PlanProps {
  plan: IPlan;
}

export default function Plan({ plan }: PlanProps) {
  return (
    <div className="p-[20px] w-[300px] h-[200px] border-[2px] border-dark-accent bg-dark-accent2 rounded-[20px] grid grid-cols-1 place-items-center">
      <h1 className="text-[30px] block">{plan.title}</h1>
      <h2 className="text-[25px] block">
        {plan.price}
        {plan.requiresPayment ? "/month" : ""}
      </h2>
      <h3 className="text-[16px] block">
        {plan.size} of cloud storage for your music library
      </h3>
    </div>
  );
}
